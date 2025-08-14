// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody
	GetRequestId() *string
}

type ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 13526224-5780-4426-8ADF-BC8B08700F23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody) SetRequestId(v string) *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTrafficMatchRuleToTrafficMarkingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
