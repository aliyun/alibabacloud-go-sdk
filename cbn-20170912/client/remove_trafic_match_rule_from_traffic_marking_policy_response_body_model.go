// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody
	GetRequestId() *string
}

type RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6DF9A765-BCD2-5C7E-8C32-C35C8A361A39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) SetRequestId(v string) *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTraficMatchRuleFromTrafficMarkingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
