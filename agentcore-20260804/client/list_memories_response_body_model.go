// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemoriesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMemoriesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMemoriesResponseBodyItems) *ListMemoriesResponseBody
	GetItems() []*ListMemoriesResponseBodyItems
	SetMaxResults(v int32) *ListMemoriesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMemoriesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMemoriesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMemoriesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMemoriesResponseBody
	GetSuccess() *bool
}

type ListMemoriesResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The memory list.
	Items []*ListMemoriesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjEwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01A00000-0000-0000-0000-000000000001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListMemoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemoriesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMemoriesResponseBody) GetItems() []*ListMemoriesResponseBodyItems {
	return s.Items
}

func (s *ListMemoriesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMemoriesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoriesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMemoriesResponseBody) SetCode(v string) *ListMemoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemoriesResponseBody) SetHttpStatusCode(v int32) *ListMemoriesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMemoriesResponseBody) SetItems(v []*ListMemoriesResponseBodyItems) *ListMemoriesResponseBody {
	s.Items = v
	return s
}

func (s *ListMemoriesResponseBody) SetMaxResults(v int32) *ListMemoriesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemoriesResponseBody) SetMessage(v string) *ListMemoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemoriesResponseBody) SetNextToken(v string) *ListMemoriesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemoriesResponseBody) SetRequestId(v string) *ListMemoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoriesResponseBody) SetSuccess(v bool) *ListMemoriesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMemoriesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMemoriesResponseBodyItems struct {
	// The memory content.
	Content *ListMemoriesResponseBodyItemsContent `json:"content,omitempty" xml:"content,omitempty" type:"Struct"`
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
	Scope *ListMemoriesResponseBodyItemsScope `json:"scope,omitempty" xml:"scope,omitempty" type:"Struct"`
	// The last update time.
	//
	// example:
	//
	// 2026-09-23T10:05:00Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
}

func (s ListMemoriesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMemoriesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponseBodyItems) GetContent() *ListMemoriesResponseBodyItemsContent {
	return s.Content
}

func (s *ListMemoriesResponseBodyItems) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListMemoriesResponseBodyItems) GetMemoryId() *string {
	return s.MemoryId
}

func (s *ListMemoriesResponseBodyItems) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *ListMemoriesResponseBodyItems) GetScope() *ListMemoriesResponseBodyItemsScope {
	return s.Scope
}

func (s *ListMemoriesResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListMemoriesResponseBodyItems) SetContent(v *ListMemoriesResponseBodyItemsContent) *ListMemoriesResponseBodyItems {
	s.Content = v
	return s
}

func (s *ListMemoriesResponseBodyItems) SetCreatedAt(v string) *ListMemoriesResponseBodyItems {
	s.CreatedAt = &v
	return s
}

func (s *ListMemoriesResponseBodyItems) SetMemoryId(v string) *ListMemoriesResponseBodyItems {
	s.MemoryId = &v
	return s
}

func (s *ListMemoriesResponseBodyItems) SetMetadata(v map[string]*string) *ListMemoriesResponseBodyItems {
	s.Metadata = v
	return s
}

func (s *ListMemoriesResponseBodyItems) SetScope(v *ListMemoriesResponseBodyItemsScope) *ListMemoriesResponseBodyItems {
	s.Scope = v
	return s
}

func (s *ListMemoriesResponseBodyItems) SetUpdatedAt(v string) *ListMemoriesResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *ListMemoriesResponseBodyItems) Validate() error {
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

type ListMemoriesResponseBodyItemsContent struct {
	// The memory text.
	//
	// example:
	//
	// The user likes hiking
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ListMemoriesResponseBodyItemsContent) String() string {
	return dara.Prettify(s)
}

func (s ListMemoriesResponseBodyItemsContent) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponseBodyItemsContent) GetText() *string {
	return s.Text
}

func (s *ListMemoriesResponseBodyItemsContent) SetText(v string) *ListMemoriesResponseBodyItemsContent {
	s.Text = &v
	return s
}

func (s *ListMemoriesResponseBodyItemsContent) Validate() error {
	return dara.Validate(s)
}

type ListMemoriesResponseBodyItemsScope struct {
	// The agent ID. The value must be 1 to 128 bytes in length and encoded in UTF-8. The value cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with an agent ID.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// The session ID. The value must be 1 to 128 bytes in length and encoded in UTF-8. The value cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with a session ID.
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The user ID. The value must be 1 to 128 bytes in length and encoded in UTF-8. The value cannot contain a forward slash (/) or an asterisk (*). This field is not returned if the memory is not associated with a user ID.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListMemoriesResponseBodyItemsScope) String() string {
	return dara.Prettify(s)
}

func (s ListMemoriesResponseBodyItemsScope) GoString() string {
	return s.String()
}

func (s *ListMemoriesResponseBodyItemsScope) GetAgentId() *string {
	return s.AgentId
}

func (s *ListMemoriesResponseBodyItemsScope) GetSessionId() *string {
	return s.SessionId
}

func (s *ListMemoriesResponseBodyItemsScope) GetUserId() *string {
	return s.UserId
}

func (s *ListMemoriesResponseBodyItemsScope) SetAgentId(v string) *ListMemoriesResponseBodyItemsScope {
	s.AgentId = &v
	return s
}

func (s *ListMemoriesResponseBodyItemsScope) SetSessionId(v string) *ListMemoriesResponseBodyItemsScope {
	s.SessionId = &v
	return s
}

func (s *ListMemoriesResponseBodyItemsScope) SetUserId(v string) *ListMemoriesResponseBodyItemsScope {
	s.UserId = &v
	return s
}

func (s *ListMemoriesResponseBodyItemsScope) Validate() error {
	return dara.Validate(s)
}
