// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcubeUpgradeTaskInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetMcubeUpgradeTaskInfoRequest
	GetAppId() *string
	SetTaskId(v int64) *GetMcubeUpgradeTaskInfoRequest
	GetTaskId() *int64
	SetTenantId(v string) *GetMcubeUpgradeTaskInfoRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *GetMcubeUpgradeTaskInfoRequest
	GetWorkspaceId() *string
}

type GetMcubeUpgradeTaskInfoRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	TaskId      *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetMcubeUpgradeTaskInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcubeUpgradeTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *GetMcubeUpgradeTaskInfoRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetMcubeUpgradeTaskInfoRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *GetMcubeUpgradeTaskInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetMcubeUpgradeTaskInfoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetMcubeUpgradeTaskInfoRequest) SetAppId(v string) *GetMcubeUpgradeTaskInfoRequest {
	s.AppId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoRequest) SetTaskId(v int64) *GetMcubeUpgradeTaskInfoRequest {
	s.TaskId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoRequest) SetTenantId(v string) *GetMcubeUpgradeTaskInfoRequest {
	s.TenantId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoRequest) SetWorkspaceId(v string) *GetMcubeUpgradeTaskInfoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetMcubeUpgradeTaskInfoRequest) Validate() error {
	return dara.Validate(s)
}
