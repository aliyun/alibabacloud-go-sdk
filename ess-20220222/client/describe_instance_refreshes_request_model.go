// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRefreshesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRefreshTaskIds(v []*string) *DescribeInstanceRefreshesRequest
	GetInstanceRefreshTaskIds() []*string
	SetMaxResults(v int32) *DescribeInstanceRefreshesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInstanceRefreshesRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeInstanceRefreshesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceRefreshesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeInstanceRefreshesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceRefreshesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceRefreshesRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *DescribeInstanceRefreshesRequest
	GetScalingGroupId() *string
}

type DescribeInstanceRefreshesRequest struct {
	// The IDs of the instance refresh tasks that you want to query.
	InstanceRefreshTaskIds []*string `json:"InstanceRefreshTaskIds,omitempty" xml:"InstanceRefreshTaskIds,omitempty" type:"Repeated"`
	// The maximum number of entries per page. Valid values: 1 to 50. Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group to which the instance refresh task belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group.
	//
	// >  When you call this operation, you must specify one of the following parameters: ScalingGroupId and InstanceRefreshTaskIds. You cannot specify both of them. If you specify neither of them, an error is reported.
	//
	// example:
	//
	// asg-bp1ffogfdauy0jw0****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeInstanceRefreshesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRefreshesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRefreshesRequest) GetInstanceRefreshTaskIds() []*string {
	return s.InstanceRefreshTaskIds
}

func (s *DescribeInstanceRefreshesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInstanceRefreshesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInstanceRefreshesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceRefreshesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceRefreshesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceRefreshesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceRefreshesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceRefreshesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeInstanceRefreshesRequest) SetInstanceRefreshTaskIds(v []*string) *DescribeInstanceRefreshesRequest {
	s.InstanceRefreshTaskIds = v
	return s
}

func (s *DescribeInstanceRefreshesRequest) SetMaxResults(v int32) *DescribeInstanceRefreshesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInstanceRefreshesRequest) SetNextToken(v string) *DescribeInstanceRefreshesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInstanceRefreshesRequest) SetOwnerAccount(v string) *DescribeInstanceRefreshesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceRefreshesRequest) SetOwnerId(v int64) *DescribeInstanceRefreshesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceRefreshesRequest) SetRegionId(v string) *DescribeInstanceRefreshesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceRefreshesRequest) SetResourceOwnerAccount(v string) *DescribeInstanceRefreshesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceRefreshesRequest) SetResourceOwnerId(v int64) *DescribeInstanceRefreshesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceRefreshesRequest) SetScalingGroupId(v string) *DescribeInstanceRefreshesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeInstanceRefreshesRequest) Validate() error {
	return dara.Validate(s)
}
