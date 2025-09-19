// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIdcProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdcName(v string) *AddIdcProbeRequest
	GetIdcName() *string
	SetIdcRegion(v string) *AddIdcProbeRequest
	GetIdcRegion() *string
	SetIntervalPeriod(v int32) *AddIdcProbeRequest
	GetIntervalPeriod() *int32
	SetIpSegments(v string) *AddIdcProbeRequest
	GetIpSegments() *string
	SetLinuxPort(v string) *AddIdcProbeRequest
	GetLinuxPort() *string
	SetPeriodUnit(v string) *AddIdcProbeRequest
	GetPeriodUnit() *string
	SetUuids(v string) *AddIdcProbeRequest
	GetUuids() *string
	SetWinPort(v string) *AddIdcProbeRequest
	GetWinPort() *string
}

type AddIdcProbeRequest struct {
	// The name of the data center.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	IdcName *string `json:"IdcName,omitempty" xml:"IdcName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hangzhou
	IdcRegion *string `json:"IdcRegion,omitempty" xml:"IdcRegion,omitempty"`
	// The scan interval.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	IntervalPeriod *int32 `json:"IntervalPeriod,omitempty" xml:"IntervalPeriod,omitempty"`
	// The settings of the CIDR block.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.XX.XX/24
	IpSegments *string `json:"IpSegments,omitempty" xml:"IpSegments,omitempty"`
	// The Linux port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40
	LinuxPort *string `json:"LinuxPort,omitempty" xml:"LinuxPort,omitempty"`
	// The unit of the scan interval. Valid values:
	//
	// 	- **day**
	//
	// 	- **hour**
	//
	// This parameter is required.
	//
	// example:
	//
	// day
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The UUID of the server. Separate multiple UUIDs with commas (,).
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 076a446d-df7d-424c-bdc5-bb5dc7f1****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
	// The Windows port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40
	WinPort *string `json:"WinPort,omitempty" xml:"WinPort,omitempty"`
}

func (s AddIdcProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s AddIdcProbeRequest) GoString() string {
	return s.String()
}

func (s *AddIdcProbeRequest) GetIdcName() *string {
	return s.IdcName
}

func (s *AddIdcProbeRequest) GetIdcRegion() *string {
	return s.IdcRegion
}

func (s *AddIdcProbeRequest) GetIntervalPeriod() *int32 {
	return s.IntervalPeriod
}

func (s *AddIdcProbeRequest) GetIpSegments() *string {
	return s.IpSegments
}

func (s *AddIdcProbeRequest) GetLinuxPort() *string {
	return s.LinuxPort
}

func (s *AddIdcProbeRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *AddIdcProbeRequest) GetUuids() *string {
	return s.Uuids
}

func (s *AddIdcProbeRequest) GetWinPort() *string {
	return s.WinPort
}

func (s *AddIdcProbeRequest) SetIdcName(v string) *AddIdcProbeRequest {
	s.IdcName = &v
	return s
}

func (s *AddIdcProbeRequest) SetIdcRegion(v string) *AddIdcProbeRequest {
	s.IdcRegion = &v
	return s
}

func (s *AddIdcProbeRequest) SetIntervalPeriod(v int32) *AddIdcProbeRequest {
	s.IntervalPeriod = &v
	return s
}

func (s *AddIdcProbeRequest) SetIpSegments(v string) *AddIdcProbeRequest {
	s.IpSegments = &v
	return s
}

func (s *AddIdcProbeRequest) SetLinuxPort(v string) *AddIdcProbeRequest {
	s.LinuxPort = &v
	return s
}

func (s *AddIdcProbeRequest) SetPeriodUnit(v string) *AddIdcProbeRequest {
	s.PeriodUnit = &v
	return s
}

func (s *AddIdcProbeRequest) SetUuids(v string) *AddIdcProbeRequest {
	s.Uuids = &v
	return s
}

func (s *AddIdcProbeRequest) SetWinPort(v string) *AddIdcProbeRequest {
	s.WinPort = &v
	return s
}

func (s *AddIdcProbeRequest) Validate() error {
	return dara.Validate(s)
}
