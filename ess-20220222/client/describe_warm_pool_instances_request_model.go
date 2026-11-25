// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWarmPoolInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeWarmPoolInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeWarmPoolInstancesRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeWarmPoolInstancesRequest
	GetRegionId() *string
	SetScalingGroupId(v string) *DescribeWarmPoolInstancesRequest
	GetScalingGroupId() *string
}

type DescribeWarmPoolInstancesRequest struct {
	// The maximum number of entries per page for a paged query. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The query token. Set the value to the NextToken value returned from the previous call. You do not need to specify this parameter for the first call. An empty NextToken value indicates that no more results are available.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeWarmPoolInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeWarmPoolInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeWarmPoolInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeWarmPoolInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeWarmPoolInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeWarmPoolInstancesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeWarmPoolInstancesRequest) SetMaxResults(v int32) *DescribeWarmPoolInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeWarmPoolInstancesRequest) SetNextToken(v string) *DescribeWarmPoolInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeWarmPoolInstancesRequest) SetRegionId(v string) *DescribeWarmPoolInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeWarmPoolInstancesRequest) SetScalingGroupId(v string) *DescribeWarmPoolInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeWarmPoolInstancesRequest) Validate() error {
	return dara.Validate(s)
}
