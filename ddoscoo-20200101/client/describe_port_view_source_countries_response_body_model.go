// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortViewSourceCountriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribePortViewSourceCountriesResponseBody
	GetRequestId() *string
	SetSourceCountrys(v []*DescribePortViewSourceCountriesResponseBodySourceCountrys) *DescribePortViewSourceCountriesResponseBody
	GetSourceCountrys() []*DescribePortViewSourceCountriesResponseBodySourceCountrys
}

type DescribePortViewSourceCountriesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the country or area from which the requests are sent.
	SourceCountrys []*DescribePortViewSourceCountriesResponseBodySourceCountrys `json:"SourceCountrys,omitempty" xml:"SourceCountrys,omitempty" type:"Repeated"`
}

func (s DescribePortViewSourceCountriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceCountriesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceCountriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortViewSourceCountriesResponseBody) GetSourceCountrys() []*DescribePortViewSourceCountriesResponseBodySourceCountrys {
	return s.SourceCountrys
}

func (s *DescribePortViewSourceCountriesResponseBody) SetRequestId(v string) *DescribePortViewSourceCountriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortViewSourceCountriesResponseBody) SetSourceCountrys(v []*DescribePortViewSourceCountriesResponseBodySourceCountrys) *DescribePortViewSourceCountriesResponseBody {
	s.SourceCountrys = v
	return s
}

func (s *DescribePortViewSourceCountriesResponseBody) Validate() error {
	if s.SourceCountrys != nil {
		for _, item := range s.SourceCountrys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePortViewSourceCountriesResponseBodySourceCountrys struct {
	// The number of requests.
	//
	// example:
	//
	// 3390671
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The abbreviation of the country or area. For example, **cn*	- indicates China and **us*	- indicates the United States.
	//
	// > For more information, see [Location parameters](https://help.aliyun.com/document_detail/167926.html).
	//
	// example:
	//
	// cn
	CountryId *string `json:"CountryId,omitempty" xml:"CountryId,omitempty"`
}

func (s DescribePortViewSourceCountriesResponseBodySourceCountrys) String() string {
	return dara.Prettify(s)
}

func (s DescribePortViewSourceCountriesResponseBodySourceCountrys) GoString() string {
	return s.String()
}

func (s *DescribePortViewSourceCountriesResponseBodySourceCountrys) GetCount() *int64 {
	return s.Count
}

func (s *DescribePortViewSourceCountriesResponseBodySourceCountrys) GetCountryId() *string {
	return s.CountryId
}

func (s *DescribePortViewSourceCountriesResponseBodySourceCountrys) SetCount(v int64) *DescribePortViewSourceCountriesResponseBodySourceCountrys {
	s.Count = &v
	return s
}

func (s *DescribePortViewSourceCountriesResponseBodySourceCountrys) SetCountryId(v string) *DescribePortViewSourceCountriesResponseBodySourceCountrys {
	s.CountryId = &v
	return s
}

func (s *DescribePortViewSourceCountriesResponseBodySourceCountrys) Validate() error {
	return dara.Validate(s)
}
