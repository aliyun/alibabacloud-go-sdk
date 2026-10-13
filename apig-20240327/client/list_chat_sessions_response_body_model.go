// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListChatSessionsResponseBody
	GetCode() *string
	SetData(v *ListChatSessionsResponseBodyData) *ListChatSessionsResponseBody
	GetData() *ListChatSessionsResponseBodyData
	SetMaxResults(v int32) *ListChatSessionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListChatSessionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListChatSessionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListChatSessionsResponseBody
	GetRequestId() *string
}

type ListChatSessionsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListChatSessionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The maximum number of entries on the current page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. An empty value indicates that the last page has been reached.
	//
	// example:
	//
	// token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListChatSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListChatSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListChatSessionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListChatSessionsResponseBody) GetData() *ListChatSessionsResponseBodyData {
	return s.Data
}

func (s *ListChatSessionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListChatSessionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListChatSessionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListChatSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListChatSessionsResponseBody) SetCode(v string) *ListChatSessionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetData(v *ListChatSessionsResponseBodyData) *ListChatSessionsResponseBody {
	s.Data = v
	return s
}

func (s *ListChatSessionsResponseBody) SetMaxResults(v int32) *ListChatSessionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetMessage(v string) *ListChatSessionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetNextToken(v string) *ListChatSessionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListChatSessionsResponseBody) SetRequestId(v string) *ListChatSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListChatSessionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListChatSessionsResponseBodyData struct {
	// The list of chat sessions. Each element in the list is named session and represents a chat session.
	Content []*ListChatSessionsResponseBodyDataContent `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListChatSessionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListChatSessionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListChatSessionsResponseBodyData) GetContent() []*ListChatSessionsResponseBodyDataContent {
	return s.Content
}

func (s *ListChatSessionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListChatSessionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChatSessionsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListChatSessionsResponseBodyData) SetContent(v []*ListChatSessionsResponseBodyDataContent) *ListChatSessionsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListChatSessionsResponseBodyData) SetPageNumber(v int32) *ListChatSessionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListChatSessionsResponseBodyData) SetPageSize(v int32) *ListChatSessionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListChatSessionsResponseBodyData) SetTotalCount(v int64) *ListChatSessionsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListChatSessionsResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListChatSessionsResponseBodyDataContent struct {
	// The name of the chat session.
	//
	// example:
	//
	// Product consultation session
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The list of API products associated with the session. Each element in the list is named product.
	Products []*string `json:"products,omitempty" xml:"products,omitempty" type:"Repeated"`
	// The ID of the chat session. You can use this ID to locate the session in API operations such as UpdateChatSession, DeleteChatSession, and ListConversations.
	//
	// example:
	//
	// ses-xxxxxxxx
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The conversation type of the session. Valid values: MODEL and AGENT. These values are defined in the talkType enumeration of the CreateChatSession request parameter, and this response parameter does not have the @enums annotation.
	//
	// example:
	//
	// MODEL
	TalkType *string `json:"talkType,omitempty" xml:"talkType,omitempty"`
}

func (s ListChatSessionsResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s ListChatSessionsResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *ListChatSessionsResponseBodyDataContent) GetName() *string {
	return s.Name
}

func (s *ListChatSessionsResponseBodyDataContent) GetProducts() []*string {
	return s.Products
}

func (s *ListChatSessionsResponseBodyDataContent) GetSessionId() *string {
	return s.SessionId
}

func (s *ListChatSessionsResponseBodyDataContent) GetTalkType() *string {
	return s.TalkType
}

func (s *ListChatSessionsResponseBodyDataContent) SetName(v string) *ListChatSessionsResponseBodyDataContent {
	s.Name = &v
	return s
}

func (s *ListChatSessionsResponseBodyDataContent) SetProducts(v []*string) *ListChatSessionsResponseBodyDataContent {
	s.Products = v
	return s
}

func (s *ListChatSessionsResponseBodyDataContent) SetSessionId(v string) *ListChatSessionsResponseBodyDataContent {
	s.SessionId = &v
	return s
}

func (s *ListChatSessionsResponseBodyDataContent) SetTalkType(v string) *ListChatSessionsResponseBodyDataContent {
	s.TalkType = &v
	return s
}

func (s *ListChatSessionsResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
