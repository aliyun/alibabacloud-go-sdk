// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubePublicTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ChangeMcubePublicTaskStatusRequest
	GetAppId() *string
	SetTaskId(v string) *ChangeMcubePublicTaskStatusRequest
	GetTaskId() *string
	SetTaskStatus(v string) *ChangeMcubePublicTaskStatusRequest
	GetTaskStatus() *string
	SetTenantId(v string) *ChangeMcubePublicTaskStatusRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ChangeMcubePublicTaskStatusRequest
	GetWorkspaceId() *string
}

type ChangeMcubePublicTaskStatusRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus  *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ChangeMcubePublicTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubePublicTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeMcubePublicTaskStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *ChangeMcubePublicTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ChangeMcubePublicTaskStatusRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *ChangeMcubePublicTaskStatusRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ChangeMcubePublicTaskStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ChangeMcubePublicTaskStatusRequest) SetAppId(v string) *ChangeMcubePublicTaskStatusRequest {
	s.AppId = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusRequest) SetTaskId(v string) *ChangeMcubePublicTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusRequest) SetTaskStatus(v string) *ChangeMcubePublicTaskStatusRequest {
	s.TaskStatus = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusRequest) SetTenantId(v string) *ChangeMcubePublicTaskStatusRequest {
	s.TenantId = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusRequest) SetWorkspaceId(v string) *ChangeMcubePublicTaskStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ChangeMcubePublicTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
