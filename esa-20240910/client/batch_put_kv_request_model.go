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
	// The key-value pairs that you want to configure at a time. The total size can be up to 2 MB (2 × 1000 × 1000).
	//
	// This parameter is required.
	KvList []*BatchPutKvRequestKvList `json:"KvList,omitempty" xml:"KvList,omitempty" type:"Repeated"`
	// The name of the namespace that you specify when you call the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
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
	// The key content.
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
