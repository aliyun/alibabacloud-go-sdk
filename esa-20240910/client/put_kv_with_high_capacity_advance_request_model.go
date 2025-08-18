// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iPutKvWithHighCapacityAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *PutKvWithHighCapacityAdvanceRequest
	GetKey() *string
	SetNamespace(v string) *PutKvWithHighCapacityAdvanceRequest
	GetNamespace() *string
	SetUrlObject(v io.Reader) *PutKvWithHighCapacityAdvanceRequest
	GetUrlObject() io.Reader
}

type PutKvWithHighCapacityAdvanceRequest struct {
	// The key name. The name can be up to 512 characters in length and cannot contain spaces or backslashes (\\\\).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the namespace that you specify when you call the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namesapce
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The download URL of the key-value pair that you want to upload. This parameter is automatically filled in when you use the SDK to call the operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	UrlObject io.Reader `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PutKvWithHighCapacityAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PutKvWithHighCapacityAdvanceRequest) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityAdvanceRequest) GetKey() *string {
	return s.Key
}

func (s *PutKvWithHighCapacityAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutKvWithHighCapacityAdvanceRequest) GetUrlObject() io.Reader {
	return s.UrlObject
}

func (s *PutKvWithHighCapacityAdvanceRequest) SetKey(v string) *PutKvWithHighCapacityAdvanceRequest {
	s.Key = &v
	return s
}

func (s *PutKvWithHighCapacityAdvanceRequest) SetNamespace(v string) *PutKvWithHighCapacityAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *PutKvWithHighCapacityAdvanceRequest) SetUrlObject(v io.Reader) *PutKvWithHighCapacityAdvanceRequest {
	s.UrlObject = v
	return s
}

func (s *PutKvWithHighCapacityAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
