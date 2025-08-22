// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchPutDcdnKvShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKvListShrink(v string) *BatchPutDcdnKvShrinkRequest
	GetKvListShrink() *string
	SetNamespace(v string) *BatchPutDcdnKvShrinkRequest
	GetNamespace() *string
}

type BatchPutDcdnKvShrinkRequest struct {
	// This parameter is required.
	KvListShrink *string `json:"KvList,omitempty" xml:"KvList,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s BatchPutDcdnKvShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchPutDcdnKvShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchPutDcdnKvShrinkRequest) GetKvListShrink() *string {
	return s.KvListShrink
}

func (s *BatchPutDcdnKvShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *BatchPutDcdnKvShrinkRequest) SetKvListShrink(v string) *BatchPutDcdnKvShrinkRequest {
	s.KvListShrink = &v
	return s
}

func (s *BatchPutDcdnKvShrinkRequest) SetNamespace(v string) *BatchPutDcdnKvShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *BatchPutDcdnKvShrinkRequest) Validate() error {
	return dara.Validate(s)
}
