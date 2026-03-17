// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProbeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *CreateProbeTaskRequest
	GetDomain() *string
	SetEnable(v bool) *CreateProbeTaskRequest
	GetEnable() *bool
	SetPacketNumber(v int32) *CreateProbeTaskRequest
	GetPacketNumber() *int32
	SetPort(v int32) *CreateProbeTaskRequest
	GetPort() *int32
	SetProbeTaskSourceAddress(v string) *CreateProbeTaskRequest
	GetProbeTaskSourceAddress() *string
	SetProtocol(v string) *CreateProbeTaskRequest
	GetProtocol() *string
	SetRegionId(v string) *CreateProbeTaskRequest
	GetRegionId() *string
	SetSagId(v string) *CreateProbeTaskRequest
	GetSagId() *string
	SetSn(v string) *CreateProbeTaskRequest
	GetSn() *string
	SetTaskName(v string) *CreateProbeTaskRequest
	GetTaskName() *string
	SetType(v string) *CreateProbeTaskRequest
	GetType() *string
}

type CreateProbeTaskRequest struct {
	// The domain name that is probed by the task. If the protocol of the probe task is ICMP or TCP, set the value to the IP address or domain name of the service that you want to probe. If the protocol of the probe task is HTTP, set the value to the URL of the service that you want to probe.
	//
	// This parameter is required.
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
	// This parameter is required.
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
	// This parameter is required.
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
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-asdfz6ac74oj5v****
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
	// The type of the probe task. Valid values:
	//
	// 	- **Internet**: probes a public network.
	//
	// 	- **Intranet**: probes a private network.
	//
	// This parameter is required.
	//
	// example:
	//
	// Internet
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateProbeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProbeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateProbeTaskRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateProbeTaskRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreateProbeTaskRequest) GetPacketNumber() *int32 {
	return s.PacketNumber
}

func (s *CreateProbeTaskRequest) GetPort() *int32 {
	return s.Port
}

func (s *CreateProbeTaskRequest) GetProbeTaskSourceAddress() *string {
	return s.ProbeTaskSourceAddress
}

func (s *CreateProbeTaskRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateProbeTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateProbeTaskRequest) GetSagId() *string {
	return s.SagId
}

func (s *CreateProbeTaskRequest) GetSn() *string {
	return s.Sn
}

func (s *CreateProbeTaskRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *CreateProbeTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateProbeTaskRequest) SetDomain(v string) *CreateProbeTaskRequest {
	s.Domain = &v
	return s
}

func (s *CreateProbeTaskRequest) SetEnable(v bool) *CreateProbeTaskRequest {
	s.Enable = &v
	return s
}

func (s *CreateProbeTaskRequest) SetPacketNumber(v int32) *CreateProbeTaskRequest {
	s.PacketNumber = &v
	return s
}

func (s *CreateProbeTaskRequest) SetPort(v int32) *CreateProbeTaskRequest {
	s.Port = &v
	return s
}

func (s *CreateProbeTaskRequest) SetProbeTaskSourceAddress(v string) *CreateProbeTaskRequest {
	s.ProbeTaskSourceAddress = &v
	return s
}

func (s *CreateProbeTaskRequest) SetProtocol(v string) *CreateProbeTaskRequest {
	s.Protocol = &v
	return s
}

func (s *CreateProbeTaskRequest) SetRegionId(v string) *CreateProbeTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProbeTaskRequest) SetSagId(v string) *CreateProbeTaskRequest {
	s.SagId = &v
	return s
}

func (s *CreateProbeTaskRequest) SetSn(v string) *CreateProbeTaskRequest {
	s.Sn = &v
	return s
}

func (s *CreateProbeTaskRequest) SetTaskName(v string) *CreateProbeTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateProbeTaskRequest) SetType(v string) *CreateProbeTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateProbeTaskRequest) Validate() error {
	return dara.Validate(s)
}
