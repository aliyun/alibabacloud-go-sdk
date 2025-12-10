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
	// The name of the permission policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// AdministratorAccess
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the permission policy. Valid values:
	//
	// 	- Custom
	//
	// 	- System
	//
	// This parameter is required.
	//
	// example:
	//
	// System
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the object to which you want to attach the permission policy.
	//
	// 	- If you want to attach the permission policy to a RAM user, specify the name in the @.onaliyun.com format. indicates the name of the RAM user, and indicates the alias of the Alibaba Cloud account to which the RAM user belongs.
	//
	// 	- If you want to attach the permission policy to a RAM user group, specify the name in the @group..onaliyun.com format. indicates the name of the RAM user group, and indicates the alias of the Alibaba Cloud account to which the RAM user group belongs.
	//
	// 	- If you want to attach the permission policy to a RAM role, specify the name in the @role..onaliyunservice.com format. indicates the name of the RAM role, and indicates the alias of the Alibaba Cloud account to which the RAM role belongs.
	//
	// >  The alias of an Alibaba Cloud account is a part of the default domain name. You can call the [GetDefaultDomain](https://help.aliyun.com/document_detail/186720.html) operation to obtain the alias of an Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice@demo.onaliyun.com
	PrincipalName *string `json:"PrincipalName,omitempty" xml:"PrincipalName,omitempty"`
	// The type of the object to which you want to attach the permission policy. Valid values:
	//
	// 	- IMSUser: RAM user
	//
	// 	- IMSGroup: RAM user group
	//
	// 	- ServiceRole: RAM role
	//
	// This parameter is required.
	//
	// example:
	//
	// IMSUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
	// The effective scope of the permission policy. Valid values:
	//
	// 	- ID of a resource group: indicates that the permission policy takes effect for the resources in the resource group.
	//
	// 	- ID of the Alibaba Cloud account to which the authorized object belongs: indicates that the permission policy takes effect for the resources within the Alibaba Cloud account.
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
