// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAsn(v int64) *DescribeCustomerGatewayResponseBody
	GetAsn() *int64
	SetAuthKey(v string) *DescribeCustomerGatewayResponseBody
	GetAuthKey() *string
	SetCreateTime(v int64) *DescribeCustomerGatewayResponseBody
	GetCreateTime() *int64
	SetCustomerGatewayId(v string) *DescribeCustomerGatewayResponseBody
	GetCustomerGatewayId() *string
	SetDescription(v string) *DescribeCustomerGatewayResponseBody
	GetDescription() *string
	SetIpAddress(v string) *DescribeCustomerGatewayResponseBody
	GetIpAddress() *string
	SetName(v string) *DescribeCustomerGatewayResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeCustomerGatewayResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeCustomerGatewayResponseBody
	GetResourceGroupId() *string
	SetTags(v *DescribeCustomerGatewayResponseBodyTags) *DescribeCustomerGatewayResponseBody
	GetTags() *DescribeCustomerGatewayResponseBodyTags
}

type DescribeCustomerGatewayResponseBody struct {
	// The autonomous system number (ASN) of the gateway device in the data center.
	//
	// example:
	//
	// 65535
	Asn *int64 `json:"Asn,omitempty" xml:"Asn,omitempty"`
	// The authentication key of the Border Gateway Protocol (BGP) routing protocol for the gateway device in the data center.
	//
	// example:
	//
	// AuthKey****
	AuthKey *string `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	// The timestamp generated when the customer gateway was created. Unit: milliseconds.
	//
	// This value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1492747187000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the customer gateway.
	//
	// example:
	//
	// cgw-bp1pvpl9r9adju6l5****
	CustomerGatewayId *string `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	// The description of the customer gateway.
	//
	// example:
	//
	// desctest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IP address of the gateway device in the data center.
	//
	// example:
	//
	// 139.32.XX.XX
	IpAddress *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	// The name of the customer gateway.
	//
	// example:
	//
	// nametest
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A0457BC9-6C0F-4437-AB9D-FB2EABC1D6A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the customer gateway belongs.
	//
	// You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query resource groups.
	//
	// example:
	//
	// rg-acfmzs372yg****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of tags added to the customer gateway.
	Tags *DescribeCustomerGatewayResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeCustomerGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewayResponseBody) GetAsn() *int64 {
	return s.Asn
}

func (s *DescribeCustomerGatewayResponseBody) GetAuthKey() *string {
	return s.AuthKey
}

func (s *DescribeCustomerGatewayResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeCustomerGatewayResponseBody) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DescribeCustomerGatewayResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeCustomerGatewayResponseBody) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeCustomerGatewayResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeCustomerGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomerGatewayResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCustomerGatewayResponseBody) GetTags() *DescribeCustomerGatewayResponseBodyTags {
	return s.Tags
}

func (s *DescribeCustomerGatewayResponseBody) SetAsn(v int64) *DescribeCustomerGatewayResponseBody {
	s.Asn = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) SetAuthKey(v string) *DescribeCustomerGatewayResponseBody {
	s.AuthKey = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) SetCreateTime(v int64) *DescribeCustomerGatewayResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) SetCustomerGatewayId(v string) *DescribeCustomerGatewayResponseBody {
	s.CustomerGatewayId = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) SetDescription(v string) *DescribeCustomerGatewayResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) SetIpAddress(v string) *DescribeCustomerGatewayResponseBody {
	s.IpAddress = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) SetName(v string) *DescribeCustomerGatewayResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) SetRequestId(v string) *DescribeCustomerGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) SetResourceGroupId(v string) *DescribeCustomerGatewayResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) SetTags(v *DescribeCustomerGatewayResponseBodyTags) *DescribeCustomerGatewayResponseBody {
	s.Tags = v
	return s
}

func (s *DescribeCustomerGatewayResponseBody) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomerGatewayResponseBodyTags struct {
	Tag []*DescribeCustomerGatewayResponseBodyTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCustomerGatewayResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewayResponseBodyTags) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewayResponseBodyTags) GetTag() []*DescribeCustomerGatewayResponseBodyTagsTag {
	return s.Tag
}

func (s *DescribeCustomerGatewayResponseBodyTags) SetTag(v []*DescribeCustomerGatewayResponseBodyTagsTag) *DescribeCustomerGatewayResponseBodyTags {
	s.Tag = v
	return s
}

func (s *DescribeCustomerGatewayResponseBodyTags) Validate() error {
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

type DescribeCustomerGatewayResponseBodyTagsTag struct {
	// The tag key.
	//
	// example:
	//
	// TagKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TagValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCustomerGatewayResponseBodyTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewayResponseBodyTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewayResponseBodyTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCustomerGatewayResponseBodyTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCustomerGatewayResponseBodyTagsTag) SetKey(v string) *DescribeCustomerGatewayResponseBodyTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBodyTagsTag) SetValue(v string) *DescribeCustomerGatewayResponseBodyTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeCustomerGatewayResponseBodyTagsTag) Validate() error {
	return dara.Validate(s)
}
