// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicyToUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *AttachPolicyToUserRequest
	GetPolicyName() *string
	SetPolicyType(v string) *AttachPolicyToUserRequest
	GetPolicyType() *string
	SetResourceGroupId(v string) *AttachPolicyToUserRequest
	GetResourceGroupId() *string
	SetUserName(v string) *AttachPolicyToUserRequest
	GetUserName() *string
}

type AttachPolicyToUserRequest struct {
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
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s AttachPolicyToUserRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicyToUserRequest) GoString() string {
	return s.String()
}

func (s *AttachPolicyToUserRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *AttachPolicyToUserRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *AttachPolicyToUserRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *AttachPolicyToUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *AttachPolicyToUserRequest) SetPolicyName(v string) *AttachPolicyToUserRequest {
	s.PolicyName = &v
	return s
}

func (s *AttachPolicyToUserRequest) SetPolicyType(v string) *AttachPolicyToUserRequest {
	s.PolicyType = &v
	return s
}

func (s *AttachPolicyToUserRequest) SetResourceGroupId(v string) *AttachPolicyToUserRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *AttachPolicyToUserRequest) SetUserName(v string) *AttachPolicyToUserRequest {
	s.UserName = &v
	return s
}

func (s *AttachPolicyToUserRequest) Validate() error {
	return dara.Validate(s)
}
