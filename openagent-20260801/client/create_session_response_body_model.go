// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateSessionResponseBodyData) *CreateSessionResponseBody
	GetData() *CreateSessionResponseBodyData
	SetRequestId(v string) *CreateSessionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSessionResponseBody
	GetSuccess() *bool
}

type CreateSessionResponseBody struct {
	// The session information of the created session.
	Data *CreateSessionResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s CreateSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionResponseBody) GetData() *CreateSessionResponseBodyData {
	return s.Data
}

func (s *CreateSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSessionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSessionResponseBody) SetData(v *CreateSessionResponseBodyData) *CreateSessionResponseBody {
	s.Data = v
	return s
}

func (s *CreateSessionResponseBody) SetRequestId(v string) *CreateSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSessionResponseBody) SetSuccess(v bool) *CreateSessionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSessionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSessionResponseBodyData struct {
	// The time when the session was created, in ISO 8601 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
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
	// The session status. The value is fixed to `IDLE` upon successful creation.
	//
	// example:
	//
	// IDLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the session was last updated, in ISO 8601 format.
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s CreateSessionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSessionResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateSessionResponseBodyData) GetPinned() *bool {
	return s.Pinned
}

func (s *CreateSessionResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateSessionResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateSessionResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateSessionResponseBodyData) SetCreatedAt(v string) *CreateSessionResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateSessionResponseBodyData) SetPinned(v bool) *CreateSessionResponseBodyData {
	s.Pinned = &v
	return s
}

func (s *CreateSessionResponseBodyData) SetSessionId(v string) *CreateSessionResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *CreateSessionResponseBodyData) SetStatus(v string) *CreateSessionResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateSessionResponseBodyData) SetUpdatedAt(v string) *CreateSessionResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateSessionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
