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
	// The end time of the creation time. The end time must be later than the start time. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The maximum time span between the start time and end time is six months.
	//
	// example:
	//
	// 2017-01-11T12:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The statistical interval. Default value: day. Valid values:
	//
	// - hour: Hour. Statistics are collected based on the natural hours of the start time and end time.
	//
	// - day: Day. Statistics are collected based on the natural days of the start time and end time.
	//
	// - week: Week. Statistics are collected based on the natural weeks of the start time and end time.
	//
	// - month: Month. Statistics are collected based on the natural months of the start time and end time.
	//
	// example:
	//
	// day
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The start time of the creation time. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The maximum time span between the start time and end time is six months.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The storage class. Valid values:
	//
	// - Standard: Standard storage.
	//
	// - IA: Infrequent access.
	//
	// - Archive: Archive storage.
	//
	// - ColdArchive: Cold archive storage.
	//
	// - SourceIA: Infrequent access for source files.
	//
	// - SourceArchive: Archive storage for source files.
	//
	// - SourceColdArchive: Cold archive storage for source files.
	//
	// - Changing: The storage type of the media asset is being changed.
	//
	// - SourceChanging: The storage type of the source file is being changed.
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
