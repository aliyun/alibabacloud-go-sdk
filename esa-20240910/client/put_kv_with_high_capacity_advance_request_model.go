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
