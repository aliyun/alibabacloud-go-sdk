// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMdsUpgradeTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMdsUpgradeTaskDetailRequest
	GetAppId() *string
	SetTaskId(v int64) *QueryMdsUpgradeTaskDetailRequest
	GetTaskId() *int64
	SetTenantId(v string) *QueryMdsUpgradeTaskDetailRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMdsUpgradeTaskDetailRequest
	GetWorkspaceId() *string
}

type QueryMdsUpgradeTaskDetailRequest struct {
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TaskId      *int64  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMdsUpgradeTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMdsUpgradeTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryMdsUpgradeTaskDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMdsUpgradeTaskDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryMdsUpgradeTaskDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMdsUpgradeTaskDetailRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMdsUpgradeTaskDetailRequest) SetAppId(v string) *QueryMdsUpgradeTaskDetailRequest {
	s.AppId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailRequest) SetTaskId(v int64) *QueryMdsUpgradeTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailRequest) SetTenantId(v string) *QueryMdsUpgradeTaskDetailRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailRequest) SetWorkspaceId(v string) *QueryMdsUpgradeTaskDetailRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMdsUpgradeTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
