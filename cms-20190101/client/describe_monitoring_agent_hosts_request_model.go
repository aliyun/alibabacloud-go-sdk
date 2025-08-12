// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunHost(v bool) *DescribeMonitoringAgentHostsRequest
	GetAliyunHost() *bool
	SetHostName(v string) *DescribeMonitoringAgentHostsRequest
	GetHostName() *string
	SetInstanceIds(v string) *DescribeMonitoringAgentHostsRequest
	GetInstanceIds() *string
	SetInstanceRegionId(v string) *DescribeMonitoringAgentHostsRequest
	GetInstanceRegionId() *string
	SetKeyWord(v string) *DescribeMonitoringAgentHostsRequest
	GetKeyWord() *string
	SetPageNumber(v int32) *DescribeMonitoringAgentHostsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMonitoringAgentHostsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeMonitoringAgentHostsRequest
	GetRegionId() *string
	SetSerialNumbers(v string) *DescribeMonitoringAgentHostsRequest
	GetSerialNumbers() *string
	SetStatus(v string) *DescribeMonitoringAgentHostsRequest
	GetStatus() *string
	SetSysomStatus(v string) *DescribeMonitoringAgentHostsRequest
	GetSysomStatus() *string
}

type DescribeMonitoringAgentHostsRequest struct {
	// Specifies whether to query Elastic Compute Service (ECS) instances that are provided by Alibaba Cloud or to query hosts that are not provided by Alibaba Cloud. Valid values:
	//
	// 	- true (default value): queries all the ECS instances that are provided by Alibaba Cloud.
	//
	// 	- false: queries all the hosts that are not provided by Alibaba Cloud.
	//
	// example:
	//
	// true
	AliyunHost *bool `json:"AliyunHost,omitempty" xml:"AliyunHost,omitempty"`
	// The name of the host.
	//
	// example:
	//
	// hostNam1
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-a3d1q1pm2f9yr29e****
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	InstanceRegionId *string `json:"InstanceRegionId,omitempty" xml:"InstanceRegionId,omitempty"`
	// The keyword that is used in fuzzy match.
	//
	// example:
	//
	// host1
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values:
	//
	// 	- 10
	//
	// 	- 20
	//
	// 	- 50
	//
	// 	- 100
	//
	// > Although Alibaba Cloud does not limit the maximum value of this parameter, we recommend that you do not set it to an excessively large value. If you set it to an excessively large value, a timeout error may occur.
	//
	// example:
	//
	// 10
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The serial number of the host.
	//
	// After the CloudMonitor agent is installed on a host, a globally unique serial number is generated. A host that is not provided by Alibaba Cloud has a serial number instead of an instance ID.
	//
	// > This parameter can be used to accurately search for a monitored host.
	//
	// example:
	//
	// a1ab31a3-1234-40f2-9e95-c8caa8f0****
	SerialNumbers *string `json:"SerialNumbers,omitempty" xml:"SerialNumbers,omitempty"`
	// The status of the hosts that you want to query. Valid values:
	//
	// 	- Running: queries the hosts that are running.
	//
	// 	- Stopped: queries the hosts that are stopped, are not installed, or fail to be installed.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status of SysOM. Valid values:
	//
	// 	- installing: SysOM is being installed.
	//
	// 	- running: SysOM is running.
	//
	// 	- stopped: SysOM is stopped.
	//
	// 	- uninstalling: SysOM is being uninstalled.
	//
	// example:
	//
	// running
	SysomStatus *string `json:"SysomStatus,omitempty" xml:"SysomStatus,omitempty"`
}

func (s DescribeMonitoringAgentHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentHostsRequest) GetAliyunHost() *bool {
	return s.AliyunHost
}

func (s *DescribeMonitoringAgentHostsRequest) GetHostName() *string {
	return s.HostName
}

func (s *DescribeMonitoringAgentHostsRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *DescribeMonitoringAgentHostsRequest) GetInstanceRegionId() *string {
	return s.InstanceRegionId
}

func (s *DescribeMonitoringAgentHostsRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *DescribeMonitoringAgentHostsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMonitoringAgentHostsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMonitoringAgentHostsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitoringAgentHostsRequest) GetSerialNumbers() *string {
	return s.SerialNumbers
}

func (s *DescribeMonitoringAgentHostsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeMonitoringAgentHostsRequest) GetSysomStatus() *string {
	return s.SysomStatus
}

func (s *DescribeMonitoringAgentHostsRequest) SetAliyunHost(v bool) *DescribeMonitoringAgentHostsRequest {
	s.AliyunHost = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetHostName(v string) *DescribeMonitoringAgentHostsRequest {
	s.HostName = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetInstanceIds(v string) *DescribeMonitoringAgentHostsRequest {
	s.InstanceIds = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetInstanceRegionId(v string) *DescribeMonitoringAgentHostsRequest {
	s.InstanceRegionId = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetKeyWord(v string) *DescribeMonitoringAgentHostsRequest {
	s.KeyWord = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetPageNumber(v int32) *DescribeMonitoringAgentHostsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetPageSize(v int32) *DescribeMonitoringAgentHostsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetRegionId(v string) *DescribeMonitoringAgentHostsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetSerialNumbers(v string) *DescribeMonitoringAgentHostsRequest {
	s.SerialNumbers = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetStatus(v string) *DescribeMonitoringAgentHostsRequest {
	s.Status = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) SetSysomStatus(v string) *DescribeMonitoringAgentHostsRequest {
	s.SysomStatus = &v
	return s
}

func (s *DescribeMonitoringAgentHostsRequest) Validate() error {
	return dara.Validate(s)
}
