// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainMax95BpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetailData(v *DescribeDomainMax95BpsDataResponseBodyDetailData) *DescribeDomainMax95BpsDataResponseBody
	GetDetailData() *DescribeDomainMax95BpsDataResponseBodyDetailData
	SetDomainName(v string) *DescribeDomainMax95BpsDataResponseBody
	GetDomainName() *string
	SetDomesticMax95Bps(v string) *DescribeDomainMax95BpsDataResponseBody
	GetDomesticMax95Bps() *string
	SetEndTime(v string) *DescribeDomainMax95BpsDataResponseBody
	GetEndTime() *string
	SetMax95Bps(v string) *DescribeDomainMax95BpsDataResponseBody
	GetMax95Bps() *string
	SetOverseasMax95Bps(v string) *DescribeDomainMax95BpsDataResponseBody
	GetOverseasMax95Bps() *string
	SetRequestId(v string) *DescribeDomainMax95BpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDomainMax95BpsDataResponseBody
	GetStartTime() *string
}

type DescribeDomainMax95BpsDataResponseBody struct {
	// Details of the 95th percentile bandwidth.
	DetailData *DescribeDomainMax95BpsDataResponseBodyDetailData `json:"DetailData,omitempty" xml:"DetailData,omitempty" type:"Struct"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The 95th percentile bandwidth in the Chinese mainland.
	//
	// example:
	//
	// 16777590.28
	DomesticMax95Bps *string `json:"DomesticMax95Bps,omitempty" xml:"DomesticMax95Bps,omitempty"`
	// The end of the time range for which the data was queried.
	//
	// example:
	//
	// 2015-12-11T21:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The 95th percentile bandwidth.
	//
	// example:
	//
	// 16777590.28
	Max95Bps *string `json:"Max95Bps,omitempty" xml:"Max95Bps,omitempty"`
	// The 95th percentile bandwidth outside the Chinese mainland.
	//
	// example:
	//
	// 0
	OverseasMax95Bps *string `json:"OverseasMax95Bps,omitempty" xml:"OverseasMax95Bps,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 3C6CCEC4-6B88-4D4A-93E4-D47B3D92CF8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range for which the data was queried.
	//
	// example:
	//
	// 2015-12-10T20:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDomainMax95BpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMax95BpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainMax95BpsDataResponseBody) GetDetailData() *DescribeDomainMax95BpsDataResponseBodyDetailData {
	return s.DetailData
}

func (s *DescribeDomainMax95BpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainMax95BpsDataResponseBody) GetDomesticMax95Bps() *string {
	return s.DomesticMax95Bps
}

func (s *DescribeDomainMax95BpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDomainMax95BpsDataResponseBody) GetMax95Bps() *string {
	return s.Max95Bps
}

func (s *DescribeDomainMax95BpsDataResponseBody) GetOverseasMax95Bps() *string {
	return s.OverseasMax95Bps
}

func (s *DescribeDomainMax95BpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainMax95BpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetDetailData(v *DescribeDomainMax95BpsDataResponseBodyDetailData) *DescribeDomainMax95BpsDataResponseBody {
	s.DetailData = v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetDomainName(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetDomesticMax95Bps(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.DomesticMax95Bps = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetEndTime(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetMax95Bps(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.Max95Bps = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetOverseasMax95Bps(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.OverseasMax95Bps = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetRequestId(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) SetStartTime(v string) *DescribeDomainMax95BpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainMax95BpsDataResponseBodyDetailData struct {
	Max95Detail []*DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail `json:"Max95Detail,omitempty" xml:"Max95Detail,omitempty" type:"Repeated"`
}

func (s DescribeDomainMax95BpsDataResponseBodyDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMax95BpsDataResponseBodyDetailData) GoString() string {
	return s.String()
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailData) GetMax95Detail() []*DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	return s.Max95Detail
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailData) SetMax95Detail(v []*DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) *DescribeDomainMax95BpsDataResponseBodyDetailData {
	s.Max95Detail = v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailData) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail struct {
	// Region of the 95th percentile bandwidth.
	//
	// example:
	//
	// CN
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// The 95th percentile bandwidth.
	//
	// example:
	//
	// 16777590.28
	Max95Bps *float32 `json:"Max95Bps,omitempty" xml:"Max95Bps,omitempty"`
	// Time of the 95th percentile bandwidth.
	//
	// example:
	//
	// 2015-12-11T21:05:00Z
	Max95BpsPeakTime *string `json:"Max95BpsPeakTime,omitempty" xml:"Max95BpsPeakTime,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2015-12-11T21:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GoString() string {
	return s.String()
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GetArea() *string {
	return s.Area
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GetMax95Bps() *float32 {
	return s.Max95Bps
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GetMax95BpsPeakTime() *string {
	return s.Max95BpsPeakTime
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) SetArea(v string) *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	s.Area = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) SetMax95Bps(v float32) *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	s.Max95Bps = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) SetMax95BpsPeakTime(v string) *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	s.Max95BpsPeakTime = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) SetTimeStamp(v string) *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDomainMax95BpsDataResponseBodyDetailDataMax95Detail) Validate() error {
	return dara.Validate(s)
}
