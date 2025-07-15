// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIpv4GatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv4GatewayModels(v []*ListIpv4GatewaysResponseBodyIpv4GatewayModels) *ListIpv4GatewaysResponseBody
	GetIpv4GatewayModels() []*ListIpv4GatewaysResponseBodyIpv4GatewayModels
	SetNextToken(v string) *ListIpv4GatewaysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIpv4GatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *ListIpv4GatewaysResponseBody
	GetTotalCount() *string
}

type ListIpv4GatewaysResponseBody struct {
	// The list of IPv4 gateways.
	Ipv4GatewayModels []*ListIpv4GatewaysResponseBodyIpv4GatewayModels `json:"Ipv4GatewayModels,omitempty" xml:"Ipv4GatewayModels,omitempty" type:"Repeated"`
	// The token that is used for the next query. Valid values:
	//
	// 	- If no value is returned for **NextToken**, no next queries are sent.
	//
	// 	- If a value of **NextToken*	- is returned, the value is the token that is used for the subsequent query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2D265800-E306-529C-8418-84B0A1D201DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIpv4GatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIpv4GatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *ListIpv4GatewaysResponseBody) GetIpv4GatewayModels() []*ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	return s.Ipv4GatewayModels
}

func (s *ListIpv4GatewaysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIpv4GatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIpv4GatewaysResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListIpv4GatewaysResponseBody) SetIpv4GatewayModels(v []*ListIpv4GatewaysResponseBodyIpv4GatewayModels) *ListIpv4GatewaysResponseBody {
	s.Ipv4GatewayModels = v
	return s
}

func (s *ListIpv4GatewaysResponseBody) SetNextToken(v string) *ListIpv4GatewaysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIpv4GatewaysResponseBody) SetRequestId(v string) *ListIpv4GatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIpv4GatewaysResponseBody) SetTotalCount(v string) *ListIpv4GatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIpv4GatewaysResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIpv4GatewaysResponseBodyIpv4GatewayModels struct {
	// Indicates whether the IPv4 gateway is activated. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The time when the IPv4 gateway was created.
	//
	// example:
	//
	// 2021-12-02T07:07:35Z
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The description of the IPv4 gateway.
	//
	// example:
	//
	// test
	Ipv4GatewayDescription *string `json:"Ipv4GatewayDescription,omitempty" xml:"Ipv4GatewayDescription,omitempty"`
	// The ID of the IPv4 gateway.
	//
	// example:
	//
	// ipv4gw-5tsnc6s4ogsedtp3k****
	Ipv4GatewayId *string `json:"Ipv4GatewayId,omitempty" xml:"Ipv4GatewayId,omitempty"`
	// The name of the IPv4 gateway.
	//
	// example:
	//
	// name
	Ipv4GatewayName *string `json:"Ipv4GatewayName,omitempty" xml:"Ipv4GatewayName,omitempty"`
	// The ID of the route table associated with the IPv4 gateway.
	//
	// example:
	//
	// vtb-5ts0ohchwkp3dydt2****
	Ipv4GatewayRouteTableId *string `json:"Ipv4GatewayRouteTableId,omitempty" xml:"Ipv4GatewayRouteTableId,omitempty"`
	// The ID of the resource group to which the IPv4 gateway belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the IPv4 gateway. Valid values:
	//
	// 	- **Creating**
	//
	// 	- **Created**
	//
	// 	- **Modifying**
	//
	// 	- **Deleting**
	//
	// 	- **Deleted**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags that are added to the resource group.
	Tags []*ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC with which the IPv4 gateways are associated.
	//
	// example:
	//
	// vpc-5tsrxlw7dv074gci4****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListIpv4GatewaysResponseBodyIpv4GatewayModels) String() string {
	return dara.Prettify(s)
}

func (s ListIpv4GatewaysResponseBodyIpv4GatewayModels) GoString() string {
	return s.String()
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetIpv4GatewayDescription() *string {
	return s.Ipv4GatewayDescription
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetIpv4GatewayId() *string {
	return s.Ipv4GatewayId
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetIpv4GatewayName() *string {
	return s.Ipv4GatewayName
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetIpv4GatewayRouteTableId() *string {
	return s.Ipv4GatewayRouteTableId
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetStatus() *string {
	return s.Status
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetTags() []*ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags {
	return s.Tags
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) GetVpcId() *string {
	return s.VpcId
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetEnabled(v bool) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.Enabled = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetGmtCreate(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.GmtCreate = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetIpv4GatewayDescription(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.Ipv4GatewayDescription = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetIpv4GatewayId(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.Ipv4GatewayId = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetIpv4GatewayName(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.Ipv4GatewayName = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetIpv4GatewayRouteTableId(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.Ipv4GatewayRouteTableId = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetResourceGroupId(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.ResourceGroupId = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetStatus(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.Status = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetTags(v []*ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.Tags = v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) SetVpcId(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModels {
	s.VpcId = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModels) Validate() error {
	return dara.Validate(s)
}

type ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags) String() string {
	return dara.Prettify(s)
}

func (s ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags) GoString() string {
	return s.String()
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags) GetKey() *string {
	return s.Key
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags) GetValue() *string {
	return s.Value
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags) SetKey(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags {
	s.Key = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags) SetValue(v string) *ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags {
	s.Value = &v
	return s
}

func (s *ListIpv4GatewaysResponseBodyIpv4GatewayModelsTags) Validate() error {
	return dara.Validate(s)
}
