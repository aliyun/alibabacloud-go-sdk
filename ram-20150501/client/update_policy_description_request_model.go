// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolicyDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewDescription(v string) *UpdatePolicyDescriptionRequest
	GetNewDescription() *string
	SetPolicyName(v string) *UpdatePolicyDescriptionRequest
	GetPolicyName() *string
}

type UpdatePolicyDescriptionRequest struct {
	// The description of the policy.
	//
	// The description must be 1 to 1,024 characters in length.
	//
	// example:
	//
	// This is a test policy.
	NewDescription *string `json:"NewDescription,omitempty" xml:"NewDescription,omitempty"`
	// The name of the policy.
	//
	// example:
	//
	// TestPolicy
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s UpdatePolicyDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolicyDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdatePolicyDescriptionRequest) GetNewDescription() *string {
	return s.NewDescription
}

func (s *UpdatePolicyDescriptionRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *UpdatePolicyDescriptionRequest) SetNewDescription(v string) *UpdatePolicyDescriptionRequest {
	s.NewDescription = &v
	return s
}

func (s *UpdatePolicyDescriptionRequest) SetPolicyName(v string) *UpdatePolicyDescriptionRequest {
	s.PolicyName = &v
	return s
}

func (s *UpdatePolicyDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
