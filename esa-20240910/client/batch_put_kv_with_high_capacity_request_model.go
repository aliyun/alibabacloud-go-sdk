// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutKvWithHighCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchPutKvWithHighCapacityRequest
	GetNamespace() *string
	SetUrl(v string) *BatchPutKvWithHighCapacityRequest
	GetUrl() *string
}

type BatchPutKvWithHighCapacityRequest struct {
	// The name of the namespace that you specify when you call the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The download URL of the key-value pairs that you want to configure. This parameter is automatically filled in when you use the SDK to call the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchPutKvWithHighCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutKvWithHighCapacityRequest) GetUrl() *string {
	return s.Url
}

func (s *BatchPutKvWithHighCapacityRequest) SetNamespace(v string) *BatchPutKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutKvWithHighCapacityRequest) SetUrl(v string) *BatchPutKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

func (s *BatchPutKvWithHighCapacityRequest) Validate() error {
	return dara.Validate(s)
}
