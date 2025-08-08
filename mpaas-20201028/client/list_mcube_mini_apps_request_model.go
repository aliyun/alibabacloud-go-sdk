// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeMiniAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeMiniAppsRequest
	GetAppId() *string
	SetKeyword(v string) *ListMcubeMiniAppsRequest
	GetKeyword() *string
	SetPageNum(v int32) *ListMcubeMiniAppsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMcubeMiniAppsRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListMcubeMiniAppsRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeMiniAppsRequest
	GetWorkspaceId() *string
}

type ListMcubeMiniAppsRequest struct {
	// This parameter is required.
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Keyword  *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeMiniAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniAppsRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniAppsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeMiniAppsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMcubeMiniAppsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMcubeMiniAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeMiniAppsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeMiniAppsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeMiniAppsRequest) SetAppId(v string) *ListMcubeMiniAppsRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeMiniAppsRequest) SetKeyword(v string) *ListMcubeMiniAppsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMcubeMiniAppsRequest) SetPageNum(v int32) *ListMcubeMiniAppsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMcubeMiniAppsRequest) SetPageSize(v int32) *ListMcubeMiniAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcubeMiniAppsRequest) SetTenantId(v string) *ListMcubeMiniAppsRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeMiniAppsRequest) SetWorkspaceId(v string) *ListMcubeMiniAppsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeMiniAppsRequest) Validate() error {
	return dara.Validate(s)
}
