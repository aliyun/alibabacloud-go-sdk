// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIstioGatewayRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateIstioGatewayRoutesResponseBody
	GetRequestId() *string
}

type CreateIstioGatewayRoutesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIstioGatewayRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIstioGatewayRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIstioGatewayRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIstioGatewayRoutesResponseBody) SetRequestId(v string) *CreateIstioGatewayRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIstioGatewayRoutesResponseBody) Validate() error {
	return dara.Validate(s)
}
