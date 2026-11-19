// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMemoryResponseBody
	GetCode() *string
	SetData(v *GetMemoryResponseBodyData) *GetMemoryResponseBody
	GetData() *GetMemoryResponseBodyData
	SetHttpStatusCode(v int32) *GetMemoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMemoryResponseBody
	GetSuccess() *bool
}

type GetMemoryResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The long-term memory details.
	Data *GetMemoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMemoryResponseBody) GetData() *GetMemoryResponseBodyData {
	return s.Data
}

func (s *GetMemoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMemoryResponseBody) SetCode(v string) *GetMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *GetMemoryResponseBody) SetData(v *GetMemoryResponseBodyData) *GetMemoryResponseBody {
	s.Data = v
	return s
}

func (s *GetMemoryResponseBody) SetHttpStatusCode(v int32) *GetMemoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMemoryResponseBody) SetMessage(v string) *GetMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *GetMemoryResponseBody) SetRequestId(v string) *GetMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMemoryResponseBody) SetSuccess(v bool) *GetMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *GetMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryResponseBodyData struct {
	// The memory content.
	Content *GetMemoryResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// The creation time. The time is in RFC 3339 format. The Z at the end of the example indicates the UTC time zone.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The long-term memory ID.
	//
	// example:
	//
	// mem-2f5f9c5a8d0e4c3b
	MemoryId *string `json:"memoryId,omitempty" xml:"memoryId,omitempty"`
	// The memory metadata.
	//
	// example:
	//
	// {"topic":"hiking"}
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The memory scope.
	Scope *GetMemoryResponseBodyDataScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The memory source. Valid values:
	//
	// - session_extraction: extracted from session events
	//
	// - content_extraction: extracted from text or messages
	//
	// - direct: written directly
	//
	// example:
	//
	// session_extraction
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The memory content type. Valid values:
	//
	// - semantic: semantic memory
	//
	// - skill: skill
	//
	// - user_profile: user preference
	//
	// - session_summary: session summary
	//
	// example:
	//
	// semantic
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The last updated time. The time is in RFC 3339 format. The Z at the end of the example indicates the UTC time zone.
	//
	// example:
	//
	// 2026-09-23T10:05:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// Indicates whether the memory is valid. A value of false indicates that the memory is revoked and is not included in valid memory retrieval.
	Valid *bool `json:"valid,omitempty" xml:"valid,omitempty"`
}

func (s GetMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBodyData) GetContent() *GetMemoryResponseBodyDataContent {
	return s.Content
}

func (s *GetMemoryResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetMemoryResponseBodyData) GetMemoryId() *string {
	return s.MemoryId
}

func (s *GetMemoryResponseBodyData) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *GetMemoryResponseBodyData) GetScope() *GetMemoryResponseBodyDataScope {
	return s.Scope
}

func (s *GetMemoryResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *GetMemoryResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetMemoryResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetMemoryResponseBodyData) GetValid() *bool {
	return s.Valid
}

func (s *GetMemoryResponseBodyData) SetContent(v *GetMemoryResponseBodyDataContent) *GetMemoryResponseBodyData {
	s.Content = v
	return s
}

func (s *GetMemoryResponseBodyData) SetCreatedAt(v string) *GetMemoryResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetMemoryId(v string) *GetMemoryResponseBodyData {
	s.MemoryId = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetMetadata(v map[string]*string) *GetMemoryResponseBodyData {
	s.Metadata = v
	return s
}

func (s *GetMemoryResponseBodyData) SetScope(v *GetMemoryResponseBodyDataScope) *GetMemoryResponseBodyData {
	s.Scope = v
	return s
}

func (s *GetMemoryResponseBodyData) SetSource(v string) *GetMemoryResponseBodyData {
	s.Source = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetType(v string) *GetMemoryResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetUpdatedAt(v string) *GetMemoryResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetMemoryResponseBodyData) SetValid(v bool) *GetMemoryResponseBodyData {
	s.Valid = &v
	return s
}

func (s *GetMemoryResponseBodyData) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMemoryResponseBodyDataContent struct {
	// The memory text.
	//
	// This parameter is required.
	//
	// example:
	//
	// The user likes hiking
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s GetMemoryResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBodyDataContent) GetText() *string {
	return s.Text
}

func (s *GetMemoryResponseBodyDataContent) SetText(v string) *GetMemoryResponseBodyDataContent {
	s.Text = &v
	return s
}

func (s *GetMemoryResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}

type GetMemoryResponseBodyDataScope struct {
	// The agent ID. The length is limited to 1 to 128 UTF-8 bytes. The value cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with an agent ID.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The session ID. The length is limited to 1 to 128 UTF-8 bytes. The value cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with a session ID.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The business team ID defined by the caller. The value cannot be empty when provided and can be up to 128 UTF-8 bytes in length.
	//
	// example:
	//
	// team-001
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The user ID. The length is limited to 1 to 128 UTF-8 bytes. The value cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with a user ID.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GetMemoryResponseBodyDataScope) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryResponseBodyDataScope) GoString() string {
	return s.String()
}

func (s *GetMemoryResponseBodyDataScope) GetAgentId() *string {
	return s.AgentId
}

func (s *GetMemoryResponseBodyDataScope) GetSessionId() *string {
	return s.SessionId
}

func (s *GetMemoryResponseBodyDataScope) GetTeamId() *string {
	return s.TeamId
}

func (s *GetMemoryResponseBodyDataScope) GetUserId() *string {
	return s.UserId
}

func (s *GetMemoryResponseBodyDataScope) SetAgentId(v string) *GetMemoryResponseBodyDataScope {
	s.AgentId = &v
	return s
}

func (s *GetMemoryResponseBodyDataScope) SetSessionId(v string) *GetMemoryResponseBodyDataScope {
	s.SessionId = &v
	return s
}

func (s *GetMemoryResponseBodyDataScope) SetTeamId(v string) *GetMemoryResponseBodyDataScope {
	s.TeamId = &v
	return s
}

func (s *GetMemoryResponseBodyDataScope) SetUserId(v string) *GetMemoryResponseBodyDataScope {
	s.UserId = &v
	return s
}

func (s *GetMemoryResponseBodyDataScope) Validate() error {
	return dara.Validate(s)
}
