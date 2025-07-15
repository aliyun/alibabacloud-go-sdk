// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServerRelatedGlobalAccelerationInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeServerRelatedGlobalAccelerationInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeServerRelatedGlobalAccelerationInstancesRequest
	GetResourceOwnerId() *int64
	SetServerId(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest
	GetServerId() *string
	SetServerType(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest
	GetServerType() *string
}

type DescribeServerRelatedGlobalAccelerationInstancesRequest struct {
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
	// The ID of the backend service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-12s3sdf****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of the backend service instance. Valid values:
	//
	// 	- **EcsInstance*	- (default): Elastic Compute Service (ECS)
	//
	// 	- **SlbInstance**: Server Load Balancer (SLB)
	//
	// example:
	//
	// EcsInstance
	ServerType *string `json:"ServerType,omitempty" xml:"ServerType,omitempty"`
}

func (s DescribeServerRelatedGlobalAccelerationInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServerRelatedGlobalAccelerationInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) GetServerType() *string {
	return s.ServerType
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) SetOwnerAccount(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) SetOwnerId(v int64) *DescribeServerRelatedGlobalAccelerationInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) SetRegionId(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) SetResourceOwnerAccount(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) SetResourceOwnerId(v int64) *DescribeServerRelatedGlobalAccelerationInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) SetServerId(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest {
	s.ServerId = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) SetServerType(v string) *DescribeServerRelatedGlobalAccelerationInstancesRequest {
	s.ServerType = &v
	return s
}

func (s *DescribeServerRelatedGlobalAccelerationInstancesRequest) Validate() error {
	return dara.Validate(s)
}
