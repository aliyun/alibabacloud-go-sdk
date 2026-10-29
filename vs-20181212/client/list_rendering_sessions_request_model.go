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
	// The cloud application ID.
	//
	// example:
	//
	// cap-b06b26edfhytbn b94a75ae1a79efc90eb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The unique ID of the end client.
	//
	// example:
	//
	// ae7990f4-203d-494b-a5ea-e0babe9fa13d
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The start time of the time range filter. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-03-31T22:20:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The page number. Pages start from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page during a paging query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the cloud application patch package. If you set this parameter to origin, only sessions that use the original version of the application are returned.
	//
	// example:
	//
	// patch-03fa76e8e13a49b6a966b063d9d309b4
	PatchId *string `json:"PatchId,omitempty" xml:"PatchId,omitempty"`
	// The project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// project-422bc38dfgh5eb44149f135ef76304f63b
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The instance ID of the cloud application service.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// session-i205217481741918129226
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The start time of the time range filter. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2025-03-25T02:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The session state. Valid values:
	//
	// - SessionStarting: The session is starting.
	//
	// - SessionStartSuspended: The session startup is suspended. You can retry by initiating a start request again.
	//
	// - SessionStarted: The session is started or in use.
	//
	// - SessionStartFailed: The session failed to start.
	//
	// - SessionAbnormal: An exception occurred after the session was successfully started.
	//
	// - SessionStopping: The session is stopping.
	//
	// - SessionStopFailed: The session failed to stop.
	//
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
