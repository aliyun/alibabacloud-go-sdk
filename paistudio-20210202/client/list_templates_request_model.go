// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabel(v string) *ListTemplatesRequest
	GetLabel() *string
	SetList(v string) *ListTemplatesRequest
	GetList() *string
	SetName(v string) *ListTemplatesRequest
	GetName() *string
	SetOrder(v string) *ListTemplatesRequest
	GetOrder() *string
	SetPageNumber(v int32) *ListTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListTemplatesRequest
	GetPageSize() *int32
	SetSortBy(v string) *ListTemplatesRequest
	GetSortBy() *string
	SetSource(v string) *ListTemplatesRequest
	GetSource() *string
	SetTagId(v string) *ListTemplatesRequest
	GetTagId() *string
	SetTemplateType(v string) *ListTemplatesRequest
	GetTemplateType() *string
	SetTypeId(v string) *ListTemplatesRequest
	GetTypeId() *string
	SetVerbose(v bool) *ListTemplatesRequest
	GetVerbose() *bool
	SetWorkspaceId(v string) *ListTemplatesRequest
	GetWorkspaceId() *string
}

type ListTemplatesRequest struct {
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// types
	List *string `json:"List,omitempty" xml:"List,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// DESC
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SortBy   *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// example:
	//
	// PaiStudio
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// template-tag-12345
	TagId        *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// example:
	//
	// template-type-12345
	TypeId      *string `json:"TypeId,omitempty" xml:"TypeId,omitempty"`
	Verbose     *bool   `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListTemplatesRequest) GetLabel() *string {
	return s.Label
}

func (s *ListTemplatesRequest) GetList() *string {
	return s.List
}

func (s *ListTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListTemplatesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTemplatesRequest) GetSortBy() *string {
	return s.SortBy
}

func (s *ListTemplatesRequest) GetSource() *string {
	return s.Source
}

func (s *ListTemplatesRequest) GetTagId() *string {
	return s.TagId
}

func (s *ListTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *ListTemplatesRequest) GetTypeId() *string {
	return s.TypeId
}

func (s *ListTemplatesRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListTemplatesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListTemplatesRequest) SetLabel(v string) *ListTemplatesRequest {
	s.Label = &v
	return s
}

func (s *ListTemplatesRequest) SetList(v string) *ListTemplatesRequest {
	s.List = &v
	return s
}

func (s *ListTemplatesRequest) SetName(v string) *ListTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListTemplatesRequest) SetOrder(v string) *ListTemplatesRequest {
	s.Order = &v
	return s
}

func (s *ListTemplatesRequest) SetPageNumber(v int32) *ListTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTemplatesRequest) SetPageSize(v int32) *ListTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTemplatesRequest) SetSortBy(v string) *ListTemplatesRequest {
	s.SortBy = &v
	return s
}

func (s *ListTemplatesRequest) SetSource(v string) *ListTemplatesRequest {
	s.Source = &v
	return s
}

func (s *ListTemplatesRequest) SetTagId(v string) *ListTemplatesRequest {
	s.TagId = &v
	return s
}

func (s *ListTemplatesRequest) SetTemplateType(v string) *ListTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *ListTemplatesRequest) SetTypeId(v string) *ListTemplatesRequest {
	s.TypeId = &v
	return s
}

func (s *ListTemplatesRequest) SetVerbose(v bool) *ListTemplatesRequest {
	s.Verbose = &v
	return s
}

func (s *ListTemplatesRequest) SetWorkspaceId(v string) *ListTemplatesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
