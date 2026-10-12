// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParametersShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdsShrink(v string) *ListParametersShrinkRequest
	GetIdsShrink() *string
	SetNamesShrink(v string) *ListParametersShrinkRequest
	GetNamesShrink() *string
	SetOwner(v string) *ListParametersShrinkRequest
	GetOwner() *string
	SetPageNumber(v int32) *ListParametersShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListParametersShrinkRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListParametersShrinkRequest
	GetProjectId() *int64
	SetScope(v string) *ListParametersShrinkRequest
	GetScope() *string
	SetSortBy(v string) *ListParametersShrinkRequest
	GetSortBy() *string
	SetType(v string) *ListParametersShrinkRequest
	GetType() *string
}

type ListParametersShrinkRequest struct {
	// The list of parameter IDs.
	IdsShrink *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The list of parameter names.
	NamesShrink *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The account ID of the owner.
	//
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the workspace. You can call the ListProjects operation to obtain the workspace ID.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The scope of the parameter. Default value: Project. Other types are not supported.
	//
	// example:
	//
	// Project
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The list of fields by which to sort the results. Supported fields include modification time and creation time. The format is "SortField+SortOrder(Desc/Asc)". You can omit Asc. Valid values:
	//
	// - ModifyTime (Desc/Asc)
	//
	// - CreateTime (Desc/Asc)
	//
	// - Name (Desc/Asc)
	//
	// example:
	//
	// ModifyTime Desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The parameter type. Valid values:
	//
	// - PlainConstant: plaintext constant.
	//
	// - SecretConstant: ciphertext constant.
	//
	// - Variable: variable.
	//
	// example:
	//
	// PlainConstant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListParametersShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListParametersShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListParametersShrinkRequest) GetIdsShrink() *string {
	return s.IdsShrink
}

func (s *ListParametersShrinkRequest) GetNamesShrink() *string {
	return s.NamesShrink
}

func (s *ListParametersShrinkRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListParametersShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListParametersShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListParametersShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListParametersShrinkRequest) GetScope() *string {
	return s.Scope
}

func (s *ListParametersShrinkRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListParametersShrinkRequest) GetType() *string {
	return s.Type
}

func (s *ListParametersShrinkRequest) SetIdsShrink(v string) *ListParametersShrinkRequest {
	s.IdsShrink = &v
	return s
}

func (s *ListParametersShrinkRequest) SetNamesShrink(v string) *ListParametersShrinkRequest {
	s.NamesShrink = &v
	return s
}

func (s *ListParametersShrinkRequest) SetOwner(v string) *ListParametersShrinkRequest {
	s.Owner = &v
	return s
}

func (s *ListParametersShrinkRequest) SetPageNumber(v int32) *ListParametersShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListParametersShrinkRequest) SetPageSize(v int32) *ListParametersShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListParametersShrinkRequest) SetProjectId(v int64) *ListParametersShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *ListParametersShrinkRequest) SetScope(v string) *ListParametersShrinkRequest {
	s.Scope = &v
	return s
}

func (s *ListParametersShrinkRequest) SetSortBy(v string) *ListParametersShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListParametersShrinkRequest) SetType(v string) *ListParametersShrinkRequest {
	s.Type = &v
	return s
}

func (s *ListParametersShrinkRequest) Validate() error {
	return dara.Validate(s)
}
