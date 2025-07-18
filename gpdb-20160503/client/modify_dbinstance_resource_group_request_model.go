// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceResourceGroupRequest
	GetDBInstanceId() *string
	SetNewResourceGroupId(v string) *ModifyDBInstanceResourceGroupRequest
	GetNewResourceGroupId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceResourceGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceResourceGroupRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyDBInstanceResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyDBInstanceResourceGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceResourceGroupRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceResourceGroupRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The ID of the resource group to which you want to move the instance. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-bp67acfmxazb4b****
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceResourceGroupRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ModifyDBInstanceResourceGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceResourceGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceResourceGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceResourceGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceResourceGroupRequest) SetDBInstanceId(v string) *ModifyDBInstanceResourceGroupRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetNewResourceGroupId(v string) *ModifyDBInstanceResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetOwnerAccount(v string) *ModifyDBInstanceResourceGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetOwnerId(v int64) *ModifyDBInstanceResourceGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetResourceGroupId(v string) *ModifyDBInstanceResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceResourceGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceResourceGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
