// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomScenePolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPolicyId(v int64) *DeleteCustomScenePolicyResponseBody
	GetPolicyId() *int64
	SetRequestId(v string) *DeleteCustomScenePolicyResponseBody
	GetRequestId() *string
}

type DeleteCustomScenePolicyResponseBody struct {
	PolicyId  *int64  `json:"PolicyId,omitempty" xml:"PolicyId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomScenePolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomScenePolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomScenePolicyResponseBody) GetPolicyId() *int64 {
	return s.PolicyId
}

func (s *DeleteCustomScenePolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomScenePolicyResponseBody) SetPolicyId(v int64) *DeleteCustomScenePolicyResponseBody {
	s.PolicyId = &v
	return s
}

func (s *DeleteCustomScenePolicyResponseBody) SetRequestId(v string) *DeleteCustomScenePolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomScenePolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
