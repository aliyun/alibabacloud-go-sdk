// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSessionsResponseBodyData) *ListSessionsResponseBody
	GetData() *ListSessionsResponseBodyData
	SetRequestId(v string) *ListSessionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSessionsResponseBody
	GetSuccess() *bool
}

type ListSessionsResponseBody struct {
	// The list of sessions and pagination information for the current page.
	Data *ListSessionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Specifies whether the request is successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSessionsResponseBody) GetData() *ListSessionsResponseBodyData {
	return s.Data
}

func (s *ListSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSessionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSessionsResponseBody) SetData(v *ListSessionsResponseBodyData) *ListSessionsResponseBody {
	s.Data = v
	return s
}

func (s *ListSessionsResponseBody) SetRequestId(v string) *ListSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSessionsResponseBody) SetSuccess(v bool) *ListSessionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSessionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSessionsResponseBodyData struct {
	// The pagination token for the next page. A value of null indicates that no more data is available.
	//
	// example:
	//
	// example-next-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The list of sessions returned on the current page.
	Sessions []*ListSessionsResponseBodyDataSessions `json:"sessions,omitempty" xml:"sessions,omitempty" type:"Repeated"`
	// The number of sessions returned on the current page.
	//
	// example:
	//
	// 1
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListSessionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSessionsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSessionsResponseBodyData) GetSessions() []*ListSessionsResponseBodyDataSessions {
	return s.Sessions
}

func (s *ListSessionsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListSessionsResponseBodyData) SetNextToken(v string) *ListSessionsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListSessionsResponseBodyData) SetSessions(v []*ListSessionsResponseBodyDataSessions) *ListSessionsResponseBodyData {
	s.Sessions = v
	return s
}

func (s *ListSessionsResponseBodyData) SetTotal(v int32) *ListSessionsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListSessionsResponseBodyData) Validate() error {
	if s.Sessions != nil {
		for _, item := range s.Sessions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSessionsResponseBodyDataSessions struct {
	// The time when the session was created, in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The ID of the most recent run.
	//
	// example:
	//
	// run_example_001
	LastRunId *string `json:"lastRunId,omitempty" xml:"lastRunId,omitempty"`
	// The status of the most recent run. Valid values:
	//
	// - `PENDING`: Pending.
	//
	// - `RUNNING`: Running.
	//
	// - `WAITING_USER`: Waiting for interrupt handling.
	//
	// - `COMPLETED`: Completed.
	//
	// - `FAILED`: Failed.
	//
	// - `CANCELLED`: Cancelled.
	//
	// - `REJECTED`: Rejected.
	//
	// example:
	//
	// COMPLETED
	LastRunStatus *string `json:"lastRunStatus,omitempty" xml:"lastRunStatus,omitempty"`
	// Specifies whether the session is pinned.
	//
	// example:
	//
	// false
	Pinned *bool `json:"pinned,omitempty" xml:"pinned,omitempty"`
	// The session ID.
	//
	// example:
	//
	// session_example_001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The session status. Valid values:
	//
	// - `IDLE`: Idle.
	//
	// - `ACTIVE`: Active.
	//
	// - `WAITING_USER_INPUT`: Waiting for user input.
	//
	// - `WAITING_APPROVAL`: Waiting for user confirmation.
	//
	// - `WAITING_OUT_OF_BAND_APPROVAL`: Waiting for out-of-band approval.
	//
	// - `CLOSED`: Closed.
	//
	// example:
	//
	// IDLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The session title.
	//
	// example:
	//
	// Cloud resource inspection example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The time when the session was last updated, in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:01:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListSessionsResponseBodyDataSessions) String() string {
	return dara.Prettify(s)
}

func (s ListSessionsResponseBodyDataSessions) GoString() string {
	return s.String()
}

func (s *ListSessionsResponseBodyDataSessions) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListSessionsResponseBodyDataSessions) GetLastRunId() *string {
	return s.LastRunId
}

func (s *ListSessionsResponseBodyDataSessions) GetLastRunStatus() *string {
	return s.LastRunStatus
}

func (s *ListSessionsResponseBodyDataSessions) GetPinned() *bool {
	return s.Pinned
}

func (s *ListSessionsResponseBodyDataSessions) GetSessionId() *string {
	return s.SessionId
}

func (s *ListSessionsResponseBodyDataSessions) GetStatus() *string {
	return s.Status
}

func (s *ListSessionsResponseBodyDataSessions) GetTitle() *string {
	return s.Title
}

func (s *ListSessionsResponseBodyDataSessions) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListSessionsResponseBodyDataSessions) SetCreatedAt(v string) *ListSessionsResponseBodyDataSessions {
	s.CreatedAt = &v
	return s
}

func (s *ListSessionsResponseBodyDataSessions) SetLastRunId(v string) *ListSessionsResponseBodyDataSessions {
	s.LastRunId = &v
	return s
}

func (s *ListSessionsResponseBodyDataSessions) SetLastRunStatus(v string) *ListSessionsResponseBodyDataSessions {
	s.LastRunStatus = &v
	return s
}

func (s *ListSessionsResponseBodyDataSessions) SetPinned(v bool) *ListSessionsResponseBodyDataSessions {
	s.Pinned = &v
	return s
}

func (s *ListSessionsResponseBodyDataSessions) SetSessionId(v string) *ListSessionsResponseBodyDataSessions {
	s.SessionId = &v
	return s
}

func (s *ListSessionsResponseBodyDataSessions) SetStatus(v string) *ListSessionsResponseBodyDataSessions {
	s.Status = &v
	return s
}

func (s *ListSessionsResponseBodyDataSessions) SetTitle(v string) *ListSessionsResponseBodyDataSessions {
	s.Title = &v
	return s
}

func (s *ListSessionsResponseBodyDataSessions) SetUpdatedAt(v string) *ListSessionsResponseBodyDataSessions {
	s.UpdatedAt = &v
	return s
}

func (s *ListSessionsResponseBodyDataSessions) Validate() error {
	return dara.Validate(s)
}
