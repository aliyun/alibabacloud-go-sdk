// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateGlobalAccelerationInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalAccelerationInstanceId(v string) *UnassociateGlobalAccelerationInstanceRequest
	GetGlobalAccelerationInstanceId() *string
	SetInstanceType(v string) *UnassociateGlobalAccelerationInstanceRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *UnassociateGlobalAccelerationInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnassociateGlobalAccelerationInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnassociateGlobalAccelerationInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnassociateGlobalAccelerationInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnassociateGlobalAccelerationInstanceRequest
	GetResourceOwnerId() *int64
}

type UnassociateGlobalAccelerationInstanceRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-1sxeedefrr33****
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	// The backend server type. Valid values:
	//
	// 	- **RemoteEcsInstance**: Elastic Compute Service (ECS) instance
	//
	// 	- **RemoteSlbInstance**: Server Load Balancer (SLB) instance
	//
	// 	- **RemoteEniInstance**: elastic network interface (ENI)
	//
	// example:
	//
	// RemoteEcsInstance
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the GA instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
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

func (s UnassociateGlobalAccelerationInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UnassociateGlobalAccelerationInstanceRequest) GoString() string {
	return s.String()
}

func (s *UnassociateGlobalAccelerationInstanceRequest) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *UnassociateGlobalAccelerationInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UnassociateGlobalAccelerationInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnassociateGlobalAccelerationInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnassociateGlobalAccelerationInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnassociateGlobalAccelerationInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnassociateGlobalAccelerationInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnassociateGlobalAccelerationInstanceRequest) SetGlobalAccelerationInstanceId(v string) *UnassociateGlobalAccelerationInstanceRequest {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceRequest) SetInstanceType(v string) *UnassociateGlobalAccelerationInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceRequest) SetOwnerAccount(v string) *UnassociateGlobalAccelerationInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceRequest) SetOwnerId(v int64) *UnassociateGlobalAccelerationInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceRequest) SetRegionId(v string) *UnassociateGlobalAccelerationInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceRequest) SetResourceOwnerAccount(v string) *UnassociateGlobalAccelerationInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceRequest) SetResourceOwnerId(v int64) *UnassociateGlobalAccelerationInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnassociateGlobalAccelerationInstanceRequest) Validate() error {
	return dara.Validate(s)
}
