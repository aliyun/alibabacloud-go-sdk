// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomScenePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *DisableCustomScenePolicyRequest
	GetPolicyId() *int64
}

type DisableCustomScenePolicyRequest struct {
	// This parameter is required.
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DisableCustomScenePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomScenePolicyRequest) GoString() string {
	return s.String()
}

func (s *DisableCustomScenePolicyRequest) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DisableCustomScenePolicyRequest) SetPolicyId(v int64) *DisableCustomScenePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DisableCustomScenePolicyRequest) Validate() error {
	return dara.Validate(s)
}
