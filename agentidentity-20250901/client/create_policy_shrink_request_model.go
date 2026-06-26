// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinitionShrink(v string) *CreatePolicyShrinkRequest
	GetDefinitionShrink() *string
	SetDescription(v string) *CreatePolicyShrinkRequest
	GetDescription() *string
	SetPolicyName(v string) *CreatePolicyShrinkRequest
	GetPolicyName() *string
	SetPolicySetName(v string) *CreatePolicyShrinkRequest
	GetPolicySetName() *string
}

type CreatePolicyShrinkRequest struct {
	DefinitionShrink *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	Description      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	PolicyName       *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
	PolicySetName    *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s CreatePolicyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyShrinkRequest) GetDefinitionShrink() *string {
	return s.DefinitionShrink
}

func (s *CreatePolicyShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyShrinkRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyShrinkRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *CreatePolicyShrinkRequest) SetDefinitionShrink(v string) *CreatePolicyShrinkRequest {
	s.DefinitionShrink = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetDescription(v string) *CreatePolicyShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPolicyName(v string) *CreatePolicyShrinkRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyShrinkRequest) SetPolicySetName(v string) *CreatePolicyShrinkRequest {
	s.PolicySetName = &v
	return s
}

func (s *CreatePolicyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
