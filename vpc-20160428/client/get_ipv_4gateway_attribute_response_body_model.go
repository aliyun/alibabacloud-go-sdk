// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpv4GatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *GetIpv4GatewayAttributeResponseBody
	GetCreateTime() *string
	SetEnabled(v bool) *GetIpv4GatewayAttributeResponseBody
	GetEnabled() *bool
	SetIpv4GatewayDescription(v string) *GetIpv4GatewayAttributeResponseBody
	GetIpv4GatewayDescription() *string
	SetIpv4GatewayId(v string) *GetIpv4GatewayAttributeResponseBody
	GetIpv4GatewayId() *string
	SetIpv4GatewayName(v string) *GetIpv4GatewayAttributeResponseBody
	GetIpv4GatewayName() *string
	SetIpv4GatewayRouteTableId(v string) *GetIpv4GatewayAttributeResponseBody
	GetIpv4GatewayRouteTableId() *string
	SetRequestId(v string) *GetIpv4GatewayAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetIpv4GatewayAttributeResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *GetIpv4GatewayAttributeResponseBody
	GetStatus() *string
	SetTags(v []*GetIpv4GatewayAttributeResponseBodyTags) *GetIpv4GatewayAttributeResponseBody
	GetTags() []*GetIpv4GatewayAttributeResponseBodyTags
	SetVpcId(v string) *GetIpv4GatewayAttributeResponseBody
	GetVpcId() *string
}

type GetIpv4GatewayAttributeResponseBody struct {
	// The time when the IPv4 gateway was created.
	//
	// example:
	//
	// 2022-02-24T09:02:36Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Indicates whether the IPv4 gateway is activated. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// The description of the IPv4 gateway.
	//
	// example:
	//
	// new
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
	// The request ID.
	//
	// example:
	//
	// 7F79A919-6FE9-50C4-967B-45705E1F9C38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// 	- **Activating**
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag list.
	Tags []*GetIpv4GatewayAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The ID of the VPC with which the IPv4 gateway is associated.
	//
	// example:
	//
	// vpc-5tsrxlw7dv074gci4****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetIpv4GatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIpv4GatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *GetIpv4GatewayAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetIpv4GatewayAttributeResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetIpv4GatewayAttributeResponseBody) GetIpv4GatewayDescription() *string {
	return s.Ipv4GatewayDescription
}

func (s *GetIpv4GatewayAttributeResponseBody) GetIpv4GatewayId() *string {
	return s.Ipv4GatewayId
}

func (s *GetIpv4GatewayAttributeResponseBody) GetIpv4GatewayName() *string {
	return s.Ipv4GatewayName
}

func (s *GetIpv4GatewayAttributeResponseBody) GetIpv4GatewayRouteTableId() *string {
	return s.Ipv4GatewayRouteTableId
}

func (s *GetIpv4GatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIpv4GatewayAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetIpv4GatewayAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetIpv4GatewayAttributeResponseBody) GetTags() []*GetIpv4GatewayAttributeResponseBodyTags {
	return s.Tags
}

func (s *GetIpv4GatewayAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *GetIpv4GatewayAttributeResponseBody) SetCreateTime(v string) *GetIpv4GatewayAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetEnabled(v bool) *GetIpv4GatewayAttributeResponseBody {
	s.Enabled = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetIpv4GatewayDescription(v string) *GetIpv4GatewayAttributeResponseBody {
	s.Ipv4GatewayDescription = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetIpv4GatewayId(v string) *GetIpv4GatewayAttributeResponseBody {
	s.Ipv4GatewayId = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetIpv4GatewayName(v string) *GetIpv4GatewayAttributeResponseBody {
	s.Ipv4GatewayName = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetIpv4GatewayRouteTableId(v string) *GetIpv4GatewayAttributeResponseBody {
	s.Ipv4GatewayRouteTableId = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetRequestId(v string) *GetIpv4GatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetResourceGroupId(v string) *GetIpv4GatewayAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetStatus(v string) *GetIpv4GatewayAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetTags(v []*GetIpv4GatewayAttributeResponseBodyTags) *GetIpv4GatewayAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) SetVpcId(v string) *GetIpv4GatewayAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetIpv4GatewayAttributeResponseBodyTags struct {
	// The key of tag N added to the resource.
	//
	// example:
	//
	// FinanceDept
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N added to the resource.
	//
	// example:
	//
	// FinanceJoshua
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetIpv4GatewayAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetIpv4GatewayAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetIpv4GatewayAttributeResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *GetIpv4GatewayAttributeResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *GetIpv4GatewayAttributeResponseBodyTags) SetKey(v string) *GetIpv4GatewayAttributeResponseBodyTags {
	s.Key = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBodyTags) SetValue(v string) *GetIpv4GatewayAttributeResponseBodyTags {
	s.Value = &v
	return s
}

func (s *GetIpv4GatewayAttributeResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
