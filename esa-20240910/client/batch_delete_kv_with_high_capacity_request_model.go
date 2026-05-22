// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteKvWithHighCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchDeleteKvWithHighCapacityRequest
	GetNamespace() *string
	SetUrl(v string) *BatchDeleteKvWithHighCapacityRequest
	GetUrl() *string
}

type BatchDeleteKvWithHighCapacityRequest struct {
	// The name of the namespace that you specify when you call the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The download URL of the key-value pairs that you want to delete. This parameter is automatically filled in when you use the SDK to call the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchDeleteKvWithHighCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteKvWithHighCapacityRequest) GetUrl() *string {
	return s.Url
}

func (s *BatchDeleteKvWithHighCapacityRequest) SetNamespace(v string) *BatchDeleteKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityRequest) SetUrl(v string) *BatchDeleteKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityRequest) Validate() error {
	return dara.Validate(s)
}
