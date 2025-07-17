// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateMetaCategoryRequest
	GetName() *string
	SetParentCategoryId(v int64) *CreateMetaCategoryRequest
	GetParentCategoryId() *int64
	SetTid(v int64) *CreateMetaCategoryRequest
	GetTid() *int64
}

type CreateMetaCategoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30000322682
	ParentCategoryId *int64 `json:"ParentCategoryId,omitempty" xml:"ParentCategoryId,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateMetaCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCategoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMetaCategoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetaCategoryRequest) GetParentCategoryId() *int64 {
	return s.ParentCategoryId
}

func (s *CreateMetaCategoryRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateMetaCategoryRequest) SetName(v string) *CreateMetaCategoryRequest {
	s.Name = &v
	return s
}

func (s *CreateMetaCategoryRequest) SetParentCategoryId(v int64) *CreateMetaCategoryRequest {
	s.ParentCategoryId = &v
	return s
}

func (s *CreateMetaCategoryRequest) SetTid(v int64) *CreateMetaCategoryRequest {
	s.Tid = &v
	return s
}

func (s *CreateMetaCategoryRequest) Validate() error {
	return dara.Validate(s)
}
