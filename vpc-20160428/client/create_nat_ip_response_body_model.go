// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv4Prefix(v string) *CreateNatIpResponseBody
	GetIpv4Prefix() *string
	SetNatIp(v string) *CreateNatIpResponseBody
	GetNatIp() *string
	SetNatIpId(v string) *CreateNatIpResponseBody
	GetNatIpId() *string
	SetNatIps(v []*CreateNatIpResponseBodyNatIps) *CreateNatIpResponseBody
	GetNatIps() []*CreateNatIpResponseBodyNatIps
	SetRequestId(v string) *CreateNatIpResponseBody
	GetRequestId() *string
}

type CreateNatIpResponseBody struct {
	// The IPv4Prefix returned by the previous API is obsolete.
	//
	// example:
	//
	// ""
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
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
	// The NatIp parameter that is returned after you create a NatIp. If you use IPv4Prefix to create a NatIp, the information about all NatIp is returned. We recommend that you use this parameter to obtain the information about a NatIp when you create a NatIp.
	NatIps []*CreateNatIpResponseBodyNatIps `json:"NatIps,omitempty" xml:"NatIps,omitempty" type:"Repeated"`
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

func (s *CreateNatIpResponseBody) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *CreateNatIpResponseBody) GetNatIp() *string {
	return s.NatIp
}

func (s *CreateNatIpResponseBody) GetNatIpId() *string {
	return s.NatIpId
}

func (s *CreateNatIpResponseBody) GetNatIps() []*CreateNatIpResponseBodyNatIps {
	return s.NatIps
}

func (s *CreateNatIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatIpResponseBody) SetIpv4Prefix(v string) *CreateNatIpResponseBody {
	s.Ipv4Prefix = &v
	return s
}

func (s *CreateNatIpResponseBody) SetNatIp(v string) *CreateNatIpResponseBody {
	s.NatIp = &v
	return s
}

func (s *CreateNatIpResponseBody) SetNatIpId(v string) *CreateNatIpResponseBody {
	s.NatIpId = &v
	return s
}

func (s *CreateNatIpResponseBody) SetNatIps(v []*CreateNatIpResponseBodyNatIps) *CreateNatIpResponseBody {
	s.NatIps = v
	return s
}

func (s *CreateNatIpResponseBody) SetRequestId(v string) *CreateNatIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNatIpResponseBody) Validate() error {
	if s.NatIps != nil {
		for _, item := range s.NatIps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateNatIpResponseBodyNatIps struct {
	// The Ipv4Prefix of the created NatIpList list is returned when Ipv4Preix is created.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 192.168.1.128/28
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
	// Returns the NatIp address of the created NatIpList list when Ipv4Preix is created.
	//
	// example:
	//
	// 192.168.2.128
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// Returns the NatIpId of the created NatIpList list when Ipv4Preix is created.
	//
	// example:
	//
	// vpcnatip-xxxxxxxx
	NatIpId *string `json:"NatIpId,omitempty" xml:"NatIpId,omitempty"`
}

func (s CreateNatIpResponseBodyNatIps) String() string {
	return dara.Prettify(s)
}

func (s CreateNatIpResponseBodyNatIps) GoString() string {
	return s.String()
}

func (s *CreateNatIpResponseBodyNatIps) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *CreateNatIpResponseBodyNatIps) GetNatIp() *string {
	return s.NatIp
}

func (s *CreateNatIpResponseBodyNatIps) GetNatIpId() *string {
	return s.NatIpId
}

func (s *CreateNatIpResponseBodyNatIps) SetIpv4Prefix(v string) *CreateNatIpResponseBodyNatIps {
	s.Ipv4Prefix = &v
	return s
}

func (s *CreateNatIpResponseBodyNatIps) SetNatIp(v string) *CreateNatIpResponseBodyNatIps {
	s.NatIp = &v
	return s
}

func (s *CreateNatIpResponseBodyNatIps) SetNatIpId(v string) *CreateNatIpResponseBodyNatIps {
	s.NatIpId = &v
	return s
}

func (s *CreateNatIpResponseBodyNatIps) Validate() error {
	return dara.Validate(s)
}
