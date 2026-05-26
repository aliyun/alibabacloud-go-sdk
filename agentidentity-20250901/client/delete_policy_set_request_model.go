// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicySetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicySetName(v string) *DeletePolicySetRequest
	GetPolicySetName() *string
}

type DeletePolicySetRequest struct {
	// example:
	//
	// default-policy-set
	PolicySetName *string `json:"PolicySetName,omitempty" xml:"PolicySetName,omitempty"`
}

func (s DeletePolicySetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicySetRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicySetRequest) GetPolicySetName() *string {
	return s.PolicySetName
}

func (s *DeletePolicySetRequest) SetPolicySetName(v string) *DeletePolicySetRequest {
	s.PolicySetName = &v
	return s
}

func (s *DeletePolicySetRequest) Validate() error {
	return dara.Validate(s)
}
