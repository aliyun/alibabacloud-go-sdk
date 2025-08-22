// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutDcdnKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKvList(v []*BatchPutDcdnKvRequestKvList) *BatchPutDcdnKvRequest
	GetKvList() []*BatchPutDcdnKvRequestKvList
	SetNamespace(v string) *BatchPutDcdnKvRequest
	GetNamespace() *string
}

type BatchPutDcdnKvRequest struct {
	// This parameter is required.
	KvList []*BatchPutDcdnKvRequestKvList `json:"KvList,omitempty" xml:"KvList,omitempty" type:"Repeated"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchPutDcdnKvRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutDcdnKvRequest) GoString() string {
	return s.String()
}

func (s *BatchPutDcdnKvRequest) GetKvList() []*BatchPutDcdnKvRequestKvList {
	return s.KvList
}

func (s *BatchPutDcdnKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutDcdnKvRequest) SetKvList(v []*BatchPutDcdnKvRequestKvList) *BatchPutDcdnKvRequest {
	s.KvList = v
	return s
}

func (s *BatchPutDcdnKvRequest) SetNamespace(v string) *BatchPutDcdnKvRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutDcdnKvRequest) Validate() error {
	return dara.Validate(s)
}

type BatchPutDcdnKvRequestKvList struct {
	Expiration    *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	ExpirationTtl *int64 `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s BatchPutDcdnKvRequestKvList) String() string {
	return dara.Prettify(s)
}

func (s BatchPutDcdnKvRequestKvList) GoString() string {
	return s.String()
}

func (s *BatchPutDcdnKvRequestKvList) GetExpiration() *int64 {
	return s.Expiration
}

func (s *BatchPutDcdnKvRequestKvList) GetExpirationTtl() *int64 {
	return s.ExpirationTtl
}

func (s *BatchPutDcdnKvRequestKvList) GetKey() *string {
	return s.Key
}

func (s *BatchPutDcdnKvRequestKvList) GetValue() *string {
	return s.Value
}

func (s *BatchPutDcdnKvRequestKvList) SetExpiration(v int64) *BatchPutDcdnKvRequestKvList {
	s.Expiration = &v
	return s
}

func (s *BatchPutDcdnKvRequestKvList) SetExpirationTtl(v int64) *BatchPutDcdnKvRequestKvList {
	s.ExpirationTtl = &v
	return s
}

func (s *BatchPutDcdnKvRequestKvList) SetKey(v string) *BatchPutDcdnKvRequestKvList {
	s.Key = &v
	return s
}

func (s *BatchPutDcdnKvRequestKvList) SetValue(v string) *BatchPutDcdnKvRequestKvList {
	s.Value = &v
	return s
}

func (s *BatchPutDcdnKvRequestKvList) Validate() error {
	return dara.Validate(s)
}
