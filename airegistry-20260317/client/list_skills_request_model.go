// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *ListSkillsRequest
	GetNamespaceId() *string
	SetOrderBy(v string) *ListSkillsRequest
	GetOrderBy() *string
	SetOwner(v string) *ListSkillsRequest
	GetOwner() *string
	SetPageNo(v int32) *ListSkillsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListSkillsRequest
	GetPageSize() *int32
	SetScope(v string) *ListSkillsRequest
	GetScope() *string
	SetSearch(v string) *ListSkillsRequest
	GetSearch() *string
	SetSkillName(v string) *ListSkillsRequest
	GetSkillName() *string
}

type ListSkillsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// example:
	//
	// download_count
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	// example:
	//
	// user123
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// PUBLIC
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// blur
	Search *string `json:"Search,omitempty" xml:"Search,omitempty"`
	// example:
	//
	// customer
	SkillName *string `json:"SkillName,omitempty" xml:"SkillName,omitempty"`
}

func (s ListSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListSkillsRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListSkillsRequest) GetOwner() *string {
	return s.Owner
}

func (s *ListSkillsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListSkillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListSkillsRequest) GetSearch() *string {
	return s.Search
}

func (s *ListSkillsRequest) GetSkillName() *string {
	return s.SkillName
}

func (s *ListSkillsRequest) SetNamespaceId(v string) *ListSkillsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListSkillsRequest) SetOrderBy(v string) *ListSkillsRequest {
	s.OrderBy = &v
	return s
}

func (s *ListSkillsRequest) SetOwner(v string) *ListSkillsRequest {
	s.Owner = &v
	return s
}

func (s *ListSkillsRequest) SetPageNo(v int32) *ListSkillsRequest {
	s.PageNo = &v
	return s
}

func (s *ListSkillsRequest) SetPageSize(v int32) *ListSkillsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillsRequest) SetScope(v string) *ListSkillsRequest {
	s.Scope = &v
	return s
}

func (s *ListSkillsRequest) SetSearch(v string) *ListSkillsRequest {
	s.Search = &v
	return s
}

func (s *ListSkillsRequest) SetSkillName(v string) *ListSkillsRequest {
	s.SkillName = &v
	return s
}

func (s *ListSkillsRequest) Validate() error {
	return dara.Validate(s)
}
