// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpAbroadCountryInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAbroadRegion(v string) *DescribeIpAbroadCountryInfosRequest
	GetAbroadRegion() *string
	SetCountry(v string) *DescribeIpAbroadCountryInfosRequest
	GetCountry() *string
	SetInstanceId(v string) *DescribeIpAbroadCountryInfosRequest
	GetInstanceId() *string
	SetLanguage(v string) *DescribeIpAbroadCountryInfosRequest
	GetLanguage() *string
	SetMaxResults(v int32) *DescribeIpAbroadCountryInfosRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeIpAbroadCountryInfosRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeIpAbroadCountryInfosRequest
	GetRegionId() *string
	SetResourceManagerResourceGroupId(v string) *DescribeIpAbroadCountryInfosRequest
	GetResourceManagerResourceGroupId() *string
}

type DescribeIpAbroadCountryInfosRequest struct {
	// The overseas region to query. You can perform a fuzzy match by using the region ID or the region name in Chinese or English.
	//
	// example:
	//
	// US-CA
	AbroadRegion *string `json:"AbroadRegion,omitempty" xml:"AbroadRegion,omitempty"`
	// The overseas country to query. You can perform a fuzzy match by using the country ID or the country name in Chinese or English.
	//
	// example:
	//
	// US
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// Instance ID of the WAF instance.
	//
	// > You can call the [DescribeInstance](https://help.aliyun.com/document_detail/433756.html) operation to query instance ID of the WAF instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xl*******005
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language used for the country and region names in the response. Valid values:
	//
	// - **en*	- (**default**): English.
	//
	// - **cn**: Simplified Chinese.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of entries to return on each page for a paging query. Valid values: 1 to 500. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The paging token that is used to query the next page. Set the value to the **NextToken*	- value returned from the previous API call. You do not need to specify this parameter for the first call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region where the WAF instance resides. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud resource group.
	//
	// example:
	//
	// rg-aekzwwk****cv5i
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
}

func (s DescribeIpAbroadCountryInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpAbroadCountryInfosRequest) GoString() string {
	return s.String()
}

func (s *DescribeIpAbroadCountryInfosRequest) GetAbroadRegion() *string {
	return s.AbroadRegion
}

func (s *DescribeIpAbroadCountryInfosRequest) GetCountry() *string {
	return s.Country
}

func (s *DescribeIpAbroadCountryInfosRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIpAbroadCountryInfosRequest) GetLanguage() *string {
	return s.Language
}

func (s *DescribeIpAbroadCountryInfosRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeIpAbroadCountryInfosRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeIpAbroadCountryInfosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeIpAbroadCountryInfosRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *DescribeIpAbroadCountryInfosRequest) SetAbroadRegion(v string) *DescribeIpAbroadCountryInfosRequest {
	s.AbroadRegion = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosRequest) SetCountry(v string) *DescribeIpAbroadCountryInfosRequest {
	s.Country = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosRequest) SetInstanceId(v string) *DescribeIpAbroadCountryInfosRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosRequest) SetLanguage(v string) *DescribeIpAbroadCountryInfosRequest {
	s.Language = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosRequest) SetMaxResults(v int32) *DescribeIpAbroadCountryInfosRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosRequest) SetNextToken(v string) *DescribeIpAbroadCountryInfosRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosRequest) SetRegionId(v string) *DescribeIpAbroadCountryInfosRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosRequest) SetResourceManagerResourceGroupId(v string) *DescribeIpAbroadCountryInfosRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosRequest) Validate() error {
	return dara.Validate(s)
}
