// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBase64(v bool) *GetKvRequest
	GetBase64() *bool
	SetKey(v string) *GetKvRequest
	GetKey() *string
	SetNamespace(v string) *GetKvRequest
	GetNamespace() *string
}

type GetKvRequest struct {
	// Specifies whether to decode the value by using Base 64. If you call the [PutKv](https://help.aliyun.com/document_detail/2850482.html) operation and set the Base64 parameter to true, set this parameter to true to read the original content.
	//
	// example:
	//
	// true
	Base64 *bool `json:"Base64,omitempty" xml:"Base64,omitempty"`
	// The key name for the query.
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
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetKvRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKvRequest) GoString() string {
	return s.String()
}

func (s *GetKvRequest) GetBase64() *bool {
	return s.Base64
}

func (s *GetKvRequest) GetKey() *string {
	return s.Key
}

func (s *GetKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetKvRequest) SetBase64(v bool) *GetKvRequest {
	s.Base64 = &v
	return s
}

func (s *GetKvRequest) SetKey(v string) *GetKvRequest {
	s.Key = &v
	return s
}

func (s *GetKvRequest) SetNamespace(v string) *GetKvRequest {
	s.Namespace = &v
	return s
}

func (s *GetKvRequest) Validate() error {
	return dara.Validate(s)
}
