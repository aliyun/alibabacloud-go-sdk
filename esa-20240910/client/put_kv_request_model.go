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
	Base64        *bool  `json:"Base64,omitempty" xml:"Base64,omitempty"`
	Expiration    *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	ExpirationTtl *int64 `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
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
