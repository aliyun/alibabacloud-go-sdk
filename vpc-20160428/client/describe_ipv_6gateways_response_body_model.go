// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpv6GatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpv6Gateways(v *DescribeIpv6GatewaysResponseBodyIpv6Gateways) *DescribeIpv6GatewaysResponseBody
	GetIpv6Gateways() *DescribeIpv6GatewaysResponseBodyIpv6Gateways
	SetPageNumber(v int32) *DescribeIpv6GatewaysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeIpv6GatewaysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeIpv6GatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIpv6GatewaysResponseBody
	GetTotalCount() *int32
}

type DescribeIpv6GatewaysResponseBody struct {
	// The information about the IPv6 gateway.
	Ipv6Gateways *DescribeIpv6GatewaysResponseBodyIpv6Gateways `json:"Ipv6Gateways,omitempty" xml:"Ipv6Gateways,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E3A06196-3E7C-490D-8F39-CB4B5A0CE8AD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIpv6GatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewaysResponseBody) GetIpv6Gateways() *DescribeIpv6GatewaysResponseBodyIpv6Gateways {
	return s.Ipv6Gateways
}

func (s *DescribeIpv6GatewaysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeIpv6GatewaysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeIpv6GatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpv6GatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIpv6GatewaysResponseBody) SetIpv6Gateways(v *DescribeIpv6GatewaysResponseBodyIpv6Gateways) *DescribeIpv6GatewaysResponseBody {
	s.Ipv6Gateways = v
	return s
}

func (s *DescribeIpv6GatewaysResponseBody) SetPageNumber(v int32) *DescribeIpv6GatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBody) SetPageSize(v int32) *DescribeIpv6GatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBody) SetRequestId(v string) *DescribeIpv6GatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBody) SetTotalCount(v int32) *DescribeIpv6GatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBody) Validate() error {
	if s.Ipv6Gateways != nil {
		if err := s.Ipv6Gateways.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIpv6GatewaysResponseBodyIpv6Gateways struct {
	Ipv6Gateway []*DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway `json:"Ipv6Gateway,omitempty" xml:"Ipv6Gateway,omitempty" type:"Repeated"`
}

func (s DescribeIpv6GatewaysResponseBodyIpv6Gateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewaysResponseBodyIpv6Gateways) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6Gateways) GetIpv6Gateway() []*DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	return s.Ipv6Gateway
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6Gateways) SetIpv6Gateway(v []*DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) *DescribeIpv6GatewaysResponseBodyIpv6Gateways {
	s.Ipv6Gateway = v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6Gateways) Validate() error {
	if s.Ipv6Gateway != nil {
		for _, item := range s.Ipv6Gateway {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway struct {
	// The status of the IPv6 gateway. Valid values:
	//
	// 	- **Normal**
	//
	// 	- **FinancialLocked**
	//
	// 	- **SecurityLocked**
	//
	// example:
	//
	// Normal
	BusinessStatus *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	// The time when the IPv6 gateway was created.
	//
	// example:
	//
	// 2020-12-20T14:51:23Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the IPv6 gateway.
	//
	// example:
	//
	// descriptionforIPv6GW
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the IPv6 gateway expires.
	//
	// example:
	//
	// 2021-12-20T14:51:23Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The billing method of the IPv6 gateway.
	//
	// Only **PostPaid*	- may be returned, which indicates that the IPv6 gateway uses the pay-as-you-go billing method.
	//
	// example:
	//
	// PostPaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The ID of the IPv6 gateway.
	//
	// example:
	//
	// ipv6gw-hp3rwmtmfhgisipv6gw-hp3rwmtmfhgis****
	Ipv6GatewayId *string `json:"Ipv6GatewayId,omitempty" xml:"Ipv6GatewayId,omitempty"`
	// The name of the IPv6 gateway.
	//
	// example:
	//
	// ipv6GW
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which the IPv6 gateway is deployed.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-bp67acfmxazb4ph****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The status of the IPv6 gateway. Valid values:
	//
	// 	- **Pending**
	//
	// 	- **Available**
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the tags.
	Tags *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the VPC to which the IPv6 gateway belongs.
	//
	// example:
	//
	// vpc-123sedrfswd23****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetDescription() *string {
	return s.Description
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetIpv6GatewayId() *string {
	return s.Ipv6GatewayId
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetName() *string {
	return s.Name
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetStatus() *string {
	return s.Status
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetTags() *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags {
	return s.Tags
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetBusinessStatus(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetCreationTime(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.CreationTime = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetDescription(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.Description = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetExpiredTime(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetInstanceChargeType(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetIpv6GatewayId(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.Ipv6GatewayId = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetName(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.Name = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetOwnerId(v int64) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.OwnerId = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetRegionId(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.RegionId = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetResourceGroupId(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetStatus(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.Status = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetTags(v *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.Tags = v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) SetVpcId(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway {
	s.VpcId = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6Gateway) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags struct {
	Tag []*DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags) GetTag() []*DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag {
	return s.Tag
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags) SetTag(v []*DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags {
	s.Tag = v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTags) Validate() error {
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

type DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag struct {
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

func (s DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag) SetKey(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag) SetValue(v string) *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeIpv6GatewaysResponseBodyIpv6GatewaysIpv6GatewayTagsTag) Validate() error {
	return dara.Validate(s)
}
