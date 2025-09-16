// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListModelsRequest
	GetName() *string
	SetPageNumber(v int32) *ListModelsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelsRequest
	GetPageSize() *int32
	SetType(v string) *ListModelsRequest
	GetType() *string
}

type ListModelsRequest struct {
	// example:
	//
	// test1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// text_embedding
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) GetName() *string {
	return s.Name
}

func (s *ListModelsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelsRequest) GetType() *string {
	return s.Type
}

func (s *ListModelsRequest) SetName(v string) *ListModelsRequest {
	s.Name = &v
	return s
}

func (s *ListModelsRequest) SetPageNumber(v int32) *ListModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelsRequest) SetPageSize(v int32) *ListModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsRequest) SetType(v string) *ListModelsRequest {
	s.Type = &v
	return s
}

func (s *ListModelsRequest) Validate() error {
	return dara.Validate(s)
}
