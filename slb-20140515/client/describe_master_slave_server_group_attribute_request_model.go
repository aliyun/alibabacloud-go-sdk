// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMasterSlaveServerGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMasterSlaveServerGroupId(v string) *DescribeMasterSlaveServerGroupAttributeRequest
	GetMasterSlaveServerGroupId() *string
	SetOwnerAccount(v string) *DescribeMasterSlaveServerGroupAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeMasterSlaveServerGroupAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeMasterSlaveServerGroupAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeMasterSlaveServerGroupAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeMasterSlaveServerGroupAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeMasterSlaveServerGroupAttributeRequest struct {
	// The ID of the primary/secondary server group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rsp-cige6j******
	MasterSlaveServerGroupId *string `json:"MasterSlaveServerGroupId,omitempty" xml:"MasterSlaveServerGroupId,omitempty"`
	OwnerAccount             *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance.
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

func (s DescribeMasterSlaveServerGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMasterSlaveServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) GetMasterSlaveServerGroupId() *string {
	return s.MasterSlaveServerGroupId
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetMasterSlaveServerGroupId(v string) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.MasterSlaveServerGroupId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetOwnerAccount(v string) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetOwnerId(v int64) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetRegionId(v string) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetResourceOwnerAccount(v string) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) SetResourceOwnerId(v int64) *DescribeMasterSlaveServerGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeMasterSlaveServerGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
