// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAutoScalingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveAutoScalingPolicyResponseBody
	GetRequestId() *string
}

type RemoveAutoScalingPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveAutoScalingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveAutoScalingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAutoScalingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveAutoScalingPolicyResponseBody) SetRequestId(v string) *RemoveAutoScalingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAutoScalingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
