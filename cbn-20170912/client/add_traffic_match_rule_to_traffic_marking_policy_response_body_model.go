// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTrafficMatchRuleToTrafficMarkingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody
	GetRequestId() *string
}

type AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0876E54E-3E36-5C31-89F0-9EE8A9266F9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody) SetRequestId(v string) *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTrafficMatchRuleToTrafficMarkingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
