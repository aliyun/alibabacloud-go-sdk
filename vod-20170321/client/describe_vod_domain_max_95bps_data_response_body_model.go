// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainMax95BpsDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetailData(v *DescribeVodDomainMax95BpsDataResponseBodyDetailData) *DescribeVodDomainMax95BpsDataResponseBody
	GetDetailData() *DescribeVodDomainMax95BpsDataResponseBodyDetailData
	SetDomainName(v string) *DescribeVodDomainMax95BpsDataResponseBody
	GetDomainName() *string
	SetDomesticMax95Bps(v string) *DescribeVodDomainMax95BpsDataResponseBody
	GetDomesticMax95Bps() *string
	SetEndTime(v string) *DescribeVodDomainMax95BpsDataResponseBody
	GetEndTime() *string
	SetMax95Bps(v string) *DescribeVodDomainMax95BpsDataResponseBody
	GetMax95Bps() *string
	SetOverseasMax95Bps(v string) *DescribeVodDomainMax95BpsDataResponseBody
	GetOverseasMax95Bps() *string
	SetRequestId(v string) *DescribeVodDomainMax95BpsDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeVodDomainMax95BpsDataResponseBody
	GetStartTime() *string
}

type DescribeVodDomainMax95BpsDataResponseBody struct {
	// Details of the 95th percentile bandwidth.
	DetailData *DescribeVodDomainMax95BpsDataResponseBodyDetailData `json:"DetailData,omitempty" xml:"DetailData,omitempty" type:"Struct"`
	// The domain name for CDN.
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
	// The end of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// >  The end time must be later than the start time.
	//
	// example:
	//
	// 2017-01-11T13:00:00Z
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
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A*****F6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodDomainMax95BpsDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainMax95BpsDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) GetDetailData() *DescribeVodDomainMax95BpsDataResponseBodyDetailData {
	return s.DetailData
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) GetDomesticMax95Bps() *string {
	return s.DomesticMax95Bps
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) GetMax95Bps() *string {
	return s.Max95Bps
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) GetOverseasMax95Bps() *string {
	return s.OverseasMax95Bps
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) SetDetailData(v *DescribeVodDomainMax95BpsDataResponseBodyDetailData) *DescribeVodDomainMax95BpsDataResponseBody {
	s.DetailData = v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) SetDomainName(v string) *DescribeVodDomainMax95BpsDataResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) SetDomesticMax95Bps(v string) *DescribeVodDomainMax95BpsDataResponseBody {
	s.DomesticMax95Bps = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) SetEndTime(v string) *DescribeVodDomainMax95BpsDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) SetMax95Bps(v string) *DescribeVodDomainMax95BpsDataResponseBody {
	s.Max95Bps = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) SetOverseasMax95Bps(v string) *DescribeVodDomainMax95BpsDataResponseBody {
	s.OverseasMax95Bps = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) SetRequestId(v string) *DescribeVodDomainMax95BpsDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) SetStartTime(v string) *DescribeVodDomainMax95BpsDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainMax95BpsDataResponseBodyDetailData struct {
	Max95Detail []*DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail `json:"Max95Detail,omitempty" xml:"Max95Detail,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainMax95BpsDataResponseBodyDetailData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainMax95BpsDataResponseBodyDetailData) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailData) GetMax95Detail() []*DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	return s.Max95Detail
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailData) SetMax95Detail(v []*DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) *DescribeVodDomainMax95BpsDataResponseBodyDetailData {
	s.Max95Detail = v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail struct {
	// The billable region where the peak 95 data was collected.
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
	// The time corresponding to the 95th percentile bandwidth peak.
	//
	// example:
	//
	// 2015-12-11T21:05:00Z
	Max95BpsPeakTime *string `json:"Max95BpsPeakTime,omitempty" xml:"Max95BpsPeakTime,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2024-01-18 10:11:32
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GetArea() *string {
	return s.Area
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GetMax95Bps() *float32 {
	return s.Max95Bps
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GetMax95BpsPeakTime() *string {
	return s.Max95BpsPeakTime
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) SetArea(v string) *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	s.Area = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) SetMax95Bps(v float32) *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	s.Max95Bps = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) SetMax95BpsPeakTime(v string) *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	s.Max95BpsPeakTime = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) SetTimeStamp(v string) *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail {
	s.TimeStamp = &v
	return s
}

func (s *DescribeVodDomainMax95BpsDataResponseBodyDetailDataMax95Detail) Validate() error {
	return dara.Validate(s)
}
