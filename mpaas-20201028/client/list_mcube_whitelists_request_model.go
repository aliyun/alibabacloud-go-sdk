// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeWhitelistsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeWhitelistsRequest
	GetAppId() *string
	SetPageNum(v int32) *ListMcubeWhitelistsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMcubeWhitelistsRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListMcubeWhitelistsRequest
	GetTenantId() *string
	SetWhitelistName(v string) *ListMcubeWhitelistsRequest
	GetWhitelistName() *string
	SetWorkspaceId(v string) *ListMcubeWhitelistsRequest
	GetWorkspaceId() *string
}

type ListMcubeWhitelistsRequest struct {
	// This parameter is required.
	AppId    *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNum  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WhitelistName *string `json:"WhitelistName,omitempty" xml:"WhitelistName,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeWhitelistsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeWhitelistsRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeWhitelistsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeWhitelistsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMcubeWhitelistsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeWhitelistsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeWhitelistsRequest) GetWhitelistName() *string {
	return s.WhitelistName
}

func (s *ListMcubeWhitelistsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeWhitelistsRequest) SetAppId(v string) *ListMcubeWhitelistsRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeWhitelistsRequest) SetPageNum(v int32) *ListMcubeWhitelistsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMcubeWhitelistsRequest) SetPageSize(v int32) *ListMcubeWhitelistsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcubeWhitelistsRequest) SetTenantId(v string) *ListMcubeWhitelistsRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeWhitelistsRequest) SetWhitelistName(v string) *ListMcubeWhitelistsRequest {
	s.WhitelistName = &v
	return s
}

func (s *ListMcubeWhitelistsRequest) SetWorkspaceId(v string) *ListMcubeWhitelistsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeWhitelistsRequest) Validate() error {
	return dara.Validate(s)
}
