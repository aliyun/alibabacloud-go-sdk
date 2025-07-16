// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserBillTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeCdnUserBillTypeRequest
	GetEndTime() *string
	SetStartTime(v string) *DescribeCdnUserBillTypeRequest
	GetStartTime() *string
}

type DescribeCdnUserBillTypeRequest struct {
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

func (s DescribeCdnUserBillTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserBillTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserBillTypeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeCdnUserBillTypeRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCdnUserBillTypeRequest) SetEndTime(v string) *DescribeCdnUserBillTypeRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeCdnUserBillTypeRequest) SetStartTime(v string) *DescribeCdnUserBillTypeRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeCdnUserBillTypeRequest) Validate() error {
	return dara.Validate(s)
}
