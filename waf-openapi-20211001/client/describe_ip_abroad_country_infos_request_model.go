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
	// example:
	//
	// US-CA
	AbroadRegion *string `json:"AbroadRegion,omitempty" xml:"AbroadRegion,omitempty"`
	// example:
	//
	// US
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// waf_elasticity-cn-0xl*******005
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0*****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
