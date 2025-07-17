// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v int64) *DeleteMetaCategoryRequest
	GetCategoryId() *int64
	SetTid(v int64) *DeleteMetaCategoryRequest
	GetTid() *int64
}

type DeleteMetaCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30000235594
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s DeleteMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetaCategoryRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *DeleteMetaCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *DeleteMetaCategoryRequest) SetCategoryId(v int64) *DeleteMetaCategoryRequest {
	s.CategoryId = &v
	return s
}

func (s *DeleteMetaCategoryRequest) SetTid(v int64) *DeleteMetaCategoryRequest {
	s.Tid = &v
	return s
}

func (s *DeleteMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
