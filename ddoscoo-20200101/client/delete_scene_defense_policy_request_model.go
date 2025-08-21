// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneDefensePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v string) *DeleteSceneDefensePolicyRequest
	GetPolicyId() *string
}

type DeleteSceneDefensePolicyRequest struct {
	// The ID of the policy that you want to delete.
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

func (s DeleteSceneDefensePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneDefensePolicyRequest) GoString() string {
	return s.String()
}

func (s *DeleteSceneDefensePolicyRequest) GetPolicyId() *string {
	return s.PolicyId
}

func (s *DeleteSceneDefensePolicyRequest) SetPolicyId(v string) *DeleteSceneDefensePolicyRequest {
	s.PolicyId = &v
	return s
}

func (s *DeleteSceneDefensePolicyRequest) Validate() error {
	return dara.Validate(s)
}
