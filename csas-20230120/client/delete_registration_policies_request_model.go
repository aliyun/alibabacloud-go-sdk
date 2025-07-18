// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistrationPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyIds(v []*string) *DeleteRegistrationPoliciesRequest
	GetPolicyIds() []*string
}

type DeleteRegistrationPoliciesRequest struct {
	// This parameter is required.
	PolicyIds []*string `json:"PolicyIds,omitempty" xml:"PolicyIds,omitempty" type:"Repeated"`
}

func (s DeleteRegistrationPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistrationPoliciesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRegistrationPoliciesRequest) GetPolicyIds() []*string {
	return s.PolicyIds
}

func (s *DeleteRegistrationPoliciesRequest) SetPolicyIds(v []*string) *DeleteRegistrationPoliciesRequest {
	s.PolicyIds = v
	return s
}

func (s *DeleteRegistrationPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
