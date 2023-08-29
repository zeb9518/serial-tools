package test

// import (
// 	"changeme/handler"
// 	"testing"
// 	"time"
// )

// func TestGetSerial(t *testing.T) {

// 	serial := &handler.SerialHandler{}
// 	list := serial.GetSerial()
// 	println(list)

// }

// func TestOpenProt(t *testing.T) {
// 	serial := &handler.SerialHandler{}
// 	port, err := serial.OpenPort("COM3")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	println(port)
// }

// // test Write
// func TestWrite(t *testing.T) {
// 	serial := &handler.SerialHandler{}
// 	port, err := serial.OpenPort("COM3")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	_, err = port.Write([]byte("hello"))
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	serial.ClosePort()
// }

// func TestRead(t *testing.T) {
// 	serial := &handler.SerialHandler{}
// 	port, err := serial.OpenPort("COM3")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	defer serial.ClosePort()
// 	buf := make([]byte, 128)
// 	go func() {
// 		for {
// 			n, err := port.Read(buf)
// 			if err != nil {
// 				t.Error(err)
// 				return
// 			}
// 			println(string(buf[:n]))
// 		}
// 	}()

// 	// 让测试持续运行一段时间（例如：10秒）
// 	time.Sleep(10 * time.Second)
// }
