// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnalysisCoreIndexRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAnalysisCoreIndexRequest
	GetAppId() *string
	SetChannel(v string) *ListAnalysisCoreIndexRequest
	GetChannel() *string
	SetEndTime(v int64) *ListAnalysisCoreIndexRequest
	GetEndTime() *int64
	SetPlatform(v string) *ListAnalysisCoreIndexRequest
	GetPlatform() *string
	SetStartTime(v int64) *ListAnalysisCoreIndexRequest
	GetStartTime() *int64
	SetTaskId(v string) *ListAnalysisCoreIndexRequest
	GetTaskId() *string
	SetTenantId(v string) *ListAnalysisCoreIndexRequest
	GetTenantId() *string
	SetType(v string) *ListAnalysisCoreIndexRequest
	GetType() *string
	SetWorkspaceId(v string) *ListAnalysisCoreIndexRequest
	GetWorkspaceId() *string
}

type ListAnalysisCoreIndexRequest struct {
	AppId       *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Channel     *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	EndTime     *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Platform    *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	StartTime   *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAnalysisCoreIndexRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisCoreIndexRequest) GoString() string {
	return s.String()
}

func (s *ListAnalysisCoreIndexRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAnalysisCoreIndexRequest) GetChannel() *string {
	return s.Channel
}

func (s *ListAnalysisCoreIndexRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListAnalysisCoreIndexRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ListAnalysisCoreIndexRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListAnalysisCoreIndexRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAnalysisCoreIndexRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ListAnalysisCoreIndexRequest) GetType() *string {
	return s.Type
}

func (s *ListAnalysisCoreIndexRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAnalysisCoreIndexRequest) SetAppId(v string) *ListAnalysisCoreIndexRequest {
	s.AppId = &v
	return s
}

func (s *ListAnalysisCoreIndexRequest) SetChannel(v string) *ListAnalysisCoreIndexRequest {
	s.Channel = &v
	return s
}

func (s *ListAnalysisCoreIndexRequest) SetEndTime(v int64) *ListAnalysisCoreIndexRequest {
	s.EndTime = &v
	return s
}

func (s *ListAnalysisCoreIndexRequest) SetPlatform(v string) *ListAnalysisCoreIndexRequest {
	s.Platform = &v
	return s
}

func (s *ListAnalysisCoreIndexRequest) SetStartTime(v int64) *ListAnalysisCoreIndexRequest {
	s.StartTime = &v
	return s
}

func (s *ListAnalysisCoreIndexRequest) SetTaskId(v string) *ListAnalysisCoreIndexRequest {
	s.TaskId = &v
	return s
}

func (s *ListAnalysisCoreIndexRequest) SetTenantId(v string) *ListAnalysisCoreIndexRequest {
	s.TenantId = &v
	return s
}

func (s *ListAnalysisCoreIndexRequest) SetType(v string) *ListAnalysisCoreIndexRequest {
	s.Type = &v
	return s
}

func (s *ListAnalysisCoreIndexRequest) SetWorkspaceId(v string) *ListAnalysisCoreIndexRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAnalysisCoreIndexRequest) Validate() error {
	return dara.Validate(s)
}
