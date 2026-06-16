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
	DdosEventList *DescribeDdosEventListResponseBodyDdosEventList `json:"DdosEventList,omitempty" xml:"DdosEventList,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// BC0907F8-A9F3-5E11-977B-D59CD98C64ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of DDoS attack events found.
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
	if s.DdosEventList != nil {
		if err := s.DdosEventList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.DdosEvent != nil {
		for _, item := range s.DdosEvent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDdosEventListResponseBodyDdosEventListDdosEvent struct {
	DdosStatus      *string `json:"DdosStatus,omitempty" xml:"DdosStatus,omitempty"`
	DdosType        *string `json:"DdosType,omitempty" xml:"DdosType,omitempty"`
	DelayTime       *int64  `json:"DelayTime,omitempty" xml:"DelayTime,omitempty"`
	EndTime         *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime       *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	UnBlackholeTime *int64  `json:"UnBlackholeTime,omitempty" xml:"UnBlackholeTime,omitempty"`
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
