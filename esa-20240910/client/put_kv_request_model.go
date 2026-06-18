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
	// Indicates whether the Value parameter uses Base64 encoding. Set this to true when storing a binary value. If this parameter is true, Value must be a valid Base64-encoded string.
	//
	// example:
	//
	// true
	Base64 *bool `json:"Base64,omitempty" xml:"Base64,omitempty"`
	// The expiration time for the key-value pair, as a Unix timestamp in seconds. This value cannot be earlier than the current time. If you specify both Expiration and ExpirationTtl, ExpirationTtl takes precedence.
	//
	// example:
	//
	// 1690081381
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The time-to-live (TTL) for the key-value pair, in seconds. If you specify both Expiration and ExpirationTtl, ExpirationTtl takes precedence.
	//
	// example:
	//
	// 3600
	ExpirationTtl *int64 `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	// The key can be up to 512 bytes long and cannot contain spaces or forward slashes (/).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the namespace. You specify this name when you call the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The value to associate with the key. The maximum size is 2 MB (2,000,000 bytes). For values that exceed this limit, use the [PutKvWithHighCapacity](https://help.aliyun.com/document_detail/2850486.html) operation.
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
