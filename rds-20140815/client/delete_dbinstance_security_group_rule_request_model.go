// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBInstanceSecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteDBInstanceSecurityGroupRuleRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *DeleteDBInstanceSecurityGroupRuleRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DeleteDBInstanceSecurityGroupRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v string) *DeleteDBInstanceSecurityGroupRuleRequest
	GetOwnerId() *string
	SetResourceGroupId(v string) *DeleteDBInstanceSecurityGroupRuleRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteDBInstanceSecurityGroupRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDBInstanceSecurityGroupRuleRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupRuleIds(v string) *DeleteDBInstanceSecurityGroupRuleRequest
	GetSecurityGroupRuleIds() *string
}

type DeleteDBInstanceSecurityGroupRuleRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SecurityGroupRuleIds *string `json:"SecurityGroupRuleIds,omitempty" xml:"SecurityGroupRuleIds,omitempty"`
}

func (s DeleteDBInstanceSecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBInstanceSecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) GetSecurityGroupRuleIds() *string {
	return s.SecurityGroupRuleIds
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) SetClientToken(v string) *DeleteDBInstanceSecurityGroupRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) SetDBInstanceId(v string) *DeleteDBInstanceSecurityGroupRuleRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) SetOwnerAccount(v string) *DeleteDBInstanceSecurityGroupRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) SetOwnerId(v string) *DeleteDBInstanceSecurityGroupRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) SetResourceGroupId(v string) *DeleteDBInstanceSecurityGroupRuleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) SetResourceOwnerAccount(v string) *DeleteDBInstanceSecurityGroupRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) SetResourceOwnerId(v int64) *DeleteDBInstanceSecurityGroupRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) SetSecurityGroupRuleIds(v string) *DeleteDBInstanceSecurityGroupRuleRequest {
	s.SecurityGroupRuleIds = &v
	return s
}

func (s *DeleteDBInstanceSecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}
