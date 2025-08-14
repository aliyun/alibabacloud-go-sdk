// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMarkingPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTrafficMarkingPolicyResponseBody
	GetRequestId() *string
	SetTrafficMarkingPolicyId(v string) *CreateTrafficMarkingPolicyResponseBody
	GetTrafficMarkingPolicyId() *string
}

type CreateTrafficMarkingPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0876E54E-3E36-5C31-89F0-9EE8A9266F9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the traffic marking policy.
	//
	// example:
	//
	// tm-u9nxup5kww5po8****
	TrafficMarkingPolicyId *string `json:"TrafficMarkingPolicyId,omitempty" xml:"TrafficMarkingPolicyId,omitempty"`
}

func (s CreateTrafficMarkingPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMarkingPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTrafficMarkingPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTrafficMarkingPolicyResponseBody) GetTrafficMarkingPolicyId() *string {
	return s.TrafficMarkingPolicyId
}

func (s *CreateTrafficMarkingPolicyResponseBody) SetRequestId(v string) *CreateTrafficMarkingPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTrafficMarkingPolicyResponseBody) SetTrafficMarkingPolicyId(v string) *CreateTrafficMarkingPolicyResponseBody {
	s.TrafficMarkingPolicyId = &v
	return s
}

func (s *CreateTrafficMarkingPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
