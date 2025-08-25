// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeMiniTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListMcubeMiniTasksRequest
	GetAppId() *string
	SetId(v string) *ListMcubeMiniTasksRequest
	GetId() *string
	SetTenantId(v string) *ListMcubeMiniTasksRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ListMcubeMiniTasksRequest
	GetWorkspaceId() *string
}

type ListMcubeMiniTasksRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListMcubeMiniTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeMiniTasksRequest) GoString() string {
	return s.String()
}

func (s *ListMcubeMiniTasksRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListMcubeMiniTasksRequest) GetId() *string {
	return s.Id
}

func (s *ListMcubeMiniTasksRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListMcubeMiniTasksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListMcubeMiniTasksRequest) SetAppId(v string) *ListMcubeMiniTasksRequest {
	s.AppId = &v
	return s
}

func (s *ListMcubeMiniTasksRequest) SetId(v string) *ListMcubeMiniTasksRequest {
	s.Id = &v
	return s
}

func (s *ListMcubeMiniTasksRequest) SetTenantId(v string) *ListMcubeMiniTasksRequest {
	s.TenantId = &v
	return s
}

func (s *ListMcubeMiniTasksRequest) SetWorkspaceId(v string) *ListMcubeMiniTasksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListMcubeMiniTasksRequest) Validate() error {
	return dara.Validate(s)
}
