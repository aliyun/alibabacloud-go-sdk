// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteKvRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeys(v []*string) *BatchDeleteKvRequest
	GetKeys() []*string
	SetNamespace(v string) *BatchDeleteKvRequest
	GetNamespace() *string
}

type BatchDeleteKvRequest struct {
	// List of keys to delete in bulk. You can delete up to 10,000 keys.
	//
	// This parameter is required.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// Name specified when you call [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchDeleteKvRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvRequest) GetKeys() []*string {
	return s.Keys
}

func (s *BatchDeleteKvRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteKvRequest) SetKeys(v []*string) *BatchDeleteKvRequest {
	s.Keys = v
	return s
}

func (s *BatchDeleteKvRequest) SetNamespace(v string) *BatchDeleteKvRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteKvRequest) Validate() error {
	return dara.Validate(s)
}
