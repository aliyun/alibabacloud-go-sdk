// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeleteDcdnKvShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeysShrink(v string) *BatchDeleteDcdnKvShrinkRequest
	GetKeysShrink() *string
	SetNamespace(v string) *BatchDeleteDcdnKvShrinkRequest
	GetNamespace() *string
}

type BatchDeleteDcdnKvShrinkRequest struct {
	// This parameter is required.
	KeysShrink *string `json:"Keys,omitempty" xml:"Keys,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchDeleteDcdnKvShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeleteDcdnKvShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteDcdnKvShrinkRequest) GetKeysShrink() *string {
	return s.KeysShrink
}

func (s *BatchDeleteDcdnKvShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchDeleteDcdnKvShrinkRequest) SetKeysShrink(v string) *BatchDeleteDcdnKvShrinkRequest {
	s.KeysShrink = &v
	return s
}

func (s *BatchDeleteDcdnKvShrinkRequest) SetNamespace(v string) *BatchDeleteDcdnKvShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *BatchDeleteDcdnKvShrinkRequest) Validate() error {
	return dara.Validate(s)
}
