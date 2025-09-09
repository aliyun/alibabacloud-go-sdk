// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDrdsInstanceId(v string) *DescribeDrdsInstanceMonitorRequest
	GetDrdsInstanceId() *string
	SetEndTime(v int64) *DescribeDrdsInstanceMonitorRequest
	GetEndTime() *int64
	SetKey(v string) *DescribeDrdsInstanceMonitorRequest
	GetKey() *string
	SetPeriodMultiple(v int32) *DescribeDrdsInstanceMonitorRequest
	GetPeriodMultiple() *int32
	SetRegionId(v string) *DescribeDrdsInstanceMonitorRequest
	GetRegionId() *string
	SetStartTime(v int64) *DescribeDrdsInstanceMonitorRequest
	GetStartTime() *int64
}

type DescribeDrdsInstanceMonitorRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// drds*********
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The end time of the query. Specify the value in the UNIX timestamp format. The timestamp must be in UTC. Unit: ms.
	//
	// >  If the time range that you specify is less than 1 hour, the monitoring data that is collected in a 1-hour period before the end time is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1603209690000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The performance monitoring metrics. You can specify one or more metrics. Separate multiple metric names with commas (,).
	//
	// >  For more information about performance monitoring metrics, see [Monitor instances](https://help.aliyun.com/document_detail/186703.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpu
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The multiple of the default time interval that you want to use to collect monitoring data. By default, the system collects monitoring data of resources at an interval of 1 minute. If you set the value of this parameter to 2, the system collects monitoring data of the instance at an interval of 2 minutes.
	//
	// example:
	//
	// 1
	PeriodMultiple *int32 `json:"PeriodMultiple,omitempty" xml:"PeriodMultiple,omitempty"`
	// The ID of the region where the instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time of the query. Specify the value in the UNIX timestamp format. The timestamp must be in UTC. Unit: ms.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1603123290000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDrdsInstanceMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceMonitorRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *DescribeDrdsInstanceMonitorRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDrdsInstanceMonitorRequest) GetKey() *string {
	return s.Key
}

func (s *DescribeDrdsInstanceMonitorRequest) GetPeriodMultiple() *int32 {
	return s.PeriodMultiple
}

func (s *DescribeDrdsInstanceMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDrdsInstanceMonitorRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDrdsInstanceMonitorRequest) SetDrdsInstanceId(v string) *DescribeDrdsInstanceMonitorRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetEndTime(v int64) *DescribeDrdsInstanceMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetKey(v string) *DescribeDrdsInstanceMonitorRequest {
	s.Key = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetPeriodMultiple(v int32) *DescribeDrdsInstanceMonitorRequest {
	s.PeriodMultiple = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetRegionId(v string) *DescribeDrdsInstanceMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) SetStartTime(v int64) *DescribeDrdsInstanceMonitorRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDrdsInstanceMonitorRequest) Validate() error {
	return dara.Validate(s)
}
