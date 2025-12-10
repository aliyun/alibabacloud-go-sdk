// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyName(v string) *DeletePolicyRequest
	GetPolicyName() *string
}

type DeletePolicyRequest struct {
	// The name of the permission policy.
	//
	// The name must be 1 to 128 characters in length and can contain letters, digits, and hyphen (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// OSS-Administrator
	PolicyName *string `json:"PolicyName,omitempty" xml:"PolicyName,omitempty"`
}

func (s DeletePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePolicyRequest) GetPolicyName() *string {
	return s.PolicyName
}

func (s *DeletePolicyRequest) SetPolicyName(v string) *DeletePolicyRequest {
	s.PolicyName = &v
	return s
}

func (s *DeletePolicyRequest) Validate() error {
	return dara.Validate(s)
}
