// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6GatewayAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBusinessStatus(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetBusinessStatus() *string
	SetCreationTime(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetCreationTime() *string
	SetDescription(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetDescription() *string
	SetExpiredTime(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetExpiredTime() *string
	SetGatewayRouteTableId(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetGatewayRouteTableId() *string
	SetInstanceChargeType(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetInstanceChargeType() *string
	SetIpv6GatewayId(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetIpv6GatewayId() *string
	SetName(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetName() *string
	SetOwnerId(v int64) *DescribeIpv6GatewayAttributeResponseBody
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetStatus() *string
	SetTags(v *DescribeIpv6GatewayAttributeResponseBodyTags) *DescribeIpv6GatewayAttributeResponseBody
	GetTags() *DescribeIpv6GatewayAttributeResponseBodyTags
	SetVpcId(v string) *DescribeIpv6GatewayAttributeResponseBody
	GetVpcId() *string
}

type DescribeIpv6GatewayAttributeResponseBody struct {
	// The business status of the IPv6 gateway. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **FinancialLocked**: financial lock.
	//
	// - **SecurityLocked**: security lock.
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the IPv6 gateway was created.
	//
	// example:
	//
	// 2018-12-05T09:21:35Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the IPv6 gateway.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expiration time of the IPv6 gateway.
	//
	// example:
	//
	// 2019-1-05T09:21:35Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The ID of the gateway route table associated with the IPv6 gateway.
	//
	//
	// > This parameter is displayed only for IPv6 gateways that are associated with a gateway route table.
	//
	// example:
	//
	// vtb-5ts0ohchwkp3dydt2****
	GatewayRouteTableId *string `json:"GatewayRouteTableId,omitempty" xml:"GatewayRouteTableId,omitempty"`
	// The billing method of the IPv6 gateway.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance ID of the IPv6 gateway.
	//
	// example:
	//
	// ipv6gw-hp3y0l3ln89j8cdvf****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	// The name of the IPv6 gateway.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account ID of the IPv6 gateway owner.
	//
	// 	Notice: This value is of the Long type. Precision loss may occur in certain programming languages. Use this value with caution.
	//
	// example:
	//
	// 2546073170691****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the IPv6 gateway.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmxazb4ph6aiy****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the IPv6 gateway. Valid values:
	//
	// - **Pending**: being configured.
	//
	// - **Available**: active.
	//
	// example:
	//
	// Available
	Status *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   *DescribeIpv6GatewayAttributeResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the VPC to which the IPv6 gateway belongs.
	//
	// example:
	//
	// vpc-123sedrfswd23****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeIpv6GatewayAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewayAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetGatewayRouteTableId() *string {
	return s.GatewayRouteTableId
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetTags() *DescribeIpv6GatewayAttributeResponseBodyTags {
	return s.Tags
}

func (s *DescribeIpv6GatewayAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetBusinessStatus(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetCreationTime(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetDescription(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetExpiredTime(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetGatewayRouteTableId(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.GatewayRouteTableId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetInstanceChargeType(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetIpv6GatewayId(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.Ipv6GatewayId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetName(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetOwnerId(v int64) *DescribeIpv6GatewayAttributeResponseBody {
	s.OwnerId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetRegionId(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetRequestId(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetResourceGroupId(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetStatus(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetTags(v *DescribeIpv6GatewayAttributeResponseBodyTags) *DescribeIpv6GatewayAttributeResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) SetVpcId(v string) *DescribeIpv6GatewayAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBody) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIpv6GatewayAttributeResponseBodyTags struct {
	Tag []*DescribeIpv6GatewayAttributeResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeIpv6GatewayAttributeResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewayAttributeResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewayAttributeResponseBodyTags) GetTag() []*DescribeIpv6GatewayAttributeResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeIpv6GatewayAttributeResponseBodyTags) SetTag(v []*DescribeIpv6GatewayAttributeResponseBodyTagsTag) *DescribeIpv6GatewayAttributeResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBodyTags) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpv6GatewayAttributeResponseBodyTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIpv6GatewayAttributeResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewayAttributeResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewayAttributeResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeIpv6GatewayAttributeResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeIpv6GatewayAttributeResponseBodyTagsTag) SetKey(v string) *DescribeIpv6GatewayAttributeResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBodyTagsTag) SetValue(v string) *DescribeIpv6GatewayAttributeResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeIpv6GatewayAttributeResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
