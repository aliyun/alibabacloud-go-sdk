// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnBlockedRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInfoList(v *DescribeDcdnBlockedRegionsResponseBodyInfoList) *DescribeDcdnBlockedRegionsResponseBody
	GetInfoList() *DescribeDcdnBlockedRegionsResponseBodyInfoList
	SetRequestId(v string) *DescribeDcdnBlockedRegionsResponseBody
	GetRequestId() *string
}

type DescribeDcdnBlockedRegionsResponseBody struct {
	// The information about the country or region.
	InfoList *DescribeDcdnBlockedRegionsResponseBodyInfoList `json:"InfoList,omitempty" xml:"InfoList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BFFCDFAD-DACC-484E-9BE6-0AF3B3A0DD23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnBlockedRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsResponseBody) GetInfoList() *DescribeDcdnBlockedRegionsResponseBodyInfoList {
	return s.InfoList
}

func (s *DescribeDcdnBlockedRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnBlockedRegionsResponseBody) SetInfoList(v *DescribeDcdnBlockedRegionsResponseBodyInfoList) *DescribeDcdnBlockedRegionsResponseBody {
	s.InfoList = v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponseBody) SetRequestId(v string) *DescribeDcdnBlockedRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnBlockedRegionsResponseBodyInfoList struct {
	InfoItem []*DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem `json:"InfoItem,omitempty" xml:"InfoItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnBlockedRegionsResponseBodyInfoList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsResponseBodyInfoList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoList) GetInfoItem() []*DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem {
	return s.InfoItem
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoList) SetInfoItem(v []*DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) *DescribeDcdnBlockedRegionsResponseBodyInfoList {
	s.InfoItem = v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoList) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem struct {
	// The larger region to which the country or region belongs.
	//
	// example:
	//
	// Middle East
	Continent *string `json:"Continent,omitempty" xml:"Continent,omitempty"`
	// The abbreviation of the name of the country or region.
	//
	// example:
	//
	// AE
	CountriesAndRegions *string `json:"CountriesAndRegions,omitempty" xml:"CountriesAndRegions,omitempty"`
	// The name of the country or region.
	//
	// example:
	//
	// United Arab Emirates
	CountriesAndRegionsName *string `json:"CountriesAndRegionsName,omitempty" xml:"CountriesAndRegionsName,omitempty"`
}

func (s DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) GetContinent() *string {
	return s.Continent
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) GetCountriesAndRegions() *string {
	return s.CountriesAndRegions
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) GetCountriesAndRegionsName() *string {
	return s.CountriesAndRegionsName
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) SetContinent(v string) *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem {
	s.Continent = &v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) SetCountriesAndRegions(v string) *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem {
	s.CountriesAndRegions = &v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) SetCountriesAndRegionsName(v string) *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem {
	s.CountriesAndRegionsName = &v
	return s
}

func (s *DescribeDcdnBlockedRegionsResponseBodyInfoListInfoItem) Validate() error {
	return dara.Validate(s)
}
