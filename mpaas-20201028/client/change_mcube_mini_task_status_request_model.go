// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubeMiniTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ChangeMcubeMiniTaskStatusRequest
	GetAppId() *string
	SetBizType(v string) *ChangeMcubeMiniTaskStatusRequest
	GetBizType() *string
	SetPackageId(v int64) *ChangeMcubeMiniTaskStatusRequest
	GetPackageId() *int64
	SetTaskId(v int64) *ChangeMcubeMiniTaskStatusRequest
	GetTaskId() *int64
	SetTaskStatus(v int64) *ChangeMcubeMiniTaskStatusRequest
	GetTaskStatus() *int64
	SetTenantId(v string) *ChangeMcubeMiniTaskStatusRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ChangeMcubeMiniTaskStatusRequest
	GetWorkspaceId() *string
}

type ChangeMcubeMiniTaskStatusRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// This parameter is required.
	PackageId *int64 `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	// This parameter is required.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	TaskStatus *int64 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ChangeMcubeMiniTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubeMiniTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeMcubeMiniTaskStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *ChangeMcubeMiniTaskStatusRequest) GetBizType() *string {
	return s.BizType
}

func (s *ChangeMcubeMiniTaskStatusRequest) GetPackageId() *int64 {
	return s.PackageId
}

func (s *ChangeMcubeMiniTaskStatusRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *ChangeMcubeMiniTaskStatusRequest) GetTaskStatus() *int64 {
	return s.TaskStatus
}

func (s *ChangeMcubeMiniTaskStatusRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ChangeMcubeMiniTaskStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ChangeMcubeMiniTaskStatusRequest) SetAppId(v string) *ChangeMcubeMiniTaskStatusRequest {
	s.AppId = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusRequest) SetBizType(v string) *ChangeMcubeMiniTaskStatusRequest {
	s.BizType = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusRequest) SetPackageId(v int64) *ChangeMcubeMiniTaskStatusRequest {
	s.PackageId = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusRequest) SetTaskId(v int64) *ChangeMcubeMiniTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusRequest) SetTaskStatus(v int64) *ChangeMcubeMiniTaskStatusRequest {
	s.TaskStatus = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusRequest) SetTenantId(v string) *ChangeMcubeMiniTaskStatusRequest {
	s.TenantId = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusRequest) SetWorkspaceId(v string) *ChangeMcubeMiniTaskStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ChangeMcubeMiniTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
