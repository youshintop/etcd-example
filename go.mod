module github.com/youshintop/etcd-example

go 1.15

require (
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/zap v1.16.0 // indirect
	k8s.io/api v0.20.0
	k8s.io/apimachinery v0.20.0
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.13+incompatible
	go.etcd.io/bbolt => go.etcd.io/bbolt v1.3.5
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200910180754-dd1b699fc489 // ae9734ed278b is the SHA for git tag v3.4.13
)
