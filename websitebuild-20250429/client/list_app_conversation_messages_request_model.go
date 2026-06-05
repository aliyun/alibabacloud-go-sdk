// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConversationMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *ListAppConversationMessagesRequest
	GetConversationId() *string
	SetMaxResults(v int32) *ListAppConversationMessagesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppConversationMessagesRequest
	GetNextToken() *string
	SetPageSize(v int32) *ListAppConversationMessagesRequest
	GetPageSize() *int32
	SetSiteId(v string) *ListAppConversationMessagesRequest
	GetSiteId() *string
	SetStartCreateTime(v string) *ListAppConversationMessagesRequest
	GetStartCreateTime() *string
}

type ListAppConversationMessagesRequest struct {
	// example:
	//
	// 81bc5a34-1d8d-4ef7-a208-7401c51b054b
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// FFh3Xqm+JgZ/U9Jyb7wdVr9LWk80Tghn5UZjbcWEVEderBcbVF+Y6PS0i8PpCL4PQZ3e0C9oEH0Asd4tJEuGtkl2WuKdiWZpEwadNydQdJPFM=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 734600510965856
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// null
	StartCreateTime *string `json:"StartCreateTime,omitempty" xml:"StartCreateTime,omitempty"`
}

func (s ListAppConversationMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppConversationMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListAppConversationMessagesRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListAppConversationMessagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppConversationMessagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppConversationMessagesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppConversationMessagesRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *ListAppConversationMessagesRequest) GetStartCreateTime() *string {
	return s.StartCreateTime
}

func (s *ListAppConversationMessagesRequest) SetConversationId(v string) *ListAppConversationMessagesRequest {
	s.ConversationId = &v
	return s
}

func (s *ListAppConversationMessagesRequest) SetMaxResults(v int32) *ListAppConversationMessagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppConversationMessagesRequest) SetNextToken(v string) *ListAppConversationMessagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppConversationMessagesRequest) SetPageSize(v int32) *ListAppConversationMessagesRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppConversationMessagesRequest) SetSiteId(v string) *ListAppConversationMessagesRequest {
	s.SiteId = &v
	return s
}

func (s *ListAppConversationMessagesRequest) SetStartCreateTime(v string) *ListAppConversationMessagesRequest {
	s.StartCreateTime = &v
	return s
}

func (s *ListAppConversationMessagesRequest) Validate() error {
	return dara.Validate(s)
}
