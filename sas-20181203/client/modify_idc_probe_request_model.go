// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIdcProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdcName(v string) *ModifyIdcProbeRequest
	GetIdcName() *string
	SetIdcRegion(v string) *ModifyIdcProbeRequest
	GetIdcRegion() *string
	SetIntervalPeriod(v int32) *ModifyIdcProbeRequest
	GetIntervalPeriod() *int32
	SetIpSegments(v string) *ModifyIdcProbeRequest
	GetIpSegments() *string
	SetLinuxPort(v string) *ModifyIdcProbeRequest
	GetLinuxPort() *string
	SetPeriodUnit(v string) *ModifyIdcProbeRequest
	GetPeriodUnit() *string
	SetStatus(v int32) *ModifyIdcProbeRequest
	GetStatus() *int32
	SetUuids(v string) *ModifyIdcProbeRequest
	GetUuids() *string
	SetWinPort(v string) *ModifyIdcProbeRequest
	GetWinPort() *string
}

type ModifyIdcProbeRequest struct {
	// The name of the data center.
	//
	// example:
	//
	// test
	IdcName *string `json:"IdcName,omitempty" xml:"IdcName,omitempty"`
	// The region ID of the data center.
	//
	// example:
	//
	// Hangzhou
	IdcRegion *string `json:"IdcRegion,omitempty" xml:"IdcRegion,omitempty"`
	// The scan interval.
	//
	// example:
	//
	// 1
	IntervalPeriod *int32 `json:"IntervalPeriod,omitempty" xml:"IntervalPeriod,omitempty"`
	// The settings of the CIDR block.
	//
	// example:
	//
	// 192.168.XX.XX/24
	IpSegments *string `json:"IpSegments,omitempty" xml:"IpSegments,omitempty"`
	// The Linux port.
	//
	// example:
	//
	// 80
	LinuxPort *string `json:"LinuxPort,omitempty" xml:"LinuxPort,omitempty"`
	// The unit of the scan interval. Valid values:
	//
	// 	- **day**
	//
	// 	- **hour**
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The usage status. Valid values:
	//
	// 	- **0**: enabled.
	//
	// 	- **1**: disabled.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server. Separate multiple UUIDs with commas (,).
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5b6d4072118f487094199cedf90c****,f6310b7976144639867beea2f346****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	// The Windows port.
	//
	// example:
	//
	// 40
	WinPort *string `json:"WinPort,omitempty" xml:"WinPort,omitempty"`
}

func (s ModifyIdcProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyIdcProbeRequest) GoString() string {
	return s.String()
}

func (s *ModifyIdcProbeRequest) GetIdcName() *string {
	return s.IdcName
}

func (s *ModifyIdcProbeRequest) GetIdcRegion() *string {
	return s.IdcRegion
}

func (s *ModifyIdcProbeRequest) GetIntervalPeriod() *int32 {
	return s.IntervalPeriod
}

func (s *ModifyIdcProbeRequest) GetIpSegments() *string {
	return s.IpSegments
}

func (s *ModifyIdcProbeRequest) GetLinuxPort() *string {
	return s.LinuxPort
}

func (s *ModifyIdcProbeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *ModifyIdcProbeRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyIdcProbeRequest) GetUuids() *string {
	return s.Uuids
}

func (s *ModifyIdcProbeRequest) GetWinPort() *string {
	return s.WinPort
}

func (s *ModifyIdcProbeRequest) SetIdcName(v string) *ModifyIdcProbeRequest {
	s.IdcName = &v
	return s
}

func (s *ModifyIdcProbeRequest) SetIdcRegion(v string) *ModifyIdcProbeRequest {
	s.IdcRegion = &v
	return s
}

func (s *ModifyIdcProbeRequest) SetIntervalPeriod(v int32) *ModifyIdcProbeRequest {
	s.IntervalPeriod = &v
	return s
}

func (s *ModifyIdcProbeRequest) SetIpSegments(v string) *ModifyIdcProbeRequest {
	s.IpSegments = &v
	return s
}

func (s *ModifyIdcProbeRequest) SetLinuxPort(v string) *ModifyIdcProbeRequest {
	s.LinuxPort = &v
	return s
}

func (s *ModifyIdcProbeRequest) SetPeriodUnit(v string) *ModifyIdcProbeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *ModifyIdcProbeRequest) SetStatus(v int32) *ModifyIdcProbeRequest {
	s.Status = &v
	return s
}

func (s *ModifyIdcProbeRequest) SetUuids(v string) *ModifyIdcProbeRequest {
	s.Uuids = &v
	return s
}

func (s *ModifyIdcProbeRequest) SetWinPort(v string) *ModifyIdcProbeRequest {
	s.WinPort = &v
	return s
}

func (s *ModifyIdcProbeRequest) Validate() error {
	return dara.Validate(s)
}
