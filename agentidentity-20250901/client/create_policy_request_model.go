// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefinition(v *Definition) *CreatePolicyRequest
	GetDefinition() *Definition
	SetDescription(v string) *CreatePolicyRequest
	GetDescription() *string
	SetPolicyName(v string) *CreatePolicyRequest
	GetPolicyName() *string
	SetPolicySetName(v string) *CreatePolicyRequest
	GetPolicySetName() *string
}

type CreatePolicyRequest struct {
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

func (s CreatePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicyRequest) GetDefinition() *Definition {
	return s.Definition
}

func (s *CreatePolicyRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *CreatePolicyRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *CreatePolicyRequest) SetDefinition(v *Definition) *CreatePolicyRequest {
	s.Definition = v
	return s
}

func (s *CreatePolicyRequest) SetDescription(v string) *CreatePolicyRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicyName(v string) *CreatePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyRequest) SetPolicySetName(v string) *CreatePolicyRequest {
	s.PolicySetName = &v
	return s
}

func (s *CreatePolicyRequest) Validate() error {
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			return err
		}
	}
	return nil
}
