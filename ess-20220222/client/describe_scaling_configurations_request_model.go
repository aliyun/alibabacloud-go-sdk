// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScalingConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeScalingConfigurationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeScalingConfigurationsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeScalingConfigurationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeScalingConfigurationsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeScalingConfigurationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeScalingConfigurationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeScalingConfigurationsRequest
	GetResourceOwnerId() *int64
	SetScalingConfigurationIds(v []*string) *DescribeScalingConfigurationsRequest
	GetScalingConfigurationIds() []*string
	SetScalingConfigurationNames(v []*string) *DescribeScalingConfigurationsRequest
	GetScalingConfigurationNames() []*string
	SetScalingGroupId(v string) *DescribeScalingConfigurationsRequest
	GetScalingGroupId() *string
}

type DescribeScalingConfigurationsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the scaling configuration list. Pages start from 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in paged queries. Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the scaling group to which the scaling configuration belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the scaling configurations to query.
	//
	// The query results include both active and inactive scaling configurations, identified by the response parameter `LifecycleState`.
	ScalingConfigurationIds []*string `json:"ScalingConfigurationIds,omitempty" xml:"ScalingConfigurationIds,omitempty" type:"Repeated"`
	// The names of the scaling configurations to query.
	//
	// The query ignores invalid scaling configuration names without returning an error.
	ScalingConfigurationNames []*string `json:"ScalingConfigurationNames,omitempty" xml:"ScalingConfigurationNames,omitempty" type:"Repeated"`
	// The ID of the scaling group. You can query all scaling configurations under this scaling group.
	//
	// example:
	//
	// asg-bp17pelvl720x3v7****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeScalingConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScalingConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScalingConfigurationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeScalingConfigurationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeScalingConfigurationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeScalingConfigurationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeScalingConfigurationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeScalingConfigurationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeScalingConfigurationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeScalingConfigurationsRequest) GetScalingConfigurationIds() []*string {
	return s.ScalingConfigurationIds
}

func (s *DescribeScalingConfigurationsRequest) GetScalingConfigurationNames() []*string {
	return s.ScalingConfigurationNames
}

func (s *DescribeScalingConfigurationsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeScalingConfigurationsRequest) SetOwnerAccount(v string) *DescribeScalingConfigurationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetOwnerId(v int64) *DescribeScalingConfigurationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetPageNumber(v int32) *DescribeScalingConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetPageSize(v int32) *DescribeScalingConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetRegionId(v string) *DescribeScalingConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetResourceOwnerAccount(v string) *DescribeScalingConfigurationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetResourceOwnerId(v int64) *DescribeScalingConfigurationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingConfigurationIds(v []*string) *DescribeScalingConfigurationsRequest {
	s.ScalingConfigurationIds = v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingConfigurationNames(v []*string) *DescribeScalingConfigurationsRequest {
	s.ScalingConfigurationNames = v
	return s
}

func (s *DescribeScalingConfigurationsRequest) SetScalingGroupId(v string) *DescribeScalingConfigurationsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeScalingConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
