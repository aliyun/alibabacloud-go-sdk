// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutAutoScalingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutAutoScalingPolicyResponseBody
	GetRequestId() *string
}

type PutAutoScalingPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutAutoScalingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutAutoScalingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *PutAutoScalingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutAutoScalingPolicyResponseBody) SetRequestId(v string) *PutAutoScalingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutAutoScalingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
