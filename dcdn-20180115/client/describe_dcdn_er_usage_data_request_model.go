// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnErUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDcdnErUsageDataRequest
	GetEndTime() *string
	SetRoutineID(v string) *DescribeDcdnErUsageDataRequest
	GetRoutineID() *string
	SetSpec(v string) *DescribeDcdnErUsageDataRequest
	GetSpec() *string
	SetSplitBy(v string) *DescribeDcdnErUsageDataRequest
	GetSplitBy() *string
	SetStartTime(v string) *DescribeDcdnErUsageDataRequest
	GetStartTime() *string
}

type DescribeDcdnErUsageDataRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > The end time must be later than the start time.
	//
	// example:
	//
	// 2018-10-31T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the routine.
	//
	// example:
	//
	// routine1.test
	RoutineID *string `json:"RoutineID,omitempty" xml:"RoutineID,omitempty"`
	// The specification of the routine. Valid values:
	//
	// 	- 5ms
	//
	// 	- 50ms
	//
	// 	- 100ms
	//
	// example:
	//
	// 50ms
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// Specifies how the results are grouped. If you set this parameter to routine, the returned results are grouped based on the routine ID. If you set this parameter to spec, the returned results are grouped based on the routine specification.
	//
	// > If you leave this parameter empty, the returned results are not grouped.
	//
	// example:
	//
	// routine
	SplitBy *string `json:"SplitBy,omitempty" xml:"SplitBy,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2018-10-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnErUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnErUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnErUsageDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnErUsageDataRequest) GetRoutineID() *string {
	return s.RoutineID
}

func (s *DescribeDcdnErUsageDataRequest) GetSpec() *string {
	return s.Spec
}

func (s *DescribeDcdnErUsageDataRequest) GetSplitBy() *string {
	return s.SplitBy
}

func (s *DescribeDcdnErUsageDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnErUsageDataRequest) SetEndTime(v string) *DescribeDcdnErUsageDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnErUsageDataRequest) SetRoutineID(v string) *DescribeDcdnErUsageDataRequest {
	s.RoutineID = &v
	return s
}

func (s *DescribeDcdnErUsageDataRequest) SetSpec(v string) *DescribeDcdnErUsageDataRequest {
	s.Spec = &v
	return s
}

func (s *DescribeDcdnErUsageDataRequest) SetSplitBy(v string) *DescribeDcdnErUsageDataRequest {
	s.SplitBy = &v
	return s
}

func (s *DescribeDcdnErUsageDataRequest) SetStartTime(v string) *DescribeDcdnErUsageDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnErUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
