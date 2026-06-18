// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutKvWithHighCapacityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *PutKvWithHighCapacityRequest
	GetKey() *string
	SetNamespace(v string) *PutKvWithHighCapacityRequest
	GetNamespace() *string
	SetUrl(v string) *PutKvWithHighCapacityRequest
	GetUrl() *string
}

type PutKvWithHighCapacityRequest struct {
	// The name of the key to set. It cannot exceed 512 characters and cannot contain spaces or backslashes (/).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name specified when calling [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namesapce
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The download link for the key-value pair to set. This parameter is automatically generated when you call the SDK. Use the SDK to call it.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxxobject.oss-cn-reginon.aliyuncs.com/9d91_xxxxxxxxxxx_158bb6e0f97c477791209bb46bd599f7
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s PutKvWithHighCapacityRequest) String() string {
	return dara.Prettify(s)
}

func (s PutKvWithHighCapacityRequest) GoString() string {
	return s.String()
}

func (s *PutKvWithHighCapacityRequest) GetKey() *string {
	return s.Key
}

func (s *PutKvWithHighCapacityRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutKvWithHighCapacityRequest) GetUrl() *string {
	return s.Url
}

func (s *PutKvWithHighCapacityRequest) SetKey(v string) *PutKvWithHighCapacityRequest {
	s.Key = &v
	return s
}

func (s *PutKvWithHighCapacityRequest) SetNamespace(v string) *PutKvWithHighCapacityRequest {
	s.Namespace = &v
	return s
}

func (s *PutKvWithHighCapacityRequest) SetUrl(v string) *PutKvWithHighCapacityRequest {
	s.Url = &v
	return s
}

func (s *PutKvWithHighCapacityRequest) Validate() error {
	return dara.Validate(s)
}
