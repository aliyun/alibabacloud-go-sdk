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
	// The information about the attack source regions.
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
	if s.Areas != nil {
		for _, item := range s.Areas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDDosEventAreaResponseBodyAreas struct {
	// The code of the attack source region, including Chinese regions (at the province or city level) and international regions (at the country level). For example, **110000*	- indicates Beijing, China, and **us*	- indicates the United States.
	//
	// > For more information about the valid values of this parameter, see [China and international regions](https://help.aliyun.com/document_detail/167926.html).
	//
	// example:
	//
	// 110000
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The number of request packets from the attack source region.
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
