// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatGatewayId(v string) *CreateNatGatewayResponseBody
	GetNatGatewayId() *string
	SetRequestId(v string) *CreateNatGatewayResponseBody
	GetRequestId() *string
}

type CreateNatGatewayResponseBody struct {
	// The ID of the NAT gateway.
	//
	// example:
	//
	// nat-5t7nh1cfm6kxiszlttr38****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNatGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateNatGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatGatewayResponseBody) SetNatGatewayId(v string) *CreateNatGatewayResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *CreateNatGatewayResponseBody) SetRequestId(v string) *CreateNatGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNatGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
