// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCopilotMessagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscending(v bool) *ListCopilotMessagesRequest
	GetAscending() *bool
	SetConversationId(v string) *ListCopilotMessagesRequest
	GetConversationId() *string
	SetLang(v string) *ListCopilotMessagesRequest
	GetLang() *string
	SetMaxResults(v int32) *ListCopilotMessagesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCopilotMessagesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListCopilotMessagesRequest
	GetRegionId() *string
	SetTraceId(v string) *ListCopilotMessagesRequest
	GetTraceId() *string
}

type ListCopilotMessagesRequest struct {
	// Used together with the SortingField parameter. A value of **true*	- specifies ascending order, and a value of **false*	- specifies descending order.
	//
	// example:
	//
	// false
	Ascending *bool `json:"Ascending,omitempty" xml:"Ascending,omitempty"`
	// The conversation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54b6c969-4b75-47be-9f43-71f88aeb****
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// The language of the returned messages. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The maximum number of results to return when you use the NextToken-based pagination method. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the position from which you want to start the query. Leave this parameter empty to query from the beginning.
	//
	// example:
	//
	// AAAAAUqcj6VO4E3ECWIrFczs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the Cloud Threat Detection and Response (CTDR) data management center resides. Specify the management center based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets reside in the Chinese mainland or Hong Kong (China).
	//
	// - ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The trace ID, which is used to query specific call information.
	//
	// example:
	//
	// 550e8400e29b41d4a71644665544****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ListCopilotMessagesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCopilotMessagesRequest) GoString() string {
	return s.String()
}

func (s *ListCopilotMessagesRequest) GetAscending() *bool {
	return s.Ascending
}

func (s *ListCopilotMessagesRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *ListCopilotMessagesRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCopilotMessagesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCopilotMessagesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCopilotMessagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCopilotMessagesRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *ListCopilotMessagesRequest) SetAscending(v bool) *ListCopilotMessagesRequest {
	s.Ascending = &v
	return s
}

func (s *ListCopilotMessagesRequest) SetConversationId(v string) *ListCopilotMessagesRequest {
	s.ConversationId = &v
	return s
}

func (s *ListCopilotMessagesRequest) SetLang(v string) *ListCopilotMessagesRequest {
	s.Lang = &v
	return s
}

func (s *ListCopilotMessagesRequest) SetMaxResults(v int32) *ListCopilotMessagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCopilotMessagesRequest) SetNextToken(v string) *ListCopilotMessagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListCopilotMessagesRequest) SetRegionId(v string) *ListCopilotMessagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCopilotMessagesRequest) SetTraceId(v string) *ListCopilotMessagesRequest {
	s.TraceId = &v
	return s
}

func (s *ListCopilotMessagesRequest) Validate() error {
	return dara.Validate(s)
}
