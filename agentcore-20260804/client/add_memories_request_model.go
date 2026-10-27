// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMemoriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *AddMemoriesRequestBody) *AddMemoriesRequest
	GetBody() *AddMemoriesRequestBody
}

type AddMemoriesRequest struct {
	// The request body for writing to memory.
	//
	// This parameter is required.
	Body *AddMemoriesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s AddMemoriesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesRequest) GoString() string {
	return s.String()
}

func (s *AddMemoriesRequest) GetBody() *AddMemoriesRequestBody {
	return s.Body
}

func (s *AddMemoriesRequest) SetBody(v *AddMemoriesRequestBody) *AddMemoriesRequest {
	s.Body = v
	return s
}

func (s *AddMemoriesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMemoriesRequestBody struct {
	// The list of messages. You must provide either this parameter or Text, but not both. The list must contain 1 to 20 messages, and the total size of the content must not exceed 32,000 UTF-8 bytes.
	Messages []*AddMemoriesRequestBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// The string key-value metadata. You can specify up to 16 keys. The maximum length of a key is 64 characters, and the maximum length of a value is 1,024 characters.
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The time reference point for the batch of messages or text. The value must be in RFC 3339 format with a time zone and support fractional seconds.
	//
	// example:
	//
	// 2023-05-08T13:56:00+08:00
	ReferenceTime *string `json:"referenceTime,omitempty" xml:"referenceTime,omitempty"`
	// The scope of the write operation. UserId, AgentId, and SessionId are all optional.
	Scope *AddMemoriesRequestBodyScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// Specifies whether to synchronously wait for memory extraction to complete. Default value: true.
	Sync *bool `json:"sync,omitempty" xml:"sync,omitempty"`
	// The text to be written directly. You must provide either this parameter or Messages, but not both. The size of the text must not exceed 32,000 UTF-8 bytes.
	//
	// example:
	//
	// The user likes hiking
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s AddMemoriesRequestBody) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesRequestBody) GoString() string {
	return s.String()
}

func (s *AddMemoriesRequestBody) GetMessages() []*AddMemoriesRequestBodyMessages {
	return s.Messages
}

func (s *AddMemoriesRequestBody) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *AddMemoriesRequestBody) GetReferenceTime() *string {
	return s.ReferenceTime
}

func (s *AddMemoriesRequestBody) GetScope() *AddMemoriesRequestBodyScope {
	return s.Scope
}

func (s *AddMemoriesRequestBody) GetSync() *bool {
	return s.Sync
}

func (s *AddMemoriesRequestBody) GetText() *string {
	return s.Text
}

func (s *AddMemoriesRequestBody) SetMessages(v []*AddMemoriesRequestBodyMessages) *AddMemoriesRequestBody {
	s.Messages = v
	return s
}

func (s *AddMemoriesRequestBody) SetMetadata(v map[string]*string) *AddMemoriesRequestBody {
	s.Metadata = v
	return s
}

func (s *AddMemoriesRequestBody) SetReferenceTime(v string) *AddMemoriesRequestBody {
	s.ReferenceTime = &v
	return s
}

func (s *AddMemoriesRequestBody) SetScope(v *AddMemoriesRequestBodyScope) *AddMemoriesRequestBody {
	s.Scope = v
	return s
}

func (s *AddMemoriesRequestBody) SetSync(v bool) *AddMemoriesRequestBody {
	s.Sync = &v
	return s
}

func (s *AddMemoriesRequestBody) SetText(v string) *AddMemoriesRequestBody {
	s.Text = &v
	return s
}

func (s *AddMemoriesRequestBody) Validate() error {
	if s.Messages != nil {
		for _, item := range s.Messages {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMemoriesRequestBodyMessages struct {
	// The content of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// The user likes hiking
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The role of the message.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s AddMemoriesRequestBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesRequestBodyMessages) GoString() string {
	return s.String()
}

func (s *AddMemoriesRequestBodyMessages) GetContent() *string {
	return s.Content
}

func (s *AddMemoriesRequestBodyMessages) GetRole() *string {
	return s.Role
}

func (s *AddMemoriesRequestBodyMessages) SetContent(v string) *AddMemoriesRequestBodyMessages {
	s.Content = &v
	return s
}

func (s *AddMemoriesRequestBodyMessages) SetRole(v string) *AddMemoriesRequestBodyMessages {
	s.Role = &v
	return s
}

func (s *AddMemoriesRequestBodyMessages) Validate() error {
	return dara.Validate(s)
}

type AddMemoriesRequestBodyScope struct {
	// The ID of the agent. The length is limited to 1 to 128 UTF-8 bytes. The ID cannot contain a forward slash (/) or an asterisk (*).
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The ID of the session. The length is limited to 1 to 128 UTF-8 bytes. The ID cannot contain a forward slash (/) or an asterisk (*).
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The ID of the user. The length is limited to 1 to 128 UTF-8 bytes. The ID cannot contain a forward slash (/) or an asterisk (*).
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddMemoriesRequestBodyScope) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesRequestBodyScope) GoString() string {
	return s.String()
}

func (s *AddMemoriesRequestBodyScope) GetAgentId() *string {
	return s.AgentId
}

func (s *AddMemoriesRequestBodyScope) GetSessionId() *string {
	return s.SessionId
}

func (s *AddMemoriesRequestBodyScope) GetUserId() *string {
	return s.UserId
}

func (s *AddMemoriesRequestBodyScope) SetAgentId(v string) *AddMemoriesRequestBodyScope {
	s.AgentId = &v
	return s
}

func (s *AddMemoriesRequestBodyScope) SetSessionId(v string) *AddMemoriesRequestBodyScope {
	s.SessionId = &v
	return s
}

func (s *AddMemoriesRequestBodyScope) SetUserId(v string) *AddMemoriesRequestBodyScope {
	s.UserId = &v
	return s
}

func (s *AddMemoriesRequestBodyScope) Validate() error {
	return dara.Validate(s)
}
