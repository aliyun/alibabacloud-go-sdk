// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushAnalysisTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryPushAnalysisTaskDetailRequest
	GetAppId() *string
	SetTaskId(v string) *QueryPushAnalysisTaskDetailRequest
	GetTaskId() *string
	SetTenantId(v string) *QueryPushAnalysisTaskDetailRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryPushAnalysisTaskDetailRequest
	GetWorkspaceId() *string
}

type QueryPushAnalysisTaskDetailRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	TaskId   *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryPushAnalysisTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryPushAnalysisTaskDetailRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryPushAnalysisTaskDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryPushAnalysisTaskDetailRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryPushAnalysisTaskDetailRequest) SetAppId(v string) *QueryPushAnalysisTaskDetailRequest {
	s.AppId = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailRequest) SetTaskId(v string) *QueryPushAnalysisTaskDetailRequest {
	s.TaskId = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailRequest) SetTenantId(v string) *QueryPushAnalysisTaskDetailRequest {
	s.TenantId = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailRequest) SetWorkspaceId(v string) *QueryPushAnalysisTaskDetailRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryPushAnalysisTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
