// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeNebulaResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeNebulaResourcesRequest
	GetAppId() *string
	SetH5Id(v string) *ListMcubeNebulaResourcesRequest
	GetH5Id() *string
	SetPageNum(v int32) *ListMcubeNebulaResourcesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMcubeNebulaResourcesRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListMcubeNebulaResourcesRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeNebulaResourcesRequest
	GetWorkspaceId() *string
}

type ListMcubeNebulaResourcesRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	H5Id        *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeNebulaResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaResourcesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeNebulaResourcesRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *ListMcubeNebulaResourcesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMcubeNebulaResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeNebulaResourcesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeNebulaResourcesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeNebulaResourcesRequest) SetAppId(v string) *ListMcubeNebulaResourcesRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeNebulaResourcesRequest) SetH5Id(v string) *ListMcubeNebulaResourcesRequest {
	s.H5Id = &v
	return s
}

func (s *ListMcubeNebulaResourcesRequest) SetPageNum(v int32) *ListMcubeNebulaResourcesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMcubeNebulaResourcesRequest) SetPageSize(v int32) *ListMcubeNebulaResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcubeNebulaResourcesRequest) SetTenantId(v string) *ListMcubeNebulaResourcesRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeNebulaResourcesRequest) SetWorkspaceId(v string) *ListMcubeNebulaResourcesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeNebulaResourcesRequest) Validate() error {
	return dara.Validate(s)
}
