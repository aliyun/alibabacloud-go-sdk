// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyGatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEgressProxyGatewayAttributeResponseBody
	GetRequestId() *string
}

type UpdateEgressProxyGatewayAttributeResponseBody struct {
	// example:
	//
	// 01A05AC2-567F-5BB5-8F2D-436C4BE402E3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEgressProxyGatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyGatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyGatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEgressProxyGatewayAttributeResponseBody) SetRequestId(v string) *UpdateEgressProxyGatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEgressProxyGatewayAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
