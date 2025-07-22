// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardTableIds(v []*string) *CreateNatGatewayResponseBody
	GetForwardTableIds() []*string
	SetNatGatewayId(v string) *CreateNatGatewayResponseBody
	GetNatGatewayId() *string
	SetRequestId(v string) *CreateNatGatewayResponseBody
	GetRequestId() *string
	SetSnatTableIds(v []*string) *CreateNatGatewayResponseBody
	GetSnatTableIds() []*string
}

type CreateNatGatewayResponseBody struct {
	ForwardTableIds []*string `json:"ForwardTableIds,omitempty" xml:"ForwardTableIds,omitempty" type:"Repeated"`
	NatGatewayId    *string   `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	RequestId       *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SnatTableIds    []*string `json:"SnatTableIds,omitempty" xml:"SnatTableIds,omitempty" type:"Repeated"`
}

func (s CreateNatGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponseBody) GetForwardTableIds() []*string {
	return s.ForwardTableIds
}

func (s *CreateNatGatewayResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateNatGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatGatewayResponseBody) GetSnatTableIds() []*string {
	return s.SnatTableIds
}

func (s *CreateNatGatewayResponseBody) SetForwardTableIds(v []*string) *CreateNatGatewayResponseBody {
	s.ForwardTableIds = v
	return s
}

func (s *CreateNatGatewayResponseBody) SetNatGatewayId(v string) *CreateNatGatewayResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *CreateNatGatewayResponseBody) SetRequestId(v string) *CreateNatGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNatGatewayResponseBody) SetSnatTableIds(v []*string) *CreateNatGatewayResponseBody {
	s.SnatTableIds = v
	return s
}

func (s *CreateNatGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
