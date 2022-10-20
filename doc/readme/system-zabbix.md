# zabbix接入配置

PrometheusAlert For zabbix采用C/S方式。需要现在zabbix-server上部署zabbixclient客户端。

- 部署zabbixclient客户端，zabbixclient需要部署在和zabbix-server一起。

```
#进入到zabbix-server所在服务器
wget https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/zabbixclient/zabbixclient -O /usr/lib/zabbix/alertscripts/zabbixclient
chown zabbix:zabbix /usr/lib/zabbix/alertscripts/zabbixclient
chmod 755 /usr/lib/zabbix/alertscripts/zabbixclient
```

- zabbixclient支持参数。

```
[root@k8s-master01 linux]# zabbixclient -h
Version 1.1 If you need help contact 244217140@qq.com or visit https://github.com/feiyu563/PrometheusAlert
Usage: zabbixclient [-h] [-t SendTarget] [-m SendMessage] [-type SendType] [-d PrometheusAlertUrl]
Example(发送告警到钉钉)：zabbixclent -t https://oapi.dingtalk.com/robot/send?access_token=xxxxx -m zabbix告警测试 -type dh -d http://127.0.0.1:8080/zabbix

Options:
  -d PrometheusAlert的地址（必须参数）
    	PrometheusAlert的地址 (default "http://127.0.0.1:8080/zabbix")
  -h	显示帮助
  -m 告警消息内容（必须参数）
    	需要发送的告警消息内容 (default "zabbix告警测试")
  -t 手机号/钉钉url/微信url/飞书url（可选参数，如无则从PrometheusAlert的app.conf文件中读取默认配置）
    	指定告警消息的接收目标的手机号/钉钉url/微信url (default "https://oapi.dingtalk.com/robot/send?access_token=xxxxx")
  -type txdx(腾讯云短信)、txdh(腾讯云电话)、alydx(阿里云短信)、alydh(阿里云电话)、hwdx(华为云短信)、rlydh(荣联云电话)、dd(钉钉)、wx(微信)、fs(飞书) （必须参数）
    	告警消息的目标类型,支持txdx(腾讯云短信)、txdh(腾讯云电话)、alydx(阿里云短信)、alydh(阿里云电话)、hwdx(华为云短信)、rlydh(荣联云电话)、dd(钉钉)、wx(微信)、fs(飞书) (default "dd")
```

Zabbix后台配置

1) 首先需要在zabbix后台新增报警媒介，进入 管理-->报警媒介-->创建媒介类型，配置如图：

![zabbix1](../images/zabbix1.png)

```
#脚本参数
    -t （可选参数）
    https://oapi.dingtalk.com/robot/send?access_token=xxxxxxxxxxxxxxxxxxxxxxxxxxx #指定告警消息的接收目标的手机号/钉钉url/微信url，注意需要与-type参数对应（可选参数）
    -m
    {ALERT.MESSAGE}   #这事zabbix内置消息变量名
    -type
    dd                #告警消息的目标类型,支持txdx(腾讯云短信)、txdh(腾讯云电话)、alydx(阿里云短信)、alydh(阿里云电话)、hwdx(华为云短信)、rlydh(荣联云电话)、dd(钉钉)、wx(微信)、fs(飞书) (default "dd")
    -d
    http://[prometheusalert-url]/zabbix/alert  #PrometheusAlert的地址
```

2) 将报警媒介分配给用户，进入 用户-->点击需要分配的用户名-->报警媒介-->添加。如图：

![zabbix2](../images/zabbix2.png)

3) 继续添加告警动作，进入 配置-->动作-->创建动作。如图：

![zabbix3](../images/zabbix3.png)

![zabbix4](../images/zabbix4.png)

最关键的配置就是消息内容的部分，可以参考以下消息模版：

