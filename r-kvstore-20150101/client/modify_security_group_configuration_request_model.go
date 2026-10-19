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
	SetSecurityToken(v string) *ModifySecurityGroupConfigurationRequest
	GetSecurityToken() *string
}

type ModifySecurityGroupConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the security groups that you want to set. You can specify up to 10 security group IDs, separated by commas (,).
	//
	// > This operation overwrites the existing settings. For valid values, query the basic information about the security groups that you created by calling [DescribeSecurityGroups](https://next.api.aliyun.com/api/Ecs/2014-05-26/DescribeSecurityGroups).
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bpcfmyiu4ekp****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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

func (s *ModifySecurityGroupConfigurationRequest) GetSecurityToken() *string {
	return s.SecurityToken
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

func (s *ModifySecurityGroupConfigurationRequest) SetSecurityToken(v string) *ModifySecurityGroupConfigurationRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifySecurityGroupConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
