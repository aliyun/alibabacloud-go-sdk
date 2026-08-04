// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *GetApiEndpointsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetApiEndpointsRequest
	GetNextToken() *string
	SetRegionId(v string) *GetApiEndpointsRequest
	GetRegionId() *string
}

type GetApiEndpointsRequest struct {
	// The maximum number of records to return in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetApiEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiEndpointsRequest) GoString() string {
	return s.String()
}

func (s *GetApiEndpointsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetApiEndpointsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetApiEndpointsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApiEndpointsRequest) SetMaxResults(v int32) *GetApiEndpointsRequest {
	s.MaxResults = &v
	return s
}

func (s *GetApiEndpointsRequest) SetNextToken(v string) *GetApiEndpointsRequest {
	s.NextToken = &v
	return s
}

func (s *GetApiEndpointsRequest) SetRegionId(v string) *GetApiEndpointsRequest {
	s.RegionId = &v
	return s
}

func (s *GetApiEndpointsRequest) Validate() error {
	return dara.Validate(s)
}
