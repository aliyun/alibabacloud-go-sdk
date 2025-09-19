// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeClusterNetworkRequest
	GetEndTime() *int64
	SetStartTime(v int64) *DescribeClusterNetworkRequest
	GetStartTime() *int64
}

type DescribeClusterNetworkRequest struct {
	// The end timestamp of the query. Unit: milliseconds.
	//
	// > The days between the start timestamp and the end timestamp cannot exceed **seven*	- days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656038940435
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start timestamp of the query. Unit: milliseconds.
	//
	// > The days between the start timestamp and the end timestamp cannot exceed **seven*	- days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1656038740435
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeClusterNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterNetworkRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNetworkRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeClusterNetworkRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeClusterNetworkRequest) SetEndTime(v int64) *DescribeClusterNetworkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeClusterNetworkRequest) SetStartTime(v int64) *DescribeClusterNetworkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeClusterNetworkRequest) Validate() error {
	return dara.Validate(s)
}
