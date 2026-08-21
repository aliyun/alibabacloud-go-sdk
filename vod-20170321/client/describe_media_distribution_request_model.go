// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMediaDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeMediaDistributionRequest
	GetEndTime() *string
	SetInterval(v string) *DescribeMediaDistributionRequest
	GetInterval() *string
	SetStartTime(v string) *DescribeMediaDistributionRequest
	GetStartTime() *string
	SetStorageClass(v string) *DescribeMediaDistributionRequest
	GetStorageClass() *string
}

type DescribeMediaDistributionRequest struct {
	// The end time of CreationTime. The end time must be later than the start time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC). The maximum time span between the start time and end time is six months.
	//
	// example:
	//
	// 2017-01-11T12:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The statistical interval. Default value: day. Valid values:
	//
	// - hour: by hour. Statistics are collected based on the calendar hours within the specified time range.
	//
	// - day: by day. Statistics are collected based on the calendar days within the specified time range.
	//
	// - week: by week. Statistics are collected based on the calendar weeks within the specified time range.
	//
	// - month: by month. Statistics are collected based on the calendar months within the specified time range.
	//
	// example:
	//
	// day
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The start time of CreationTime. Format: yyyy-MM-ddTHH:mm:ssZ (UTC). The maximum time span between the start time and end time is six months.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The storage class. Valid values:
	//
	// - Standard: standard storage.
	//
	// - IA: Infrequent Access.
	//
	// - Archive: Archive storage.
	//
	// - ColdArchive: Cold Archive storage.
	//
	// - SourceIA: Infrequent Access for source files.
	//
	// - SourceArchive: Archive storage for source files.
	//
	// - SourceColdArchive: Cold Archive storage for source files.
	//
	// - Changing: the media asset storage class is being changed.
	//
	// - SourceChanging: the source file storage class is being changed.
	//
	// example:
	//
	// Standard
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s DescribeMediaDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMediaDistributionRequest) GoString() string {
	return s.String()
}

func (s *DescribeMediaDistributionRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMediaDistributionRequest) GetInterval() *string {
	return s.Interval
}

func (s *DescribeMediaDistributionRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMediaDistributionRequest) GetStorageClass() *string {
	return s.StorageClass
}

func (s *DescribeMediaDistributionRequest) SetEndTime(v string) *DescribeMediaDistributionRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMediaDistributionRequest) SetInterval(v string) *DescribeMediaDistributionRequest {
	s.Interval = &v
	return s
}

func (s *DescribeMediaDistributionRequest) SetStartTime(v string) *DescribeMediaDistributionRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMediaDistributionRequest) SetStorageClass(v string) *DescribeMediaDistributionRequest {
	s.StorageClass = &v
	return s
}

func (s *DescribeMediaDistributionRequest) Validate() error {
	return dara.Validate(s)
}
