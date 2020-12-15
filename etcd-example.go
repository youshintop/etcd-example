package main

import (
	etcdv3 "go.etcd.io/etcd/clientv3"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
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
