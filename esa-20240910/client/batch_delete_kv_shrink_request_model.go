// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteKvShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeysShrink(v string) *BatchDeleteKvShrinkRequest
	GetKeysShrink() *string
	SetNamespace(v string) *BatchDeleteKvShrinkRequest
	GetNamespace() *string
}

type BatchDeleteKvShrinkRequest struct {
	// The keys that you want to delete. You can delete a maximum of 10,000 key-value pairs at a time.
	//
	// This parameter is required.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// The name of the namespace that you specify when you call the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchDeleteKvShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteKvShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteKvShrinkRequest) GetKeysShrink() *string {
	return s.KeysShrink
}

func (s *BatchDeleteKvShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteKvShrinkRequest) SetKeysShrink(v string) *BatchDeleteKvShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *BatchDeleteKvShrinkRequest) SetNamespace(v string) *BatchDeleteKvShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteKvShrinkRequest) Validate() error {
	return dara.Validate(s)
}
