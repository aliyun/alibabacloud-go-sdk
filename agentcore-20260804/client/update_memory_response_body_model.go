// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateMemoryResponseBody
	GetCode() *string
	SetData(v *UpdateMemoryResponseBodyData) *UpdateMemoryResponseBody
	GetData() *UpdateMemoryResponseBodyData
	SetHttpStatusCode(v int32) *UpdateMemoryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateMemoryResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateMemoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateMemoryResponseBody
	GetSuccess() *bool
}

type UpdateMemoryResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The updated memory.
	Data *UpdateMemoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s UpdateMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateMemoryResponseBody) GetData() *UpdateMemoryResponseBodyData {
	return s.Data
}

func (s *UpdateMemoryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateMemoryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMemoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateMemoryResponseBody) SetCode(v string) *UpdateMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetData(v *UpdateMemoryResponseBodyData) *UpdateMemoryResponseBody {
	s.Data = v
	return s
}

func (s *UpdateMemoryResponseBody) SetHttpStatusCode(v int32) *UpdateMemoryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetMessage(v string) *UpdateMemoryResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetRequestId(v string) *UpdateMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMemoryResponseBody) SetSuccess(v bool) *UpdateMemoryResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateMemoryResponseBodyData struct {
	// The memory content.
	Content *UpdateMemoryResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// The creation time. The time is in RFC 3339 format. The Z at the end of the example indicates the UTC time zone.
	//
	// example:
	//
	// 2026-09-23T10:00:00Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// The ID of the long-term memory.
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
	Scope *UpdateMemoryResponseBodyDataScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The memory source. Valid values: session_extraction: Extracted from session events. content_extraction: Extracted from text or messages. direct: Written directly.
	//
	// example:
	//
	// session_extraction
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The memory content type. Valid values: semantic: Semantic memory. skill: Skill. user_profile: User preference. session_summary: Session summary.
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

func (s UpdateMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBodyData) GetContent() *UpdateMemoryResponseBodyDataContent {
	return s.Content
}

func (s *UpdateMemoryResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateMemoryResponseBodyData) GetMemoryId() *string {
	return s.MemoryId
}

func (s *UpdateMemoryResponseBodyData) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *UpdateMemoryResponseBodyData) GetScope() *UpdateMemoryResponseBodyDataScope {
	return s.Scope
}

func (s *UpdateMemoryResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *UpdateMemoryResponseBodyData) GetType() *string {
	return s.Type
}

func (s *UpdateMemoryResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateMemoryResponseBodyData) GetValid() *bool {
	return s.Valid
}

func (s *UpdateMemoryResponseBodyData) SetContent(v *UpdateMemoryResponseBodyDataContent) *UpdateMemoryResponseBodyData {
	s.Content = v
	return s
}

func (s *UpdateMemoryResponseBodyData) SetCreatedAt(v string) *UpdateMemoryResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateMemoryResponseBodyData) SetMemoryId(v string) *UpdateMemoryResponseBodyData {
	s.MemoryId = &v
	return s
}

func (s *UpdateMemoryResponseBodyData) SetMetadata(v map[string]*string) *UpdateMemoryResponseBodyData {
	s.Metadata = v
	return s
}

func (s *UpdateMemoryResponseBodyData) SetScope(v *UpdateMemoryResponseBodyDataScope) *UpdateMemoryResponseBodyData {
	s.Scope = v
	return s
}

func (s *UpdateMemoryResponseBodyData) SetSource(v string) *UpdateMemoryResponseBodyData {
	s.Source = &v
	return s
}

func (s *UpdateMemoryResponseBodyData) SetType(v string) *UpdateMemoryResponseBodyData {
	s.Type = &v
	return s
}

func (s *UpdateMemoryResponseBodyData) SetUpdatedAt(v string) *UpdateMemoryResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateMemoryResponseBodyData) SetValid(v bool) *UpdateMemoryResponseBodyData {
	s.Valid = &v
	return s
}

func (s *UpdateMemoryResponseBodyData) Validate() error {
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

type UpdateMemoryResponseBodyDataContent struct {
	// The memory text.
	//
	// This parameter is required.
	//
	// example:
	//
	// The user enjoys hiking
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s UpdateMemoryResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBodyDataContent) GetText() *string {
	return s.Text
}

func (s *UpdateMemoryResponseBodyDataContent) SetText(v string) *UpdateMemoryResponseBodyDataContent {
	s.Text = &v
	return s
}

func (s *UpdateMemoryResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}

type UpdateMemoryResponseBodyDataScope struct {
	// The agent ID. The length is limited to 1 to 128 UTF-8 bytes. The ID cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with an agent ID.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The session ID. The length is limited to 1 to 128 UTF-8 bytes. The ID cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with a session ID.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The business team ID defined by the caller. It cannot be empty when provided and can be up to 128 UTF-8 bytes in length.
	//
	// example:
	//
	// team-001
	TeamId *string `json:"teamId,omitempty" xml:"teamId,omitempty"`
	// The user ID. The length is limited to 1 to 128 UTF-8 bytes. The ID cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with a user ID.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateMemoryResponseBodyDataScope) String() string {
	return dara.Prettify(s)
}

func (s UpdateMemoryResponseBodyDataScope) GoString() string {
	return s.String()
}

func (s *UpdateMemoryResponseBodyDataScope) GetAgentId() *string {
	return s.AgentId
}

func (s *UpdateMemoryResponseBodyDataScope) GetSessionId() *string {
	return s.SessionId
}

func (s *UpdateMemoryResponseBodyDataScope) GetTeamId() *string {
	return s.TeamId
}

func (s *UpdateMemoryResponseBodyDataScope) GetUserId() *string {
	return s.UserId
}

func (s *UpdateMemoryResponseBodyDataScope) SetAgentId(v string) *UpdateMemoryResponseBodyDataScope {
	s.AgentId = &v
	return s
}

func (s *UpdateMemoryResponseBodyDataScope) SetSessionId(v string) *UpdateMemoryResponseBodyDataScope {
	s.SessionId = &v
	return s
}

func (s *UpdateMemoryResponseBodyDataScope) SetTeamId(v string) *UpdateMemoryResponseBodyDataScope {
	s.TeamId = &v
	return s
}

func (s *UpdateMemoryResponseBodyDataScope) SetUserId(v string) *UpdateMemoryResponseBodyDataScope {
	s.UserId = &v
	return s
}

func (s *UpdateMemoryResponseBodyDataScope) Validate() error {
	return dara.Validate(s)
}
