// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *DetachPolicyRequest
	GetPolicyName() *string
	SetPolicyType(v string) *DetachPolicyRequest
	GetPolicyType() *string
	SetPrincipalName(v string) *DetachPolicyRequest
	GetPrincipalName() *string
	SetPrincipalType(v string) *DetachPolicyRequest
	GetPrincipalType() *string
	SetResourceGroupId(v string) *DetachPolicyRequest
	GetResourceGroupId() *string
}

type DetachPolicyRequest struct {
	// The name of the permission policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS-Administrator
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
	// Custom
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	// The name of the object to which you want to attach the permission policy.
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
	// The ID of the resource group or the ID of the Alibaba Cloud account to which the resource group belongs.
	//
	// This parameter specifies the resource group or Alibaba Cloud account for which you want to revoke permissions.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DetachPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DetachPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetachPolicyRequest) GetPrincipalName() *string {
	return s.PrincipalName
}

func (s *DetachPolicyRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *DetachPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DetachPolicyRequest) SetPolicyName(v string) *DetachPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *DetachPolicyRequest) SetPolicyType(v string) *DetachPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachPolicyRequest) SetPrincipalName(v string) *DetachPolicyRequest {
	s.PrincipalName = &v
	return s
}

func (s *DetachPolicyRequest) SetPrincipalType(v string) *DetachPolicyRequest {
	s.PrincipalType = &v
	return s
}

func (s *DetachPolicyRequest) SetResourceGroupId(v string) *DetachPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachPolicyRequest) Validate() error {
	return dara.Validate(s)
}
