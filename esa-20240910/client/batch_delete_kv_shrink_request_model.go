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
	// This parameter is required.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// This parameter is required.
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
