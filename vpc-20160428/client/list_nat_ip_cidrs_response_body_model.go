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
	// The list of NAT CIDR blocks.
	NatIpCidrs []*ListNatIpCidrsResponseBodyNatIpCidrs `json:"NatIpCidrs,omitempty" xml:"NatIpCidrs,omitempty" type:"Repeated"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// - If **NextToken*	- is empty, no subsequent requests exist.
	//
	// - If **NextToken*	- is returned, the value indicates the token for the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7479A224-4A28-4895-9604-11F48BCE6A88
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of NAT CIDR block entries returned.
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
	// The time when the NAT CIDR block was created.
	//
	// example:
	//
	// 2021-06-28T20:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the NAT CIDR block is the default NAT CIDR block. Valid values:
	//
	// - **true**: The NAT CIDR block is the default NAT CIDR block.
	//
	// - **false**: The NAT CIDR block is not the default NAT CIDR block.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// The instance ID of the VPC NAT gateway to which the NAT CIDR block belongs.
	//
	// example:
	//
	// ngw-gw8v16wgvtq26vh59****
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	// The NAT CIDR block.
	//
	// example:
	//
	// 172.16.0.0/24
	NatIpCidr *string `json:"NatIpCidr,omitempty" xml:"NatIpCidr,omitempty"`
	// The description of the NAT CIDR block.
	//
	// example:
	//
	// test
	NatIpCidrDescription *string `json:"NatIpCidrDescription,omitempty" xml:"NatIpCidrDescription,omitempty"`
	// The instance ID of the NAT CIDR block.
	//
	// example:
	//
	// vpcnatcidr-gw8ov42ei6xh1jys2****
	NatIpCidrId *string `json:"NatIpCidrId,omitempty" xml:"NatIpCidrId,omitempty"`
	// The name of the NAT CIDR block.
	//
	// example:
	//
	// Name
	NatIpCidrName *string `json:"NatIpCidrName,omitempty" xml:"NatIpCidrName,omitempty"`
	// The status of the NAT CIDR block. The value is **Available**, which indicates that the NAT CIDR block is available.
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
