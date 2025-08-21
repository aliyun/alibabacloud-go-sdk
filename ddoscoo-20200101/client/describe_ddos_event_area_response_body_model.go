// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDosEventAreaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAreas(v []*DescribeDDosEventAreaResponseBodyAreas) *DescribeDDosEventAreaResponseBody
	GetAreas() []*DescribeDDosEventAreaResponseBodyAreas
	SetRequestId(v string) *DescribeDDosEventAreaResponseBody
	GetRequestId() *string
}

type DescribeDDosEventAreaResponseBody struct {
	// The information about the source region from which the volumetric attack was initiated.
	Areas []*DescribeDDosEventAreaResponseBodyAreas `json:"Areas,omitempty" xml:"Areas,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 11710C9F-BC5E-481A-BEC5-C6D8FBFCA827
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDDosEventAreaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventAreaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAreaResponseBody) GetAreas() []*DescribeDDosEventAreaResponseBodyAreas {
	return s.Areas
}

func (s *DescribeDDosEventAreaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDosEventAreaResponseBody) SetAreas(v []*DescribeDDosEventAreaResponseBodyAreas) *DescribeDDosEventAreaResponseBody {
	s.Areas = v
	return s
}

func (s *DescribeDDosEventAreaResponseBody) SetRequestId(v string) *DescribeDDosEventAreaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDosEventAreaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDDosEventAreaResponseBodyAreas struct {
	// The code or ID of the source region. For more information, see [Codes of administrative regions in China and codes of countries and areas](https://help.aliyun.com/document_detail/167926.html). For example, **110000*	- indicates Beijing, China, and **us*	- indicates the United States.
	//
	// example:
	//
	// 110000
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The number of request packets that were sent from the source region.
	//
	// example:
	//
	// 228
	InPkts *int64 `json:"InPkts,omitempty" xml:"InPkts,omitempty"`
}

func (s DescribeDDosEventAreaResponseBodyAreas) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDosEventAreaResponseBodyAreas) GoString() string {
	return s.String()
}

func (s *DescribeDDosEventAreaResponseBodyAreas) GetArea() *string {
	return s.Area
}

func (s *DescribeDDosEventAreaResponseBodyAreas) GetInPkts() *int64 {
	return s.InPkts
}

func (s *DescribeDDosEventAreaResponseBodyAreas) SetArea(v string) *DescribeDDosEventAreaResponseBodyAreas {
	s.Area = &v
	return s
}

func (s *DescribeDDosEventAreaResponseBodyAreas) SetInPkts(v int64) *DescribeDDosEventAreaResponseBodyAreas {
	s.InPkts = &v
	return s
}

func (s *DescribeDDosEventAreaResponseBodyAreas) Validate() error {
	return dara.Validate(s)
}
