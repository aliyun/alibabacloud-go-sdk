// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLiveInputsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListMediaLiveInputsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListMediaLiveInputsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMediaLiveInputsRequest
	GetNextToken() *string
	SetSkip(v int32) *ListMediaLiveInputsRequest
	GetSkip() *int32
	SetSortOrder(v string) *ListMediaLiveInputsRequest
	GetSortOrder() *string
	SetTypes(v string) *ListMediaLiveInputsRequest
	GetTypes() *string
}

type ListMediaLiveInputsRequest struct {
	// The keyword of the query. You can perform a fuzzy search on input ID or name.
	//
	// example:
	//
	// 123
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: If you do not specify this parameter or if you set a value smaller than 10, the default value is 10. If you set a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to be skipped in the query. If the number of entries you attempt to skip exceeds the number of entries that meet the condition, an empty list is returned.
	//
	// example:
	//
	// 20
	Skip *int32 `json:"Skip,omitempty" xml:"Skip,omitempty"`
	// The sorting order of the inputs by creation time. Default value: asc. Valid values: desc and asc. asc indicates the ascending order, and desc indicates the descending order.
	//
	// example:
	//
	// desc
	SortOrder *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	// The type of inputs you want to query. You can separate multiple input types with commas (,) in a JSON array.
	//
	// example:
	//
	// ["RTMP_PUSH","SRT_PULL"]
	Types *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListMediaLiveInputsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveInputsRequest) GoString() string {
	return s.String()
}

func (s *ListMediaLiveInputsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMediaLiveInputsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMediaLiveInputsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMediaLiveInputsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *ListMediaLiveInputsRequest) GetSortOrder() *string {
	return s.SortOrder
}

func (s *ListMediaLiveInputsRequest) GetTypes() *string {
	return s.Types
}

func (s *ListMediaLiveInputsRequest) SetKeyword(v string) *ListMediaLiveInputsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMediaLiveInputsRequest) SetMaxResults(v int32) *ListMediaLiveInputsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMediaLiveInputsRequest) SetNextToken(v string) *ListMediaLiveInputsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMediaLiveInputsRequest) SetSkip(v int32) *ListMediaLiveInputsRequest {
	s.Skip = &v
	return s
}

func (s *ListMediaLiveInputsRequest) SetSortOrder(v string) *ListMediaLiveInputsRequest {
	s.SortOrder = &v
	return s
}

func (s *ListMediaLiveInputsRequest) SetTypes(v string) *ListMediaLiveInputsRequest {
	s.Types = &v
	return s
}

func (s *ListMediaLiveInputsRequest) Validate() error {
	return dara.Validate(s)
}
