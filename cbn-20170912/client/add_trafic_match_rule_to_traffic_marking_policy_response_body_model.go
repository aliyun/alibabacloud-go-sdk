// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTraficMatchRuleToTrafficMarkingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody
	GetRequestId() *string
}

type AddTraficMatchRuleToTrafficMarkingPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0876E54E-3E36-5C31-89F0-9EE8A9266F9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) SetRequestId(v string) *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTraficMatchRuleToTrafficMarkingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
