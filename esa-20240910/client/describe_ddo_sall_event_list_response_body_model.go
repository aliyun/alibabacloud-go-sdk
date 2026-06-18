// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDDoSAllEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataList(v []*DescribeDDoSAllEventListResponseBodyDataList) *DescribeDDoSAllEventListResponseBody
	GetDataList() []*DescribeDDoSAllEventListResponseBodyDataList
	SetPageNumber(v int32) *DescribeDDoSAllEventListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDDoSAllEventListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDDoSAllEventListResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DescribeDDoSAllEventListResponseBody
	GetSiteId() *int64
	SetTotalCount(v int32) *DescribeDDoSAllEventListResponseBody
	GetTotalCount() *int32
}

type DescribeDDoSAllEventListResponseBody struct {
	// A list of DDoS attack event details.
	DataList []*DescribeDDoSAllEventListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
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
	// The site ID.
	//
	// example:
	//
	// 7096621098****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDDoSAllEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSAllEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDDoSAllEventListResponseBody) GetDataList() []*DescribeDDoSAllEventListResponseBodyDataList {
	return s.DataList
}

func (s *DescribeDDoSAllEventListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDDoSAllEventListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDDoSAllEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDDoSAllEventListResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DescribeDDoSAllEventListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDDoSAllEventListResponseBody) SetDataList(v []*DescribeDDoSAllEventListResponseBodyDataList) *DescribeDDoSAllEventListResponseBody {
	s.DataList = v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetPageNumber(v int32) *DescribeDDoSAllEventListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetPageSize(v int32) *DescribeDDoSAllEventListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetRequestId(v string) *DescribeDDoSAllEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetSiteId(v int64) *DescribeDDoSAllEventListResponseBody {
	s.SiteId = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) SetTotalCount(v int32) *DescribeDDoSAllEventListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBody) Validate() error {
	if s.DataList != nil {
		for _, item := range s.DataList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDDoSAllEventListResponseBodyDataList struct {
	// The peak bits per second (Bps) of a volumetric attack.
	//
	// example:
	//
	// 800
	Bps *int64 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	// The attack region. Valid values:
	//
	// - **domestic**: Chinese mainland.
	//
	// - **global**: Global.
	//
	// - **overseas**: global (excluding Chinese mainland).
	//
	// example:
	//
	// domestic
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The peak connections per second (Cps) of a connection-based attack.
	//
	// example:
	//
	// 50
	Cps *int64 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The end time of the DDoS attack event.
	//
	// The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is in UTC.
	//
	// example:
	//
	// 2023-02-12T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The event ID.
	//
	// example:
	//
	// web-cc_1
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The event result. Valid values:
	//
	// - **clean**: The attack was successfully cleaned.
	//
	// - **ratelimit**: Rate limiting was applied.
	//
	// - **blackhole**: Blackhole filtering was triggered.
	//
	// example:
	//
	// clean
	EventResult *string `json:"EventResult,omitempty" xml:"EventResult,omitempty"`
	// The type of the DDoS attack event. Valid values:
	//
	// - **web-cc**: A web resource exhaustion attack.
	//
	// - **cc**: A connection-based attack.
	//
	// - **traffic**: A volumetric attack.
	//
	// example:
	//
	// web-cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The peak packets per second (Pps) of a volumetric attack.
	//
	// example:
	//
	// 12000
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The peak queries per second (Qps) of a web resource exhaustion attack.
	//
	// example:
	//
	// 7692
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The start time of the DDoS attack event.
	//
	// The time is in the `yyyy-MM-ddTHH:mm:ssZ` format and is in UTC.
	//
	// example:
	//
	// 2023-02-12T15:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The attack target.
	//
	// example:
	//
	// example.com
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The ID of the attack target.
	//
	// example:
	//
	// 000000000155****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s DescribeDDoSAllEventListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDDoSAllEventListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetBps() *int64 {
	return s.Bps
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetCoverage() *string {
	return s.Coverage
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetCps() *int64 {
	return s.Cps
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetEventId() *string {
	return s.EventId
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetEventResult() *string {
	return s.EventResult
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetEventType() *string {
	return s.EventType
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetPps() *int64 {
	return s.Pps
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetQps() *int64 {
	return s.Qps
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetTarget() *string {
	return s.Target
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) GetTargetId() *string {
	return s.TargetId
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetBps(v int64) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Bps = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetCoverage(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Coverage = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetCps(v int64) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Cps = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetEndTime(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.EndTime = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetEventId(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.EventId = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetEventResult(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.EventResult = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetEventType(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.EventType = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetPps(v int64) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Pps = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetQps(v int64) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Qps = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetStartTime(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.StartTime = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetTarget(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.Target = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) SetTargetId(v string) *DescribeDDoSAllEventListResponseBodyDataList {
	s.TargetId = &v
	return s
}

func (s *DescribeDDoSAllEventListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
