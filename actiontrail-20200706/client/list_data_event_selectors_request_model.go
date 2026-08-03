// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataEventSelectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListDataEventSelectorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataEventSelectorsRequest
	GetNextToken() *string
}

type ListDataEventSelectorsRequest struct {
	// The maximum number of results to return.
	//
	// - Valid values: 1 to 100.
	//
	// - Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// - You do not need to specify this parameter for the first request.
	//
	// - You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// VjE6dLbnNpVmbz06****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListDataEventSelectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventSelectorsRequest) GoString() string {
	return s.String()
}

func (s *ListDataEventSelectorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataEventSelectorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataEventSelectorsRequest) SetMaxResults(v int32) *ListDataEventSelectorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataEventSelectorsRequest) SetNextToken(v string) *ListDataEventSelectorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataEventSelectorsRequest) Validate() error {
	return dara.Validate(s)
}
