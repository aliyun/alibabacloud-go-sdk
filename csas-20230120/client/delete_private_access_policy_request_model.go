// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateAccessPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *DeletePrivateAccessPolicyRequest
	GetPolicyId() *string
}

type DeletePrivateAccessPolicyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-policy-867ef4007c8a****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeletePrivateAccessPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateAccessPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeletePrivateAccessPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeletePrivateAccessPolicyRequest) SetPolicyId(v string) *DeletePrivateAccessPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeletePrivateAccessPolicyRequest) Validate() error {
	return dara.Validate(s)
}
