// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListChatSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListChatSessionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListChatSessionsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListChatSessionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListChatSessionsRequest
	GetPageSize() *int32
}

type ListChatSessionsRequest struct {
	// The maximum number of entries per page. The default value and value range are not specified in the API definition.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Pass the nextToken returned in the previous response to obtain the next page. Do not specify this parameter for the first query.
	//
	// example:
	//
	// token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number. Pages start from 1.
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
}

func (s ListChatSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListChatSessionsRequest) GoString() string {
	return s.String()
}

func (s *ListChatSessionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListChatSessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListChatSessionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListChatSessionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListChatSessionsRequest) SetMaxResults(v int32) *ListChatSessionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListChatSessionsRequest) SetNextToken(v string) *ListChatSessionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListChatSessionsRequest) SetPageNumber(v int32) *ListChatSessionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListChatSessionsRequest) SetPageSize(v int32) *ListChatSessionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListChatSessionsRequest) Validate() error {
	return dara.Validate(s)
}
