// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotProbeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHoneypotProbeResponseBody
	GetCode() *string
	SetData(v *GetHoneypotProbeResponseBodyData) *GetHoneypotProbeResponseBody
	GetData() *GetHoneypotProbeResponseBodyData
	SetHttpStatusCode(v int32) *GetHoneypotProbeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHoneypotProbeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHoneypotProbeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHoneypotProbeResponseBody
	GetSuccess() *bool
}

type GetHoneypotProbeResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the probe.
	Data *GetHoneypotProbeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6550E0E6-FD6C-5F39-AB5E-35B30DCA97B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHoneypotProbeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotProbeResponseBody) GoString() string {
	return s.String()
}

func (s *GetHoneypotProbeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHoneypotProbeResponseBody) GetData() *GetHoneypotProbeResponseBodyData {
	return s.Data
}

func (s *GetHoneypotProbeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHoneypotProbeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHoneypotProbeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHoneypotProbeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHoneypotProbeResponseBody) SetCode(v string) *GetHoneypotProbeResponseBody {
	s.Code = &v
	return s
}

func (s *GetHoneypotProbeResponseBody) SetData(v *GetHoneypotProbeResponseBodyData) *GetHoneypotProbeResponseBody {
	s.Data = v
	return s
}

func (s *GetHoneypotProbeResponseBody) SetHttpStatusCode(v int32) *GetHoneypotProbeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHoneypotProbeResponseBody) SetMessage(v string) *GetHoneypotProbeResponseBody {
	s.Message = &v
	return s
}

func (s *GetHoneypotProbeResponseBody) SetRequestId(v string) *GetHoneypotProbeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHoneypotProbeResponseBody) SetSuccess(v bool) *GetHoneypotProbeResponseBody {
	s.Success = &v
	return s
}

