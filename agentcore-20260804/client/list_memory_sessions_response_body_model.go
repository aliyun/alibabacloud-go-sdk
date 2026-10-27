// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemorySessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemorySessionsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMemorySessionsResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMemorySessionsResponseBodyItems) *ListMemorySessionsResponseBody
	GetItems() []*ListMemorySessionsResponseBodyItems
	SetMaxResults(v int32) *ListMemorySessionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMemorySessionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMemorySessionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMemorySessionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMemorySessionsResponseBody
	GetSuccess() *bool
}

type ListMemorySessionsResponseBody struct {
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
	// The list of sessions.
	Items []*ListMemorySessionsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page.
	//
	// example:
	//
	// 100
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
	// eyJvZmZzZXQiOjEwMH0=
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

func (s ListMemorySessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemorySessionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemorySessionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMemorySessionsResponseBody) GetItems() []*ListMemorySessionsResponseBodyItems {
	return s.Items
}

func (s *ListMemorySessionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemorySessionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMemorySessionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemorySessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemorySessionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMemorySessionsResponseBody) SetCode(v string) *ListMemorySessionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemorySessionsResponseBody) SetHttpStatusCode(v int32) *ListMemorySessionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMemorySessionsResponseBody) SetItems(v []*ListMemorySessionsResponseBodyItems) *ListMemorySessionsResponseBody {
	s.Items = v
	return s
}

func (s *ListMemorySessionsResponseBody) SetMaxResults(v int32) *ListMemorySessionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemorySessionsResponseBody) SetMessage(v string) *ListMemorySessionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemorySessionsResponseBody) SetNextToken(v string) *ListMemorySessionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemorySessionsResponseBody) SetRequestId(v string) *ListMemorySessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemorySessionsResponseBody) SetSuccess(v bool) *ListMemorySessionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMemorySessionsResponseBody) Validate() error {
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

type ListMemorySessionsResponseBodyItems struct {
	// The agent ID. This field is not returned if the session is not associated with an agent ID.
	//
	// example:
	//
	// agent-001
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// Session ID
	//
	// example:
	//
	// session-001
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The user ID. This field is not returned if the session is not associated with a user ID.
	//
	// example:
	//
	// user-001
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s ListMemorySessionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMemorySessionsResponseBodyItems) GetAgentId() *string {
	return s.AgentId
}

func (s *ListMemorySessionsResponseBodyItems) GetSessionId() *string {
	return s.SessionId
}

func (s *ListMemorySessionsResponseBodyItems) GetUserId() *string {
	return s.UserId
}

func (s *ListMemorySessionsResponseBodyItems) SetAgentId(v string) *ListMemorySessionsResponseBodyItems {
	s.AgentId = &v
	return s
}

func (s *ListMemorySessionsResponseBodyItems) SetSessionId(v string) *ListMemorySessionsResponseBodyItems {
	s.SessionId = &v
	return s
}

func (s *ListMemorySessionsResponseBodyItems) SetUserId(v string) *ListMemorySessionsResponseBodyItems {
	s.UserId = &v
	return s
}

func (s *ListMemorySessionsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
