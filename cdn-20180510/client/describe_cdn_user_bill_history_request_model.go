// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeCdnUserBillHistoryRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeCdnUserBillHistoryRequest
	GetStartTime() *string
}

type DescribeCdnUserBillHistoryRequest struct {
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

func (s DescribeCdnUserBillHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnUserBillHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnUserBillHistoryRequest) SetEndTime(v string) *DescribeCdnUserBillHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserBillHistoryRequest) SetStartTime(v string) *DescribeCdnUserBillHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillHistoryRequest) Validate() error {
	return dara.Validate(s)
}
