// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistrationPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *GetRegistrationPolicyRequest
	GetPolicyId() *string
}

type GetRegistrationPolicyRequest struct {
	// The ID of the device registration policy. Valid values are obtained from the following sources:
	//
	// - [ListRegistrationPolicies](~~ListRegistrationPolicies~~): Queries device registration policies in batches.
	//
	// - [GetRegistrationPolicy](~~GetRegistrationPolicy~~): Queries the details of a device registration policy.
	//
	// - [CreateRegistrationPolicy](~~CreateRegistrationPolicy~~): Creates a device registration policy.
	//
	// - [UpdateRegistrationPolicy](~~UpdateRegistrationPolicy~~): Updates a device registration policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// reg-policy-dcbfd33cb004****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s GetRegistrationPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegistrationPolicyRequest) GoString() string {
	return s.String()
}

func (s *GetRegistrationPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *GetRegistrationPolicyRequest) SetPolicyId(v string) *GetRegistrationPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *GetRegistrationPolicyRequest) Validate() error {
	return dara.Validate(s)
}
