<script setup>
import { reactive, onMounted, toRefs, h } from 'vue'
import * as wailsApp from '../../wailsjs/go/main/App'
import { EventsOn } from '../../wailsjs/runtime'
import { useMessage, NButton, NDivider, NSelect, NSpace } from 'naive-ui'
import { Play12Filled, PlugDisconnected28Filled, ArrowReset20Filled } from '@vicons/fluent'

const message = useMessage()

const state = reactive({
  portName: null,
  portList: [],
  isConnect: false,
})

const { portName, portList } = toRefs(state)

function getSerialPortList() {
  wailsApp.GetSerial().then(result => {
    if (result.length !== 0) {
      state.portList = result.map((item, index) => {
        return {
          label: item,
          value: item,
        }
      })
    }
  })
}

onMounted(() => {
  getSerialPortList()
})


// 连接 
function handleOpen() {
  const { portName } = state
  if (!portName) {
    alert('请选择串口')
    return
  }
  wailsApp.OpenPort(portName).then(result => {
    console.log(result)
    state.isConnect = true
  }).catch(err => {
    console.log(err)
    handleSerialError(err)
  })
}

// 断开
function handleClose() {
  wailsApp.ClosePort().then(result => {
    console.log(result)
    state.isConnect = false
  })
}

EventsOn("onSerialError", handleSerialError)

function handleSerialError(err) {
  const errMap = new Map()
  errMap.set("Serial port busy", { type: 'warning', message: "串口已被占用" })
  errMap.set("Access is denied.", { type: 'error', message: "串口被拒绝访问" })
  if (errMap.has(err)) {
    state.isConnect = false
    const value = errMap.get(err)
    message[value.type](value.message)
  }
}


</script>

<template>
  <div>
    <div style="padding: 0 20px; margin-top: 10px;">
      <n-space item-style="display: flex; align-item: center;">
        <span style="margin: auto;">串口号：</span>
        <n-select v-model:value="portName" :options="portList" style="width: 200px;" :disabled="state.isConnect"
          placeholder="请选择串口" />
        <n-button type="info" v-if="!state.isConnect" :render-icon="() => h(Play12Filled)"
          @click="handleOpen">连接</n-button>
        <n-button type="error" v-else :render-icon="() => h(PlugDisconnected28Filled)" @click="handleClose">断开</n-button>
        <n-button :render-icon="() => h(ArrowReset20Filled)" @click="getSerialPortList">刷新</n-button>
      </n-space>
    </div>
    <n-divider title-placement="left">设备状态</n-divider>
    <h1>Dashboard</h1>
  </div>
</template>

<style lang='less' scoped></style>
