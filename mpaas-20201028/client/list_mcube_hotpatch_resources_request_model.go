// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeHotpatchResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeHotpatchResourcesRequest
	GetAppId() *string
	SetPageNum(v int32) *ListMcubeHotpatchResourcesRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMcubeHotpatchResourcesRequest
	GetPageSize() *int32
	SetTenantId(v string) *ListMcubeHotpatchResourcesRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeHotpatchResourcesRequest
	GetWorkspaceId() *string
}

type ListMcubeHotpatchResourcesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZXCXMAHQ
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeHotpatchResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchResourcesRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeHotpatchResourcesRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMcubeHotpatchResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMcubeHotpatchResourcesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeHotpatchResourcesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeHotpatchResourcesRequest) SetAppId(v string) *ListMcubeHotpatchResourcesRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeHotpatchResourcesRequest) SetPageNum(v int32) *ListMcubeHotpatchResourcesRequest {
	s.PageNum = &v
	return s
}

func (s *ListMcubeHotpatchResourcesRequest) SetPageSize(v int32) *ListMcubeHotpatchResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListMcubeHotpatchResourcesRequest) SetTenantId(v string) *ListMcubeHotpatchResourcesRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeHotpatchResourcesRequest) SetWorkspaceId(v string) *ListMcubeHotpatchResourcesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeHotpatchResourcesRequest) Validate() error {
	return dara.Validate(s)
}
