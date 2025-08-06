// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodStatisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeVodStatisRequest
	GetEndTime() *string
	SetOwnerId(v int64) *DescribeVodStatisRequest
	GetOwnerId() *int64
	SetStartTime(v string) *DescribeVodStatisRequest
	GetStartTime() *string
}

type DescribeVodStatisRequest struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeVodStatisRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStatisRequest) GoString() string {
	return s.String()
}

func (s *DescribeVodStatisRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeVodStatisRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVodStatisRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeVodStatisRequest) SetEndTime(v string) *DescribeVodStatisRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeVodStatisRequest) SetOwnerId(v int64) *DescribeVodStatisRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVodStatisRequest) SetStartTime(v string) *DescribeVodStatisRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeVodStatisRequest) Validate() error {
	return dara.Validate(s)
}
