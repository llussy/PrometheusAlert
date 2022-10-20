# SonarQube 接入配置
 
首先使用管理员或者具有告警配置权限的帐号登录进SonarQube管理页面，登录后进入`Administration Configuration`配置。
 
![sonar1](../images/sonar1.png)
 
选择`Configuration`--->`Webhooks`--->`Create`，并填入相关配置信息
 
![sonar2](../images/sonar2.png)
 
其中`URL`项，请填写PrometheusAlert的Url，参考如下：

```
#下面的地址是使用自定义模版中已经默认集成的SonarQube的钉钉模版
http://[YOUR-PrometheusAlert-URL]/prometheusalert?type=dd&tpl=sonar&ddurl=https://oapi.dingtalk.com/robot/send?access_token=xxxxx
```

如需修改默认的SonarQube的钉钉模版或者使用其他的模版可通过PrometheusAlert的Dashboard进行操作。

![sonar4](../images/sonar4.png)

![sonar5](../images/sonar5.png)

最终告警效果:
 
![sonar3](../images/sonar3.png)