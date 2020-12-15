package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/clientv3"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/apimachinery/pkg/runtime/serializer"
	"time"
)

var Scheme = runtime.NewScheme()
var Codecs = serializer.NewCodecFactory(Scheme)
var inMediaType = "application/vnd.kubernetes.probobuf"
var outMediaType = "application/json"

func init() {
	corev1.AddToScheme(Scheme)
}

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{
			"https://172.26.68.112:2379",
		},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	resp, err := cli.Get(context.TODO(), "/registry/pods/default/nginx")
	if err != nil {
		panic(err)
	}
	kv := resp.Kvs[0]

	inCodec := newCodec(inMediaType)
	outCodec := newCodec(outMediaType)
	obj, err := runtime.Decode(inCodec, kv.Value)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decode ---")
	fmt.Println(obj)

	encoded, err := runtime.Encode(outCodec, obj)
	if err != nil {
		panic(err)
	}

	fmt.Println("Encode ---")
	fmt.Println(string(encoded))

}

func newCodec(mediaTypes string) runtime.Codec {
	info,ok :=	runtime.SerializerInfoForMediaType(Codecs.SupportedMediaTypes(), mediaTypes)
	if !ok {
		panic(fmt.Errorf("no Serializers registered for %v", mediaTypes))
	}

	cfactory := serializer.NewCodecFactory(Scheme)

	gv, err := schema.ParseGroupVersion("v1")
	if err != nil {
		panic(err)
	}

	encoder :=cfactory.EncoderForVersion(info.Serializer, gv)
	decoder := cfactory.DecoderToVersion(info.Serializer, gv)
	return cfactory.CodecForVersions(encoder, decoder, gv, gv)
}
