// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKey(v string) *ListTagValuesRequest
	GetKey() *string
	SetMaxResults(v int32) *ListTagValuesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagValuesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagValuesRequest
	GetRegionId() *string
	SetResourceType(v string) *ListTagValuesRequest
	GetResourceType() *string
}

type ListTagValuesRequest struct {
	// The tag key to query.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The maximum number of results on each page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request.
	//
	// example:
	//
	// 3272h923879hsaksad
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the tagged resource.
	//
	// example:
	//
	// template
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) GetKey() *string {
	return s.Key
}

func (s *ListTagValuesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagValuesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
