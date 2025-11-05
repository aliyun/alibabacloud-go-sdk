// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DescribeDiskMonitorDataRequest
	GetDiskId() *string
	SetEndTime(v string) *DescribeDiskMonitorDataRequest
	GetEndTime() *string
	SetPeriod(v int64) *DescribeDiskMonitorDataRequest
	GetPeriod() *int64
	SetRegionId(v string) *DescribeDiskMonitorDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeDiskMonitorDataRequest
	GetStartTime() *string
	SetType(v string) *DescribeDiskMonitorDataRequest
	GetType() *string
}

type DescribeDiskMonitorDataRequest struct {
	// The ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp67acfmxazb4p****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The end of the time range during which you want to query the near real-time monitoring data of the disk. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-06-01T05:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The interval at which the near real-time monitoring data is collected. Unit: seconds. Valid values:
	//
	// 	- 5
	//
	// 	- 60
	//
	// Default value: 5.
	//
	// example:
	//
	// 5
	Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range during which you want to query the near real-time monitoring data of the disk. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-06-01T03:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The type of the monitoring data. Valid values:
	//
	// 	- basic: baseline performance data.
	//
	// 	- pro: burst performance data, such as burst I/O operations.
	//
	// example:
	//
	// basic
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiskMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiskMonitorDataRequest) GetPeriod() *int64 {
	return s.Period
}

func (s *DescribeDiskMonitorDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiskMonitorDataRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDiskMonitorDataRequest) SetDiskId(v string) *DescribeDiskMonitorDataRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetEndTime(v string) *DescribeDiskMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetPeriod(v int64) *DescribeDiskMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetRegionId(v string) *DescribeDiskMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetStartTime(v string) *DescribeDiskMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetType(v string) *DescribeDiskMonitorDataRequest {
	s.Type = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
