// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyToGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *AttachPolicyToGroupRequest
	GetGroupName() *string
	SetPolicyName(v string) *AttachPolicyToGroupRequest
	GetPolicyName() *string
	SetPolicyType(v string) *AttachPolicyToGroupRequest
	GetPolicyType() *string
	SetResourceGroupId(v string) *AttachPolicyToGroupRequest
	GetResourceGroupId() *string
}

type AttachPolicyToGroupRequest struct {
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

func (s AttachPolicyToGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyToGroupRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyToGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *AttachPolicyToGroupRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *AttachPolicyToGroupRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AttachPolicyToGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AttachPolicyToGroupRequest) SetGroupName(v string) *AttachPolicyToGroupRequest {
	s.GroupName = &v
	return s
}

func (s *AttachPolicyToGroupRequest) SetPolicyName(v string) *AttachPolicyToGroupRequest {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyToGroupRequest) SetPolicyType(v string) *AttachPolicyToGroupRequest {
	s.PolicyType = &v
	return s
}

func (s *AttachPolicyToGroupRequest) SetResourceGroupId(v string) *AttachPolicyToGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AttachPolicyToGroupRequest) Validate() error {
	return dara.Validate(s)
}
