// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomScenePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *DeleteCustomScenePolicyRequest
	GetPolicyId() *int64
}

type DeleteCustomScenePolicyRequest struct {
	// This parameter is required.
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DeleteCustomScenePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomScenePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomScenePolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DeleteCustomScenePolicyRequest) SetPolicyId(v int64) *DeleteCustomScenePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteCustomScenePolicyRequest) Validate() error {
	return dara.Validate(s)
}
