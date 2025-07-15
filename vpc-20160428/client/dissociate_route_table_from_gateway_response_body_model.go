// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateRouteTableFromGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateRouteTableFromGatewayResponseBody
	GetRequestId() *string
}

type DissociateRouteTableFromGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C5644C9A-7480-13B6-AECB-30FF142E3724
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateRouteTableFromGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateRouteTableFromGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateRouteTableFromGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateRouteTableFromGatewayResponseBody) SetRequestId(v string) *DissociateRouteTableFromGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateRouteTableFromGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
