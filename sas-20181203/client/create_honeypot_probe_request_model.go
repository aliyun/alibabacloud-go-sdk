// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHoneypotProbeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArp(v bool) *CreateHoneypotProbeRequest
	GetArp() *bool
	SetBusinessGroupId(v string) *CreateHoneypotProbeRequest
	GetBusinessGroupId() *string
	SetControlNodeId(v string) *CreateHoneypotProbeRequest
	GetControlNodeId() *string
	SetDisplayName(v string) *CreateHoneypotProbeRequest
	GetDisplayName() *string
	SetHoneypotBindList(v []*CreateHoneypotProbeRequestHoneypotBindList) *CreateHoneypotProbeRequest
	GetHoneypotBindList() []*CreateHoneypotProbeRequestHoneypotBindList
	SetPing(v bool) *CreateHoneypotProbeRequest
	GetPing() *bool
	SetProbeType(v string) *CreateHoneypotProbeRequest
	GetProbeType() *string
	SetProbeVersion(v string) *CreateHoneypotProbeRequest
	GetProbeVersion() *string
	SetProxyIp(v string) *CreateHoneypotProbeRequest
	GetProxyIp() *string
	SetUuid(v string) *CreateHoneypotProbeRequest
	GetUuid() *string
	SetVpcId(v string) *CreateHoneypotProbeRequest
	GetVpcId() *string
}

type CreateHoneypotProbeRequest struct {
	// Specifies whether to enable Address Resolution Protocol (ARP) spoofing. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Arp *bool `json:"Arp,omitempty" xml:"Arp,omitempty"`
	// The ID of the business group.
	//
	// example:
	//
	// 2022011817324588686
	BusinessGroupId *string `json:"BusinessGroupId,omitempty" xml:"BusinessGroupId,omitempty"`
	// The ID of the management node.
	//
	// > You can call the [ListHoneypotNode](~~ListHoneypotNode~~) operation to query the IDs of management nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// a882e590-b87b-45a6-87b9-d0a3e5a0****
	ControlNodeId *string `json:"ControlNodeId,omitempty" xml:"ControlNodeId,omitempty"`
	// The name of the probe.
	//
	// This parameter is required.
	//
	// example:
	//
	// testHoneyPotProbe
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The configuration of the probe.
	HoneypotBindList []*CreateHoneypotProbeRequestHoneypotBindList `json:"HoneypotBindList,omitempty" xml:"HoneypotBindList,omitempty" type:"Repeated"`
	// Specifies whether to enable ping scan. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Ping *bool `json:"Ping,omitempty" xml:"Ping,omitempty"`
	// The type of the probe. Valid values:
	//
	// 	- **host_probe**: host probe
	//
	// 	- **vpc_black_hole_probe**: virtual private cloud (VPC) probe
	//
	// This parameter is required.
	//
	// example:
	//
	// host_probe
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	// The version of the probe.
	//
	// example:
	//
	// 0.0.0
	ProbeVersion *string `json:"ProbeVersion,omitempty" xml:"ProbeVersion,omitempty"`
	// The IP address of the proxy.
	//
	// example:
	//
	// 192.168.XX.XX
	ProxyIp *string `json:"ProxyIp,omitempty" xml:"ProxyIp,omitempty"`
	// The UUID of the instance.
	//
	// > If **ProbeType*	- is set to **host_probe**, this parameter is required.
	//
	// example:
	//
	// e4af3620-6895-4e2f-a641-a9d8fb53****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the VPC.
	//
	// > If **ProbeType*	- is set to **vpc_black_hole_probe**, this parameter is required. You can call the [DescribeVpcHoneyPotList](~~DescribeVpcHoneyPotList~~) operation to query the IDs of VPCs.
	//
	// example:
	//
	// vpc-zm0asrkpv1q8gnk7mn4dn
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateHoneypotProbeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeRequest) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeRequest) GetArp() *bool {
	return s.Arp
}

func (s *CreateHoneypotProbeRequest) GetBusinessGroupId() *string {
	return s.BusinessGroupId
}

func (s *CreateHoneypotProbeRequest) GetControlNodeId() *string {
	return s.ControlNodeId
}

func (s *CreateHoneypotProbeRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateHoneypotProbeRequest) GetHoneypotBindList() []*CreateHoneypotProbeRequestHoneypotBindList {
	return s.HoneypotBindList
}

func (s *CreateHoneypotProbeRequest) GetPing() *bool {
	return s.Ping
}

func (s *CreateHoneypotProbeRequest) GetProbeType() *string {
	return s.ProbeType
}

func (s *CreateHoneypotProbeRequest) GetProbeVersion() *string {
	return s.ProbeVersion
}

func (s *CreateHoneypotProbeRequest) GetProxyIp() *string {
	return s.ProxyIp
}

func (s *CreateHoneypotProbeRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CreateHoneypotProbeRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateHoneypotProbeRequest) SetArp(v bool) *CreateHoneypotProbeRequest {
	s.Arp = &v
	return s
}

func (s *CreateHoneypotProbeRequest) SetBusinessGroupId(v string) *CreateHoneypotProbeRequest {
	s.BusinessGroupId = &v
	return s
}

func (s *CreateHoneypotProbeRequest) SetControlNodeId(v string) *CreateHoneypotProbeRequest {
	s.ControlNodeId = &v
	return s
}

