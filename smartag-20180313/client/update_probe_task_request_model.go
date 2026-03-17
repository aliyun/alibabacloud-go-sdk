// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProbeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *UpdateProbeTaskRequest
	GetDomain() *string
	SetEnable(v bool) *UpdateProbeTaskRequest
	GetEnable() *bool
	SetPacketNumber(v int32) *UpdateProbeTaskRequest
	GetPacketNumber() *int32
	SetPort(v int32) *UpdateProbeTaskRequest
	GetPort() *int32
	SetProbeTaskId(v string) *UpdateProbeTaskRequest
	GetProbeTaskId() *string
	SetProbeTaskSourceAddress(v string) *UpdateProbeTaskRequest
	GetProbeTaskSourceAddress() *string
	SetProtocol(v string) *UpdateProbeTaskRequest
	GetProtocol() *string
	SetRegionId(v string) *UpdateProbeTaskRequest
	GetRegionId() *string
	SetSagId(v string) *UpdateProbeTaskRequest
	GetSagId() *string
	SetSn(v string) *UpdateProbeTaskRequest
	GetSn() *string
	SetTaskName(v string) *UpdateProbeTaskRequest
	GetTaskName() *string
}

type UpdateProbeTaskRequest struct {
	// The domain name that is probed by the task.
	//
	// If the protocol of the probe task is ICMP or TCP, set the value to the IP address or domain name of the service that you want to probe. If the protocol of the probe task is HTTP, set the value to the URL of the service that you want to probe.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Specifies whether to enable the probe task. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The number of probe packets transmitted by the probe task per minute.
	//
	// Valid values: **1*	- to **60**.
	//
	// > This parameter is required if the protocol of the probe task is ICMP. Ignore this parameter if the protocol of the probe task is not ICMP.
	//
	// example:
	//
	// 10
	PacketNumber *int32 `json:"PacketNumber,omitempty" xml:"PacketNumber,omitempty"`
	// The port that is probed by the task.
	//
	// > This parameter is required if the protocol of the probe task is TCP. Ignore this parameter if the protocol of the probe task is not TCP.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the probe task.
	//
	// This parameter is required.
	//
	// example:
	//
	// probe-****
	ProbeTaskId *string `json:"ProbeTaskId,omitempty" xml:"ProbeTaskId,omitempty"`
	// The source address of the probe task.
	//
	// > This parameter is required if the task probes private networks.
	//
	// example:
	//
	// 192.168.1.1
	ProbeTaskSourceAddress *string `json:"ProbeTaskSourceAddress,omitempty" xml:"ProbeTaskSourceAddress,omitempty"`
	// The protocol of the probe task. Valid values:
	//
	// 	- **ICMP**
	//
	// 	- **TCP**
	//
	// 	- **HTTP**
	//
	// > Tasks that probe private networks support only ICMP and TCP.
	//
	// example:
	//
	// ICMP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID of the SAG instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Smart Access Gateway (SAG) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-****
	SagId *string `json:"SagId,omitempty" xml:"SagId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag****
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The name of the probe task.
	//
	// example:
	//
	// test-ping
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
}

func (s UpdateProbeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateProbeTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateProbeTaskRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateProbeTaskRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdateProbeTaskRequest) GetPacketNumber() *int32 {
	return s.PacketNumber
}

func (s *UpdateProbeTaskRequest) GetPort() *int32 {
	return s.Port
}

func (s *UpdateProbeTaskRequest) GetProbeTaskId() *string {
	return s.ProbeTaskId
}

func (s *UpdateProbeTaskRequest) GetProbeTaskSourceAddress() *string {
	return s.ProbeTaskSourceAddress
}

func (s *UpdateProbeTaskRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateProbeTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateProbeTaskRequest) GetSagId() *string {
	return s.SagId
}

func (s *UpdateProbeTaskRequest) GetSn() *string {
	return s.Sn
}

func (s *UpdateProbeTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *UpdateProbeTaskRequest) SetDomain(v string) *UpdateProbeTaskRequest {
	s.Domain = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetEnable(v bool) *UpdateProbeTaskRequest {
	s.Enable = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetPacketNumber(v int32) *UpdateProbeTaskRequest {
	s.PacketNumber = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetPort(v int32) *UpdateProbeTaskRequest {
	s.Port = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetProbeTaskId(v string) *UpdateProbeTaskRequest {
	s.ProbeTaskId = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetProbeTaskSourceAddress(v string) *UpdateProbeTaskRequest {
	s.ProbeTaskSourceAddress = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetProtocol(v string) *UpdateProbeTaskRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetRegionId(v string) *UpdateProbeTaskRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetSagId(v string) *UpdateProbeTaskRequest {
	s.SagId = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetSn(v string) *UpdateProbeTaskRequest {
	s.Sn = &v
	return s
}

func (s *UpdateProbeTaskRequest) SetTaskName(v string) *UpdateProbeTaskRequest {
	s.TaskName = &v
	return s
}

func (s *UpdateProbeTaskRequest) Validate() error {
	return dara.Validate(s)
}
