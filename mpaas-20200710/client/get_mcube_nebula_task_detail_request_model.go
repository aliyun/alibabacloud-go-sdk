// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeNebulaTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMcubeNebulaTaskDetailRequest
	GetAppId() *string
	SetTaskId(v int64) *GetMcubeNebulaTaskDetailRequest
	GetTaskId() *int64
	SetTenantId(v string) *GetMcubeNebulaTaskDetailRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetMcubeNebulaTaskDetailRequest
	GetWorkspaceId() *string
}

type GetMcubeNebulaTaskDetailRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TaskId      *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMcubeNebulaTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeNebulaTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *GetMcubeNebulaTaskDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMcubeNebulaTaskDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetMcubeNebulaTaskDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMcubeNebulaTaskDetailRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMcubeNebulaTaskDetailRequest) SetAppId(v string) *GetMcubeNebulaTaskDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailRequest) SetTaskId(v int64) *GetMcubeNebulaTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailRequest) SetTenantId(v string) *GetMcubeNebulaTaskDetailRequest {
	s.TenantId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailRequest) SetWorkspaceId(v string) *GetMcubeNebulaTaskDetailRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetMcubeNebulaTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
