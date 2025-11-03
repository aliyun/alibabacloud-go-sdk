// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNatIpCidrsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatIpCidrs(v []*ListNatIpCidrsResponseBodyNatIpCidrs) *ListNatIpCidrsResponseBody
	GetNatIpCidrs() []*ListNatIpCidrsResponseBodyNatIpCidrs
	SetNextToken(v string) *ListNatIpCidrsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListNatIpCidrsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListNatIpCidrsResponseBody
	GetTotalCount() *string
}

type ListNatIpCidrsResponseBody struct {
	// The CIDR blocks of the NAT gateway.
	NatIpCidrs []*ListNatIpCidrsResponseBodyNatIpCidrs `json:"NatIpCidrs,omitempty" xml:"NatIpCidrs,omitempty" type:"Repeated"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If the value of **NextToken*	- is not returned, it indicates that no next query is to be sent.
	//
	// 	- If the value of **NextToken*	- is returned, the value indicates the token that is used for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7479A224-4A28-4895-9604-11F48BCE6A88
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of CIDR blocks that are returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNatIpCidrsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNatIpCidrsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNatIpCidrsResponseBody) GetNatIpCidrs() []*ListNatIpCidrsResponseBodyNatIpCidrs {
	return s.NatIpCidrs
}

func (s *ListNatIpCidrsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNatIpCidrsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNatIpCidrsResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListNatIpCidrsResponseBody) SetNatIpCidrs(v []*ListNatIpCidrsResponseBodyNatIpCidrs) *ListNatIpCidrsResponseBody {
	s.NatIpCidrs = v
	return s
}

func (s *ListNatIpCidrsResponseBody) SetNextToken(v string) *ListNatIpCidrsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNatIpCidrsResponseBody) SetRequestId(v string) *ListNatIpCidrsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNatIpCidrsResponseBody) SetTotalCount(v string) *ListNatIpCidrsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNatIpCidrsResponseBody) Validate() error {
	if s.NatIpCidrs != nil {
		for _, item := range s.NatIpCidrs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNatIpCidrsResponseBodyNatIpCidrs struct {
	// The time when the CIDR block was created.
	//
	// example:
	//
	// 2021-06-28T20:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the CIDR block is the default CIDR block of the NAT gateway. Valid values:
	//
	// 	- **true**: The CIDR block is the default CIDR block of the NAT gateway.
	//
	// 	- **false**: The CIDR block is not the default CIDR block of the NAT gateway.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The ID of the VPC NAT gateway.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The CIDR block of the NAT gateway.
	//
	// example:
	//
	// 172.16.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The description of the CIDR block of the NAT gateway.
	//
	// example:
	//
	// test
	NatIpCidrDescription *string `json:"NatIpCidrDescription,omitempty" xml:"NatIpCidrDescription,omitempty"`
	// The ID of the CIDR block of the NAT gateway.
	//
	// example:
	//
	// vpcnatcidr-gw8ov42ei6xh1jys2****
	NatIpCidrId *string `json:"NatIpCidrId,omitempty" xml:"NatIpCidrId,omitempty"`
	// The name of the CIDR block of the NAT gateway.
	//
	// example:
	//
	// Name
	NatIpCidrName *string `json:"NatIpCidrName,omitempty" xml:"NatIpCidrName,omitempty"`
	// The status of the CIDR block of the NAT gateway. If **Available*	- is returned, it indicates that the CIDR block is available.
	//
	// example:
	//
	// Available
	NatIpCidrStatus *string `json:"NatIpCidrStatus,omitempty" xml:"NatIpCidrStatus,omitempty"`
}

func (s ListNatIpCidrsResponseBodyNatIpCidrs) String() string {
	return dara.Prettify(s)
}

func (s ListNatIpCidrsResponseBodyNatIpCidrs) GoString() string {
	return s.String()
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) GetNatIpCidr() *string {
	return s.NatIpCidr
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) GetNatIpCidrDescription() *string {
	return s.NatIpCidrDescription
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) GetNatIpCidrId() *string {
	return s.NatIpCidrId
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) GetNatIpCidrName() *string {
	return s.NatIpCidrName
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) GetNatIpCidrStatus() *string {
	return s.NatIpCidrStatus
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) SetCreationTime(v string) *ListNatIpCidrsResponseBodyNatIpCidrs {
	s.CreationTime = &v
	return s
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) SetIsDefault(v bool) *ListNatIpCidrsResponseBodyNatIpCidrs {
	s.IsDefault = &v
	return s
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) SetNatGatewayId(v string) *ListNatIpCidrsResponseBodyNatIpCidrs {
	s.NatGatewayId = &v
	return s
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) SetNatIpCidr(v string) *ListNatIpCidrsResponseBodyNatIpCidrs {
	s.NatIpCidr = &v
	return s
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) SetNatIpCidrDescription(v string) *ListNatIpCidrsResponseBodyNatIpCidrs {
	s.NatIpCidrDescription = &v
	return s
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) SetNatIpCidrId(v string) *ListNatIpCidrsResponseBodyNatIpCidrs {
	s.NatIpCidrId = &v
	return s
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) SetNatIpCidrName(v string) *ListNatIpCidrsResponseBodyNatIpCidrs {
	s.NatIpCidrName = &v
	return s
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) SetNatIpCidrStatus(v string) *ListNatIpCidrsResponseBodyNatIpCidrs {
	s.NatIpCidrStatus = &v
	return s
}

func (s *ListNatIpCidrsResponseBodyNatIpCidrs) Validate() error {
	return dara.Validate(s)
}
