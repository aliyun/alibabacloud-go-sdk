// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIpAbroadCountryInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAbroadInfos(v []*DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) *DescribeIpAbroadCountryInfosResponseBody
	GetAbroadInfos() []*DescribeIpAbroadCountryInfosResponseBodyAbroadInfos
	SetMaxResults(v int32) *DescribeIpAbroadCountryInfosResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeIpAbroadCountryInfosResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeIpAbroadCountryInfosResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeIpAbroadCountryInfosResponseBody
	GetTotalCount() *int32
}

type DescribeIpAbroadCountryInfosResponseBody struct {
	AbroadInfos []*DescribeIpAbroadCountryInfosResponseBodyAbroadInfos `json:"AbroadInfos,omitempty" xml:"AbroadInfos,omitempty" type:"Repeated"`
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
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeIpAbroadCountryInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpAbroadCountryInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIpAbroadCountryInfosResponseBody) GetAbroadInfos() []*DescribeIpAbroadCountryInfosResponseBodyAbroadInfos {
	return s.AbroadInfos
}

func (s *DescribeIpAbroadCountryInfosResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeIpAbroadCountryInfosResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeIpAbroadCountryInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIpAbroadCountryInfosResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeIpAbroadCountryInfosResponseBody) SetAbroadInfos(v []*DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) *DescribeIpAbroadCountryInfosResponseBody {
	s.AbroadInfos = v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBody) SetMaxResults(v int32) *DescribeIpAbroadCountryInfosResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBody) SetNextToken(v string) *DescribeIpAbroadCountryInfosResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBody) SetRequestId(v string) *DescribeIpAbroadCountryInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBody) SetTotalCount(v int32) *DescribeIpAbroadCountryInfosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBody) Validate() error {
	if s.AbroadInfos != nil {
		for _, item := range s.AbroadInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpAbroadCountryInfosResponseBodyAbroadInfos struct {
	// example:
	//
	// 北美洲
	Continent *string `json:"Continent,omitempty" xml:"Continent,omitempty"`
	// example:
	//
	// US
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// example:
	//
	// 美国
	CountryName *string                                                       `json:"CountryName,omitempty" xml:"CountryName,omitempty"`
	Regions     []*DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) GoString() string {
	return s.String()
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) GetContinent() *string {
	return s.Continent
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) GetCountry() *string {
	return s.Country
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) GetCountryName() *string {
	return s.CountryName
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) GetRegions() []*DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions {
	return s.Regions
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) SetContinent(v string) *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos {
	s.Continent = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) SetCountry(v string) *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos {
	s.Country = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) SetCountryName(v string) *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos {
	s.CountryName = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) SetRegions(v []*DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions) *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos {
	s.Regions = v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfos) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions struct {
	// example:
	//
	// US-CA
	AbroadRegionId *string `json:"AbroadRegionId,omitempty" xml:"AbroadRegionId,omitempty"`
	// example:
	//
	// 加利福尼亚州
	AbroadRegionName *string `json:"AbroadRegionName,omitempty" xml:"AbroadRegionName,omitempty"`
}

func (s DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions) GoString() string {
	return s.String()
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions) GetAbroadRegionId() *string {
	return s.AbroadRegionId
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions) GetAbroadRegionName() *string {
	return s.AbroadRegionName
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions) SetAbroadRegionId(v string) *DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions {
	s.AbroadRegionId = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions) SetAbroadRegionName(v string) *DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions {
	s.AbroadRegionName = &v
	return s
}

func (s *DescribeIpAbroadCountryInfosResponseBodyAbroadInfosRegions) Validate() error {
	return dara.Validate(s)
}
