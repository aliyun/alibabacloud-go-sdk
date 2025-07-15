// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalAccelerationInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalAccelerationInstanceId(v string) *DeleteGlobalAccelerationInstanceRequest
	GetGlobalAccelerationInstanceId() *string
	SetOwnerAccount(v string) *DeleteGlobalAccelerationInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteGlobalAccelerationInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteGlobalAccelerationInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteGlobalAccelerationInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteGlobalAccelerationInstanceRequest
	GetResourceOwnerId() *int64
}

type DeleteGlobalAccelerationInstanceRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-asdfsl22s****
	GlobalAccelerationInstanceId *string `json:"GlobalAccelerationInstanceId,omitempty" xml:"GlobalAccelerationInstanceId,omitempty"`
	OwnerAccount                 *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s DeleteGlobalAccelerationInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalAccelerationInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGlobalAccelerationInstanceRequest) GetGlobalAccelerationInstanceId() *string {
	return s.GlobalAccelerationInstanceId
}

func (s *DeleteGlobalAccelerationInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteGlobalAccelerationInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteGlobalAccelerationInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGlobalAccelerationInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteGlobalAccelerationInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteGlobalAccelerationInstanceRequest) SetGlobalAccelerationInstanceId(v string) *DeleteGlobalAccelerationInstanceRequest {
	s.GlobalAccelerationInstanceId = &v
	return s
}

func (s *DeleteGlobalAccelerationInstanceRequest) SetOwnerAccount(v string) *DeleteGlobalAccelerationInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteGlobalAccelerationInstanceRequest) SetOwnerId(v int64) *DeleteGlobalAccelerationInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteGlobalAccelerationInstanceRequest) SetRegionId(v string) *DeleteGlobalAccelerationInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGlobalAccelerationInstanceRequest) SetResourceOwnerAccount(v string) *DeleteGlobalAccelerationInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteGlobalAccelerationInstanceRequest) SetResourceOwnerId(v int64) *DeleteGlobalAccelerationInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteGlobalAccelerationInstanceRequest) Validate() error {
	return dara.Validate(s)
}
