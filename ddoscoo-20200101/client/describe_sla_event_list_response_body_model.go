// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlaEventListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSlaEventListResponseBody
	GetRequestId() *string
	SetSlaEvent(v []*DescribeSlaEventListResponseBodySlaEvent) *DescribeSlaEventListResponseBody
	GetSlaEvent() []*DescribeSlaEventListResponseBodySlaEvent
	SetTotal(v int64) *DescribeSlaEventListResponseBody
	GetTotal() *int64
}

type DescribeSlaEventListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The destination rate limit events.
	SlaEvent []*DescribeSlaEventListResponseBodySlaEvent `json:"SlaEvent,omitempty" xml:"SlaEvent,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeSlaEventListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlaEventListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlaEventListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlaEventListResponseBody) GetSlaEvent() []*DescribeSlaEventListResponseBodySlaEvent {
	return s.SlaEvent
}

func (s *DescribeSlaEventListResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeSlaEventListResponseBody) SetRequestId(v string) *DescribeSlaEventListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlaEventListResponseBody) SetSlaEvent(v []*DescribeSlaEventListResponseBodySlaEvent) *DescribeSlaEventListResponseBody {
	s.SlaEvent = v
	return s
}

func (s *DescribeSlaEventListResponseBody) SetTotal(v int64) *DescribeSlaEventListResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeSlaEventListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSlaEventListResponseBodySlaEvent struct {
	// The end of the time range. Unit: seconds.
	//
	// example:
	//
	// 1671886740
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IP address of the instance.
	//
	// example:
	//
	// 203.107.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The region to which the destination IP address belongs. Valid values:
	//
	// 	- **cn**: a region in the Chinese mainland
	//
	// 	- **cn-hongkong**: China (Hong Kong)
	//
	// example:
	//
	// cn
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The beginning of the time range. Unit: seconds.
	//
	// example:
	//
	// 1678080840
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlaEventListResponseBodySlaEvent) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlaEventListResponseBodySlaEvent) GoString() string {
	return s.String()
}

func (s *DescribeSlaEventListResponseBodySlaEvent) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSlaEventListResponseBodySlaEvent) GetIp() *string {
	return s.Ip
}

func (s *DescribeSlaEventListResponseBodySlaEvent) GetRegion() *string {
	return s.Region
}

func (s *DescribeSlaEventListResponseBodySlaEvent) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSlaEventListResponseBodySlaEvent) SetEndTime(v int64) *DescribeSlaEventListResponseBodySlaEvent {
	s.EndTime = &v
	return s
}

func (s *DescribeSlaEventListResponseBodySlaEvent) SetIp(v string) *DescribeSlaEventListResponseBodySlaEvent {
	s.Ip = &v
	return s
}

func (s *DescribeSlaEventListResponseBodySlaEvent) SetRegion(v string) *DescribeSlaEventListResponseBodySlaEvent {
	s.Region = &v
	return s
}

func (s *DescribeSlaEventListResponseBodySlaEvent) SetStartTime(v int64) *DescribeSlaEventListResponseBodySlaEvent {
	s.StartTime = &v
	return s
}

func (s *DescribeSlaEventListResponseBodySlaEvent) Validate() error {
	return dara.Validate(s)
}
