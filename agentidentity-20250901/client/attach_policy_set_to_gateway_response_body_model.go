// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPolicySetToGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachPolicySetToGatewayResponseBody
	GetRequestId() *string
}

type AttachPolicySetToGatewayResponseBody struct {
	// example:
	//
	// 2A48EB1D-D645-5758-91AF-EDF8E36E257B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachPolicySetToGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachPolicySetToGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *AttachPolicySetToGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachPolicySetToGatewayResponseBody) SetRequestId(v string) *AttachPolicySetToGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachPolicySetToGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
