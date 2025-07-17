// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotificationConfigurationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeNotificationConfigurationsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeNotificationConfigurationsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeNotificationConfigurationsRequest
	GetResourceOwnerAccount() *string
	SetScalingGroupId(v string) *DescribeNotificationConfigurationsRequest
	GetScalingGroupId() *string
}

type DescribeNotificationConfigurationsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp1igpak5ft1flyp****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s DescribeNotificationConfigurationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotificationConfigurationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNotificationConfigurationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNotificationConfigurationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNotificationConfigurationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNotificationConfigurationsRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *DescribeNotificationConfigurationsRequest) SetOwnerId(v int64) *DescribeNotificationConfigurationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNotificationConfigurationsRequest) SetRegionId(v string) *DescribeNotificationConfigurationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNotificationConfigurationsRequest) SetResourceOwnerAccount(v string) *DescribeNotificationConfigurationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNotificationConfigurationsRequest) SetScalingGroupId(v string) *DescribeNotificationConfigurationsRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *DescribeNotificationConfigurationsRequest) Validate() error {
	return dara.Validate(s)
}
