// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMarkingPolicyAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTrafficMarkingPolicyAttributeResponseBody
	GetRequestId() *string
}

type UpdateTrafficMarkingPolicyAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 60BB11B2-7BF4-54DC-BCC9-F706E1EB02AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficMarkingPolicyAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMarkingPolicyAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMarkingPolicyAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrafficMarkingPolicyAttributeResponseBody) SetRequestId(v string) *UpdateTrafficMarkingPolicyAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrafficMarkingPolicyAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
