// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutManagedScalingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PutManagedScalingPolicyResponseBody
	GetRequestId() *string
}

type PutManagedScalingPolicyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PutManagedScalingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutManagedScalingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *PutManagedScalingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutManagedScalingPolicyResponseBody) SetRequestId(v string) *PutManagedScalingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutManagedScalingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
