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
	// The DDoS attack events.
	DataList []*DescribeDDoSAllEventListResponseBodyDataList `json:"DataList,omitempty" xml:"DataList,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
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
	// The website ID.
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
	// The peak of volumetric attacks. Unit: bit/s.
	//
	// example:
	//
	// 800
	Bps      *int64  `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Coverage *string `json:"Coverage,omitempty" xml:"Coverage,omitempty"`
	// The peak of connection flood attacks. Unit: connections per seconds (CPS).
	//
	// example:
	//
	// 50
	Cps *int64 `json:"Cps,omitempty" xml:"Cps,omitempty"`
	// The time when the DDoS attack ends.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2023-02-12T15:59:59Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The attack event ID.
	//
	// example:
	//
	// web-cc_1
	EventId     *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	EventResult *string `json:"EventResult,omitempty" xml:"EventResult,omitempty"`
	// The type of DDoS attacks that was queried. Valid values:
	//
	// 	- **web-cc**: web resource exhaustion attacks.
	//
	// 	- **cc**: connection flood attacks.
	//
	// 	- **traffic**: volumetric attacks.
	//
	// example:
	//
	// web-cc
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The peak of volumetric attacks. Unit: packets per second (PPS).
	//
	// example:
	//
	// 12000
	Pps *int64 `json:"Pps,omitempty" xml:"Pps,omitempty"`
	// The peak QPS of web resource exhaustion attacks.
	//
	// example:
	//
	// 7692
	Qps *int64 `json:"Qps,omitempty" xml:"Qps,omitempty"`
	// The time when the DDoS attack starts.
	//
	// The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
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
	// The ID of the web resource exhaustion attack target.
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
