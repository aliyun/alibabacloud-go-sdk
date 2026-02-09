// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMcubeHotpatchTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryMcubeHotpatchTaskDetailRequest
	GetAppId() *string
	SetTaskId(v int64) *QueryMcubeHotpatchTaskDetailRequest
	GetTaskId() *int64
	SetTenantId(v string) *QueryMcubeHotpatchTaskDetailRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryMcubeHotpatchTaskDetailRequest
	GetWorkspaceId() *string
}

type QueryMcubeHotpatchTaskDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6071959
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ZXCXMAHQ
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// default
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryMcubeHotpatchTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMcubeHotpatchTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryMcubeHotpatchTaskDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryMcubeHotpatchTaskDetailRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *QueryMcubeHotpatchTaskDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryMcubeHotpatchTaskDetailRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryMcubeHotpatchTaskDetailRequest) SetAppId(v string) *QueryMcubeHotpatchTaskDetailRequest {
	s.AppId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailRequest) SetTaskId(v int64) *QueryMcubeHotpatchTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailRequest) SetTenantId(v string) *QueryMcubeHotpatchTaskDetailRequest {
	s.TenantId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailRequest) SetWorkspaceId(v string) *QueryMcubeHotpatchTaskDetailRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryMcubeHotpatchTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
