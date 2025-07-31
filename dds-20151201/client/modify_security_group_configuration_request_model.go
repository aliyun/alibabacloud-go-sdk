// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifySecurityGroupConfigurationRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifySecurityGroupConfigurationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySecurityGroupConfigurationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySecurityGroupConfigurationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySecurityGroupConfigurationRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *ModifySecurityGroupConfigurationRequest
	GetSecurityGroupId() *string
}

type ModifySecurityGroupConfigurationRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bpxxxxxxxx
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the ECS security group.
	//
	// > 	- You can bind up to 10 ECS security groups to an ApsaraDB for MongoDB instance.
	//
	// > 	- You can call the [DescribeSecurityGroup](https://help.aliyun.com/document_detail/25556.html) operation of ECS to query the security groups in the specified region.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bpxxxxxxxx
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
}

func (s ModifySecurityGroupConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupConfigurationRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySecurityGroupConfigurationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySecurityGroupConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySecurityGroupConfigurationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySecurityGroupConfigurationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityGroupConfigurationRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifySecurityGroupConfigurationRequest) SetDBInstanceId(v string) *ModifySecurityGroupConfigurationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetOwnerAccount(v string) *ModifySecurityGroupConfigurationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetOwnerId(v int64) *ModifySecurityGroupConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetResourceOwnerAccount(v string) *ModifySecurityGroupConfigurationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetResourceOwnerId(v int64) *ModifySecurityGroupConfigurationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) SetSecurityGroupId(v string) *ModifySecurityGroupConfigurationRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
