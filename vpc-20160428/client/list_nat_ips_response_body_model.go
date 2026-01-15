// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNatIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatIps(v []*ListNatIpsResponseBodyNatIps) *ListNatIpsResponseBody
	GetNatIps() []*ListNatIpsResponseBodyNatIps
	SetNextToken(v string) *ListNatIpsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListNatIpsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListNatIpsResponseBody
	GetTotalCount() *string
}

type ListNatIpsResponseBody struct {
	// The list of IP addresses of the NAT gateway.
	NatIps []*ListNatIpsResponseBodyNatIps `json:"NatIps,omitempty" xml:"NatIps,omitempty" type:"Repeated"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If the value of **NextToken*	- is not returned, it indicates that no next query is to be sent.
	//
	// 	- If the value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7E78CEEA-BF8F-44D1-9DCD-D9141135B71E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of IP addresses that are returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNatIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNatIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNatIpsResponseBody) GetNatIps() []*ListNatIpsResponseBodyNatIps {
	return s.NatIps
}

func (s *ListNatIpsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNatIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNatIpsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListNatIpsResponseBody) SetNatIps(v []*ListNatIpsResponseBodyNatIps) *ListNatIpsResponseBody {
	s.NatIps = v
	return s
}

func (s *ListNatIpsResponseBody) SetNextToken(v string) *ListNatIpsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNatIpsResponseBody) SetRequestId(v string) *ListNatIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNatIpsResponseBody) SetTotalCount(v string) *ListNatIpsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNatIpsResponseBody) Validate() error {
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

type ListNatIpsResponseBodyNatIps struct {
	// The IP prefix address range to which the queried NAT IP address belongs.
	//
	// example:
	//
	// 192.168.0.0/28
	Ipv4Prefix *string `json:"Ipv4Prefix,omitempty" xml:"Ipv4Prefix,omitempty"`
	// Indicates whether the IP address is the default IP address of the NAT gateway. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the Virtual Private Cloud (VPC) NAT gateway to which the IP address is assigned.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The IP address.
	//
	// example:
	//
	// 192.168.0.126
	NatIp *string `json:"NatIp,omitempty" xml:"NatIp,omitempty"`
	// The CIDR block to which the IP address belongs.
	//
	// example:
	//
	// 192.168.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The description of the IP address.
	//
	// example:
	//
	// default
	NatIpDescription *string `json:"NatIpDescription,omitempty" xml:"NatIpDescription,omitempty"`
	// The ID of the IP address.
	//
	// example:
	//
	// vpcnatip-gw8a863sut1zijxh0****
	NatIpId *string `json:"NatIpId,omitempty" xml:"NatIpId,omitempty"`
	// The name of the IP address.
	//
	// example:
	//
	// default
	NatIpName *string `json:"NatIpName,omitempty" xml:"NatIpName,omitempty"`
	// The status of the IP address. Valid values:
	//
	// 	- **Available**: available
	//
	// 	- **Deleted**: deleted
	//
	// 	- **Deleting**: deleting
	//
	// 	- **Creating**: creating
	//
	// 	- **Associated**: specified in an SNAT or DNAT entry
	//
	// 	- **Associating**: being specified in an SNAT or DNAT entry
	//
	// example:
	//
	// Available
	NatIpStatus *string `json:"NatIpStatus,omitempty" xml:"NatIpStatus,omitempty"`
}

func (s ListNatIpsResponseBodyNatIps) String() string {
	return dara.Prettify(s)
}

func (s ListNatIpsResponseBodyNatIps) GoString() string {
	return s.String()
}

func (s *ListNatIpsResponseBodyNatIps) GetIpv4Prefix() *string {
	return s.Ipv4Prefix
}

func (s *ListNatIpsResponseBodyNatIps) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListNatIpsResponseBodyNatIps) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ListNatIpsResponseBodyNatIps) GetNatIp() *string {
	return s.NatIp
}

func (s *ListNatIpsResponseBodyNatIps) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *ListNatIpsResponseBodyNatIps) GetNatIpDescription() *string {
	return s.NatIpDescription
}

func (s *ListNatIpsResponseBodyNatIps) GetNatIpId() *string {
	return s.NatIpId
}

func (s *ListNatIpsResponseBodyNatIps) GetNatIpName() *string {
	return s.NatIpName
}

func (s *ListNatIpsResponseBodyNatIps) GetNatIpStatus() *string {
	return s.NatIpStatus
}

func (s *ListNatIpsResponseBodyNatIps) SetIpv4Prefix(v string) *ListNatIpsResponseBodyNatIps {
	s.Ipv4Prefix = &v
	return s
}

func (s *ListNatIpsResponseBodyNatIps) SetIsDefault(v bool) *ListNatIpsResponseBodyNatIps {
	s.IsDefault = &v
	return s
}

func (s *ListNatIpsResponseBodyNatIps) SetNatGatewayId(v string) *ListNatIpsResponseBodyNatIps {
	s.NatGatewayId = &v
	return s
}

func (s *ListNatIpsResponseBodyNatIps) SetNatIp(v string) *ListNatIpsResponseBodyNatIps {
	s.NatIp = &v
	return s
}

func (s *ListNatIpsResponseBodyNatIps) SetNatIpCidr(v string) *ListNatIpsResponseBodyNatIps {
	s.NatIpCidr = &v
	return s
}

func (s *ListNatIpsResponseBodyNatIps) SetNatIpDescription(v string) *ListNatIpsResponseBodyNatIps {
	s.NatIpDescription = &v
	return s
}

func (s *ListNatIpsResponseBodyNatIps) SetNatIpId(v string) *ListNatIpsResponseBodyNatIps {
	s.NatIpId = &v
	return s
}

func (s *ListNatIpsResponseBodyNatIps) SetNatIpName(v string) *ListNatIpsResponseBodyNatIps {
	s.NatIpName = &v
	return s
}

func (s *ListNatIpsResponseBodyNatIps) SetNatIpStatus(v string) *ListNatIpsResponseBodyNatIps {
	s.NatIpStatus = &v
	return s
}

func (s *ListNatIpsResponseBodyNatIps) Validate() error {
	return dara.Validate(s)
}
