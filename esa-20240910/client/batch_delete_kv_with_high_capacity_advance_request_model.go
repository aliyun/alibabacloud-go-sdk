// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iBatchDeleteKvWithHighCapacityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchDeleteKvWithHighCapacityAdvanceRequest
	GetNamespace() *string
	SetUrlObject(v io.Reader) *BatchDeleteKvWithHighCapacityAdvanceRequest
	GetUrlObject() io.Reader
}

type BatchDeleteKvWithHighCapacityAdvanceRequest struct {
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
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchDeleteKvWithHighCapacityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *BatchDeleteKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *BatchDeleteKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *BatchDeleteKvWithHighCapacityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
