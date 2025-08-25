// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushAnalysisCoreIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *QueryPushAnalysisCoreIndexRequest
	GetAppId() *string
	SetChannel(v string) *QueryPushAnalysisCoreIndexRequest
	GetChannel() *string
	SetEndTime(v int64) *QueryPushAnalysisCoreIndexRequest
	GetEndTime() *int64
	SetPlatform(v string) *QueryPushAnalysisCoreIndexRequest
	GetPlatform() *string
	SetStartTime(v int64) *QueryPushAnalysisCoreIndexRequest
	GetStartTime() *int64
	SetTaskId(v string) *QueryPushAnalysisCoreIndexRequest
	GetTaskId() *string
	SetTenantId(v string) *QueryPushAnalysisCoreIndexRequest
	GetTenantId() *string
	SetType(v string) *QueryPushAnalysisCoreIndexRequest
	GetType() *string
	SetWorkspaceId(v string) *QueryPushAnalysisCoreIndexRequest
	GetWorkspaceId() *string
}

type QueryPushAnalysisCoreIndexRequest struct {
	// This parameter is required.
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// This parameter is required.
	EndTime  *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// This parameter is required.
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type      *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s QueryPushAnalysisCoreIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryPushAnalysisCoreIndexRequest) GoString() string {
	return s.String()
}

func (s *QueryPushAnalysisCoreIndexRequest) GetAppId() *string {
	return s.AppId
}

func (s *QueryPushAnalysisCoreIndexRequest) GetChannel() *string {
	return s.Channel
}

func (s *QueryPushAnalysisCoreIndexRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryPushAnalysisCoreIndexRequest) GetPlatform() *string {
	return s.Platform
}

func (s *QueryPushAnalysisCoreIndexRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryPushAnalysisCoreIndexRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *QueryPushAnalysisCoreIndexRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryPushAnalysisCoreIndexRequest) GetType() *string {
	return s.Type
}

func (s *QueryPushAnalysisCoreIndexRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *QueryPushAnalysisCoreIndexRequest) SetAppId(v string) *QueryPushAnalysisCoreIndexRequest {
	s.AppId = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexRequest) SetChannel(v string) *QueryPushAnalysisCoreIndexRequest {
	s.Channel = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexRequest) SetEndTime(v int64) *QueryPushAnalysisCoreIndexRequest {
	s.EndTime = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexRequest) SetPlatform(v string) *QueryPushAnalysisCoreIndexRequest {
	s.Platform = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexRequest) SetStartTime(v int64) *QueryPushAnalysisCoreIndexRequest {
	s.StartTime = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexRequest) SetTaskId(v string) *QueryPushAnalysisCoreIndexRequest {
	s.TaskId = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexRequest) SetTenantId(v string) *QueryPushAnalysisCoreIndexRequest {
	s.TenantId = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexRequest) SetType(v string) *QueryPushAnalysisCoreIndexRequest {
	s.Type = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexRequest) SetWorkspaceId(v string) *QueryPushAnalysisCoreIndexRequest {
	s.WorkspaceId = &v
	return s
}

func (s *QueryPushAnalysisCoreIndexRequest) Validate() error {
	return dara.Validate(s)
}
