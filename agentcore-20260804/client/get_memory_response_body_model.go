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
	// The creation time.
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
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The memory scope.
	Scope *GetMemoryResponseBodyDataScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The last update time.
	//
	// example:
	//
	// 2026-09-23T10:05:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
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

func (s *GetMemoryResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
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

func (s *GetMemoryResponseBodyData) SetUpdatedAt(v string) *GetMemoryResponseBodyData {
	s.UpdatedAt = &v
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
	// The agent ID. The length is limited to 1 to 128 UTF-8 bytes. It cannot contain / or *. This field is not returned if the memory is not associated with an agent ID.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The session ID. The length is limited to 1 to 128 UTF-8 bytes. It cannot contain / or *. This field is not returned if the memory is not associated with a session ID.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The user ID. The length is limited to 1 to 128 UTF-8 bytes. It cannot contain / or *. This field is not returned if the memory is not associated with a user ID.
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

func (s *GetMemoryResponseBodyDataScope) SetUserId(v string) *GetMemoryResponseBodyDataScope {
	s.UserId = &v
	return s
}

func (s *GetMemoryResponseBodyDataScope) Validate() error {
	return dara.Validate(s)
}
