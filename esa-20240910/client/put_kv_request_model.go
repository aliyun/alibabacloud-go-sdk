// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBase64(v bool) *PutKvRequest
	GetBase64() *bool
	SetExpiration(v int64) *PutKvRequest
	GetExpiration() *int64
	SetExpirationTtl(v int64) *PutKvRequest
	GetExpirationTtl() *int64
	SetKey(v string) *PutKvRequest
	GetKey() *string
	SetNamespace(v string) *PutKvRequest
	GetNamespace() *string
	SetValue(v string) *PutKvRequest
	GetValue() *string
}

type PutKvRequest struct {
	// Specifies whether the content of the key is Base64-encoded. Set this parameter to true if you want to store the key content in binary format. When this parameter is set to true, the Value parameter must be Base64-encoded.
	//
	// example:
	//
	// true
	Base64 *bool `json:"Base64,omitempty" xml:"Base64,omitempty"`
	// The time when the key-value pair expires, which cannot be earlier than the current time. The value is a timestamp in seconds. If you specify both Expiration and ExpirationTtl, only ExpirationTtl takes effect.
	//
	// example:
	//
	// 1690081381
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The relative expiration time. Unit: seconds. If you specify both Expiration and ExpirationTtl, only ExpirationTtl takes effect.
	//
	// example:
	//
	// 3600
	ExpirationTtl *int64 `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
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
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The content of the key, which can be up to 2 MB (2 × 1000 × 1000). If the content is larger than 2 MB, call [PutKvWithHighCapacity](https://help.aliyun.com/document_detail/2850486.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutKvRequest) String() string {
	return dara.Prettify(s)
}

func (s PutKvRequest) GoString() string {
	return s.String()
}

func (s *PutKvRequest) GetBase64() *bool {
	return s.Base64
}

func (s *PutKvRequest) GetExpiration() *int64 {
	return s.Expiration
}

func (s *PutKvRequest) GetExpirationTtl() *int64 {
	return s.ExpirationTtl
}

func (s *PutKvRequest) GetKey() *string {
	return s.Key
}

func (s *PutKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutKvRequest) GetValue() *string {
	return s.Value
}

func (s *PutKvRequest) SetBase64(v bool) *PutKvRequest {
	s.Base64 = &v
	return s
}

func (s *PutKvRequest) SetExpiration(v int64) *PutKvRequest {
	s.Expiration = &v
	return s
}

func (s *PutKvRequest) SetExpirationTtl(v int64) *PutKvRequest {
	s.ExpirationTtl = &v
	return s
}

func (s *PutKvRequest) SetKey(v string) *PutKvRequest {
	s.Key = &v
	return s
}

func (s *PutKvRequest) SetNamespace(v string) *PutKvRequest {
	s.Namespace = &v
	return s
}

func (s *PutKvRequest) SetValue(v string) *PutKvRequest {
	s.Value = &v
	return s
}

func (s *PutKvRequest) Validate() error {
	return dara.Validate(s)
}
