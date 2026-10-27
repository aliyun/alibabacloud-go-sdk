// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMemoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SearchMemoriesResponseBody
	GetCode() *string
	SetData(v *SearchMemoriesResponseBodyData) *SearchMemoriesResponseBody
	GetData() *SearchMemoriesResponseBodyData
	SetHttpStatusCode(v int32) *SearchMemoriesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SearchMemoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *SearchMemoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SearchMemoriesResponseBody
	GetSuccess() *bool
}

type SearchMemoriesResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The retrieval results.
	Data *SearchMemoriesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s SearchMemoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *SearchMemoriesResponseBody) GetData() *SearchMemoriesResponseBodyData {
	return s.Data
}

func (s *SearchMemoriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SearchMemoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SearchMemoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchMemoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SearchMemoriesResponseBody) SetCode(v string) *SearchMemoriesResponseBody {
	s.Code = &v
	return s
}

func (s *SearchMemoriesResponseBody) SetData(v *SearchMemoriesResponseBodyData) *SearchMemoriesResponseBody {
	s.Data = v
	return s
}

func (s *SearchMemoriesResponseBody) SetHttpStatusCode(v int32) *SearchMemoriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SearchMemoriesResponseBody) SetMessage(v string) *SearchMemoriesResponseBody {
	s.Message = &v
	return s
}

func (s *SearchMemoriesResponseBody) SetRequestId(v string) *SearchMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchMemoriesResponseBody) SetSuccess(v bool) *SearchMemoriesResponseBody {
	s.Success = &v
	return s
}

func (s *SearchMemoriesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMemoriesResponseBodyData struct {
	// The list of retrieval results.
	Memories []*SearchMemoriesResponseBodyDataMemories `json:"memories,omitempty" xml:"memories,omitempty" type:"Repeated"`
}

func (s SearchMemoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBodyData) GetMemories() []*SearchMemoriesResponseBodyDataMemories {
	return s.Memories
}

func (s *SearchMemoriesResponseBodyData) SetMemories(v []*SearchMemoriesResponseBodyDataMemories) *SearchMemoriesResponseBodyData {
	s.Memories = v
	return s
}

func (s *SearchMemoriesResponseBodyData) Validate() error {
	if s.Memories != nil {
		for _, item := range s.Memories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchMemoriesResponseBodyDataMemories struct {
	// The matched memory.
	Memory *SearchMemoriesResponseBodyDataMemoriesMemory `json:"memory,omitempty" xml:"memory,omitempty" type:"Struct"`
	// The relevance score.
	//
	// example:
	//
	// 0.86
	Score *float32 `json:"score,omitempty" xml:"score,omitempty"`
	// The vector similarity.
	//
	// example:
	//
	// 0.91
	Similarity *float32 `json:"similarity,omitempty" xml:"similarity,omitempty"`
}

func (s SearchMemoriesResponseBodyDataMemories) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBodyDataMemories) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBodyDataMemories) GetMemory() *SearchMemoriesResponseBodyDataMemoriesMemory {
	return s.Memory
}

func (s *SearchMemoriesResponseBodyDataMemories) GetScore() *float32 {
	return s.Score
}

func (s *SearchMemoriesResponseBodyDataMemories) GetSimilarity() *float32 {
	return s.Similarity
}

func (s *SearchMemoriesResponseBodyDataMemories) SetMemory(v *SearchMemoriesResponseBodyDataMemoriesMemory) *SearchMemoriesResponseBodyDataMemories {
	s.Memory = v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemories) SetScore(v float32) *SearchMemoriesResponseBodyDataMemories {
	s.Score = &v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemories) SetSimilarity(v float32) *SearchMemoriesResponseBodyDataMemories {
	s.Similarity = &v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemories) Validate() error {
	if s.Memory != nil {
		if err := s.Memory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchMemoriesResponseBodyDataMemoriesMemory struct {
	// The memory content.
	Content *SearchMemoriesResponseBodyDataMemoriesMemoryContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
	// The time when the memory was created.
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
	Scope *SearchMemoriesResponseBodyDataMemoriesMemoryScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The time when the memory was last updated.
	//
	// example:
	//
	// 2026-09-23T10:05:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s SearchMemoriesResponseBodyDataMemoriesMemory) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBodyDataMemoriesMemory) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) GetContent() *SearchMemoriesResponseBodyDataMemoriesMemoryContent {
	return s.Content
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) GetMemoryId() *string {
	return s.MemoryId
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) GetScope() *SearchMemoriesResponseBodyDataMemoriesMemoryScope {
	return s.Scope
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) SetContent(v *SearchMemoriesResponseBodyDataMemoriesMemoryContent) *SearchMemoriesResponseBodyDataMemoriesMemory {
	s.Content = v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) SetCreatedAt(v string) *SearchMemoriesResponseBodyDataMemoriesMemory {
	s.CreatedAt = &v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) SetMemoryId(v string) *SearchMemoriesResponseBodyDataMemoriesMemory {
	s.MemoryId = &v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) SetMetadata(v map[string]*string) *SearchMemoriesResponseBodyDataMemoriesMemory {
	s.Metadata = v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) SetScope(v *SearchMemoriesResponseBodyDataMemoriesMemoryScope) *SearchMemoriesResponseBodyDataMemoriesMemory {
	s.Scope = v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) SetUpdatedAt(v string) *SearchMemoriesResponseBodyDataMemoriesMemory {
	s.UpdatedAt = &v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemory) Validate() error {
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

type SearchMemoriesResponseBodyDataMemoriesMemoryContent struct {
	// The memory text.
	//
	// example:
	//
	// The user likes hiking
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s SearchMemoriesResponseBodyDataMemoriesMemoryContent) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBodyDataMemoriesMemoryContent) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryContent) GetText() *string {
	return s.Text
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryContent) SetText(v string) *SearchMemoriesResponseBodyDataMemoriesMemoryContent {
	s.Text = &v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryContent) Validate() error {
	return dara.Validate(s)
}

type SearchMemoriesResponseBodyDataMemoriesMemoryScope struct {
	// The agent ID. The ID must be 1 to 128 bytes in length and encoded in UTF-8. The ID cannot contain a forward slash (/) or an asterisk (*). If the memory is not associated with an agent ID, this field is not returned.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The session ID. The ID must be 1 to 128 bytes in length and encoded in UTF-8. The ID cannot contain a forward slash (/) or an asterisk (*). If the memory is not associated with a session ID, this field is not returned.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The user ID. The ID must be 1 to 128 bytes in length and encoded in UTF-8. The ID cannot contain a forward slash (/) or an asterisk (*). If the memory is not associated with a user ID, this field is not returned.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s SearchMemoriesResponseBodyDataMemoriesMemoryScope) String() string {
	return dara.Prettify(s)
}

func (s SearchMemoriesResponseBodyDataMemoriesMemoryScope) GoString() string {
	return s.String()
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryScope) GetAgentId() *string {
	return s.AgentId
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryScope) GetSessionId() *string {
	return s.SessionId
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryScope) GetUserId() *string {
	return s.UserId
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryScope) SetAgentId(v string) *SearchMemoriesResponseBodyDataMemoriesMemoryScope {
	s.AgentId = &v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryScope) SetSessionId(v string) *SearchMemoriesResponseBodyDataMemoriesMemoryScope {
	s.SessionId = &v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryScope) SetUserId(v string) *SearchMemoriesResponseBodyDataMemoriesMemoryScope {
	s.UserId = &v
	return s
}

func (s *SearchMemoriesResponseBodyDataMemoriesMemoryScope) Validate() error {
	return dara.Validate(s)
}
