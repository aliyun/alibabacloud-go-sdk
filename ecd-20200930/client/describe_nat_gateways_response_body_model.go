// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeNatGatewaysResponseBody
	GetMaxResults() *int32
	SetNatGateways(v []*DescribeNatGatewaysResponseBodyNatGateways) *DescribeNatGatewaysResponseBody
	GetNatGateways() []*DescribeNatGatewaysResponseBodyNatGateways
	SetNextToken(v string) *DescribeNatGatewaysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeNatGatewaysResponseBody
	GetRequestId() *string
}

type DescribeNatGatewaysResponseBody struct {
	MaxResults  *int32                                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NatGateways []*DescribeNatGatewaysResponseBodyNatGateways `json:"NatGateways,omitempty" xml:"NatGateways,omitempty" type:"Repeated"`
	NextToken   *string                                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeNatGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeNatGatewaysResponseBody) GetNatGateways() []*DescribeNatGatewaysResponseBodyNatGateways {
	return s.NatGateways
}

func (s *DescribeNatGatewaysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeNatGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatGatewaysResponseBody) SetMaxResults(v int32) *DescribeNatGatewaysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetNatGateways(v []*DescribeNatGatewaysResponseBodyNatGateways) *DescribeNatGatewaysResponseBody {
	s.NatGateways = v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetNextToken(v string) *DescribeNatGatewaysResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetRequestId(v string) *DescribeNatGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGateways struct {
	ForwardTableIds []*string                                            `json:"ForwardTableIds,omitempty" xml:"ForwardTableIds,omitempty" type:"Repeated"`
	IpLists         []*DescribeNatGatewaysResponseBodyNatGatewaysIpLists `json:"IpLists,omitempty" xml:"IpLists,omitempty" type:"Repeated"`
	Name            *string                                              `json:"Name,omitempty" xml:"Name,omitempty"`
	NatGatewayId    *string                                              `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	SnatTableIds    []*string                                            `json:"SnatTableIds,omitempty" xml:"SnatTableIds,omitempty" type:"Repeated"`
	Status          *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId           *string                                              `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGateways) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetForwardTableIds() []*string {
	return s.ForwardTableIds
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetIpLists() []*DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	return s.IpLists
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetName() *string {
	return s.Name
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetSnatTableIds() []*string {
	return s.SnatTableIds
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetStatus() *string {
	return s.Status
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetForwardTableIds(v []*string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.ForwardTableIds = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetIpLists(v []*DescribeNatGatewaysResponseBodyNatGatewaysIpLists) *DescribeNatGatewaysResponseBodyNatGateways {
	s.IpLists = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetName(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.Name = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetNatGatewayId(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetSnatTableIds(v []*string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.SnatTableIds = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetStatus(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.Status = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetVpcId(v string) *DescribeNatGatewaysResponseBodyNatGateways {
	s.VpcId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysIpLists struct {
	AllocationId     *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	IpAddress        *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	SnatEntryEnabled *string `json:"SnatEntryEnabled,omitempty" xml:"SnatEntryEnabled,omitempty"`
	UsingStatus      *string `json:"UsingStatus,omitempty" xml:"UsingStatus,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysIpLists) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GetSnatEntryEnabled() *string {
	return s.SnatEntryEnabled
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) GetUsingStatus() *string {
	return s.UsingStatus
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) SetAllocationId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	s.AllocationId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) SetIpAddress(v string) *DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	s.IpAddress = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) SetPrivateIpAddress(v string) *DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) SetSnatEntryEnabled(v string) *DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	s.SnatEntryEnabled = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) SetUsingStatus(v string) *DescribeNatGatewaysResponseBodyNatGatewaysIpLists {
	s.UsingStatus = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysIpLists) Validate() error {
	return dara.Validate(s)
}
