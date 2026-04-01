// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLCollectorRetentionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *ModifySQLCollectorRetentionRequest
	GetConfigValue() *string
	SetDBInstanceId(v string) *ModifySQLCollectorRetentionRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifySQLCollectorRetentionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySQLCollectorRetentionRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifySQLCollectorRetentionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifySQLCollectorRetentionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySQLCollectorRetentionRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifySQLCollectorRetentionRequest
	GetSecurityToken() *string
}

type ModifySQLCollectorRetentionRequest struct {
	// The log retention period that is allowed by the SQL Explorer feature on the instance. Valid values:
	//
	// 	- 30: 30 days
	//
	// 	- 180: 180 days
	//
	// 	- 365: one year
	//
	// 	- 1095: three years
	//
	// 	- 1825: five years
	//
	// This parameter is required.
	//
	// example:
	//
	// 365
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to obtain the resource group ID.
	//
	// example:
	//
	// rg-acfmyxxxx
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifySQLCollectorRetentionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLCollectorRetentionRequest) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorRetentionRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ModifySQLCollectorRetentionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySQLCollectorRetentionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySQLCollectorRetentionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySQLCollectorRetentionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifySQLCollectorRetentionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySQLCollectorRetentionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySQLCollectorRetentionRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifySQLCollectorRetentionRequest) SetConfigValue(v string) *ModifySQLCollectorRetentionRequest {
	s.ConfigValue = &v
	return s
}

func (s *ModifySQLCollectorRetentionRequest) SetDBInstanceId(v string) *ModifySQLCollectorRetentionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySQLCollectorRetentionRequest) SetOwnerAccount(v string) *ModifySQLCollectorRetentionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySQLCollectorRetentionRequest) SetOwnerId(v int64) *ModifySQLCollectorRetentionRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySQLCollectorRetentionRequest) SetResourceGroupId(v string) *ModifySQLCollectorRetentionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySQLCollectorRetentionRequest) SetResourceOwnerAccount(v string) *ModifySQLCollectorRetentionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySQLCollectorRetentionRequest) SetResourceOwnerId(v int64) *ModifySQLCollectorRetentionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySQLCollectorRetentionRequest) SetSecurityToken(v string) *ModifySQLCollectorRetentionRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifySQLCollectorRetentionRequest) Validate() error {
	return dara.Validate(s)
}
