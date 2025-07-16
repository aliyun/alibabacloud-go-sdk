// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlockedRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInfoList(v *DescribeBlockedRegionsResponseBodyInfoList) *DescribeBlockedRegionsResponseBody
	GetInfoList() *DescribeBlockedRegionsResponseBodyInfoList
	SetRequestId(v string) *DescribeBlockedRegionsResponseBody
	GetRequestId() *string
}

type DescribeBlockedRegionsResponseBody struct {
	// The information returned.
	InfoList *DescribeBlockedRegionsResponseBodyInfoList `json:"InfoList,omitempty" xml:"InfoList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BFFCDFAD-DACC-484E-9BE6-0AF3B3A0DD23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBlockedRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockedRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBlockedRegionsResponseBody) GetInfoList() *DescribeBlockedRegionsResponseBodyInfoList {
	return s.InfoList
}

func (s *DescribeBlockedRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBlockedRegionsResponseBody) SetInfoList(v *DescribeBlockedRegionsResponseBodyInfoList) *DescribeBlockedRegionsResponseBody {
	s.InfoList = v
	return s
}

func (s *DescribeBlockedRegionsResponseBody) SetRequestId(v string) *DescribeBlockedRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBlockedRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeBlockedRegionsResponseBodyInfoList struct {
	InfoItem []*DescribeBlockedRegionsResponseBodyInfoListInfoItem `json:"InfoItem,omitempty" xml:"InfoItem,omitempty" type:"Repeated"`
}

func (s DescribeBlockedRegionsResponseBodyInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockedRegionsResponseBodyInfoList) GoString() string {
	return s.String()
}

func (s *DescribeBlockedRegionsResponseBodyInfoList) GetInfoItem() []*DescribeBlockedRegionsResponseBodyInfoListInfoItem {
	return s.InfoItem
}

func (s *DescribeBlockedRegionsResponseBodyInfoList) SetInfoItem(v []*DescribeBlockedRegionsResponseBodyInfoListInfoItem) *DescribeBlockedRegionsResponseBodyInfoList {
	s.InfoItem = v
	return s
}

func (s *DescribeBlockedRegionsResponseBodyInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeBlockedRegionsResponseBodyInfoListInfoItem struct {
	// The district to which the country or region belongs.
	//
	// example:
	//
	// Asia
	Continent *string `json:"Continent,omitempty" xml:"Continent,omitempty"`
	// The abbreviation of the name of the country or region.
	//
	// example:
	//
	// AF
	CountriesAndRegions *string `json:"CountriesAndRegions,omitempty" xml:"CountriesAndRegions,omitempty"`
	// The name of the country or region.
	//
	// example:
	//
	// Afghanistan
	CountriesAndRegionsName *string `json:"CountriesAndRegionsName,omitempty" xml:"CountriesAndRegionsName,omitempty"`
}

func (s DescribeBlockedRegionsResponseBodyInfoListInfoItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlockedRegionsResponseBodyInfoListInfoItem) GoString() string {
	return s.String()
}

func (s *DescribeBlockedRegionsResponseBodyInfoListInfoItem) GetContinent() *string {
	return s.Continent
}

func (s *DescribeBlockedRegionsResponseBodyInfoListInfoItem) GetCountriesAndRegions() *string {
	return s.CountriesAndRegions
}

func (s *DescribeBlockedRegionsResponseBodyInfoListInfoItem) GetCountriesAndRegionsName() *string {
	return s.CountriesAndRegionsName
}

func (s *DescribeBlockedRegionsResponseBodyInfoListInfoItem) SetContinent(v string) *DescribeBlockedRegionsResponseBodyInfoListInfoItem {
	s.Continent = &v
	return s
}

func (s *DescribeBlockedRegionsResponseBodyInfoListInfoItem) SetCountriesAndRegions(v string) *DescribeBlockedRegionsResponseBodyInfoListInfoItem {
	s.CountriesAndRegions = &v
	return s
}

func (s *DescribeBlockedRegionsResponseBodyInfoListInfoItem) SetCountriesAndRegionsName(v string) *DescribeBlockedRegionsResponseBodyInfoListInfoItem {
	s.CountriesAndRegionsName = &v
	return s
}

func (s *DescribeBlockedRegionsResponseBodyInfoListInfoItem) Validate() error {
	return dara.Validate(s)
}
