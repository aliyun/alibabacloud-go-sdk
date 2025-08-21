// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDdosEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDdosEventList(v *DescribeDdosEventListResponseBodyDdosEventList) *DescribeDdosEventListResponseBody
	GetDdosEventList() *DescribeDdosEventListResponseBodyDdosEventList
	SetRequestId(v string) *DescribeDdosEventListResponseBody
	GetRequestId() *string
	SetTotal(v int32) *DescribeDdosEventListResponseBody
	GetTotal() *int32
}

type DescribeDdosEventListResponseBody struct {
	// The details of the DDoS attack events.
	DdosEventList *DescribeDdosEventListResponseBodyDdosEventList `json:"DdosEventList,omitempty" xml:"DdosEventList,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// BC0907F8-A9F3-5E11-977B-D59CD98C64ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DDoS attack events.
	//
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeDdosEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBody) GetDdosEventList() *DescribeDdosEventListResponseBodyDdosEventList {
	return s.DdosEventList
}

func (s *DescribeDdosEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDdosEventListResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *DescribeDdosEventListResponseBody) SetDdosEventList(v *DescribeDdosEventListResponseBodyDdosEventList) *DescribeDdosEventListResponseBody {
	s.DdosEventList = v
	return s
}

func (s *DescribeDdosEventListResponseBody) SetRequestId(v string) *DescribeDdosEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDdosEventListResponseBody) SetTotal(v int32) *DescribeDdosEventListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeDdosEventListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosEventListResponseBodyDdosEventList struct {
	DdosEvent []*DescribeDdosEventListResponseBodyDdosEventListDdosEvent `json:"DdosEvent,omitempty" xml:"DdosEvent,omitempty" type:"Repeated"`
}

func (s DescribeDdosEventListResponseBodyDdosEventList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosEventListResponseBodyDdosEventList) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) GetDdosEvent() []*DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	return s.DdosEvent
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) SetDdosEvent(v []*DescribeDdosEventListResponseBodyDdosEventListDdosEvent) *DescribeDdosEventListResponseBodyDdosEventList {
	s.DdosEvent = v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventList) Validate() error {
	return dara.Validate(s)
}

type DescribeDdosEventListResponseBodyDdosEventListDdosEvent struct {
	// The status of the DDoS attack event. Valid values:
	//
	// 	- **mitigating**: indicates that traffic scrubbing is in progress.
	//
	// 	- **blackholed**: indicates that blackhole filtering is triggered for the asset.
	//
	// 	- **normal**: indicates that the DDoS attack event ends.
	//
	// example:
	//
	// normal
	DdosStatus *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	// The type of the DDoS attack event. Valid values:
	//
	// 	- **defense**: an attack event that triggers traffic scrubbing
	//
	// 	- **blackhole**: an attack event that triggers blackhole filtering
	//
	// example:
	//
	// blackhole
	DdosType *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	// The time of the last attack. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// > This parameter is returned only when the asset is attacked multiple times within a DDoS attack event.
	//
	// example:
	//
	// 1637817679000
	DelayTime *int64 `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	// The end time of the DDoS attack event. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1637817679000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the DDoS attack event. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1637812279000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when blackhole filtering is deactivated. The value is a UNIX timestamp. Unit: milliseconds.
	//
	// > This parameter is returned only when the value of the **DdosType*	- parameter is **blackhole**.
	//
	// example:
	//
	// 1637814079000
	UnBlackholeTime *int64 `json:"UnBlackholeTime,omitempty" xml:"UnBlackholeTime,omitempty"`
}

func (s DescribeDdosEventListResponseBodyDdosEventListDdosEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeDdosEventListResponseBodyDdosEventListDdosEvent) GoString() string {
	return s.String()
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) GetDdosStatus() *string {
	return s.DdosStatus
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) GetDdosType() *string {
	return s.DdosType
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) GetDelayTime() *int64 {
	return s.DelayTime
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) GetUnBlackholeTime() *int64 {
	return s.UnBlackholeTime
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetDdosStatus(v string) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.DdosStatus = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetDdosType(v string) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.DdosType = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetDelayTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.DelayTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetEndTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.EndTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetStartTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.StartTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) SetUnBlackholeTime(v int64) *DescribeDdosEventListResponseBodyDdosEventListDdosEvent {
	s.UnBlackholeTime = &v
	return s
}

func (s *DescribeDdosEventListResponseBodyDdosEventListDdosEvent) Validate() error {
	return dara.Validate(s)
}
