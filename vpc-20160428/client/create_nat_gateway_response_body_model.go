// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardTableIds(v *CreateNatGatewayResponseBodyForwardTableIds) *CreateNatGatewayResponseBody
	GetForwardTableIds() *CreateNatGatewayResponseBodyForwardTableIds
	SetFullNatTableIds(v *CreateNatGatewayResponseBodyFullNatTableIds) *CreateNatGatewayResponseBody
	GetFullNatTableIds() *CreateNatGatewayResponseBodyFullNatTableIds
	SetNatGatewayId(v string) *CreateNatGatewayResponseBody
	GetNatGatewayId() *string
	SetRequestId(v string) *CreateNatGatewayResponseBody
	GetRequestId() *string
	SetSnatTableIds(v *CreateNatGatewayResponseBodySnatTableIds) *CreateNatGatewayResponseBody
	GetSnatTableIds() *CreateNatGatewayResponseBodySnatTableIds
}

type CreateNatGatewayResponseBody struct {
	// A list of DNAT entries.
	ForwardTableIds *CreateNatGatewayResponseBodyForwardTableIds `json:"ForwardTableIds,omitempty" xml:"ForwardTableIds,omitempty" type:"Struct"`
	// A list of FULLNAT entries.
	FullNatTableIds *CreateNatGatewayResponseBodyFullNatTableIds `json:"FullNatTableIds,omitempty" xml:"FullNatTableIds,omitempty" type:"Struct"`
	// The ID of the NAT gateway.
	//
	// example:
	//
	// ngw-112za33e4****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2315DEB7-5E92-423A-91F7-4C1EC9AD97C3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A list of SNAT entries.
	SnatTableIds *CreateNatGatewayResponseBodySnatTableIds `json:"SnatTableIds,omitempty" xml:"SnatTableIds,omitempty" type:"Struct"`
}

func (s CreateNatGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponseBody) GetForwardTableIds() *CreateNatGatewayResponseBodyForwardTableIds {
	return s.ForwardTableIds
}

func (s *CreateNatGatewayResponseBody) GetFullNatTableIds() *CreateNatGatewayResponseBodyFullNatTableIds {
	return s.FullNatTableIds
}

func (s *CreateNatGatewayResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateNatGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatGatewayResponseBody) GetSnatTableIds() *CreateNatGatewayResponseBodySnatTableIds {
	return s.SnatTableIds
}

func (s *CreateNatGatewayResponseBody) SetForwardTableIds(v *CreateNatGatewayResponseBodyForwardTableIds) *CreateNatGatewayResponseBody {
	s.ForwardTableIds = v
	return s
}

func (s *CreateNatGatewayResponseBody) SetFullNatTableIds(v *CreateNatGatewayResponseBodyFullNatTableIds) *CreateNatGatewayResponseBody {
	s.FullNatTableIds = v
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

func (s *CreateNatGatewayResponseBody) SetSnatTableIds(v *CreateNatGatewayResponseBodySnatTableIds) *CreateNatGatewayResponseBody {
	s.SnatTableIds = v
	return s
}

func (s *CreateNatGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateNatGatewayResponseBodyForwardTableIds struct {
	ForwardTableId []*string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty" type:"Repeated"`
}

func (s CreateNatGatewayResponseBodyForwardTableIds) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponseBodyForwardTableIds) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponseBodyForwardTableIds) GetForwardTableId() []*string {
	return s.ForwardTableId
}

func (s *CreateNatGatewayResponseBodyForwardTableIds) SetForwardTableId(v []*string) *CreateNatGatewayResponseBodyForwardTableIds {
	s.ForwardTableId = v
	return s
}

func (s *CreateNatGatewayResponseBodyForwardTableIds) Validate() error {
	return dara.Validate(s)
}

type CreateNatGatewayResponseBodyFullNatTableIds struct {
	FullNatTableId []*string `json:"FullNatTableId,omitempty" xml:"FullNatTableId,omitempty" type:"Repeated"`
}

func (s CreateNatGatewayResponseBodyFullNatTableIds) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponseBodyFullNatTableIds) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponseBodyFullNatTableIds) GetFullNatTableId() []*string {
	return s.FullNatTableId
}

func (s *CreateNatGatewayResponseBodyFullNatTableIds) SetFullNatTableId(v []*string) *CreateNatGatewayResponseBodyFullNatTableIds {
	s.FullNatTableId = v
	return s
}

func (s *CreateNatGatewayResponseBodyFullNatTableIds) Validate() error {
	return dara.Validate(s)
}

type CreateNatGatewayResponseBodySnatTableIds struct {
	SnatTableId []*string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty" type:"Repeated"`
}

func (s CreateNatGatewayResponseBodySnatTableIds) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponseBodySnatTableIds) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponseBodySnatTableIds) GetSnatTableId() []*string {
	return s.SnatTableId
}

func (s *CreateNatGatewayResponseBodySnatTableIds) SetSnatTableId(v []*string) *CreateNatGatewayResponseBodySnatTableIds {
	s.SnatTableId = v
	return s
}

func (s *CreateNatGatewayResponseBodySnatTableIds) Validate() error {
	return dara.Validate(s)
}
