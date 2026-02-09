// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeHotpatchTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeHotpatchTasksRequest
	GetAppId() *string
	SetId(v int64) *ListMcubeHotpatchTasksRequest
	GetId() *int64
	SetTenantId(v string) *ListMcubeHotpatchTasksRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeHotpatchTasksRequest
	GetWorkspaceId() *string
}

type ListMcubeHotpatchTasksRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1387
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s ListMcubeHotpatchTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeHotpatchTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeHotpatchTasksRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeHotpatchTasksRequest) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeHotpatchTasksRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeHotpatchTasksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeHotpatchTasksRequest) SetAppId(v string) *ListMcubeHotpatchTasksRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeHotpatchTasksRequest) SetId(v int64) *ListMcubeHotpatchTasksRequest {
	s.Id = &v
	return s
}

func (s *ListMcubeHotpatchTasksRequest) SetTenantId(v string) *ListMcubeHotpatchTasksRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeHotpatchTasksRequest) SetWorkspaceId(v string) *ListMcubeHotpatchTasksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeHotpatchTasksRequest) Validate() error {
	return dara.Validate(s)
}
