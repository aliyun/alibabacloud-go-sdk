// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSessionResponseBodyData) *GetSessionResponseBody
	GetData() *GetSessionResponseBodyData
	SetRequestId(v string) *GetSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSessionResponseBody
	GetSuccess() *bool
}

type GetSessionResponseBody struct {
	// The returned data.
	Data *GetSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSessionResponseBody) GoString() string {
	return s.String()
}

func (s *GetSessionResponseBody) GetData() *GetSessionResponseBodyData {
	return s.Data
}

func (s *GetSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSessionResponseBody) SetData(v *GetSessionResponseBodyData) *GetSessionResponseBody {
	s.Data = v
	return s
}

func (s *GetSessionResponseBody) SetRequestId(v string) *GetSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSessionResponseBody) SetSuccess(v bool) *GetSessionResponseBody {
	s.Success = &v
	return s
}

func (s *GetSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSessionResponseBodyData struct {
	// The time when the session was created, in ISO 8601 format.
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
	// The status of the most recent run. Valid values: `PENDING` (pending execution), `RUNNING` (running), `WAITING_USER` (waiting for interrupt handling), `COMPLETED` (completed), `FAILED` (failed), `CANCELLED` (cancelled), and `REJECTED` (rejected).
	//
	// example:
	//
	// COMPLETED
	LastRunStatus *string `json:"lastRunStatus,omitempty" xml:"lastRunStatus,omitempty"`
	// Indicates whether the session is pinned.
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
	// The session status. Valid values: `IDLE` (idle), `ACTIVE` (active), `WAITING_USER_INPUT` (waiting for user input), `WAITING_APPROVAL` (waiting for user confirmation), `WAITING_OUT_OF_BAND_APPROVAL` (waiting for out-of-band approval), and `CLOSED` (closed).
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
	// The time when the session was last updated, in ISO 8601 format.
	//
	// example:
	//
	// 2026-01-01T00:01:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s GetSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSessionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetSessionResponseBodyData) GetLastRunId() *string {
	return s.LastRunId
}

func (s *GetSessionResponseBodyData) GetLastRunStatus() *string {
	return s.LastRunStatus
}

func (s *GetSessionResponseBodyData) GetPinned() *bool {
	return s.Pinned
}

func (s *GetSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *GetSessionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetSessionResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *GetSessionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetSessionResponseBodyData) SetCreatedAt(v string) *GetSessionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetSessionResponseBodyData) SetLastRunId(v string) *GetSessionResponseBodyData {
	s.LastRunId = &v
	return s
}

func (s *GetSessionResponseBodyData) SetLastRunStatus(v string) *GetSessionResponseBodyData {
	s.LastRunStatus = &v
	return s
}

func (s *GetSessionResponseBodyData) SetPinned(v bool) *GetSessionResponseBodyData {
	s.Pinned = &v
	return s
}

func (s *GetSessionResponseBodyData) SetSessionId(v string) *GetSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *GetSessionResponseBodyData) SetStatus(v string) *GetSessionResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetSessionResponseBodyData) SetTitle(v string) *GetSessionResponseBodyData {
	s.Title = &v
	return s
}

func (s *GetSessionResponseBodyData) SetUpdatedAt(v string) *GetSessionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
