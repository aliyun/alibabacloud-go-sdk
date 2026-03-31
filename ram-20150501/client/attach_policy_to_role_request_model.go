// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyToRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *AttachPolicyToRoleRequest
	GetPolicyName() *string
	SetPolicyType(v string) *AttachPolicyToRoleRequest
	GetPolicyType() *string
	SetResourceGroupId(v string) *AttachPolicyToRoleRequest
	GetResourceGroupId() *string
	SetRoleName(v string) *AttachPolicyToRoleRequest
	GetRoleName() *string
}

type AttachPolicyToRoleRequest struct {
	// The name of the policy.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// The type of the policy. Valid values: `System` and `Custom`.
	//
	// example:
	//
	// Custom
	PolicyType      *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The name of the RAM role.
	//
	// example:
	//
	// OSSAdminRole
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s AttachPolicyToRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyToRoleRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyToRoleRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *AttachPolicyToRoleRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AttachPolicyToRoleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AttachPolicyToRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *AttachPolicyToRoleRequest) SetPolicyName(v string) *AttachPolicyToRoleRequest {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyToRoleRequest) SetPolicyType(v string) *AttachPolicyToRoleRequest {
	s.PolicyType = &v
	return s
}

func (s *AttachPolicyToRoleRequest) SetResourceGroupId(v string) *AttachPolicyToRoleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AttachPolicyToRoleRequest) SetRoleName(v string) *AttachPolicyToRoleRequest {
	s.RoleName = &v
	return s
}

func (s *AttachPolicyToRoleRequest) Validate() error {
	return dara.Validate(s)
}
