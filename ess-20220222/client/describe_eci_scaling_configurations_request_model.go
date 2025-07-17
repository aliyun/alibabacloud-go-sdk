// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEciScalingConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeEciScalingConfigurationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEciScalingConfigurationsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeEciScalingConfigurationsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeEciScalingConfigurationsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeEciScalingConfigurationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeEciScalingConfigurationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEciScalingConfigurationsRequest
	GetResourceOwnerId() *int64
	SetScalingConfigurationIds(v []*string) *DescribeEciScalingConfigurationsRequest
	GetScalingConfigurationIds() []*string
	SetScalingConfigurationNames(v []*string) *DescribeEciScalingConfigurationsRequest
	GetScalingConfigurationNames() []*string
	SetScalingGroupId(v string) *DescribeEciScalingConfigurationsRequest
	GetScalingGroupId() *string
}

type DescribeEciScalingConfigurationsRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 50.
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
	// The IDs of the scaling configurations that you want to query. You can specify 1 to 10 scaling configuration IDs.
	//
	// The IDs of active and inactive scaling configurations are displayed in the query results. You can distinguish between active and inactive scaling configurations based on the value of `LifecycleState`.
	ScalingConfigurationIds []*string `json:"ScalingConfigurationIds,omitempty" xml:"ScalingConfigurationIds,omitempty" type:"Repeated"`
	// The names of the scaling configurations that you want to query. You can specify 1 to 10 scaling configuration names.
	//
	// The names of inactive scaling configurations are not displayed in the query results, and no error is reported.
	ScalingConfigurationNames []*string `json:"ScalingConfigurationNames,omitempty" xml:"ScalingConfigurationNames,omitempty" type:"Repeated"`
	// The ID of the scaling group. You can use the ID to query all scaling configurations in the scaling group.
	//
	// example:
	//
	// asg-bp17pelvl720x3v7****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeEciScalingConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEciScalingConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeEciScalingConfigurationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEciScalingConfigurationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEciScalingConfigurationsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeEciScalingConfigurationsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeEciScalingConfigurationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEciScalingConfigurationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEciScalingConfigurationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEciScalingConfigurationsRequest) GetScalingConfigurationIds() []*string {
	return s.ScalingConfigurationIds
}

func (s *DescribeEciScalingConfigurationsRequest) GetScalingConfigurationNames() []*string {
	return s.ScalingConfigurationNames
}

func (s *DescribeEciScalingConfigurationsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeEciScalingConfigurationsRequest) SetOwnerAccount(v string) *DescribeEciScalingConfigurationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetOwnerId(v int64) *DescribeEciScalingConfigurationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetPageNumber(v int32) *DescribeEciScalingConfigurationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetPageSize(v int32) *DescribeEciScalingConfigurationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetRegionId(v string) *DescribeEciScalingConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetResourceOwnerAccount(v string) *DescribeEciScalingConfigurationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetResourceOwnerId(v int64) *DescribeEciScalingConfigurationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetScalingConfigurationIds(v []*string) *DescribeEciScalingConfigurationsRequest {
	s.ScalingConfigurationIds = v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetScalingConfigurationNames(v []*string) *DescribeEciScalingConfigurationsRequest {
	s.ScalingConfigurationNames = v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) SetScalingGroupId(v string) *DescribeEciScalingConfigurationsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeEciScalingConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
