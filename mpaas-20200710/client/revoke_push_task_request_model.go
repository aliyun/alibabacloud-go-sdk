// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePushTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RevokePushTaskRequest
	GetAppId() *string
	SetTaskId(v string) *RevokePushTaskRequest
	GetTaskId() *string
	SetTenantId(v string) *RevokePushTaskRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *RevokePushTaskRequest
	GetWorkspaceId() *string
}

type RevokePushTaskRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RevokePushTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokePushTaskRequest) GoString() string {
	return s.String()
}

func (s *RevokePushTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *RevokePushTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RevokePushTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RevokePushTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RevokePushTaskRequest) SetAppId(v string) *RevokePushTaskRequest {
	s.AppId = &v
	return s
}

func (s *RevokePushTaskRequest) SetTaskId(v string) *RevokePushTaskRequest {
	s.TaskId = &v
	return s
}

func (s *RevokePushTaskRequest) SetTenantId(v string) *RevokePushTaskRequest {
	s.TenantId = &v
	return s
}

func (s *RevokePushTaskRequest) SetWorkspaceId(v string) *RevokePushTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RevokePushTaskRequest) Validate() error {
	return dara.Validate(s)
}
