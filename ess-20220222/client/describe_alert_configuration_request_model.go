// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAlertConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeAlertConfigurationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAlertConfigurationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAlertConfigurationRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DescribeAlertConfigurationRequest
	GetScalingGroupId() *string
}

type DescribeAlertConfigurationRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeAlertConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAlertConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeAlertConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAlertConfigurationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAlertConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAlertConfigurationRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeAlertConfigurationRequest) SetOwnerId(v int64) *DescribeAlertConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAlertConfigurationRequest) SetRegionId(v string) *DescribeAlertConfigurationRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAlertConfigurationRequest) SetResourceOwnerAccount(v string) *DescribeAlertConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAlertConfigurationRequest) SetScalingGroupId(v string) *DescribeAlertConfigurationRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeAlertConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
