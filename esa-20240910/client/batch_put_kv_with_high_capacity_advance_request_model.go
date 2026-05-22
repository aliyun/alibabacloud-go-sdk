// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iBatchPutKvWithHighCapacityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *BatchPutKvWithHighCapacityAdvanceRequest
	GetNamespace() *string
	SetUrlObject(v io.Reader) *BatchPutKvWithHighCapacityAdvanceRequest
	GetUrlObject() io.Reader
}

type BatchPutKvWithHighCapacityAdvanceRequest struct {
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
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s BatchPutKvWithHighCapacityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *BatchPutKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *BatchPutKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *BatchPutKvWithHighCapacityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
