// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchType(v string) *ListTagKeysRequest
	GetMatchType() *string
	SetMaxResults(v int32) *ListTagKeysRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagKeysRequest
	GetNextToken() *string
	SetTagKey(v string) *ListTagKeysRequest
	GetTagKey() *string
}

type ListTagKeysRequest struct {
	// The matching mode. Valid values:
	//
	// 	- Equals: equal match
	//
	// 	- Prefix: match by prefix
	//
	// example:
	//
	// Equals
	MatchType *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 20.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// If the total number of entries returned for the current request exceeds the value of the `MaxResults` parameter, the entries are truncated. In this case, you can use the `token` to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// AAAAAUYb00R0gHZBE8FVDeoh2ME93VeeEPUHs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tag key.
	//
	// example:
	//
	// test_key
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *ListTagKeysRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagKeysRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagKeysRequest) SetMatchType(v string) *ListTagKeysRequest {
	s.MatchType = &v
	return s
}

func (s *ListTagKeysRequest) SetMaxResults(v int32) *ListTagKeysRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetTagKey(v string) *ListTagKeysRequest {
	s.TagKey = &v
	return s
}

func (s *ListTagKeysRequest) Validate() error {
	return dara.Validate(s)
}
