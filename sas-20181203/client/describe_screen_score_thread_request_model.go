// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScreenScoreThreadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeScreenScoreThreadRequest
	GetEndTime() *int64
	SetSource(v int32) *DescribeScreenScoreThreadRequest
	GetSource() *int32
	SetStartTime(v int64) *DescribeScreenScoreThreadRequest
	GetStartTime() *int64
}

type DescribeScreenScoreThreadRequest struct {
	// The end of the time range to query. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1668064495000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Source of security score, default is Cloud Security Center if left empty. Enum values:
	//
	// - 0:Cloud Security Center.
	//
	// - 1:Yaochi Console.
	//
	// example:
	//
	// 0
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// The beginning of the time range to query. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1651290987000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeScreenScoreThreadRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScreenScoreThreadRequest) GoString() string {
	return s.String()
}

func (s *DescribeScreenScoreThreadRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeScreenScoreThreadRequest) GetSource() *int32 {
	return s.Source
}

func (s *DescribeScreenScoreThreadRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeScreenScoreThreadRequest) SetEndTime(v int64) *DescribeScreenScoreThreadRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeScreenScoreThreadRequest) SetSource(v int32) *DescribeScreenScoreThreadRequest {
	s.Source = &v
	return s
}

func (s *DescribeScreenScoreThreadRequest) SetStartTime(v int64) *DescribeScreenScoreThreadRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeScreenScoreThreadRequest) Validate() error {
	return dara.Validate(s)
}
