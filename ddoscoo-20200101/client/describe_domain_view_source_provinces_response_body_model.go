// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainViewSourceProvincesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDomainViewSourceProvincesResponseBody
	GetRequestId() *string
	SetSourceProvinces(v []*DescribeDomainViewSourceProvincesResponseBodySourceProvinces) *DescribeDomainViewSourceProvincesResponseBody
	GetSourceProvinces() []*DescribeDomainViewSourceProvincesResponseBodySourceProvinces
}

type DescribeDomainViewSourceProvincesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array consisting of the details of the administrative region in China from which the requests are sent.
	SourceProvinces []*DescribeDomainViewSourceProvincesResponseBodySourceProvinces `json:"SourceProvinces,omitempty" xml:"SourceProvinces,omitempty" type:"Repeated"`
}

func (s DescribeDomainViewSourceProvincesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewSourceProvincesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceProvincesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainViewSourceProvincesResponseBody) GetSourceProvinces() []*DescribeDomainViewSourceProvincesResponseBodySourceProvinces {
	return s.SourceProvinces
}

func (s *DescribeDomainViewSourceProvincesResponseBody) SetRequestId(v string) *DescribeDomainViewSourceProvincesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponseBody) SetSourceProvinces(v []*DescribeDomainViewSourceProvincesResponseBodySourceProvinces) *DescribeDomainViewSourceProvincesResponseBody {
	s.SourceProvinces = v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponseBody) Validate() error {
	if s.SourceProvinces != nil {
		for _, item := range s.SourceProvinces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainViewSourceProvincesResponseBodySourceProvinces struct {
	// The total number of requests.
	//
	// example:
	//
	// 3390671
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the region inside China. For more information, see the **Codes of administrative regions in China*	- section of the [Codes of administrative regions in China and codes of countries and areas](https://help.aliyun.com/document_detail/167926.html) topic. For example, **110000*	- indicates Beijing, and **120000*	- indicates Tianjin.
	//
	// example:
	//
	// 440000
	ProvinceId *string `json:"ProvinceId,omitempty" xml:"ProvinceId,omitempty"`
}

func (s DescribeDomainViewSourceProvincesResponseBodySourceProvinces) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainViewSourceProvincesResponseBodySourceProvinces) GoString() string {
	return s.String()
}

func (s *DescribeDomainViewSourceProvincesResponseBodySourceProvinces) GetCount() *int64 {
	return s.Count
}

func (s *DescribeDomainViewSourceProvincesResponseBodySourceProvinces) GetProvinceId() *string {
	return s.ProvinceId
}

func (s *DescribeDomainViewSourceProvincesResponseBodySourceProvinces) SetCount(v int64) *DescribeDomainViewSourceProvincesResponseBodySourceProvinces {
	s.Count = &v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponseBodySourceProvinces) SetProvinceId(v string) *DescribeDomainViewSourceProvincesResponseBodySourceProvinces {
	s.ProvinceId = &v
	return s
}

func (s *DescribeDomainViewSourceProvincesResponseBodySourceProvinces) Validate() error {
	return dara.Validate(s)
}
