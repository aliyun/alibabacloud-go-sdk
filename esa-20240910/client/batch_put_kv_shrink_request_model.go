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
	// This parameter is required.
	KvListShrink *string `json:"KvList,omitempty" xml:"KvList,omitempty"`
	// This parameter is required.
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
