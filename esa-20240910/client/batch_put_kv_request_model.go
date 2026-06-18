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
	// The list of key-value pairs to set. The total size cannot exceed 2 MB (2 × 1,000 × 1,000).
	//
	// This parameter is required.
	KvList []*BatchPutKvRequestKvList `json:"KvList,omitempty" xml:"KvList,omitempty" type:"Repeated"`
	// The name specified when you call [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
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
	// The expiration time. This is a UNIX timestamp in seconds and cannot be earlier than the current time. If you set both Expiration and ExpirationTtl, ExpirationTtl takes precedence.
	//
	// example:
	//
	// 1690081381
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// The time-to-live (TTL). This is a relative time in seconds. If you set both Expiration and ExpirationTtl, ExpirationTtl takes precedence.
	//
	// example:
	//
	// 3600
	ExpirationTtl *int64 `json:"ExpirationTtl,omitempty" xml:"ExpirationTtl,omitempty"`
	// The name of the key. The key can be up to 512 characters long and cannot contain spaces or backslashes (/).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the key.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_value
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
