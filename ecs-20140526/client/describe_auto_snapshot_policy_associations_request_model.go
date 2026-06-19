// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyAssociationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetAutoSnapshotPolicyId() *string
	SetDiskId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetDiskId() *string
	SetMaxResults(v int32) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoSnapshotPolicyAssociationsRequest
	GetResourceOwnerId() *int64
}

type DescribeAutoSnapshotPolicyAssociationsRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// - You can specify only one of AutoSnapshotPolicyId and DiskId.
	//
	// example:
	//
	// sp-bp12quk7gqhhuu1f****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The cloud disk ID.
	//
	// - You can specify only one of AutoSnapshotPolicyId and DiskId.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The maximum number of entries per page for paging queries. Maximum value: 100.
	//
	// Default value:
	//
	// - If this parameter is not set or is set to a value less than 10, the default value is 10.
	//
	// - If the value is set to a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the NextToken value returned in the previous API call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the automatic snapshot policy. You can call [DescribeRegions](https://www.alibabacloud.com/help/en/ecs/developer-reference/api-ecs-2014-05-26-describeregions) to query the most recent list of Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAutoSnapshotPolicyAssociationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyAssociationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetDiskId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetMaxResults(v int32) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetNextToken(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetOwnerAccount(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetOwnerId(v int64) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetRegionId(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetResourceOwnerAccount(v string) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) SetResourceOwnerId(v int64) *DescribeAutoSnapshotPolicyAssociationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyAssociationsRequest) Validate() error {
	return dara.Validate(s)
}
