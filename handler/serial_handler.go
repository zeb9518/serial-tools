package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"go.bug.st/serial"
)

type SerialHandler struct {
}

var (
	serialPort serial.Port
)

// GetSerial is a mock function for serial.GetSerial
func (s *SerialHandler) GetSerial() []string {
	ports, err := serial.GetPortsList()
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return ports
}

// Open is a mock function for serial.OpenPort
func (s *SerialHandler) OpenPort(ctx context.Context, portName string) error {
	port, err := serial.Open(portName, &serial.Mode{
		BaudRate: 9600,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	})
	if err != nil {
		return err
	}
	serialPort = port
	errCh := make(chan error)
	// 读取模式状态位
	go func() {
		defer close(errCh) // 关闭通道
		for serialPort != nil {
			// 检查通道是否有错误，有错误则中断
			if err := <-errCh; err != nil {
				s.emitSerialError(ctx, err) // 通道有错误，发送错误
				return
			}
			// 获取模式状态位
			_, err := port.GetModemStatusBits()
			if err != nil {
				fmt.Println("获取模式状态位 Error:", port, err)
				errCh <- err // 发送错误到通道
				return
			}
			// 等待
			time.Sleep(time.Second * 1)
		}
	}()
	// 读取数据
	go func() {
		buf := make([]byte, 256)
		port.ResetInputBuffer() // 清空缓冲区
		for serialPort != nil {
			n, err := port.Read(buf)
			if err != nil {
				errCh <- err // 发送错误到通道
				return
			}
			if n == 0 {
				continue
			}
			s.emitSerialData(ctx, buf[:n])
			fmt.Println("Read:", string(buf[:n]))
		}
	}()

	return nil
}

// Close is a mock function for serial.Close
func (s *SerialHandler) ClosePort() {
	if serialPort != nil {
		serialPort.Close()
		serialPort = nil
	}
}

func (s *SerialHandler) emitSerialError(ctx context.Context, err error) {
	fmt.Println("Error:", err)
	runtime.EventsEmit(ctx, "onSerialError", err.Error())
}
func (s *SerialHandler) emitSerialData(ctx context.Context, data []byte) {
	dataStr := string(data)
	runtime.EventsEmit(ctx, "onSerialRead", dataStr)
}

// Read is a mock function for serial.Read
func (s *SerialHandler) Write([]byte) (int, error) {
	return 0, nil
}
