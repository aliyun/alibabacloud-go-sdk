// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinitionShrink(v string) *UpdatePolicyShrinkRequest
	GetDefinitionShrink() *string
	SetDescription(v string) *UpdatePolicyShrinkRequest
	GetDescription() *string
	SetPolicyName(v string) *UpdatePolicyShrinkRequest
	GetPolicyName() *string
	SetPolicySetName(v string) *UpdatePolicyShrinkRequest
	GetPolicySetName() *string
}

type UpdatePolicyShrinkRequest struct {
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// rate-limit-policy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s UpdatePolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyShrinkRequest) GetDefinitionShrink() *string {
	return s.DefinitionShrink
}

func (s *UpdatePolicyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdatePolicyShrinkRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *UpdatePolicyShrinkRequest) SetDefinitionShrink(v string) *UpdatePolicyShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetDescription(v string) *UpdatePolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetPolicyName(v string) *UpdatePolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) SetPolicySetName(v string) *UpdatePolicyShrinkRequest {
	s.PolicySetName = &v
	return s
}

func (s *UpdatePolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
