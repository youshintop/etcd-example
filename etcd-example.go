package main

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	etcdv3 "youshintop/etcd-example/v3"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var inMediaType = "application/vnd.kubernetes.probobuf"
var outMediaType = "application/json"

func init() {
	corev1.AddToScheme(Scheme)
}

func main() {
	clientv3.New()
}
