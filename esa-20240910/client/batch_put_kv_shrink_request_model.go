// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutKvShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKvListShrink(v string) *BatchPutKvShrinkRequest
	GetKvListShrink() *string
	SetNamespace(v string) *BatchPutKvShrinkRequest
	GetNamespace() *string
}

type BatchPutKvShrinkRequest struct {
	// The list of key-value pairs to set. The total size cannot exceed 2 MB (2 × 1,000 × 1,000).
	//
	// This parameter is required.
	KvListShrink *string `json:"KvList,omitempty" xml:"KvList,omitempty"`
	// The name specified when you call [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchPutKvShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutKvShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchPutKvShrinkRequest) GetKvListShrink() *string {
	return s.KvListShrink
}

func (s *BatchPutKvShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutKvShrinkRequest) SetKvListShrink(v string) *BatchPutKvShrinkRequest {
	s.KvListShrink = &v
	return s
}

func (s *BatchPutKvShrinkRequest) SetNamespace(v string) *BatchPutKvShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutKvShrinkRequest) Validate() error {
	return dara.Validate(s)
}