func (s *CreateHoneypotProbeRequest) SetDisplayName(v string) *CreateHoneypotProbeRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateHoneypotProbeRequest) SetHoneypotBindList(v []*CreateHoneypotProbeRequestHoneypotBindList) *CreateHoneypotProbeRequest {
	s.HoneypotBindList = v
	return s
}

func (s *CreateHoneypotProbeRequest) SetPing(v bool) *CreateHoneypotProbeRequest {
	s.Ping = &v
	return s
}

func (s *CreateHoneypotProbeRequest) SetProbeType(v string) *CreateHoneypotProbeRequest {
	s.ProbeType = &v
	return s
}

func (s *CreateHoneypotProbeRequest) SetProbeVersion(v string) *CreateHoneypotProbeRequest {
	s.ProbeVersion = &v
	return s
}

func (s *CreateHoneypotProbeRequest) SetProxyIp(v string) *CreateHoneypotProbeRequest {
	s.ProxyIp = &v
	return s
}

func (s *CreateHoneypotProbeRequest) SetUuid(v string) *CreateHoneypotProbeRequest {
	s.Uuid = &v
	return s
}

func (s *CreateHoneypotProbeRequest) SetVpcId(v string) *CreateHoneypotProbeRequest {
	s.VpcId = &v
	return s
}

func (s *CreateHoneypotProbeRequest) Validate() error {
	if s.HoneypotBindList != nil {
		for _, item := range s.HoneypotBindList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHoneypotProbeRequestHoneypotBindList struct {
	// The listener ports.
	BindPortList []*CreateHoneypotProbeRequestHoneypotBindListBindPortList `json:"BindPortList,omitempty" xml:"BindPortList,omitempty" type:"Repeated"`
	// The ID of the honeypot.
	//
	// > You can call the [ListHoneypot](~~ListHoneypot~~) operation to query the IDs of honeypots.
	//
	// example:
	//
	// 1a5eda2d40f92ac87d6b63e1a5ad4b76fe0d4110c4a3e2fa85438a29ae55****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
}

func (s CreateHoneypotProbeRequestHoneypotBindList) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeRequestHoneypotBindList) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeRequestHoneypotBindList) GetBindPortList() []*CreateHoneypotProbeRequestHoneypotBindListBindPortList {
	return s.BindPortList
}

func (s *CreateHoneypotProbeRequestHoneypotBindList) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *CreateHoneypotProbeRequestHoneypotBindList) SetBindPortList(v []*CreateHoneypotProbeRequestHoneypotBindListBindPortList) *CreateHoneypotProbeRequestHoneypotBindList {
	s.BindPortList = v
	return s
}

func (s *CreateHoneypotProbeRequestHoneypotBindList) SetHoneypotId(v string) *CreateHoneypotProbeRequestHoneypotBindList {
	s.HoneypotId = &v
	return s
}

func (s *CreateHoneypotProbeRequestHoneypotBindList) Validate() error {
	if s.BindPortList != nil {
		for _, item := range s.BindPortList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHoneypotProbeRequestHoneypotBindListBindPortList struct {
	// Specifies whether to bind a port. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	BindPort *bool `json:"BindPort,omitempty" xml:"BindPort,omitempty"`
	// The end of the port range.
	//
	// example:
	//
	// 90
	EndPort *int32 `json:"EndPort,omitempty" xml:"EndPort,omitempty"`
	// Specifies whether the port is a fixed port. Valid values:
	//
	// 	- **0**: no
	//
	// 	- **1**: yes
	//
	// example:
	//
	// 0
	Fixed *bool `json:"Fixed,omitempty" xml:"Fixed,omitempty"`
	// The start of the port range.
	//
	// example:
	//
	// 80
	StartPort *int32 `json:"StartPort,omitempty" xml:"StartPort,omitempty"`
	// The destination port.
	//
	// > If **HoneypotId*	- is specified, this parameter is required.
	//
	// example:
	//
	// 80
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s CreateHoneypotProbeRequestHoneypotBindListBindPortList) String() string {
	return dara.Prettify(s)
}

func (s CreateHoneypotProbeRequestHoneypotBindListBindPortList) GoString() string {
	return s.String()
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) GetBindPort() *bool {
	return s.BindPort
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) GetEndPort() *int32 {
	return s.EndPort
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) GetFixed() *bool {
	return s.Fixed
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) GetStartPort() *int32 {
	return s.StartPort
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) SetBindPort(v bool) *CreateHoneypotProbeRequestHoneypotBindListBindPortList {
	s.BindPort = &v
	return s
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) SetEndPort(v int32) *CreateHoneypotProbeRequestHoneypotBindListBindPortList {
	s.EndPort = &v
	return s
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) SetFixed(v bool) *CreateHoneypotProbeRequestHoneypotBindListBindPortList {
	s.Fixed = &v
	return s
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) SetStartPort(v int32) *CreateHoneypotProbeRequestHoneypotBindListBindPortList {
	s.StartPort = &v
	return s
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) SetTargetPort(v int32) *CreateHoneypotProbeRequestHoneypotBindListBindPortList {
	s.TargetPort = &v
	return s
}

func (s *CreateHoneypotProbeRequestHoneypotBindListBindPortList) Validate() error {
	return dara.Validate(s)
}
