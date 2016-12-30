# services
## ectd
	- Etcd是一个高可用的 Key/Value 存储系统，主要用于分享配置和服务发现。
	- 简单：支持 curl 方式的用户 API (HTTP+JSON)
	- 安全：可选 SSL 客户端证书认证
	- 快速：单实例可达每秒 1000 次写操作
	- 可靠：使用 Raft 实现分布式
## registrator
	http://gliderlabs.com/registrator/latest/
	一个服务发现系统主要由三部分组成：
		- 注册器(registrator)：根据服务运行状态，注册/注销服务。主要要解决的问题是，何时发起注册/注销动作。
		- 注册表(registry)：存储服务信息。常见的解决方案有zookeeper、etcd、cousul等。
		- 发现机制(discovery)：从注册表读取服务信息，给用户封装访问接口。
	Registrator是独立的服务注册器，不需要依赖任何集群管理系统：
		- 通过docker socket直接监听容器event，根据容器启动/停止等event来注册/注销服务
		- 每个容器的每个exposed端口对应不同的服务
		- 支持可插拔的registry backend，默认支持Consul, etcd and SkyDNS
		- 自身也是docker化的，可以容器方式启动
		- 用户可自定义配置，如服务TTL（time-to-live）、服务名称、服务tag等