```
1.钉钉消息模版

//操作消息模版

## [PrometheusAlert告警平台告警信息](https://zabbix.nb.cn)
#### {TRIGGER.NAME}

###### 故障时间：{EVENT.DATE} {EVENT.TIME}
###### 告警级别：{TRIGGER.SEVERITY}
###### 故障前状态：{ITEM.LASTVALUE}
###### 故障事件ID：{EVENT.ID}
###### 故障主机IP：{HOST.IP}
###### 故障主机名：{HOST.NAME}
###### 故障时长：{EVENT.AGE}
###### 故障是否确认：{EVENT.ACK.STATUS}

`{EVENT.STATUS}`

![PrometheusAlert](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)

---------------------------------------------------------------

//恢复消息模版

## [PrometheusAlert告警平台恢复信息](https://zabbix.nb.cn)
#### {TRIGGER.NAME} 已经恢复

###### 恢复时间：{EVENT.RECOVERY.DATE} {EVENT.RECOVERY.TIME}
###### 故障前状态：{ITEM.LASTVALUE}
###### 故障事件ID：{EVENT.ID}
###### 故障主机IP：{HOST.IP}
###### 故障主机名：{HOST.NAME}
###### 故障时长：{EVENT.AGE}
###### 故障是否确认：{EVENT.ACK.STATUS}

`{EVENT.STATUS}`

![PrometheusAlert](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)

---------------------------------------------------------------

//确认消息模版

## [PrometheusAlert告警平台确认信息](https://zabbix.nb.cn)
#### {USER.FULLNAME} 已经确认故障原因

###### 确认时间：`{ACK.DATE} {ACK.TIME}`
###### 故障前状态：{ITEM.LASTVALUE}
###### 故障事件ID：{EVENT.ID}
###### 故障主机IP：{HOST.IP}
###### 故障主机名：{HOST.NAME}
###### 故障时长：{EVENT.AGE}
###### 故障原因：{ACK.MESSAGE}
###### 故障是否确认：{EVENT.ACK.STATUS}

`{EVENT.STATUS}`

![PrometheusAlert](https://raw.githubusercontent.com/feiyu563/PrometheusAlert/master/doc/alert-center.png)

·····························································

2.微信消息模版

//操作消息模版

[PrometheusAlert告警平台告警信息](https://zabbix.nb.cn)
>**{TRIGGER.NAME}**
>`故障时间：`{EVENT.DATE} {EVENT.TIME}
`告警级别：`{TRIGGER.SEVERITY}
`故障前状态：`{ITEM.LASTVALUE}
`故障事件ID:`{EVENT.ID}
`故障主机IP：`{HOST.IP}
`故障主机名：`{HOST.NAME}
`故障时长：`{EVENT.AGE}
`故障是否确认：`{EVENT.ACK.STATUS}
**{EVENT.STATUS}**

---------------------------------------------------------------

//恢复消息模版

[PrometheusAlert告警平台恢复信息](https://zabbix.nb.cn)
>**{TRIGGER.NAME} 已经恢复**
>`恢复时间：`{EVENT.RECOVERY.DATE} {EVENT.RECOVERY.TIME}
`告警级别：`{TRIGGER.SEVERITY}
`故障前状态：`{ITEM.LASTVALUE}
`故障主机IP：`{HOST.IP}
`故障主机名：`{HOST.NAME}
`故障时长：`{EVENT.AGE}
`故障是否确认：`{EVENT.ACK.STATUS}
**{EVENT.STATUS}**

---------------------------------------------------------------

//确认消息模版

[PrometheusAlert告警平台确认信息](https://zabbix.nb.cn)
>**{USER.FULLNAME} 已经确认故障原因**
>`确认时间：`{ACK.DATE} {ACK.TIME}
`告警级别：`{TRIGGER.SEVERITY}
`故障前状态：`{ITEM.LASTVALUE}
`故障主机IP：`{HOST.IP}
`故障主机名：`{HOST.NAME}
`故障时长：`{EVENT.AGE}
`故障原因：`{ACK.MESSAGE}
`故障是否确认：`{EVENT.ACK.STATUS}
**{EVENT.STATUS}**

·····························································

3.短信和电话消息模版

//操作消息模版

故障主机IP：{HOST.IP} {TRIGGER.NAME}

---------------------------------------------------------------

//恢复消息模版

故障主机IP：{HOST.IP} {TRIGGER.NAME} 已经恢复

---------------------------------------------------------------

//确认消息模版

故障主机IP：{HOST.IP} {USER.FULLNAME} 已经确认故障原因
```


4) 最终告警效果:

![zabbix5](../images/zabbix5.png)

5) zabbixclient其他用法

其中命令行工具zabbixclient不仅支持作为监控系统Zabbix的告警中心代理客户端，同时也可作为一个单独的命令行客户端工具来使用，zabbixclient的存在，使其他监控系统能够更容易得接入PrometheusAlert告警中心，zabbixclient支持的参数如下：

```
./zabbixclient -t https://oapi.dingtalk.com/robot/send?access_token=xxxxxxx -m 'PrometheusAlert告警平台告警测试' -type dd -d http://dingding.datafountain.cn/zabbix
```
