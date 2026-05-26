// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolicySetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreatePolicySetRequest
	GetDescription() *string
	SetPolicySetName(v string) *CreatePolicySetRequest
	GetPolicySetName() *string
}

type CreatePolicySetRequest struct {
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s CreatePolicySetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePolicySetRequest) GoString() string {
	return s.String()
}

func (s *CreatePolicySetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePolicySetRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *CreatePolicySetRequest) SetDescription(v string) *CreatePolicySetRequest {
	s.Description = &v
	return s
}

func (s *CreatePolicySetRequest) SetPolicySetName(v string) *CreatePolicySetRequest {
	s.PolicySetName = &v
	return s
}

func (s *CreatePolicySetRequest) Validate() error {
	return dara.Validate(s)
}
