// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeNebulaAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeNebulaAppsRequest
	GetAppId() *string
	SetKeyword(v string) *ListMcubeNebulaAppsRequest
	GetKeyword() *string
	SetPageNum(v int32) *ListMcubeNebulaAppsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMcubeNebulaAppsRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListMcubeNebulaAppsRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeNebulaAppsRequest
	GetWorkspaceId() *string
}

type ListMcubeNebulaAppsRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Keyword     *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeNebulaAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaAppsRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaAppsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeNebulaAppsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMcubeNebulaAppsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMcubeNebulaAppsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeNebulaAppsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeNebulaAppsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeNebulaAppsRequest) SetAppId(v string) *ListMcubeNebulaAppsRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeNebulaAppsRequest) SetKeyword(v string) *ListMcubeNebulaAppsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMcubeNebulaAppsRequest) SetPageNum(v int32) *ListMcubeNebulaAppsRequest {
	s.PageNum = &v
	return s
}

func (s *ListMcubeNebulaAppsRequest) SetPageSize(v int32) *ListMcubeNebulaAppsRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcubeNebulaAppsRequest) SetTenantId(v string) *ListMcubeNebulaAppsRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeNebulaAppsRequest) SetWorkspaceId(v string) *ListMcubeNebulaAppsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeNebulaAppsRequest) Validate() error {
	return dara.Validate(s)
}
