// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomerGateways(v *DescribeCustomerGatewaysResponseBodyCustomerGateways) *DescribeCustomerGatewaysResponseBody
	GetCustomerGateways() *DescribeCustomerGatewaysResponseBodyCustomerGateways
	SetPageNumber(v int32) *DescribeCustomerGatewaysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCustomerGatewaysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCustomerGatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCustomerGatewaysResponseBody
	GetTotalCount() *int32
}

type DescribeCustomerGatewaysResponseBody struct {
	CustomerGateways *DescribeCustomerGatewaysResponseBodyCustomerGateways `json:"CustomerGateways,omitempty" xml:"CustomerGateways,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E82612A9-CB90-4D7E-B394-1DB7F6509B29
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCustomerGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewaysResponseBody) GetCustomerGateways() *DescribeCustomerGatewaysResponseBodyCustomerGateways {
	return s.CustomerGateways
}

func (s *DescribeCustomerGatewaysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCustomerGatewaysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustomerGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomerGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCustomerGatewaysResponseBody) SetCustomerGateways(v *DescribeCustomerGatewaysResponseBodyCustomerGateways) *DescribeCustomerGatewaysResponseBody {
	s.CustomerGateways = v
	return s
}

func (s *DescribeCustomerGatewaysResponseBody) SetPageNumber(v int32) *DescribeCustomerGatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBody) SetPageSize(v int32) *DescribeCustomerGatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBody) SetRequestId(v string) *DescribeCustomerGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBody) SetTotalCount(v int32) *DescribeCustomerGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBody) Validate() error {
	if s.CustomerGateways != nil {
		if err := s.CustomerGateways.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomerGatewaysResponseBodyCustomerGateways struct {
	CustomerGateway []*DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway `json:"CustomerGateway,omitempty" xml:"CustomerGateway,omitempty" type:"Repeated"`
}

func (s DescribeCustomerGatewaysResponseBodyCustomerGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewaysResponseBodyCustomerGateways) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGateways) GetCustomerGateway() []*DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	return s.CustomerGateway
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGateways) SetCustomerGateway(v []*DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) *DescribeCustomerGatewaysResponseBodyCustomerGateways {
	s.CustomerGateway = v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGateways) Validate() error {
	if s.CustomerGateway != nil {
		for _, item := range s.CustomerGateway {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway struct {
	Asn               *int64                                                                   `json:"Asn,omitempty" xml:"Asn,omitempty"`
	AuthKey           *string                                                                  `json:"AuthKey,omitempty" xml:"AuthKey,omitempty"`
	CreateTime        *int64                                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CustomerGatewayId *string                                                                  `json:"CustomerGatewayId,omitempty" xml:"CustomerGatewayId,omitempty"`
	Description       *string                                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	IpAddress         *string                                                                  `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
	Name              *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceGroupId   *string                                                                  `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags              *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GetAsn() *int64 {
	return s.Asn
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GetAuthKey() *string {
	return s.AuthKey
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GetCustomerGatewayId() *string {
	return s.CustomerGatewayId
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GetDescription() *string {
	return s.Description
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GetIpAddress() *string {
	return s.IpAddress
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GetName() *string {
	return s.Name
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) GetTags() *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags {
	return s.Tags
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) SetAsn(v int64) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	s.Asn = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) SetAuthKey(v string) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	s.AuthKey = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) SetCreateTime(v int64) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	s.CreateTime = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) SetCustomerGatewayId(v string) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	s.CustomerGatewayId = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) SetDescription(v string) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	s.Description = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) SetIpAddress(v string) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	s.IpAddress = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) SetName(v string) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	s.Name = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) SetResourceGroupId(v string) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) SetTags(v *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway {
	s.Tags = v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGateway) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags struct {
	Tag []*DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags) GetTag() []*DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag {
	return s.Tag
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags) SetTag(v []*DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags {
	s.Tag = v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTags) Validate() error {
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

type DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag) SetKey(v string) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag) SetValue(v string) *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeCustomerGatewaysResponseBodyCustomerGatewaysCustomerGatewayTagsTag) Validate() error {
	return dara.Validate(s)
}
