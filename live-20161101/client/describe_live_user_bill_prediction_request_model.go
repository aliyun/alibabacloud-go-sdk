// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserBillPredictionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeLiveUserBillPredictionRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeLiveUserBillPredictionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveUserBillPredictionRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeLiveUserBillPredictionRequest
	GetStartTime() *string
}

type DescribeLiveUserBillPredictionRequest struct {
	// The end time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2015-12-01T05:45:00Z
	EndTime  *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the ISO 8601 standard in the *yyyy-MM-dd**THH:mm:ssZ	- format. The time must be in UTC.
	//
	// example:
	//
	// 2015-12-01T05:40:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeLiveUserBillPredictionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserBillPredictionRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserBillPredictionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLiveUserBillPredictionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveUserBillPredictionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveUserBillPredictionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeLiveUserBillPredictionRequest) SetEndTime(v string) *DescribeLiveUserBillPredictionRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeLiveUserBillPredictionRequest) SetOwnerId(v int64) *DescribeLiveUserBillPredictionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveUserBillPredictionRequest) SetRegionId(v string) *DescribeLiveUserBillPredictionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveUserBillPredictionRequest) SetStartTime(v string) *DescribeLiveUserBillPredictionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeLiveUserBillPredictionRequest) Validate() error {
	return dara.Validate(s)
}
