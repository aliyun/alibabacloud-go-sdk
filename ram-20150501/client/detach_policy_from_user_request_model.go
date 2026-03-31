// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicyFromUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *DetachPolicyFromUserRequest
	GetPolicyName() *string
	SetPolicyType(v string) *DetachPolicyFromUserRequest
	GetPolicyType() *string
	SetResourceGroupId(v string) *DetachPolicyFromUserRequest
	GetResourceGroupId() *string
	SetUserName(v string) *DetachPolicyFromUserRequest
	GetUserName() *string
}

type DetachPolicyFromUserRequest struct {
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

func (s DetachPolicyFromUserRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicyFromUserRequest) GoString() string {
	return s.String()
}

func (s *DetachPolicyFromUserRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DetachPolicyFromUserRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DetachPolicyFromUserRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DetachPolicyFromUserRequest) GetUserName() *string {
	return s.UserName
}

func (s *DetachPolicyFromUserRequest) SetPolicyName(v string) *DetachPolicyFromUserRequest {
	s.PolicyName = &v
	return s
}

func (s *DetachPolicyFromUserRequest) SetPolicyType(v string) *DetachPolicyFromUserRequest {
	s.PolicyType = &v
	return s
}

func (s *DetachPolicyFromUserRequest) SetResourceGroupId(v string) *DetachPolicyFromUserRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DetachPolicyFromUserRequest) SetUserName(v string) *DetachPolicyFromUserRequest {
	s.UserName = &v
	return s
}

func (s *DetachPolicyFromUserRequest) Validate() error {
	return dara.Validate(s)
}
