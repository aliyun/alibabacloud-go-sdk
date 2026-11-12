// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateSessionResponseBodyData) *UpdateSessionResponseBody
	GetData() *UpdateSessionResponseBodyData
	SetRequestId(v string) *UpdateSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateSessionResponseBody
	GetSuccess() *bool
}

type UpdateSessionResponseBody struct {
	// The session information.
	Data *UpdateSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. You can use the ID to locate and troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful. A value of true indicates success.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSessionResponseBody) GetData() *UpdateSessionResponseBodyData {
	return s.Data
}

func (s *UpdateSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateSessionResponseBody) SetData(v *UpdateSessionResponseBodyData) *UpdateSessionResponseBody {
	s.Data = v
	return s
}

func (s *UpdateSessionResponseBody) SetRequestId(v string) *UpdateSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSessionResponseBody) SetSuccess(v bool) *UpdateSessionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateSessionResponseBodyData struct {
	// The time when the session was created. The time is in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The ID of the last run.
	//
	// example:
	//
	// run_example_001
	LastRunId *string `json:"lastRunId,omitempty" xml:"lastRunId,omitempty"`
	// The status of the last run. Valid values:
	//
	// - `PENDING`: pending
	//
	// - `RUNNING`: running
	//
	// - `WAITING_USER`: waiting for interrupt handling
	//
	// - `COMPLETED`: completed
	//
	// - `FAILED`: failed
	//
	// - `CANCELLED`: cancelled
	//
	// - `REJECTED`: rejected
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
	// The ID of the session.
	//
	// example:
	//
	// session_example_001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The status of the session. Valid values:
	//
	// - `IDLE`: idle
	//
	// - `ACTIVE`: running
	//
	// - `WAITING_USER_INPUT`: waiting for user input
	//
	// - `WAITING_APPROVAL`: waiting for user confirmation
	//
	// - `WAITING_OUT_OF_BAND_APPROVAL`: waiting for out-of-band approval
	//
	// - `CLOSED`: closed
	//
	// example:
	//
	// IDLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The title of the session.
	//
	// example:
	//
	// Cloud resource inspection example
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The time when the session was last updated. The time is in RFC 3339 format.
	//
	// example:
	//
	// 2026-01-01T00:01:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s UpdateSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateSessionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateSessionResponseBodyData) GetLastRunId() *string {
	return s.LastRunId
}

func (s *UpdateSessionResponseBodyData) GetLastRunStatus() *string {
	return s.LastRunStatus
}

func (s *UpdateSessionResponseBodyData) GetPinned() *bool {
	return s.Pinned
}

func (s *UpdateSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateSessionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateSessionResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *UpdateSessionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateSessionResponseBodyData) SetCreatedAt(v string) *UpdateSessionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateSessionResponseBodyData) SetLastRunId(v string) *UpdateSessionResponseBodyData {
	s.LastRunId = &v
	return s
}

func (s *UpdateSessionResponseBodyData) SetLastRunStatus(v string) *UpdateSessionResponseBodyData {
	s.LastRunStatus = &v
	return s
}

func (s *UpdateSessionResponseBodyData) SetPinned(v bool) *UpdateSessionResponseBodyData {
	s.Pinned = &v
	return s
}

func (s *UpdateSessionResponseBodyData) SetSessionId(v string) *UpdateSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *UpdateSessionResponseBodyData) SetStatus(v string) *UpdateSessionResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateSessionResponseBodyData) SetTitle(v string) *UpdateSessionResponseBodyData {
	s.Title = &v
	return s
}

func (s *UpdateSessionResponseBodyData) SetUpdatedAt(v string) *UpdateSessionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
