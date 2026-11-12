// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeMonitorDataRequest
	GetClientToken() *string
	SetEndTime(v string) *DescribeMonitorDataRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeMonitorDataRequest
	GetInstanceId() *string
	SetLength(v string) *DescribeMonitorDataRequest
	GetLength() *string
	SetMetricName(v string) *DescribeMonitorDataRequest
	GetMetricName() *string
	SetNextToken(v string) *DescribeMonitorDataRequest
	GetNextToken() *string
	SetPeriod(v string) *DescribeMonitorDataRequest
	GetPeriod() *string
	SetRegionId(v string) *DescribeMonitorDataRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribeMonitorDataRequest
	GetStartTime() *string
}

type DescribeMonitorDataRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end of the time range to query. The following formats are supported:
	//
	// - UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 January 1, 1970.
	//
	// - Time format: YYYY-MM-DDThh:mm:ssZ.
	//
	// > The interval between the start time and the end time is less than or equal to 31 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-08T08:04:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ad1ae67295445f598017499dc****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. Valid values: 1 to 1440.
	//
	// example:
	//
	// 100
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The name of the metric. Valid values:
	//
	// - MEMORY_ACTUALUSEDSPACE: the memory usage. Unit: bytes.
	//
	// - DISKUSAGE_USED: the disk usage. Unit: bytes.
	//
	// - CPU_UTILIZATION: the CPU usage, in percentage.
	//
	// - VPC_PUBLICIP_INTERNETOUT_RATE: the outbound bandwidth. Unit: bits/s.
	//
	// - VPC_PUBLICIP_INTERNETIN_RATE: the inbound bandwidth. Unit: bits/s.
	//
	// - DISK_READ_IOPS: the read IOPS of the disk. Unit: count/s.
	//
	// - DISK_WRITE_IOPS: the write IOPS of the disk. Unit: count/s.
	//
	// - FLOW_USED: the traffic usage. Unit: bytes.
	//
	// This parameter is required.
	//
	// example:
	//
	// DISKUSAGE_USED
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nOc1nj4M9UaAZ/I8db***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The interval at which the monitoring data is queried. Valid values: 60, 300, and 900. Unit: seconds.
	//
	// > If MetricName is set to FLOW_USED, set Period to 3600 (one hour). In other cases, set Period based on your business requirements.
	//
	// This parameter is required.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The beginning of the time range to query. The following formats are supported:
	//
	// - UNIX timestamp: the number of milliseconds that have elapsed since 00:00:00 January 1, 1970.
	//
	// - Time format: YYYY-MM-DDThh:mm:ssZ.
	//
	// > 	- The specified time range includes the end time and excludes the start time. The start time must be earlier than the end time.
	//
	// >
	//
	// > 	- The interval between the start time and the end time is less than or equal to 31 days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2022-09-07T04:04:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorDataRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeMonitorDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitorDataRequest) GetLength() *string {
	return s.Length
}

func (s *DescribeMonitorDataRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *DescribeMonitorDataRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeMonitorDataRequest) GetPeriod() *string {
	return s.Period
}

func (s *DescribeMonitorDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeMonitorDataRequest) SetClientToken(v string) *DescribeMonitorDataRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetEndTime(v string) *DescribeMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetInstanceId(v string) *DescribeMonitorDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetLength(v string) *DescribeMonitorDataRequest {
	s.Length = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetMetricName(v string) *DescribeMonitorDataRequest {
	s.MetricName = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetNextToken(v string) *DescribeMonitorDataRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetPeriod(v string) *DescribeMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetRegionId(v string) *DescribeMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorDataRequest) SetStartTime(v string) *DescribeMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