func (s *GetHoneypotProbeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHoneypotProbeResponseBodyData struct {
	// Indicates whether address resolution protocol (ARP) is enabled for the check type.
	//
	// example:
	//
	// true
	Arp *bool `json:"Arp,omitempty" xml:"Arp,omitempty"`
	// An array consisting of the IP addresses that can be monitored.
	CanListenIpList []*string `json:"CanListenIpList,omitempty" xml:"CanListenIpList,omitempty" type:"Repeated"`
	// The CIDR blocks of the probe deployed in a virtual private cloud (VPC).
	CidrList []*string `json:"CidrList,omitempty" xml:"CidrList,omitempty" type:"Repeated"`
	// The information about the management node.
	ControlNode *GetHoneypotProbeResponseBodyDataControlNode `json:"ControlNode,omitempty" xml:"ControlNode,omitempty" type:"Struct"`
	// The CPU utilization.
	//
	// example:
	//
	// 0.51
	CpuLoad *float64 `json:"CpuLoad,omitempty" xml:"CpuLoad,omitempty"`
	// The time when the probe was deployed.
	//
	// example:
	//
	// 1669363825000
	DeployTime *int64 `json:"DeployTime,omitempty" xml:"DeployTime,omitempty"`
	// The name of the probe.
	//
	// example:
	//
	// test-probe
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The ports that the honeypot monitors.
	HoneyPotProbeScanPort *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort `json:"HoneyPotProbeScanPort,omitempty" xml:"HoneyPotProbeScanPort,omitempty" type:"Struct"`
	// The honeypots that are bound to the probe.
	HoneypotProbeBindList []*GetHoneypotProbeResponseBodyDataHoneypotProbeBindList `json:"HoneypotProbeBindList,omitempty" xml:"HoneypotProbeBindList,omitempty" type:"Repeated"`
	// The IP address of the server on which the probe is deployed.
	//
	// example:
	//
	// 33.53.XX.XX
	HostIp *string `json:"HostIp,omitempty" xml:"HostIp,omitempty"`
	// An array consisting of the IP addresses that can be monitored.
	ListenIpList []*string `json:"ListenIpList,omitempty" xml:"ListenIpList,omitempty" type:"Repeated"`
	// The memory usage.
	//
	// example:
	//
	// 1.94
	MemoryLoad *float64 `json:"MemoryLoad,omitempty" xml:"MemoryLoad,omitempty"`
	// The operating system of the server on which the probe is deployed. Valid values:
	//
	// 	- windows
	//
	// 	- linux
	//
	// example:
	//
	// windows
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// Indicates whether ping is enabled for the check type.
	//
	// example:
	//
	// false
	Ping *bool `json:"Ping,omitempty" xml:"Ping,omitempty"`
	// The ID of the probe.
	//
	// example:
	//
	// 40f6501d-45ec-4bf0-b813-0072ceb4****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
	// The type of the probe. Valid values:
	//
	// 	- **host_probe**: host probe
	//
	// 	- **vpc_black_hole_probe**: virtual private cloud (VPC) probe
	//
	// example:
	//
	// host_probe
	ProbeType *string `json:"ProbeType,omitempty" xml:"ProbeType,omitempty"`
	// The version of the probe.
	//
	// example:
	//
	// 18060096
	ProbeVersion *string `json:"ProbeVersion,omitempty" xml:"ProbeVersion,omitempty"`
	// The IP address of the proxy server.
	//
	// example:
	//
	// 47.108.XX.XX
	ProxyIp *string `json:"ProxyIp,omitempty" xml:"ProxyIp,omitempty"`
	// The status of the probe. Valid values:
	//
	// 	- **installed**: installed
	//
	// 	- **install_failed**: installation failed
	//
	// 	- **online**: online
	//
	// 	- **offline**: offline
	//
	// 	- **unnormal**: abnormal
	//
	// 	- **unprobe**: unauthorized
	//
	// 	- **uninstalling**: being uninstalled
	//
	// 	- **uninstalled**: uninstalled
	//
	// 	- **uninstall_failed**: uninstallation failed
	//
	// 	- **not_exist**: not installed
	//
	// example:
	//
	// online
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the asset on which the host probe is deployed.
	//
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The ID of the VPC in which the probe is deployed.
	//
	// example:
	//
	// vpc-2vchkxmf2j9yjt3x2****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetHoneypotProbeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotProbeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHoneypotProbeResponseBodyData) GetArp() *bool {
	return s.Arp
}

func (s *GetHoneypotProbeResponseBodyData) GetCanListenIpList() []*string {
	return s.CanListenIpList
}

func (s *GetHoneypotProbeResponseBodyData) GetCidrList() []*string {
	return s.CidrList
}

func (s *GetHoneypotProbeResponseBodyData) GetControlNode() *GetHoneypotProbeResponseBodyDataControlNode {
	return s.ControlNode
}

func (s *GetHoneypotProbeResponseBodyData) GetCpuLoad() *float64 {
	return s.CpuLoad
}

func (s *GetHoneypotProbeResponseBodyData) GetDeployTime() *int64 {
	return s.DeployTime
}

func (s *GetHoneypotProbeResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetHoneypotProbeResponseBodyData) GetHoneyPotProbeScanPort() *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort {
	return s.HoneyPotProbeScanPort
}

func (s *GetHoneypotProbeResponseBodyData) GetHoneypotProbeBindList() []*GetHoneypotProbeResponseBodyDataHoneypotProbeBindList {
	return s.HoneypotProbeBindList
}

func (s *GetHoneypotProbeResponseBodyData) GetHostIp() *string {
	return s.HostIp
}

func (s *GetHoneypotProbeResponseBodyData) GetListenIpList() []*string {
	return s.ListenIpList
}

func (s *GetHoneypotProbeResponseBodyData) GetMemoryLoad() *float64 {
	return s.MemoryLoad
}

func (s *GetHoneypotProbeResponseBodyData) GetOsType() *string {
	return s.OsType
}

func (s *GetHoneypotProbeResponseBodyData) GetPing() *bool {
	return s.Ping
}

func (s *GetHoneypotProbeResponseBodyData) GetProbeId() *string {
	return s.ProbeId
}

func (s *GetHoneypotProbeResponseBodyData) GetProbeType() *string {
	return s.ProbeType
}

func (s *GetHoneypotProbeResponseBodyData) GetProbeVersion() *string {
	return s.ProbeVersion
}

func (s *GetHoneypotProbeResponseBodyData) GetProxyIp() *string {
	return s.ProxyIp
}

func (s *GetHoneypotProbeResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetHoneypotProbeResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetHoneypotProbeResponseBodyData) GetVpcId() *string {
	return s.VpcId
}

func (s *GetHoneypotProbeResponseBodyData) SetArp(v bool) *GetHoneypotProbeResponseBodyData {
	s.Arp = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetCanListenIpList(v []*string) *GetHoneypotProbeResponseBodyData {
	s.CanListenIpList = v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetCidrList(v []*string) *GetHoneypotProbeResponseBodyData {
	s.CidrList = v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetControlNode(v *GetHoneypotProbeResponseBodyDataControlNode) *GetHoneypotProbeResponseBodyData {
	s.ControlNode = v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetCpuLoad(v float64) *GetHoneypotProbeResponseBodyData {
	s.CpuLoad = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetDeployTime(v int64) *GetHoneypotProbeResponseBodyData {
	s.DeployTime = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetDisplayName(v string) *GetHoneypotProbeResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetHoneyPotProbeScanPort(v *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) *GetHoneypotProbeResponseBodyData {
	s.HoneyPotProbeScanPort = v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetHoneypotProbeBindList(v []*GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) *GetHoneypotProbeResponseBodyData {
	s.HoneypotProbeBindList = v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetHostIp(v string) *GetHoneypotProbeResponseBodyData {
	s.HostIp = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetListenIpList(v []*string) *GetHoneypotProbeResponseBodyData {
	s.ListenIpList = v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetMemoryLoad(v float64) *GetHoneypotProbeResponseBodyData {
	s.MemoryLoad = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetOsType(v string) *GetHoneypotProbeResponseBodyData {
	s.OsType = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetPing(v bool) *GetHoneypotProbeResponseBodyData {
	s.Ping = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetProbeId(v string) *GetHoneypotProbeResponseBodyData {
	s.ProbeId = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetProbeType(v string) *GetHoneypotProbeResponseBodyData {
	s.ProbeType = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetProbeVersion(v string) *GetHoneypotProbeResponseBodyData {
	s.ProbeVersion = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetProxyIp(v string) *GetHoneypotProbeResponseBodyData {
	s.ProxyIp = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetStatus(v int32) *GetHoneypotProbeResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetUuid(v string) *GetHoneypotProbeResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) SetVpcId(v string) *GetHoneypotProbeResponseBodyData {
	s.VpcId = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyData) Validate() error {
	if s.ControlNode != nil {
		if err := s.ControlNode.Validate(); err != nil {
			return err
		}
	}
	if s.HoneyPotProbeScanPort != nil {
		if err := s.HoneyPotProbeScanPort.Validate(); err != nil {
			return err
		}
	}
	if s.HoneypotProbeBindList != nil {
		for _, item := range s.HoneypotProbeBindList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHoneypotProbeResponseBodyDataControlNode struct {
	// The instance ID of the management node.
	//
	// example:
	//
	// i-bp19ijepxytwtzrk****
	EcsInstanceId *string `json:"EcsInstanceId,omitempty" xml:"EcsInstanceId,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// 37a15ff1-3475-4897-aa6c-f7fd9122****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the management node.
	//
	// example:
	//
	// online-honeypot
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s GetHoneypotProbeResponseBodyDataControlNode) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotProbeResponseBodyDataControlNode) GoString() string {
	return s.String()
}

func (s *GetHoneypotProbeResponseBodyDataControlNode) GetEcsInstanceId() *string {
	return s.EcsInstanceId
}

func (s *GetHoneypotProbeResponseBodyDataControlNode) GetNodeId() *string {
	return s.NodeId
}

func (s *GetHoneypotProbeResponseBodyDataControlNode) GetNodeName() *string {
	return s.NodeName
}

func (s *GetHoneypotProbeResponseBodyDataControlNode) SetEcsInstanceId(v string) *GetHoneypotProbeResponseBodyDataControlNode {
	s.EcsInstanceId = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataControlNode) SetNodeId(v string) *GetHoneypotProbeResponseBodyDataControlNode {
	s.NodeId = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataControlNode) SetNodeName(v string) *GetHoneypotProbeResponseBodyDataControlNode {
	s.NodeName = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataControlNode) Validate() error {
	return dara.Validate(s)
}

type GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort struct {
	// The unique ID of the service that is monitored.
	//
	// example:
	//
	// 15389
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ports that are monitored.
	//
	// example:
	//
	// {
	//
	//     "tcp": "1-65535",
	//
	//     "udp": "1-65535"
	//
	// }
	Ports *string `json:"Ports,omitempty" xml:"Ports,omitempty"`
	// The ID of the probe.
	//
	// example:
	//
	// a46f5162-c70d-4e26-8ddf-7435feca****
	ProbeId *string `json:"ProbeId,omitempty" xml:"ProbeId,omitempty"`
	// The IP addresses that are monitored.
	ServiceIpList []*string `json:"ServiceIpList,omitempty" xml:"ServiceIpList,omitempty" type:"Repeated"`
	// The monitoring status. Valid values:
	//
	// 	- **1**: abnormal
	//
	// 	- **3**: normal
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) GoString() string {
	return s.String()
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) GetId() *int64 {
	return s.Id
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) GetPorts() *string {
	return s.Ports
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) GetProbeId() *string {
	return s.ProbeId
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) GetServiceIpList() []*string {
	return s.ServiceIpList
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) GetStatus() *int32 {
	return s.Status
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) SetId(v int64) *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort {
	s.Id = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) SetPorts(v string) *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort {
	s.Ports = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) SetProbeId(v string) *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort {
	s.ProbeId = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) SetServiceIpList(v []*string) *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort {
	s.ServiceIpList = v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) SetStatus(v int32) *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort {
	s.Status = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneyPotProbeScanPort) Validate() error {
	return dara.Validate(s)
}

type GetHoneypotProbeResponseBodyDataHoneypotProbeBindList struct {
	// The unique ID of the honeypot that is bound to the probe.
	//
	// example:
	//
	// 45378f64-d7b4-4a53-9c48-4303eb4b****
	BindId *string `json:"BindId,omitempty" xml:"BindId,omitempty"`
	// The ports that are bound to the probe.
	BindPortList []*GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList `json:"BindPortList,omitempty" xml:"BindPortList,omitempty" type:"Repeated"`
	// The honeypot ID.
	//
	// example:
	//
	// 913347774a3b3c378c6a50f66de23dfa097765214ec3f0526b01c67bf59c****
	HoneypotId *string `json:"HoneypotId,omitempty" xml:"HoneypotId,omitempty"`
	// The IP addresses that are monitored.
	ServiceIpList []*string `json:"ServiceIpList,omitempty" xml:"ServiceIpList,omitempty" type:"Repeated"`
	// The status of the honeypot that is bound to the probe. Valid values:
	//
	// 	- **1**: abnormal
	//
	// 	- **3**: normal
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) GoString() string {
	return s.String()
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) GetBindId() *string {
	return s.BindId
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) GetBindPortList() []*GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	return s.BindPortList
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) GetHoneypotId() *string {
	return s.HoneypotId
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) GetServiceIpList() []*string {
	return s.ServiceIpList
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) GetStatus() *int32 {
	return s.Status
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) SetBindId(v string) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList {
	s.BindId = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) SetBindPortList(v []*GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList {
	s.BindPortList = v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) SetHoneypotId(v string) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList {
	s.HoneypotId = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) SetServiceIpList(v []*string) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList {
	s.ServiceIpList = v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) SetStatus(v int32) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList {
	s.Status = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindList) Validate() error {
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

type GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList struct {
	// Indicates whether the port is bound.
	//
	// example:
	//
	// false
	BindPort *bool `json:"BindPort,omitempty" xml:"BindPort,omitempty"`
	// The end port on which the probe monitors.
	//
	// example:
	//
	// 80
	EndPort *int32 `json:"EndPort,omitempty" xml:"EndPort,omitempty"`
	// The error that is returned if an error occurred in the port of the honeypot that is bound to the probe.
	//
	// example:
	//
	// portmap failed
	Err *string `json:"Err,omitempty" xml:"Err,omitempty"`
	// Indicates whether the port is a fixed port.
	//
	// example:
	//
	// false
	Fixed *bool `json:"Fixed,omitempty" xml:"Fixed,omitempty"`
	// The unique ID of the port binding record.
	//
	// example:
	//
	// 2512
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message that is returned if an error occurred in the port of the honeypot that is bound to the probe.
	//
	// example:
	//
	// listen 22 tcp4 failed
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The type of the protocol.
	//
	// example:
	//
	// tcp
	Proto *string `json:"Proto,omitempty" xml:"Proto,omitempty"`
	// The start port on which the probe monitors.
	//
	// example:
	//
	// 22
	StartPort *int32 `json:"StartPort,omitempty" xml:"StartPort,omitempty"`
	// The status of the port of the honeypot that is bound to the probe. Valid values:
	//
	// 	- **1**: abnormal
	//
	// 	- **3**: normal
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The destination port.
	//
	// example:
	//
	// 80
	TargetPort *int32 `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
}

func (s GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GoString() string {
	return s.String()
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetBindPort() *bool {
	return s.BindPort
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetEndPort() *int32 {
	return s.EndPort
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetErr() *string {
	return s.Err
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetFixed() *bool {
	return s.Fixed
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetId() *int64 {
	return s.Id
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetMsg() *string {
	return s.Msg
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetProto() *string {
	return s.Proto
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetStartPort() *int32 {
	return s.StartPort
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetStatus() *int32 {
	return s.Status
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetBindPort(v bool) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.BindPort = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetEndPort(v int32) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.EndPort = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetErr(v string) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.Err = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetFixed(v bool) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.Fixed = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetId(v int64) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.Id = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetMsg(v string) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.Msg = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetProto(v string) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.Proto = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetStartPort(v int32) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.StartPort = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetStatus(v int32) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.Status = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) SetTargetPort(v int32) *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList {
	s.TargetPort = &v
	return s
}

func (s *GetHoneypotProbeResponseBodyDataHoneypotProbeBindListBindPortList) Validate() error {
	return dara.Validate(s)
}
