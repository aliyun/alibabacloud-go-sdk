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
	// The write memory request.
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
	// Specifies whether to extract memories. If omitted or set to true, memories are extracted from text or messages. If set to false, memories are directly written.
	Extract *bool `json:"extract,omitempty" xml:"extract,omitempty"`
	// The list of memories to be directly written when extract is set to false. The list can contain 1 to 100 memories. You cannot provide text, messages, referenceTime, or request-level metadata at the same time.
	Memories []*AddMemoriesRequestBodyMemories `json:"memories,omitempty" xml:"memories,omitempty" type:"Repeated"`
	// The list of messages. You must provide either text or messages, but not both. The list can contain 1 to 20 messages, and the total length of the content cannot exceed 32,000 UTF-8 bytes.
	Messages []*AddMemoriesRequestBodyMessages `json:"messages,omitempty" xml:"messages,omitempty" type:"Repeated"`
	// The string key-value metadata. A maximum of 16 keys are supported. A key can be up to 64 characters in length, and a value can be up to 1,024 characters in length.
	//
	// example:
	//
	// {"topic":"hiking"}
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The time reference point for the batch of messages or text. It uses the RFC 3339 format with a time zone and supports fractional seconds.
	//
	// example:
	//
	// 2023-05-08T13:56:00+08:00
	ReferenceTime *string `json:"referenceTime,omitempty" xml:"referenceTime,omitempty"`
	// The write scope. UserId, AgentId, and SessionId are all optional.
	Scope *AddMemoriesRequestBodyScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// Specifies whether to synchronously wait for the memory extraction to complete. Default value: true.
	Sync *bool `json:"sync,omitempty" xml:"sync,omitempty"`
	// The text to be directly written. You must provide either text or messages, but not both. The text cannot exceed 32,000 UTF-8 bytes.
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

func (s *AddMemoriesRequestBody) GetExtract() *bool {
	return s.Extract
}

func (s *AddMemoriesRequestBody) GetMemories() []*AddMemoriesRequestBodyMemories {
	return s.Memories
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

func (s *AddMemoriesRequestBody) SetExtract(v bool) *AddMemoriesRequestBody {
	s.Extract = &v
	return s
}

func (s *AddMemoriesRequestBody) SetMemories(v []*AddMemoriesRequestBodyMemories) *AddMemoriesRequestBody {
	s.Memories = v
	return s
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
	if s.Memories != nil {
		for _, item := range s.Memories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type AddMemoriesRequestBodyMemories struct {
	// The string key-value tags set by the caller. A maximum of 16 tags are supported. A key can be up to 64 Unicode characters in length, and a value can be up to 1,024 Unicode characters in length. Empty strings are allowed. Nested objects or arrays are not supported.
	//
	// example:
	//
	// {"topic":"hiking"}
	Metadata map[string]*string `json:"metadata,omitempty" xml:"metadata,omitempty"`
	// The content of the memory to be directly written. It cannot be empty and can be up to 32,000 UTF-8 bytes in length.
	//
	// example:
	//
	// The user likes hiking
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
	// The content type of the memory. Valid values: semantic for semantic memory, skill for skill, user_profile for user preference, and session_summary for session summary.
	//
	// example:
	//
	// semantic
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AddMemoriesRequestBodyMemories) String() string {
	return dara.Prettify(s)
}

func (s AddMemoriesRequestBodyMemories) GoString() string {
	return s.String()
}

func (s *AddMemoriesRequestBodyMemories) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *AddMemoriesRequestBodyMemories) GetText() *string {
	return s.Text
}

func (s *AddMemoriesRequestBodyMemories) GetType() *string {
	return s.Type
}

func (s *AddMemoriesRequestBodyMemories) SetMetadata(v map[string]*string) *AddMemoriesRequestBodyMemories {
	s.Metadata = v
	return s
}

func (s *AddMemoriesRequestBodyMemories) SetText(v string) *AddMemoriesRequestBodyMemories {
	s.Text = &v
	return s
}

func (s *AddMemoriesRequestBodyMemories) SetType(v string) *AddMemoriesRequestBodyMemories {
	s.Type = &v
	return s
}

func (s *AddMemoriesRequestBodyMemories) Validate() error {
	return dara.Validate(s)
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
	// The reference time of the message in RFC 3339 format with a time zone. If omitted, the request-level referenceTime is used. If both are omitted, the time of the current request is used.
	//
	// example:
	//
	// 2023-05-08T13:56:00+08:00
	ReferenceTime *string `json:"referenceTime,omitempty" xml:"referenceTime,omitempty"`
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

func (s *AddMemoriesRequestBodyMessages) GetReferenceTime() *string {
	return s.ReferenceTime
}

func (s *AddMemoriesRequestBodyMessages) GetRole() *string {
	return s.Role
}

func (s *AddMemoriesRequestBodyMessages) SetContent(v string) *AddMemoriesRequestBodyMessages {
	s.Content = &v
	return s
}

func (s *AddMemoriesRequestBodyMessages) SetReferenceTime(v string) *AddMemoriesRequestBodyMessages {
	s.ReferenceTime = &v
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
	// The agent ID. The length is limited to 1 to 128 UTF-8 bytes. It cannot contain a forward slash (/) or an asterisk (*).
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The session ID. The length is limited to 1 to 128 UTF-8 bytes. It cannot contain a forward slash (/) or an asterisk (*).
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
	// The user ID. The length is limited to 1 to 128 UTF-8 bytes. It cannot contain a forward slash (/) or an asterisk (*).
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

func (s *AddMemoriesRequestBodyScope) GetTeamId() *string {
	return s.TeamId
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

func (s *AddMemoriesRequestBodyScope) SetTeamId(v string) *AddMemoriesRequestBodyScope {
	s.TeamId = &v
	return s
}

func (s *AddMemoriesRequestBodyScope) SetUserId(v string) *AddMemoriesRequestBodyScope {
	s.UserId = &v
	return s
}

func (s *AddMemoriesRequestBodyScope) Validate() error {
	return dara.Validate(s)
}
