// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountSecurityPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyAccountSecurityPolicyRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyAccountSecurityPolicyRequest
	GetDBInstanceId() *string
	SetGroupPolicy(v string) *ModifyAccountSecurityPolicyRequest
	GetGroupPolicy() *string
	SetOwnerAccount(v string) *ModifyAccountSecurityPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountSecurityPolicyRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyAccountSecurityPolicyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyAccountSecurityPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountSecurityPolicyRequest
	GetResourceOwnerId() *int64
}

type ModifyAccountSecurityPolicyRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/2628785.html) operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-bp1ibu****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The custom password policy for the account of the ApsaraDB RDS for SQL Server instance. The following policies are supported:
	//
	// 	- `{"account security policy": {"MaximumPasswordAge": Specify the maximum password age}}`: You can configure only the maximum password age. After the maximum password age is reached, you must change the password.
	//
	// 	- `{"accountSecurityPolicy": {"MaximumPasswordAge": Specify the minimum password age}}`: You can configure only the minimum password age. During the specified period, you cannot change the password.
	//
	// 	- `{"accountSecurityPolicy": {"MaximumPasswordAge": Specify the maximum password age, "MinimumPasswordAge": Specify the minimum password age}}`: You can configure the maximum and minimum password age at the same time.
	//
	// >  The minimum password age cannot be greater than the maximum password age. Valid values for the minimum password age: 0 to 998. Valid values for the maximum password age: 0 to 999.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"accountSecurityPolicy": {"MaximumPasswordAge": 30, "MinimumPasswordAge": 20}}
	GroupPolicy  *string `json:"GroupPolicy,omitempty" xml:"GroupPolicy,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAccountSecurityPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountSecurityPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountSecurityPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAccountSecurityPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAccountSecurityPolicyRequest) GetGroupPolicy() *string {
	return s.GroupPolicy
}

func (s *ModifyAccountSecurityPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountSecurityPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountSecurityPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyAccountSecurityPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountSecurityPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountSecurityPolicyRequest) SetClientToken(v string) *ModifyAccountSecurityPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAccountSecurityPolicyRequest) SetDBInstanceId(v string) *ModifyAccountSecurityPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountSecurityPolicyRequest) SetGroupPolicy(v string) *ModifyAccountSecurityPolicyRequest {
	s.GroupPolicy = &v
	return s
}

func (s *ModifyAccountSecurityPolicyRequest) SetOwnerAccount(v string) *ModifyAccountSecurityPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountSecurityPolicyRequest) SetOwnerId(v int64) *ModifyAccountSecurityPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountSecurityPolicyRequest) SetResourceGroupId(v string) *ModifyAccountSecurityPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyAccountSecurityPolicyRequest) SetResourceOwnerAccount(v string) *ModifyAccountSecurityPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountSecurityPolicyRequest) SetResourceOwnerId(v int64) *ModifyAccountSecurityPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountSecurityPolicyRequest) Validate() error {
	return dara.Validate(s)
}
