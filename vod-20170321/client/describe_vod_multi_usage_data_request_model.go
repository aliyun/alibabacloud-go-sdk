// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodMultiUsageDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeVodMultiUsageDataRequest
	GetOwnerId() *int64
	SetTimePoint(v string) *DescribeVodMultiUsageDataRequest
	GetTimePoint() *string
}

type DescribeVodMultiUsageDataRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	TimePoint *string `json:"TimePoint,omitempty" xml:"TimePoint,omitempty"`
}

func (s DescribeVodMultiUsageDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodMultiUsageDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodMultiUsageDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodMultiUsageDataRequest) GetTimePoint() *string {
	return s.TimePoint
}

func (s *DescribeVodMultiUsageDataRequest) SetOwnerId(v int64) *DescribeVodMultiUsageDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodMultiUsageDataRequest) SetTimePoint(v string) *DescribeVodMultiUsageDataRequest {
	s.TimePoint = &v
	return s
}

func (s *DescribeVodMultiUsageDataRequest) Validate() error {
	return dara.Validate(s)
}
