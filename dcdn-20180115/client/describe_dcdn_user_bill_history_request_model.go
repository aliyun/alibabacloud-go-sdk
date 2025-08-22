// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserBillHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeDcdnUserBillHistoryRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeDcdnUserBillHistoryRequest
	GetStartTime() *string
}

type DescribeDcdnUserBillHistoryRequest struct {
	// The end of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The end time must be later than the start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-10-31T16:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// The minimum data granularity is 5 minutes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-09-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDcdnUserBillHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserBillHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserBillHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDcdnUserBillHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDcdnUserBillHistoryRequest) SetEndTime(v string) *DescribeDcdnUserBillHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryRequest) SetStartTime(v string) *DescribeDcdnUserBillHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDcdnUserBillHistoryRequest) Validate() error {
	return dara.Validate(s)
}
