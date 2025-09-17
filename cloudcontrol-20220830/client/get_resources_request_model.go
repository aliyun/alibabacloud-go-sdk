// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v map[string]interface{}) *GetResourcesRequest
	GetFilter() map[string]interface{}
	SetMaxResults(v int32) *GetResourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *GetResourcesRequest
	GetRegionId() *string
}

type GetResourcesRequest struct {
	// The filter condition. The JSON format. You can use some resource properties as filter conditions.
	Filter map[string]interface{} `json:"filter,omitempty" xml:"filter,omitempty"`
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

func (s GetResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourcesRequest) GoString() string {
	return s.String()
}

func (s *GetResourcesRequest) GetFilter() map[string]interface{} {
	return s.Filter
}

func (s *GetResourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetResourcesRequest) SetFilter(v map[string]interface{}) *GetResourcesRequest {
	s.Filter = v
	return s
}

func (s *GetResourcesRequest) SetMaxResults(v int32) *GetResourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResourcesRequest) SetNextToken(v string) *GetResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourcesRequest) SetRegionId(v string) *GetResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *GetResourcesRequest) Validate() error {
	return dara.Validate(s)
}
