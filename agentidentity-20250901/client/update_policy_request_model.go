// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v *Definition) *UpdatePolicyRequest
	GetDefinition() *Definition
	SetDescription(v string) *UpdatePolicyRequest
	GetDescription() *string
	SetPolicyName(v string) *UpdatePolicyRequest
	GetPolicyName() *string
	SetPolicySetName(v string) *UpdatePolicyRequest
	GetPolicySetName() *string
}

type UpdatePolicyRequest struct {
	Definition *Definition `json:"Definition,omitempty" xml:"Definition,omitempty"`
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

func (s UpdatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyRequest) GetDefinition() *Definition {
	return s.Definition
}

func (s *UpdatePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdatePolicyRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *UpdatePolicyRequest) SetDefinition(v *Definition) *UpdatePolicyRequest {
	s.Definition = v
	return s
}

func (s *UpdatePolicyRequest) SetDescription(v string) *UpdatePolicyRequest {
	s.Description = &v
	return s
}

func (s *UpdatePolicyRequest) SetPolicyName(v string) *UpdatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyRequest) SetPolicySetName(v string) *UpdatePolicyRequest {
	s.PolicySetName = &v
	return s
}

func (s *UpdatePolicyRequest) Validate() error {
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
