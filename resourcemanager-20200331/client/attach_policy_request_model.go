// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *AttachPolicyRequest
	GetPolicyName() *string
	SetPolicyType(v string) *AttachPolicyRequest
	GetPolicyType() *string
	SetPrincipalName(v string) *AttachPolicyRequest
	GetPrincipalName() *string
	SetPrincipalType(v string) *AttachPolicyRequest
	GetPrincipalType() *string
	SetResourceGroupId(v string) *AttachPolicyRequest
	GetResourceGroupId() *string
}

type AttachPolicyRequest struct {
	// Policy Name of the access policy.
	//
	// Policy Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// AdministratorAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the access policy. Valid values:
	//
	// - Custom: custom access policy.
	//
	// - System: system access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the authorization object. The format varies based on the object type:
	//
	// - Resource Access Management (RAM) user: <UserName>@<AccountAlias>.onaliyun.com, where <UserName> is the RAM username and <AccountAlias> is the account alias.
	//
	// - RAM user group: <GroupName>@group.<AccountAlias>.onaliyun.com, where <GroupName> is the RAM user group name and <AccountAlias> is the account alias.
	//
	// - RAM role: <RoleName>@role.<AccountAlias>.onaliyunservice.com, where <RoleName> is the RAM role name and <AccountAlias> is the account alias.
	//
	// > The account alias is part of the default domain name. Invoke [GetDefaultDomain](https://help.aliyun.com/document_detail/186720.html) to obtain the account alias.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice@demo.onaliyun.com
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The authorization object type. Valid values:
	//
	// - IMSUser: Resource Access Management (RAM) user.
	//
	// - IMSGroup: RAM user group.
	//
	// - ServiceRole: RAM role.
	//
	// This parameter is required.
	//
	// example:
	//
	// IMSUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The authorization scope. Valid values:
	//
	// - Resource group ID: The permissions take effect within the specified resource group.
	//
	// - Alibaba Cloud account ID: The ID of the Alibaba Cloud account to which the principal belongs. The permissions take effect within the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s AttachPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *AttachPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AttachPolicyRequest) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *AttachPolicyRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *AttachPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AttachPolicyRequest) SetPolicyName(v string) *AttachPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyRequest) SetPolicyType(v string) *AttachPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *AttachPolicyRequest) SetPrincipalName(v string) *AttachPolicyRequest {
	s.PrincipalName = &v
	return s
}

func (s *AttachPolicyRequest) SetPrincipalType(v string) *AttachPolicyRequest {
	s.PrincipalType = &v
	return s
}

func (s *AttachPolicyRequest) SetResourceGroupId(v string) *AttachPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AttachPolicyRequest) Validate() error {
	return dara.Validate(s)
}
