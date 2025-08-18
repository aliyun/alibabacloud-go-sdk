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
	// The keys that you want to delete. You can delete a maximum of 10,000 key-value pairs at a time.
	//
	// This parameter is required.
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// The name of the namespace that you specify when you call the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
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
