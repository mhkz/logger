#### 日志库需求分析
##### 日志库产生的背景
- 程序运行是黑盒
- 日志是程序运行的外在表现
- 通过日志，可以知道程序的健康状态
##### 需求分析
- Debugger 级别 用来调试程序，日志最详细， 对程序性能影响比较大
- Trace 级别 用来最终问题
- Info 级别 打印比较重要的信息，比如访问日志
- Warn 级别 警告日志，说明程序出现了潜在问题
- Error级别 错误日志，程序运行发生错误，但不影响程序运行
- Fatal 级别 严重错误日志，会导致程序退出
##### 存储位置
- 直接输入到控制台
- 打印到文件里
- 存到网络库 kafka
##### 为何使用接口
- 定义日志库的规范或者标准
- 易扩展
- 可维护性
#### 日志库的设计
- 打印各个level的日志
- 设置级别
- 控制台打印

### How use it ?
You can refer it
```
func initLogger(logPath, logName string, level string)(err error)  {
	m := make(map[string]string)
	m["log_path"] = logPath
	m["log_name"] = "user_server"
	m["log_level"] = level
	err = logger.InitLogger("console", m)
	if err !=nil {
		return
	}
	logger.Debug("init logger success")
	return
}
```