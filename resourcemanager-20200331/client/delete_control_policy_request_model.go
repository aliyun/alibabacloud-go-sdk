// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteControlPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *DeleteControlPolicyRequest
	GetPolicyId() *string
}

type DeleteControlPolicyRequest struct {
	// The ID of the control policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// cp-SImPt8GCEwiq****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeleteControlPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteControlPolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteControlPolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeleteControlPolicyRequest) SetPolicyId(v string) *DeleteControlPolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteControlPolicyRequest) Validate() error {
	return dara.Validate(s)
}
