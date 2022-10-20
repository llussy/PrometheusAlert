# 企业微信告警配置

打开企业微信,进入企业微信群中,选择群设置-->群机器人-->添加，可参下图：

![wx1](../images/wx1.png)

![wx2](../images/wx2.png)

复制图中的Webhook地址，并填入PrometheusAlert配置文件app.conf中对应配置项即可。

 **PS: 企业微信机器人目前已经支持 `@某人` ,使用该功能需要通过企业微信管理后台取得对应用户的帐号，如下图：**

![wx2](../images/wx3.png)


企业微信机器人目前支持的markdown语法是如下的子集：

```
标题 （支持1至6级标题，注意#与文字中间要有空格）
# 标题一
## 标题二
### 标题三
#### 标题四
##### 标题五
###### 标题六

加粗
**bold**

链接
[这是一个链接](http://work.weixin.qq.com/api/doc)

行内代码段（暂不支持跨行）
`code`

引用
> 引用文字

字体颜色(只支持3种内置颜色)
<font color="info">绿色</font>
<font color="comment">灰色</font>
<font color="warning">橙红色</font>
```

企业微信机器人相关配置：

```
#---------------------↓webhook-----------------------
#是否开启微信告警通道,可同时开始多个通道0为关闭,1为开启
open-weixin=1
#默认企业微信机器人地址
wxurl=https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxxxx
```


**如何使用**

以Prometheus配合自定义模板为例：

Prometheus配置参考：

```
global:
  resolve_timeout: 5m
route:
  group_by: ['instance']
  group_wait: 10m
  group_interval: 10s
  repeat_interval: 10m
  receiver: 'web.hook.prometheusalert'
receivers:
- name: 'web.hook.prometheusalert'
  webhook_configs:
  - url: 'http://[prometheusalert_url]:8080/prometheusalert?type=wx&tpl=prometheus-wx&wxurl=微信机器人地址,微信机器人地址2&at=zhangsan,lisi'
```