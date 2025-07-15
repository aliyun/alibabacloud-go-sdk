// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatIp(v string) *CreateNatIpResponseBody
	GetNatIp() *string
	SetNatIpId(v string) *CreateNatIpResponseBody
	GetNatIpId() *string
	SetRequestId(v string) *CreateNatIpResponseBody
	GetRequestId() *string
}

type CreateNatIpResponseBody struct {
	// The NAT IP address.
	//
	// example:
	//
	// 192.168.0.34
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The ID of the NAT IP address.
	//
	// example:
	//
	// vpcnatip-gw8y7q3cpk3fggs8****
	NatIpId *string `json:"NatIpId,omitempty" xml:"NatIpId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E9AD97A0-5338-43F8-8A80-5E274CCBA11B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNatIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNatIpResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatIpResponseBody) GetNatIp() *string {
	return s.NatIp
}

func (s *CreateNatIpResponseBody) GetNatIpId() *string {
	return s.NatIpId
}

func (s *CreateNatIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatIpResponseBody) SetNatIp(v string) *CreateNatIpResponseBody {
	s.NatIp = &v
	return s
}

func (s *CreateNatIpResponseBody) SetNatIpId(v string) *CreateNatIpResponseBody {
	s.NatIpId = &v
	return s
}

func (s *CreateNatIpResponseBody) SetRequestId(v string) *CreateNatIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNatIpResponseBody) Validate() error {
	return dara.Validate(s)
}
