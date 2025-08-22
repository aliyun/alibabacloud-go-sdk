// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnWafGeoInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*DescribeDcdnWafGeoInfoResponseBodyContent) *DescribeDcdnWafGeoInfoResponseBody
	GetContent() []*DescribeDcdnWafGeoInfoResponseBodyContent
	SetRequestId(v string) *DescribeDcdnWafGeoInfoResponseBody
	GetRequestId() *string
}

type DescribeDcdnWafGeoInfoResponseBody struct {
	// The type of information about the country or region.
	Content []*DescribeDcdnWafGeoInfoResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 66A98669-CC6E-4F3E-80A6-3014697B11AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnWafGeoInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGeoInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGeoInfoResponseBody) GetContent() []*DescribeDcdnWafGeoInfoResponseBodyContent {
	return s.Content
}

func (s *DescribeDcdnWafGeoInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnWafGeoInfoResponseBody) SetContent(v []*DescribeDcdnWafGeoInfoResponseBodyContent) *DescribeDcdnWafGeoInfoResponseBody {
	s.Content = v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponseBody) SetRequestId(v string) *DescribeDcdnWafGeoInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafGeoInfoResponseBodyContent struct {
	// The information about the country or region.
	Continents []*DescribeDcdnWafGeoInfoResponseBodyContentContinents `json:"Continents,omitempty" xml:"Continents,omitempty" type:"Repeated"`
	// The type of the region.
	//
	// 	- CN: China
	//
	// 	- Other: outside China
	//
	// example:
	//
	// CN
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDcdnWafGeoInfoResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGeoInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContent) GetContinents() []*DescribeDcdnWafGeoInfoResponseBodyContentContinents {
	return s.Continents
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContent) GetType() *string {
	return s.Type
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContent) SetContinents(v []*DescribeDcdnWafGeoInfoResponseBodyContentContinents) *DescribeDcdnWafGeoInfoResponseBodyContent {
	s.Continents = v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContent) SetType(v string) *DescribeDcdnWafGeoInfoResponseBodyContent {
	s.Type = &v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContent) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafGeoInfoResponseBodyContentContinents struct {
	// The district to which the country or region belongs.
	//
	// example:
	//
	// China
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region information.
	Regions []*DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeDcdnWafGeoInfoResponseBodyContentContinents) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGeoInfoResponseBodyContentContinents) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinents) GetName() *string {
	return s.Name
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinents) GetRegions() []*DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions {
	return s.Regions
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinents) SetName(v string) *DescribeDcdnWafGeoInfoResponseBodyContentContinents {
	s.Name = &v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinents) SetRegions(v []*DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions) *DescribeDcdnWafGeoInfoResponseBodyContentContinents {
	s.Regions = v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinents) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions struct {
	// The name of the country or region.
	//
	// example:
	//
	// Beijing
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The code of the country or region.
	//
	// example:
	//
	// 110000
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions) GoString() string {
	return s.String()
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions) GetName() *string {
	return s.Name
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions) GetValue() *string {
	return s.Value
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions) SetName(v string) *DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions {
	s.Name = &v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions) SetValue(v string) *DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions {
	s.Value = &v
	return s
}

func (s *DescribeDcdnWafGeoInfoResponseBodyContentContinentsRegions) Validate() error {
	return dara.Validate(s)
}
