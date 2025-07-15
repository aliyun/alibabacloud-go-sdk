// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatIpCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatIpCidrId(v string) *CreateNatIpCidrResponseBody
	GetNatIpCidrId() *string
	SetRequestId(v string) *CreateNatIpCidrResponseBody
	GetRequestId() *string
}

type CreateNatIpCidrResponseBody struct {
	// The ID of the NAT CIDR block.
	//
	// example:
	//
	// vpcnatcidr-gw8lhqtvdn4qnea****
	NatIpCidrId *string `json:"NatIpCidrId,omitempty" xml:"NatIpCidrId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7021BEB1-210F-48A9-AB82-BE9A9110BB89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNatIpCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNatIpCidrResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatIpCidrResponseBody) GetNatIpCidrId() *string {
	return s.NatIpCidrId
}

func (s *CreateNatIpCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatIpCidrResponseBody) SetNatIpCidrId(v string) *CreateNatIpCidrResponseBody {
	s.NatIpCidrId = &v
	return s
}

func (s *CreateNatIpCidrResponseBody) SetRequestId(v string) *CreateNatIpCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNatIpCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
