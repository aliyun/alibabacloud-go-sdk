// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *ListParametersRequest
	GetIds() []*int64
	SetNames(v []*string) *ListParametersRequest
	GetNames() []*string
	SetOwner(v string) *ListParametersRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListParametersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListParametersRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListParametersRequest
	GetProjectId() *int64
	SetScope(v string) *ListParametersRequest
	GetScope() *string
	SetSortBy(v string) *ListParametersRequest
	GetSortBy() *string
	SetType(v string) *ListParametersRequest
	GetType() *string
}

type ListParametersRequest struct {
	Ids   []*int64  `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
	Names []*string `json:"Names,omitempty" xml:"Names,omitempty" type:"Repeated"`
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// Project
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// ModifyTime Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// PlainConstant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListParametersRequest) GoString() string {
	return s.String()
}

func (s *ListParametersRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *ListParametersRequest) GetNames() []*string {
	return s.Names
}

func (s *ListParametersRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListParametersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListParametersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListParametersRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListParametersRequest) GetScope() *string {
	return s.Scope
}

func (s *ListParametersRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListParametersRequest) GetType() *string {
	return s.Type
}

func (s *ListParametersRequest) SetIds(v []*int64) *ListParametersRequest {
	s.Ids = v
	return s
}

func (s *ListParametersRequest) SetNames(v []*string) *ListParametersRequest {
	s.Names = v
	return s
}

func (s *ListParametersRequest) SetOwner(v string) *ListParametersRequest {
	s.Owner = &v
	return s
}

func (s *ListParametersRequest) SetPageNumber(v int32) *ListParametersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListParametersRequest) SetPageSize(v int32) *ListParametersRequest {
	s.PageSize = &v
	return s
}

func (s *ListParametersRequest) SetProjectId(v int64) *ListParametersRequest {
	s.ProjectId = &v
	return s
}

func (s *ListParametersRequest) SetScope(v string) *ListParametersRequest {
	s.Scope = &v
	return s
}

func (s *ListParametersRequest) SetSortBy(v string) *ListParametersRequest {
	s.SortBy = &v
	return s
}

func (s *ListParametersRequest) SetType(v string) *ListParametersRequest {
	s.Type = &v
	return s
}

func (s *ListParametersRequest) Validate() error {
	return dara.Validate(s)
}
