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
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// HMS
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// example:
	//
	// 1745337419999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// ALIPUBE5C3F6D091419_ANDROID-default
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// 1745251200000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// 23876
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// LZFPEFIM
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// SIMPLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// default
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
