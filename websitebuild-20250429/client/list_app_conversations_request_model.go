// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppConversationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBotId(v string) *ListAppConversationsRequest
	GetBotId() *string
	SetEndModifyTime(v string) *ListAppConversationsRequest
	GetEndModifyTime() *string
	SetMaxResults(v int32) *ListAppConversationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAppConversationsRequest
	GetNextToken() *string
	SetPageNum(v int32) *ListAppConversationsRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListAppConversationsRequest
	GetPageSize() *int32
	SetSiteId(v string) *ListAppConversationsRequest
	GetSiteId() *string
	SetStartModifyTime(v string) *ListAppConversationsRequest
	GetStartModifyTime() *string
}

type ListAppConversationsRequest struct {
	// Bot ID
	//
	// example:
	//
	// Zero2
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// End modification time (ISO 8601 format)
	//
	// example:
	//
	// 20201212
	EndModifyTime *string `json:"EndModifyTime,omitempty" xml:"EndModifyTime,omitempty"`
	// The number of entries to return in each query result.
	//
	// Valid values: 10 to 100. Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Token indicating the start of the next query. It is empty when there is no next query.
	//
	// example:
	//
	// 0l45bkwM022Dt+rOvPi/oQ==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Page number
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Number of entries per page
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Site ID
	//
	// example:
	//
	// 1168642640022064
	SiteId *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Start modification time (ISO 8601 format)
	//
	// example:
	//
	// 20200101
	StartModifyTime *string `json:"StartModifyTime,omitempty" xml:"StartModifyTime,omitempty"`
}

func (s ListAppConversationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppConversationsRequest) GoString() string {
	return s.String()
}

func (s *ListAppConversationsRequest) GetBotId() *string {
	return s.BotId
}

func (s *ListAppConversationsRequest) GetEndModifyTime() *string {
	return s.EndModifyTime
}

func (s *ListAppConversationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAppConversationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAppConversationsRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListAppConversationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppConversationsRequest) GetSiteId() *string {
	return s.SiteId
}

func (s *ListAppConversationsRequest) GetStartModifyTime() *string {
	return s.StartModifyTime
}

func (s *ListAppConversationsRequest) SetBotId(v string) *ListAppConversationsRequest {
	s.BotId = &v
	return s
}

func (s *ListAppConversationsRequest) SetEndModifyTime(v string) *ListAppConversationsRequest {
	s.EndModifyTime = &v
	return s
}

func (s *ListAppConversationsRequest) SetMaxResults(v int32) *ListAppConversationsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAppConversationsRequest) SetNextToken(v string) *ListAppConversationsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAppConversationsRequest) SetPageNum(v int32) *ListAppConversationsRequest {
	s.PageNum = &v
	return s
}

func (s *ListAppConversationsRequest) SetPageSize(v int32) *ListAppConversationsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppConversationsRequest) SetSiteId(v string) *ListAppConversationsRequest {
	s.SiteId = &v
	return s
}

func (s *ListAppConversationsRequest) SetStartModifyTime(v string) *ListAppConversationsRequest {
	s.StartModifyTime = &v
	return s
}

func (s *ListAppConversationsRequest) Validate() error {
	return dara.Validate(s)
}
