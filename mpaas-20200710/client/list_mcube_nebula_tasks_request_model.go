// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeNebulaTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeNebulaTasksRequest
	GetAppId() *string
	SetId(v int64) *ListMcubeNebulaTasksRequest
	GetId() *int64
	SetTenantId(v string) *ListMcubeNebulaTasksRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeNebulaTasksRequest
	GetWorkspaceId() *string
}

type ListMcubeNebulaTasksRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Id          *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeNebulaTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeNebulaTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeNebulaTasksRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeNebulaTasksRequest) GetId() *int64 {
	return s.Id
}

func (s *ListMcubeNebulaTasksRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeNebulaTasksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeNebulaTasksRequest) SetAppId(v string) *ListMcubeNebulaTasksRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeNebulaTasksRequest) SetId(v int64) *ListMcubeNebulaTasksRequest {
	s.Id = &v
	return s
}

func (s *ListMcubeNebulaTasksRequest) SetTenantId(v string) *ListMcubeNebulaTasksRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeNebulaTasksRequest) SetWorkspaceId(v string) *ListMcubeNebulaTasksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeNebulaTasksRequest) Validate() error {
	return dara.Validate(s)
}
