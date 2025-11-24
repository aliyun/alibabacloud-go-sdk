// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIstioGatewayRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIstioGatewayRoutesResponseBody
	GetRequestId() *string
}

type UpdateIstioGatewayRoutesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIstioGatewayRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIstioGatewayRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIstioGatewayRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIstioGatewayRoutesResponseBody) SetRequestId(v string) *UpdateIstioGatewayRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIstioGatewayRoutesResponseBody) Validate() error {
	return dara.Validate(s)
}
