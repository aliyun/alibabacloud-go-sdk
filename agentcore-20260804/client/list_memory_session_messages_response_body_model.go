// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemorySessionMessagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemorySessionMessagesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListMemorySessionMessagesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListMemorySessionMessagesResponseBodyItems) *ListMemorySessionMessagesResponseBody
	GetItems() []*ListMemorySessionMessagesResponseBodyItems
	SetMaxResults(v int32) *ListMemorySessionMessagesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListMemorySessionMessagesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListMemorySessionMessagesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMemorySessionMessagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMemorySessionMessagesResponseBody
	GetSuccess() *bool
}

type ListMemorySessionMessagesResponseBody struct {
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
	// The message list.
	Items []*ListMemorySessionMessagesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of records per page.
	//
	// example:
	//
	// 200
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
	// eyJvZmZzZXQiOjIwMH0=
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

func (s ListMemorySessionMessagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionMessagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemorySessionMessagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemorySessionMessagesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMemorySessionMessagesResponseBody) GetItems() []*ListMemorySessionMessagesResponseBodyItems {
	return s.Items
}

func (s *ListMemorySessionMessagesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemorySessionMessagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMemorySessionMessagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemorySessionMessagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemorySessionMessagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMemorySessionMessagesResponseBody) SetCode(v string) *ListMemorySessionMessagesResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemorySessionMessagesResponseBody) SetHttpStatusCode(v int32) *ListMemorySessionMessagesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMemorySessionMessagesResponseBody) SetItems(v []*ListMemorySessionMessagesResponseBodyItems) *ListMemorySessionMessagesResponseBody {
	s.Items = v
	return s
}

func (s *ListMemorySessionMessagesResponseBody) SetMaxResults(v int32) *ListMemorySessionMessagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMemorySessionMessagesResponseBody) SetMessage(v string) *ListMemorySessionMessagesResponseBody {
	s.Message = &v
	return s
}

func (s *ListMemorySessionMessagesResponseBody) SetNextToken(v string) *ListMemorySessionMessagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMemorySessionMessagesResponseBody) SetRequestId(v string) *ListMemorySessionMessagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemorySessionMessagesResponseBody) SetSuccess(v bool) *ListMemorySessionMessagesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMemorySessionMessagesResponseBody) Validate() error {
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

type ListMemorySessionMessagesResponseBodyItems struct {
	// The message content.
	//
	// example:
	//
	// The user likes hiking
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The message role.
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
}

func (s ListMemorySessionMessagesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionMessagesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListMemorySessionMessagesResponseBodyItems) GetContent() *string {
	return s.Content
}

func (s *ListMemorySessionMessagesResponseBodyItems) GetRole() *string {
	return s.Role
}

func (s *ListMemorySessionMessagesResponseBodyItems) SetContent(v string) *ListMemorySessionMessagesResponseBodyItems {
	s.Content = &v
	return s
}

func (s *ListMemorySessionMessagesResponseBodyItems) SetRole(v string) *ListMemorySessionMessagesResponseBodyItems {
	s.Role = &v
	return s
}

func (s *ListMemorySessionMessagesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
