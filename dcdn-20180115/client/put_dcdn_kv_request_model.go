// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutDcdnKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpiration(v int64) *PutDcdnKvRequest
	GetExpiration() *int64
	SetExpirationTtl(v int64) *PutDcdnKvRequest
	GetExpirationTtl() *int64
	SetKey(v string) *PutDcdnKvRequest
	GetKey() *string
	SetNamespace(v string) *PutDcdnKvRequest
	GetNamespace() *string
	SetValue(v string) *PutDcdnKvRequest
	GetValue() *string
}

type PutDcdnKvRequest struct {
	// The time when the key expires.Example: "1690081381".
	//
	// example:
	//
	// 1690081381
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The time when the key expires.Example: "3600".
	//
	// example:
	//
	// 3600
	ExpirationTtl *int64 `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	// The key. The key can be up to 512 characters in length, and cannot contain spaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// key1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The value of the key. The maximum size is 2 MB (2 x 1000 x 1000 bytes).
	//
	// This parameter is required.
	//
	// example:
	//
	// value1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s PutDcdnKvRequest) String() string {
	return dara.Prettify(s)
}

func (s PutDcdnKvRequest) GoString() string {
	return s.String()
}

func (s *PutDcdnKvRequest) GetExpiration() *int64 {
	return s.Expiration
}

func (s *PutDcdnKvRequest) GetExpirationTtl() *int64 {
	return s.ExpirationTtl
}

func (s *PutDcdnKvRequest) GetKey() *string {
	return s.Key
}

func (s *PutDcdnKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *PutDcdnKvRequest) GetValue() *string {
	return s.Value
}

func (s *PutDcdnKvRequest) SetExpiration(v int64) *PutDcdnKvRequest {
	s.Expiration = &v
	return s
}

func (s *PutDcdnKvRequest) SetExpirationTtl(v int64) *PutDcdnKvRequest {
	s.ExpirationTtl = &v
	return s
}

func (s *PutDcdnKvRequest) SetKey(v string) *PutDcdnKvRequest {
	s.Key = &v
	return s
}

func (s *PutDcdnKvRequest) SetNamespace(v string) *PutDcdnKvRequest {
	s.Namespace = &v
	return s
}

func (s *PutDcdnKvRequest) SetValue(v string) *PutDcdnKvRequest {
	s.Value = &v
	return s
}

func (s *PutDcdnKvRequest) Validate() error {
	return dara.Validate(s)
}
