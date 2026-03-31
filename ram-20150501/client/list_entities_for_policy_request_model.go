// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesForPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *ListEntitiesForPolicyRequest
	GetPolicyName() *string
	SetPolicyType(v string) *ListEntitiesForPolicyRequest
	GetPolicyType() *string
}

type ListEntitiesForPolicyRequest struct {
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
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ListEntitiesForPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesForPolicyRequest) GoString() string {
	return s.String()
}

func (s *ListEntitiesForPolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *ListEntitiesForPolicyRequest) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ListEntitiesForPolicyRequest) SetPolicyName(v string) *ListEntitiesForPolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *ListEntitiesForPolicyRequest) SetPolicyType(v string) *ListEntitiesForPolicyRequest {
	s.PolicyType = &v
	return s
}

func (s *ListEntitiesForPolicyRequest) Validate() error {
	return dara.Validate(s)
}
