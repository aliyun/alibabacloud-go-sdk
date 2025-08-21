// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSceneDefensePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *DisableSceneDefensePolicyRequest
	GetPolicyId() *string
}

type DisableSceneDefensePolicyRequest struct {
	// The ID of the policy that you want to disable.
	//
	// > You can call the [DescribeSceneDefensePolicies](https://help.aliyun.com/document_detail/159382.html) operation to query the IDs of all policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// 321a-fd31-df51-****
	PolicyId *string `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
}

func (s DisableSceneDefensePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneDefensePolicyRequest) GoString() string {
	return s.String()
}

func (s *DisableSceneDefensePolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DisableSceneDefensePolicyRequest) SetPolicyId(v string) *DisableSceneDefensePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DisableSceneDefensePolicyRequest) Validate() error {
	return dara.Validate(s)
}
