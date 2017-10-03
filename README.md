# MakeUseragent 浏览器User Agent生成器
##demo
go1.8测试通过
```python
	agent := MakeUseragent.New()
	fmt.Println(agent.Android())
	fmt.Println(agent.Baidu_box_app())
	fmt.Println(agent.Chrome_pc())
	fmt.Println(agent.Chrome_wap())
	fmt.Println(agent.Firefox_pc())
	fmt.Println(agent.Iphone())
	fmt.Println(agent.Pc())
	fmt.Println(agent.Pc_linux())
	fmt.Println(agent.Pc_mac())
	fmt.Println(agent.Pc_windows())
	fmt.Println(agent.Wap())
	fmt.Println(agent.Wechat())
	fmt.Println(agent.Random())
```
##运行效果
```
Mozilla/5.0 (Linux; Android 7.4.8; SM-E7000 Build/KTU84P) AppleWebKit/217.66 (KHTML, like Gecko) Chrome/29.5.8448.10 Mobile Safari/217.66
Mozilla/5.0 (iPhone; CPU iPhone OS 6_5_4 like Mac OS X) AppleWebKit/865.72 (KHTML, like Gecko) Mobile/2Q2T1P baiduboxapp/0_19.0.7.7_enohpi_5087_338/2.01_4C2%258enohPi/1099a/LNK6OPDYR89T3AS6EE7OZDSGZX8702IN9DQC94Y69UH7VEDT2VI/1
Mozilla/5.0 (Windows NT 9.3; WOW64) AppleWebKit/368.53 (KHTML, like Gecko) Chrome/37.5.9919.29 Safari/368.53
Mozilla/5.0 (iPhone; CPU iPhone OS 8_3_5 like Mac OS X) AppleWebKit/633.30 (KHTML, like Gecko) Version/9.0 Mobile/73RUQU Safari/633.30
Mozilla/5.0 (Windows NT 8.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0
Mozilla/5.0 (iPhone; CPU iPhone OS 6_5_4 like Mac OS X) AppleWebKit/148.0 (KHTML, like Gecko) Mobile/ITL86A MicroMessenger/6.3.0.8106 NetType/4G Language/zh_CN
Mozilla/5.0 (Macintosh; Intel Mac OS X 8_1_9) AppleWebKit/745.1 (KHTML, like Gecko) Chrome/20.0.2689.64 Safari/745.1
Mozilla/5.0 (X11; NetBSD amd64; rv:46.0) Gecko/20100101 Firefox/46.0
Mozilla/5.0 (Macintosh; Intel Mac OS X 7_6_9) AppleWebKit/688.11 (KHTML, like Gecko) Chrome/42.8.9246.72 Safari/688.11
Mozilla/5.0 (Windows NT 9.7; WOW64; rv:30.0) Gecko/20100101 Firefox/30.0
Mozilla/5.0 (Linux; U; Android 7.4.0; zh-cn; A31t Build/KTU84P) AppleWebKit/522.95 (KHTML, like Gecko) Version/4.0 UCBrowser/1.0.0.100 U3/0.8.0 Mobile Safari/522.95 AliApp(TB/6.6.4) WindVane/8.0.0 1080X1920 GCanvas/1.4.2.21
Mozilla/5.0 (Linux; Android 8.6.0; MI NOTE LTE Build/KTU84P) AppleWebKit/221.21 (KHTML, like Gecko) Version/4.0 Chrome/37.2.3971.16 Mobile MQQBrowser/6.2 TBS/022725 Safari/221.21 MicroMessenger/6.6.1.232 NetType/3G Language/zh_CN
Mozilla/5.0 (iPhone; CPU iPhone OS 12_5_1 like Mac OS X) AppleWebKit/949.96 (KHTML, like Gecko) Version/9.0 Mobile/DH8A63 Safari/949.96
```
