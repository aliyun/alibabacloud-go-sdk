// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyFromRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *DetachPolicyFromRoleRequest
	GetPolicyName() *string
	SetPolicyType(v string) *DetachPolicyFromRoleRequest
	GetPolicyType() *string
	SetResourceGroupId(v string) *DetachPolicyFromRoleRequest
	GetResourceGroupId() *string
	SetRoleName(v string) *DetachPolicyFromRoleRequest
	GetRoleName() *string
}

type DetachPolicyFromRoleRequest struct {
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

func (s DetachPolicyFromRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyFromRoleRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromRoleRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DetachPolicyFromRoleRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetachPolicyFromRoleRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DetachPolicyFromRoleRequest) GetRoleName() *string {
	return s.RoleName
}

func (s *DetachPolicyFromRoleRequest) SetPolicyName(v string) *DetachPolicyFromRoleRequest {
	s.PolicyName = &v
	return s
}

func (s *DetachPolicyFromRoleRequest) SetPolicyType(v string) *DetachPolicyFromRoleRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachPolicyFromRoleRequest) SetResourceGroupId(v string) *DetachPolicyFromRoleRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachPolicyFromRoleRequest) SetRoleName(v string) *DetachPolicyFromRoleRequest {
	s.RoleName = &v
	return s
}

func (s *DetachPolicyFromRoleRequest) Validate() error {
	return dara.Validate(s)
}
