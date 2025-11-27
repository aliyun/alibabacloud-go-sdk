// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountCheckPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountCheckPolicyRequest
	GetAccountName() *string
	SetCheckPolicy(v bool) *ModifyAccountCheckPolicyRequest
	GetCheckPolicy() *bool
	SetClientToken(v string) *ModifyAccountCheckPolicyRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyAccountCheckPolicyRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyAccountCheckPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAccountCheckPolicyRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifyAccountCheckPolicyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifyAccountCheckPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAccountCheckPolicyRequest
	GetResourceOwnerId() *int64
}

type ModifyAccountCheckPolicyRequest struct {
	// The account username.
	//
	// This parameter is required.
	//
	// example:
	//
	// DatabaseTest
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Specifies whether to apply the password policy
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	CheckPolicy *bool `json:"CheckPolicy,omitempty" xml:"CheckPolicy,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters. If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOC****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. For more information about resource groups, see related documentation.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyAccountCheckPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountCheckPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountCheckPolicyRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountCheckPolicyRequest) GetCheckPolicy() *bool {
	return s.CheckPolicy
}

func (s *ModifyAccountCheckPolicyRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAccountCheckPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAccountCheckPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAccountCheckPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAccountCheckPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyAccountCheckPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAccountCheckPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAccountCheckPolicyRequest) SetAccountName(v string) *ModifyAccountCheckPolicyRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountCheckPolicyRequest) SetCheckPolicy(v bool) *ModifyAccountCheckPolicyRequest {
	s.CheckPolicy = &v
	return s
}

func (s *ModifyAccountCheckPolicyRequest) SetClientToken(v string) *ModifyAccountCheckPolicyRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAccountCheckPolicyRequest) SetDBInstanceId(v string) *ModifyAccountCheckPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountCheckPolicyRequest) SetOwnerAccount(v string) *ModifyAccountCheckPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAccountCheckPolicyRequest) SetOwnerId(v int64) *ModifyAccountCheckPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAccountCheckPolicyRequest) SetResourceGroupId(v string) *ModifyAccountCheckPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyAccountCheckPolicyRequest) SetResourceOwnerAccount(v string) *ModifyAccountCheckPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAccountCheckPolicyRequest) SetResourceOwnerId(v int64) *ModifyAccountCheckPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAccountCheckPolicyRequest) Validate() error {
	return dara.Validate(s)
}
