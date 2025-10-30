// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccountFactoryBaselinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAccountFactoryBaselinesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAccountFactoryBaselinesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListAccountFactoryBaselinesRequest
	GetRegionId() *string
}

type ListAccountFactoryBaselinesRequest struct {
	// The maximum number of entries to return on each page.
	//
	// Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// AAAAALHWGpGoYCcYMxiFfmlhvh62Xr2DzYbz/SAfc*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAccountFactoryBaselinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAccountFactoryBaselinesRequest) GoString() string {
	return s.String()
}

func (s *ListAccountFactoryBaselinesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAccountFactoryBaselinesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccountFactoryBaselinesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAccountFactoryBaselinesRequest) SetMaxResults(v int32) *ListAccountFactoryBaselinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAccountFactoryBaselinesRequest) SetNextToken(v string) *ListAccountFactoryBaselinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAccountFactoryBaselinesRequest) SetRegionId(v string) *ListAccountFactoryBaselinesRequest {
	s.RegionId = &v
	return s
}

func (s *ListAccountFactoryBaselinesRequest) Validate() error {
	return dara.Validate(s)
}
