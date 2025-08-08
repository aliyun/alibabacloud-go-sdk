// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcdpAimRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcdpAimRequest
	GetAppId() *string
	SetEmptyTag(v string) *ListMcdpAimRequest
	GetEmptyTag() *string
	SetKeyword(v string) *ListMcdpAimRequest
	GetKeyword() *string
	SetName(v string) *ListMcdpAimRequest
	GetName() *string
	SetPageNo(v int64) *ListMcdpAimRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListMcdpAimRequest
	GetPageSize() *int64
	SetSort(v string) *ListMcdpAimRequest
	GetSort() *string
	SetSortField(v string) *ListMcdpAimRequest
	GetSortField() *string
	SetTenantId(v string) *ListMcdpAimRequest
	GetTenantId() *string
	SetType(v string) *ListMcdpAimRequest
	GetType() *string
	SetWorkspaceId(v string) *ListMcdpAimRequest
	GetWorkspaceId() *string
}

type ListMcdpAimRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	EmptyTag    *string `json:"EmptyTag,omitempty" xml:"EmptyTag,omitempty"`
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PageNo      *int64  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int64  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Sort        *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
	SortField   *string `json:"SortField,omitempty" xml:"SortField,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcdpAimRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcdpAimRequest) GoString() string {
	return s.String()
}

func (s *ListMcdpAimRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcdpAimRequest) GetEmptyTag() *string {
	return s.EmptyTag
}

func (s *ListMcdpAimRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMcdpAimRequest) GetName() *string {
	return s.Name
}

func (s *ListMcdpAimRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListMcdpAimRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListMcdpAimRequest) GetSort() *string {
	return s.Sort
}

func (s *ListMcdpAimRequest) GetSortField() *string {
	return s.SortField
}

func (s *ListMcdpAimRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcdpAimRequest) GetType() *string {
	return s.Type
}

func (s *ListMcdpAimRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcdpAimRequest) SetAppId(v string) *ListMcdpAimRequest {
	s.AppId = &v
	return s
}

func (s *ListMcdpAimRequest) SetEmptyTag(v string) *ListMcdpAimRequest {
	s.EmptyTag = &v
	return s
}

func (s *ListMcdpAimRequest) SetKeyword(v string) *ListMcdpAimRequest {
	s.Keyword = &v
	return s
}

func (s *ListMcdpAimRequest) SetName(v string) *ListMcdpAimRequest {
	s.Name = &v
	return s
}

func (s *ListMcdpAimRequest) SetPageNo(v int64) *ListMcdpAimRequest {
	s.PageNo = &v
	return s
}

func (s *ListMcdpAimRequest) SetPageSize(v int64) *ListMcdpAimRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcdpAimRequest) SetSort(v string) *ListMcdpAimRequest {
	s.Sort = &v
	return s
}

func (s *ListMcdpAimRequest) SetSortField(v string) *ListMcdpAimRequest {
	s.SortField = &v
	return s
}

func (s *ListMcdpAimRequest) SetTenantId(v string) *ListMcdpAimRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcdpAimRequest) SetType(v string) *ListMcdpAimRequest {
	s.Type = &v
	return s
}

func (s *ListMcdpAimRequest) SetWorkspaceId(v string) *ListMcdpAimRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcdpAimRequest) Validate() error {
	return dara.Validate(s)
}
