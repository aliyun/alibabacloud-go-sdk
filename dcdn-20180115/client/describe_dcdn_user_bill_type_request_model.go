// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserBillTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDcdnUserBillTypeRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnUserBillTypeRequest
	GetStartTime() *string
}

type DescribeDcdnUserBillTypeRequest struct {
	// The end of the time range to query.
	//
	// The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-10-31T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query.
	//
	// Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnUserBillTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillTypeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnUserBillTypeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnUserBillTypeRequest) SetEndTime(v string) *DescribeDcdnUserBillTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserBillTypeRequest) SetStartTime(v string) *DescribeDcdnUserBillTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnUserBillTypeRequest) Validate() error {
	return dara.Validate(s)
}
