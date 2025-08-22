// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnErUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDcdnErUsageDataResponseBody
	GetEndTime() *string
	SetErAccData(v *DescribeDcdnErUsageDataResponseBodyErAccData) *DescribeDcdnErUsageDataResponseBody
	GetErAccData() *DescribeDcdnErUsageDataResponseBodyErAccData
	SetRequestId(v string) *DescribeDcdnErUsageDataResponseBody
	GetRequestId() *string
	SetStartTime(v string) *DescribeDcdnErUsageDataResponseBody
	GetStartTime() *string
}

type DescribeDcdnErUsageDataResponseBody struct {
	// The end of the time range during which data was queried.
	//
	// example:
	//
	// 2018-10-31T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of the data returned.
	ErAccData *DescribeDcdnErUsageDataResponseBodyErAccData `json:"ErAccData,omitempty" xml:"ErAccData,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C370DAF1-C838-4288-A1A0-9A87633D248E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The start of the time range during which data was queried.
	//
	// example:
	//
	// 2018-10-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnErUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnErUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnErUsageDataResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnErUsageDataResponseBody) GetErAccData() *DescribeDcdnErUsageDataResponseBodyErAccData {
	return s.ErAccData
}

func (s *DescribeDcdnErUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnErUsageDataResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnErUsageDataResponseBody) SetEndTime(v string) *DescribeDcdnErUsageDataResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnErUsageDataResponseBody) SetErAccData(v *DescribeDcdnErUsageDataResponseBodyErAccData) *DescribeDcdnErUsageDataResponseBody {
	s.ErAccData = v
	return s
}

func (s *DescribeDcdnErUsageDataResponseBody) SetRequestId(v string) *DescribeDcdnErUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnErUsageDataResponseBody) SetStartTime(v string) *DescribeDcdnErUsageDataResponseBody {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnErUsageDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnErUsageDataResponseBodyErAccData struct {
	ErAccItem []*DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem `json:"ErAccItem,omitempty" xml:"ErAccItem,omitempty" type:"Repeated"`
}

func (s DescribeDcdnErUsageDataResponseBodyErAccData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnErUsageDataResponseBodyErAccData) GoString() string {
	return s.String()
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccData) GetErAccItem() []*DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem {
	return s.ErAccItem
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccData) SetErAccItem(v []*DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) *DescribeDcdnErUsageDataResponseBodyErAccData {
	s.ErAccItem = v
	return s
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccData) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem struct {
	// The number of requests.
	//
	// example:
	//
	// 125
	ErAcc *int64 `json:"ErAcc,omitempty" xml:"ErAcc,omitempty"`
	// The ID of the routine. This parameter is returned only when SplitBy is set to routine.
	//
	// example:
	//
	// routine1.test
	Routine *string `json:"Routine,omitempty" xml:"Routine,omitempty"`
	// The specification of the routine. This parameter is returned only when SplitBy is set to spec.
	//
	// example:
	//
	// 50ms
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The timestamp of the returned data.
	//
	// example:
	//
	// 2018-10-30T13:00:00Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) GoString() string {
	return s.String()
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) GetErAcc() *int64 {
	return s.ErAcc
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) GetRoutine() *string {
	return s.Routine
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) GetSpec() *string {
	return s.Spec
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) SetErAcc(v int64) *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem {
	s.ErAcc = &v
	return s
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) SetRoutine(v string) *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem {
	s.Routine = &v
	return s
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) SetSpec(v string) *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem {
	s.Spec = &v
	return s
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) SetTimeStamp(v string) *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem {
	s.TimeStamp = &v
	return s
}

func (s *DescribeDcdnErUsageDataResponseBodyErAccDataErAccItem) Validate() error {
	return dara.Validate(s)
}
