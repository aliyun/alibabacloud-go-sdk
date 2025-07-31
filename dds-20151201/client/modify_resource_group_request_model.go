// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyResourceGroupRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyResourceGroupRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyResourceGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ModifyResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyResourceGroupRequest
	GetResourceOwnerId() *int64
}

type ModifyResourceGroupRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1366caac83****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/61933.html) operation to query the region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. For more information, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-acfmyiu4ekp****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyResourceGroupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyResourceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyResourceGroupRequest) SetDBInstanceId(v string) *ModifyResourceGroupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetOwnerAccount(v string) *ModifyResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetOwnerId(v int64) *ModifyResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetRegionId(v string) *ModifyResourceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetResourceGroupId(v string) *ModifyResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
