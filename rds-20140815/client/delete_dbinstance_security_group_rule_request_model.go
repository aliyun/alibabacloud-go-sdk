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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC******
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/2628785.html) operation to query the IDs of instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp15i4hn07r******
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the security group rule. You can call the [DescribeDBInstanceSecurityGroupRule](https://help.aliyun.com/document_detail/2834044.html) to obtain the ID of the security group rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// sgr-2ze17u******
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
