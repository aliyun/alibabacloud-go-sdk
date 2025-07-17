// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebalanceInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *RebalanceInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RebalanceInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RebalanceInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RebalanceInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RebalanceInstancesRequest
	GetResourceOwnerId() *int64
	SetScalingGroupId(v string) *RebalanceInstancesRequest
	GetScalingGroupId() *string
}

type RebalanceInstancesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the scaling group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the scaling group.
	//
	// This parameter is required.
	//
	// example:
	//
	// asg-bp18p2yfxow2dloq****
	ScalingGroupId *string `json:"ScalingGroupId,omitempty" xml:"ScalingGroupId,omitempty"`
}

func (s RebalanceInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RebalanceInstancesRequest) GoString() string {
	return s.String()
}

func (s *RebalanceInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RebalanceInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RebalanceInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebalanceInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RebalanceInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RebalanceInstancesRequest) GetScalingGroupId() *string {
	return s.ScalingGroupId
}

func (s *RebalanceInstancesRequest) SetOwnerAccount(v string) *RebalanceInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RebalanceInstancesRequest) SetOwnerId(v int64) *RebalanceInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RebalanceInstancesRequest) SetRegionId(v string) *RebalanceInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RebalanceInstancesRequest) SetResourceOwnerAccount(v string) *RebalanceInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RebalanceInstancesRequest) SetResourceOwnerId(v int64) *RebalanceInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RebalanceInstancesRequest) SetScalingGroupId(v string) *RebalanceInstancesRequest {
	s.ScalingGroupId = &v
	return s
}

func (s *RebalanceInstancesRequest) Validate() error {
	return dara.Validate(s)
}
