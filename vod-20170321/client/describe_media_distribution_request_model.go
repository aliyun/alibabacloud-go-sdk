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
	// The end of the time range to query. The end time must be later than the start time. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The maximum time range to query is 6 months.
	//
	// example:
	//
	// 2017-01-11T12:59:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The statistical interval. Default value: day. Valid values:
	//
	// 	- hour: natural hour of the start and end time.
	//
	// 	- day: natural day of the start and end time.
	//
	// 	- week: natural week of the start and end time.
	//
	// 	- month: natural month of the start and end time.
	//
	// example:
	//
	// day
	Interval *string `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The beginning of the time range to query. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. The maximum time range to query is 6 months.
	//
	// example:
	//
	// 2017-01-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The hierarchical storage type. Valid values:
	//
	// 	- Standard
	//
	// 	- IA
	//
	// 	- Archive
	//
	// 	- ColdArchive
	//
	// 	- SourceIA
	//
	// 	- SourceArchive
	//
	// 	- SourceColdArchive
	//
	// 	- Changing
	//
	// 	- SourceChanging
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
