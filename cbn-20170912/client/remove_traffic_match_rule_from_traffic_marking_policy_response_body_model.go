// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody
	GetRequestId() *string
}

type RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6DF9A765-BCD2-5C7E-8C32-C35C8A361A39
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody) SetRequestId(v string) *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveTrafficMatchRuleFromTrafficMarkingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
