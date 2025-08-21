// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewSourceCountriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainViewSourceCountriesResponseBody
	GetRequestId() *string
	SetSourceCountrys(v []*DescribeDomainViewSourceCountriesResponseBodySourceCountrys) *DescribeDomainViewSourceCountriesResponseBody
	GetSourceCountrys() []*DescribeDomainViewSourceCountriesResponseBodySourceCountrys
}

type DescribeDomainViewSourceCountriesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the country or area from which the requests are sent.
	SourceCountrys []*DescribeDomainViewSourceCountriesResponseBodySourceCountrys `json:"SourceCountrys,omitempty" xml:"SourceCountrys,omitempty" type:"Repeated"`
}

func (s DescribeDomainViewSourceCountriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewSourceCountriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceCountriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainViewSourceCountriesResponseBody) GetSourceCountrys() []*DescribeDomainViewSourceCountriesResponseBodySourceCountrys {
	return s.SourceCountrys
}

func (s *DescribeDomainViewSourceCountriesResponseBody) SetRequestId(v string) *DescribeDomainViewSourceCountriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponseBody) SetSourceCountrys(v []*DescribeDomainViewSourceCountriesResponseBodySourceCountrys) *DescribeDomainViewSourceCountriesResponseBody {
	s.SourceCountrys = v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainViewSourceCountriesResponseBodySourceCountrys struct {
	// The total number of requests.
	//
	// example:
	//
	// 3390671
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The abbreviation of the country or area. For more information, see the **Codes of countries and areas*	- section of the [Codes of administrative regions in China and codes of countries and areas](https://help.aliyun.com/document_detail/167926.html) topic. For example, **cn*	- indicates China, and **us*	- indicates the United States.
	//
	// example:
	//
	// cn
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
}

func (s DescribeDomainViewSourceCountriesResponseBodySourceCountrys) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewSourceCountriesResponseBodySourceCountrys) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceCountriesResponseBodySourceCountrys) GetCount() *int64 {
	return s.Count
}

func (s *DescribeDomainViewSourceCountriesResponseBodySourceCountrys) GetCountryId() *string {
	return s.CountryId
}

func (s *DescribeDomainViewSourceCountriesResponseBodySourceCountrys) SetCount(v int64) *DescribeDomainViewSourceCountriesResponseBodySourceCountrys {
	s.Count = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponseBodySourceCountrys) SetCountryId(v string) *DescribeDomainViewSourceCountriesResponseBodySourceCountrys {
	s.CountryId = &v
	return s
}

func (s *DescribeDomainViewSourceCountriesResponseBodySourceCountrys) Validate() error {
	return dara.Validate(s)
}
