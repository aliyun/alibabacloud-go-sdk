// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyFromGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *DetachPolicyFromGroupRequest
	GetGroupName() *string
	SetPolicyName(v string) *DetachPolicyFromGroupRequest
	GetPolicyName() *string
	SetPolicyType(v string) *DetachPolicyFromGroupRequest
	GetPolicyType() *string
	SetResourceGroupId(v string) *DetachPolicyFromGroupRequest
	GetResourceGroupId() *string
}

type DetachPolicyFromGroupRequest struct {
	// The name of the RAM user group.
	//
	// example:
	//
	// dev
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
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
}

func (s DetachPolicyFromGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyFromGroupRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DetachPolicyFromGroupRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DetachPolicyFromGroupRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetachPolicyFromGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DetachPolicyFromGroupRequest) SetGroupName(v string) *DetachPolicyFromGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DetachPolicyFromGroupRequest) SetPolicyName(v string) *DetachPolicyFromGroupRequest {
	s.PolicyName = &v
	return s
}

func (s *DetachPolicyFromGroupRequest) SetPolicyType(v string) *DetachPolicyFromGroupRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachPolicyFromGroupRequest) SetResourceGroupId(v string) *DetachPolicyFromGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachPolicyFromGroupRequest) Validate() error {
	return dara.Validate(s)
}
