// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushAnalysisTaskListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryPushAnalysisTaskListRequest
	GetAppId() *string
	SetPageNumber(v int32) *QueryPushAnalysisTaskListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryPushAnalysisTaskListRequest
	GetPageSize() *int32
	SetStartTime(v int64) *QueryPushAnalysisTaskListRequest
	GetStartTime() *int64
	SetTaskId(v string) *QueryPushAnalysisTaskListRequest
	GetTaskId() *string
	SetTaskName(v string) *QueryPushAnalysisTaskListRequest
	GetTaskName() *string
	SetTenantId(v string) *QueryPushAnalysisTaskListRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *QueryPushAnalysisTaskListRequest
	GetWorkspaceId() *string
}

type QueryPushAnalysisTaskListRequest struct {
	// This parameter is required.
	AppId      *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName  *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryPushAnalysisTaskListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisTaskListRequest) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisTaskListRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryPushAnalysisTaskListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryPushAnalysisTaskListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryPushAnalysisTaskListRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryPushAnalysisTaskListRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryPushAnalysisTaskListRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *QueryPushAnalysisTaskListRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryPushAnalysisTaskListRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryPushAnalysisTaskListRequest) SetAppId(v string) *QueryPushAnalysisTaskListRequest {
	s.AppId = &v
	return s
}

func (s *QueryPushAnalysisTaskListRequest) SetPageNumber(v int32) *QueryPushAnalysisTaskListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryPushAnalysisTaskListRequest) SetPageSize(v int32) *QueryPushAnalysisTaskListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryPushAnalysisTaskListRequest) SetStartTime(v int64) *QueryPushAnalysisTaskListRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPushAnalysisTaskListRequest) SetTaskId(v string) *QueryPushAnalysisTaskListRequest {
	s.TaskId = &v
	return s
}

func (s *QueryPushAnalysisTaskListRequest) SetTaskName(v string) *QueryPushAnalysisTaskListRequest {
	s.TaskName = &v
	return s
}

func (s *QueryPushAnalysisTaskListRequest) SetTenantId(v string) *QueryPushAnalysisTaskListRequest {
	s.TenantId = &v
	return s
}

func (s *QueryPushAnalysisTaskListRequest) SetWorkspaceId(v string) *QueryPushAnalysisTaskListRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryPushAnalysisTaskListRequest) Validate() error {
	return dara.Validate(s)
}
