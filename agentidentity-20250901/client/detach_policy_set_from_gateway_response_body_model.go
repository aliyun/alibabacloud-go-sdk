// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachPolicySetFromGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachPolicySetFromGatewayResponseBody
	GetRequestId() *string
}

type DetachPolicySetFromGatewayResponseBody struct {
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachPolicySetFromGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachPolicySetFromGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DetachPolicySetFromGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachPolicySetFromGatewayResponseBody) SetRequestId(v string) *DetachPolicySetFromGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachPolicySetFromGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
