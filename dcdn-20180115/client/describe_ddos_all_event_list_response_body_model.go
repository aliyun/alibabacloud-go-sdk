// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosAllEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeDdosAllEventListResponseBodyDataList) *DescribeDdosAllEventListResponseBody
	GetDataList() []*DescribeDdosAllEventListResponseBodyDataList
	SetPageNumber(v int32) *DescribeDdosAllEventListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDdosAllEventListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDdosAllEventListResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeDdosAllEventListResponseBody
	GetTotal() *int64
}

type DescribeDdosAllEventListResponseBody struct {
	// The list of events.
	DataList []*DescribeDdosAllEventListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned on each page. Default value: **10**. Valid values: 5, 10, and 20.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D73A4243-CFBD-5110-876F-09237E77ECBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 20
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosAllEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosAllEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosAllEventListResponseBody) GetDataList() []*DescribeDdosAllEventListResponseBodyDataList {
	return s.DataList
}

func (s *DescribeDdosAllEventListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDdosAllEventListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDdosAllEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDdosAllEventListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeDdosAllEventListResponseBody) SetDataList(v []*DescribeDdosAllEventListResponseBodyDataList) *DescribeDdosAllEventListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeDdosAllEventListResponseBody) SetPageNumber(v int32) *DescribeDdosAllEventListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBody) SetPageSize(v int32) *DescribeDdosAllEventListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBody) SetRequestId(v string) *DescribeDdosAllEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBody) SetTotal(v int64) *DescribeDdosAllEventListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosAllEventListResponseBodyDataList struct {
	// The peak attack traffic of volumetric attacks. Unit: bit/s.
	//
	// example:
	//
	// 800
	Bps *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The peak of connection flood attacks. Unit: connections per seconds (CPS).
	//
	// example:
	//
	// 50
	Cps *int64 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The end of the time range during which data was queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-26T15:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event ID.
	//
	// example:
	//
	// 28069
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The type of the DDoS attack event that was queried. Valid values:
	//
	// 	- **web-cc**: web resource exhaustion attacks
	//
	// 	- **cc**: connection flood attacks
	//
	// 	- **traffic**: volumetric attacks
	//
	// 	- If you do not configure this parameter, DDoS attack events of all types are queried.
	//
	// example:
	//
	// web-cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The peak attack traffic of volumetric attacks. Unit: packets per second (PPS).
	//
	// example:
	//
	// 12000
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The peak of web resource exhaustion attacks. Unit: queries per second (QPS).
	//
	// example:
	//
	// 7692
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The beginning of the time range during which data was queried. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-10-09T10:03:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The attack target.
	//
	// example:
	//
	// example.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s DescribeDdosAllEventListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosAllEventListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeDdosAllEventListResponseBodyDataList) GetBps() *int64 {
	return s.Bps
}

func (s *DescribeDdosAllEventListResponseBodyDataList) GetCps() *int64 {
	return s.Cps
}

func (s *DescribeDdosAllEventListResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDdosAllEventListResponseBodyDataList) GetEventId() *string {
	return s.EventId
}

func (s *DescribeDdosAllEventListResponseBodyDataList) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDdosAllEventListResponseBodyDataList) GetPps() *int64 {
	return s.Pps
}

func (s *DescribeDdosAllEventListResponseBodyDataList) GetQps() *int64 {
	return s.Qps
}

func (s *DescribeDdosAllEventListResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDdosAllEventListResponseBodyDataList) GetTarget() *string {
	return s.Target
}

func (s *DescribeDdosAllEventListResponseBodyDataList) SetBps(v int64) *DescribeDdosAllEventListResponseBodyDataList {
	s.Bps = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBodyDataList) SetCps(v int64) *DescribeDdosAllEventListResponseBodyDataList {
	s.Cps = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBodyDataList) SetEndTime(v string) *DescribeDdosAllEventListResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBodyDataList) SetEventId(v string) *DescribeDdosAllEventListResponseBodyDataList {
	s.EventId = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBodyDataList) SetEventType(v string) *DescribeDdosAllEventListResponseBodyDataList {
	s.EventType = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBodyDataList) SetPps(v int64) *DescribeDdosAllEventListResponseBodyDataList {
	s.Pps = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBodyDataList) SetQps(v int64) *DescribeDdosAllEventListResponseBodyDataList {
	s.Qps = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBodyDataList) SetStartTime(v string) *DescribeDdosAllEventListResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBodyDataList) SetTarget(v string) *DescribeDdosAllEventListResponseBodyDataList {
	s.Target = &v
	return s
}

func (s *DescribeDdosAllEventListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
