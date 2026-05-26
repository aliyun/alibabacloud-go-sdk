// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicySetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdatePolicySetRequest
	GetDescription() *string
	SetPolicySetName(v string) *UpdatePolicySetRequest
	GetPolicySetName() *string
}

type UpdatePolicySetRequest struct {
	// example:
	//
	// example description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s UpdatePolicySetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicySetRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicySetRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdatePolicySetRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *UpdatePolicySetRequest) SetDescription(v string) *UpdatePolicySetRequest {
	s.Description = &v
	return s
}

func (s *UpdatePolicySetRequest) SetPolicySetName(v string) *UpdatePolicySetRequest {
	s.PolicySetName = &v
	return s
}

func (s *UpdatePolicySetRequest) Validate() error {
	return dara.Validate(s)
}
