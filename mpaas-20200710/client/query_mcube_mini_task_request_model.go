// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeMiniTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMcubeMiniTaskRequest
	GetAppId() *string
	SetTaskId(v int64) *QueryMcubeMiniTaskRequest
	GetTaskId() *int64
	SetTenantId(v string) *QueryMcubeMiniTaskRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMcubeMiniTaskRequest
	GetWorkspaceId() *string
}

type QueryMcubeMiniTaskRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMcubeMiniTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeMiniTaskRequest) GoString() string {
	return s.String()
}

func (s *QueryMcubeMiniTaskRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMcubeMiniTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryMcubeMiniTaskRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMcubeMiniTaskRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMcubeMiniTaskRequest) SetAppId(v string) *QueryMcubeMiniTaskRequest {
	s.AppId = &v
	return s
}

func (s *QueryMcubeMiniTaskRequest) SetTaskId(v int64) *QueryMcubeMiniTaskRequest {
	s.TaskId = &v
	return s
}

func (s *QueryMcubeMiniTaskRequest) SetTenantId(v string) *QueryMcubeMiniTaskRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMcubeMiniTaskRequest) SetWorkspaceId(v string) *QueryMcubeMiniTaskRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMcubeMiniTaskRequest) Validate() error {
	return dara.Validate(s)
}
