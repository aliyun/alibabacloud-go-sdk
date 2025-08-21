// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListRenderingSessionsRequest
	GetAppId() *string
	SetClientId(v string) *ListRenderingSessionsRequest
	GetClientId() *string
	SetEndTime(v string) *ListRenderingSessionsRequest
	GetEndTime() *string
	SetPageNumber(v int32) *ListRenderingSessionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRenderingSessionsRequest
	GetPageSize() *int32
	SetPatchId(v string) *ListRenderingSessionsRequest
	GetPatchId() *string
	SetProjectId(v string) *ListRenderingSessionsRequest
	GetProjectId() *string
	SetRenderingInstanceId(v string) *ListRenderingSessionsRequest
	GetRenderingInstanceId() *string
	SetSessionId(v string) *ListRenderingSessionsRequest
	GetSessionId() *string
	SetStartTime(v string) *ListRenderingSessionsRequest
	GetStartTime() *string
	SetState(v string) *ListRenderingSessionsRequest
	GetState() *string
}

type ListRenderingSessionsRequest struct {
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// ae7990f4-203d-494b-a5ea-e0babe9fa13d
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId           *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// example:
	//
	// session-i205217481741918129226
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SessionStarting
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ListRenderingSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListRenderingSessionsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListRenderingSessionsRequest) GetClientId() *string {
	return s.ClientId
}

func (s *ListRenderingSessionsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListRenderingSessionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRenderingSessionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRenderingSessionsRequest) GetPatchId() *string {
	return s.PatchId
}

func (s *ListRenderingSessionsRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *ListRenderingSessionsRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ListRenderingSessionsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *ListRenderingSessionsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListRenderingSessionsRequest) GetState() *string {
	return s.State
}

func (s *ListRenderingSessionsRequest) SetAppId(v string) *ListRenderingSessionsRequest {
	s.AppId = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetClientId(v string) *ListRenderingSessionsRequest {
	s.ClientId = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetEndTime(v string) *ListRenderingSessionsRequest {
	s.EndTime = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetPageNumber(v int32) *ListRenderingSessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetPageSize(v int32) *ListRenderingSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetPatchId(v string) *ListRenderingSessionsRequest {
	s.PatchId = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetProjectId(v string) *ListRenderingSessionsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetRenderingInstanceId(v string) *ListRenderingSessionsRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetSessionId(v string) *ListRenderingSessionsRequest {
	s.SessionId = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetStartTime(v string) *ListRenderingSessionsRequest {
	s.StartTime = &v
	return s
}

func (s *ListRenderingSessionsRequest) SetState(v string) *ListRenderingSessionsRequest {
	s.State = &v
	return s
}

func (s *ListRenderingSessionsRequest) Validate() error {
	return dara.Validate(s)
}
