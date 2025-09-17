// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilterShrink(v string) *GetResourcesShrinkRequest
	GetFilterShrink() *string
	SetMaxResults(v int32) *GetResourcesShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetResourcesShrinkRequest
	GetNextToken() *string
	SetRegionId(v string) *GetResourcesShrinkRequest
	GetRegionId() *string
}

type GetResourcesShrinkRequest struct {
	// The filter condition. The JSON format. You can use some resource properties as filter conditions.
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If you leave this parameter empty, the query starts from the beginning.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the region. This parameter is required if the cloud product is deployed in a region.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetResourcesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetResourcesShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *GetResourcesShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetResourcesShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResourcesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetResourcesShrinkRequest) SetFilterShrink(v string) *GetResourcesShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *GetResourcesShrinkRequest) SetMaxResults(v int32) *GetResourcesShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResourcesShrinkRequest) SetNextToken(v string) *GetResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourcesShrinkRequest) SetRegionId(v string) *GetResourcesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *GetResourcesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
