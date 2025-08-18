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
	// The policy ID, which can be obtained by calling the [DescribeCustomScenePolicies](https://help.aliyun.com/document_detail/2850508.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
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
