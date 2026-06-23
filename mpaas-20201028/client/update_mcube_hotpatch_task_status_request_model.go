// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMcubeHotpatchTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateMcubeHotpatchTaskStatusRequest
	GetAppId() *string
	SetBizType(v string) *UpdateMcubeHotpatchTaskStatusRequest
	GetBizType() *string
	SetPackageId(v int64) *UpdateMcubeHotpatchTaskStatusRequest
	GetPackageId() *int64
	SetTaskId(v int64) *UpdateMcubeHotpatchTaskStatusRequest
	GetTaskId() *int64
	SetTaskStatus(v int64) *UpdateMcubeHotpatchTaskStatusRequest
	GetTaskStatus() *int64
	SetTenantId(v string) *UpdateMcubeHotpatchTaskStatusRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *UpdateMcubeHotpatchTaskStatusRequest
	GetWorkspaceId() *string
}

type UpdateMcubeHotpatchTaskStatusRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	BizType   *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PackageId *int64  `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	TaskId    *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	TaskStatus *int64 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateMcubeHotpatchTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMcubeHotpatchTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) GetBizType() *string {
	return s.BizType
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) GetPackageId() *int64 {
	return s.PackageId
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) GetTaskStatus() *int64 {
	return s.TaskStatus
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) SetAppId(v string) *UpdateMcubeHotpatchTaskStatusRequest {
	s.AppId = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) SetBizType(v string) *UpdateMcubeHotpatchTaskStatusRequest {
	s.BizType = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) SetPackageId(v int64) *UpdateMcubeHotpatchTaskStatusRequest {
	s.PackageId = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) SetTaskId(v int64) *UpdateMcubeHotpatchTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) SetTaskStatus(v int64) *UpdateMcubeHotpatchTaskStatusRequest {
	s.TaskStatus = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) SetTenantId(v string) *UpdateMcubeHotpatchTaskStatusRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) SetWorkspaceId(v string) *UpdateMcubeHotpatchTaskStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateMcubeHotpatchTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
