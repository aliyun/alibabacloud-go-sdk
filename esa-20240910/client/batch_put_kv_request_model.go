// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKvList(v []*BatchPutKvRequestKvList) *BatchPutKvRequest
	GetKvList() []*BatchPutKvRequestKvList
	SetNamespace(v string) *BatchPutKvRequest
	GetNamespace() *string
}

type BatchPutKvRequest struct {
	// This parameter is required.
	KvList []*BatchPutKvRequestKvList `json:"KvList,omitempty" xml:"KvList,omitempty" type:"Repeated"`
	// This parameter is required.
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchPutKvRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvRequest) GetKvList() []*BatchPutKvRequestKvList {
	return s.KvList
}

func (s *BatchPutKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutKvRequest) SetKvList(v []*BatchPutKvRequestKvList) *BatchPutKvRequest {
	s.KvList = v
	return s
}

func (s *BatchPutKvRequest) SetNamespace(v string) *BatchPutKvRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutKvRequest) Validate() error {
	if s.KvList != nil {
		for _, item := range s.KvList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchPutKvRequestKvList struct {
	Expiration    *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	ExpirationTtl *int64 `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	// This parameter is required.
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// This parameter is required.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s BatchPutKvRequestKvList) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvRequestKvList) GoString() string {
	return s.String()
}

func (s *BatchPutKvRequestKvList) GetExpiration() *int64 {
	return s.Expiration
}

func (s *BatchPutKvRequestKvList) GetExpirationTtl() *int64 {
	return s.ExpirationTtl
}

func (s *BatchPutKvRequestKvList) GetKey() *string {
	return s.Key
}

func (s *BatchPutKvRequestKvList) GetValue() *string {
	return s.Value
}

func (s *BatchPutKvRequestKvList) SetExpiration(v int64) *BatchPutKvRequestKvList {
	s.Expiration = &v
	return s
}

func (s *BatchPutKvRequestKvList) SetExpirationTtl(v int64) *BatchPutKvRequestKvList {
	s.ExpirationTtl = &v
	return s
}

func (s *BatchPutKvRequestKvList) SetKey(v string) *BatchPutKvRequestKvList {
	s.Key = &v
	return s
}

func (s *BatchPutKvRequestKvList) SetValue(v string) *BatchPutKvRequestKvList {
	s.Value = &v
	return s
}

func (s *BatchPutKvRequestKvList) Validate() error {
	return dara.Validate(s)
}
