// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeMcubeNebulaTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ChangeMcubeNebulaTaskStatusRequest
	GetAppId() *string
	SetBizType(v string) *ChangeMcubeNebulaTaskStatusRequest
	GetBizType() *string
	SetPackageId(v string) *ChangeMcubeNebulaTaskStatusRequest
	GetPackageId() *string
	SetTaskId(v string) *ChangeMcubeNebulaTaskStatusRequest
	GetTaskId() *string
	SetTaskStatus(v int32) *ChangeMcubeNebulaTaskStatusRequest
	GetTaskStatus() *int32
	SetTenantId(v string) *ChangeMcubeNebulaTaskStatusRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *ChangeMcubeNebulaTaskStatusRequest
	GetWorkspaceId() *string
}

type ChangeMcubeNebulaTaskStatusRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	PackageId   *string `json:"PackageId,omitempty" xml:"PackageId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus  *int32  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ChangeMcubeNebulaTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeMcubeNebulaTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *ChangeMcubeNebulaTaskStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *ChangeMcubeNebulaTaskStatusRequest) GetBizType() *string {
	return s.BizType
}

func (s *ChangeMcubeNebulaTaskStatusRequest) GetPackageId() *string {
	return s.PackageId
}

func (s *ChangeMcubeNebulaTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ChangeMcubeNebulaTaskStatusRequest) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *ChangeMcubeNebulaTaskStatusRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ChangeMcubeNebulaTaskStatusRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ChangeMcubeNebulaTaskStatusRequest) SetAppId(v string) *ChangeMcubeNebulaTaskStatusRequest {
	s.AppId = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusRequest) SetBizType(v string) *ChangeMcubeNebulaTaskStatusRequest {
	s.BizType = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusRequest) SetPackageId(v string) *ChangeMcubeNebulaTaskStatusRequest {
	s.PackageId = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusRequest) SetTaskId(v string) *ChangeMcubeNebulaTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusRequest) SetTaskStatus(v int32) *ChangeMcubeNebulaTaskStatusRequest {
	s.TaskStatus = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusRequest) SetTenantId(v string) *ChangeMcubeNebulaTaskStatusRequest {
	s.TenantId = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusRequest) SetWorkspaceId(v string) *ChangeMcubeNebulaTaskStatusRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ChangeMcubeNebulaTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
