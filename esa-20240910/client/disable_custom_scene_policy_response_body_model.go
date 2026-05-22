// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCustomScenePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *DisableCustomScenePolicyResponseBody
	GetPolicyId() *int64
	SetRequestId(v string) *DisableCustomScenePolicyResponseBody
	GetRequestId() *string
}

type DisableCustomScenePolicyResponseBody struct {
	// The ID of the disabled policy.
	//
	// example:
	//
	// 100001
	PolicyId *int64 `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableCustomScenePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableCustomScenePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DisableCustomScenePolicyResponseBody) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DisableCustomScenePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableCustomScenePolicyResponseBody) SetPolicyId(v int64) *DisableCustomScenePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *DisableCustomScenePolicyResponseBody) SetRequestId(v string) *DisableCustomScenePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableCustomScenePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
