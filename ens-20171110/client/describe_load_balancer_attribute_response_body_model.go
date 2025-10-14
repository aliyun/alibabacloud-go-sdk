// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *DescribeLoadBalancerAttributeResponseBody
	GetAddress() *string
	SetAddressIPVersion(v string) *DescribeLoadBalancerAttributeResponseBody
	GetAddressIPVersion() *string
	SetAddressType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetAddressType() *string
	SetBackendServers(v []*DescribeLoadBalancerAttributeResponseBodyBackendServers) *DescribeLoadBalancerAttributeResponseBody
	GetBackendServers() []*DescribeLoadBalancerAttributeResponseBodyBackendServers
	SetBandwidth(v int32) *DescribeLoadBalancerAttributeResponseBody
	GetBandwidth() *int32
	SetCreateTime(v string) *DescribeLoadBalancerAttributeResponseBody
	GetCreateTime() *string
	SetEndTime(v string) *DescribeLoadBalancerAttributeResponseBody
	GetEndTime() *string
	SetEnsRegionId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetEnsRegionId() *string
	SetListenerPorts(v []*string) *DescribeLoadBalancerAttributeResponseBody
	GetListenerPorts() []*string
	SetListenerPortsAndProtocols(v []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) *DescribeLoadBalancerAttributeResponseBody
	GetListenerPortsAndProtocols() []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols
	SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerName() *string
	SetLoadBalancerSpec(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerSpec() *string
	SetLoadBalancerStatus(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerStatus() *string
	SetLoadBalancerType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerType() *string
	SetNetworkId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetNetworkId() *string
	SetPayType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetPayType() *string
	SetRequestId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetRequestId() *string
	SetVSwitchId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetVSwitchId() *string
}

type DescribeLoadBalancerAttributeResponseBody struct {
	// The IP address that the Edge Load Balancer (ELB) instance uses to provide services.
	//
	// example:
	//
	// 192.168XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The IP version of the ELB instance. Valid values: ipv4 and ipv6.
	//
	// example:
	//
	// ipv4
	AddressIPVersion *string `json:"AddressIPVersion,omitempty" xml:"AddressIPVersion,omitempty"`
	AddressType      *string `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	// The list of backend servers.
	BackendServers []*DescribeLoadBalancerAttributeResponseBodyBackendServers `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Repeated"`
	// The peak bandwidth of the ELB. The default value is -1, which indicates that the bandwidth is unlimited.
	//
	// example:
	//
	// -1
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The time when the ELB instance was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-05-21T12:22:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the ELB instance was disabled.
	//
	// example:
	//
	// 2020-05-21T12:22:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the ENS node.
	//
	// example:
	//
	// cn-chengdu-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The frontend ports that are used by the ELB instance.
	ListenerPorts []*string `json:"ListenerPorts,omitempty" xml:"ListenerPorts,omitempty" type:"Repeated"`
	// The frontend ports and protocols that are used by the ELB instance.
	ListenerPortsAndProtocols []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols `json:"ListenerPortsAndProtocols,omitempty" xml:"ListenerPortsAndProtocols,omitempty" type:"Repeated"`
	// The ID of the ELB instance.
	//
	// example:
	//
	// lb-5rcvo1n1t3hykfhhjwjgqp****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The name of the ELB instance.
	//
	// example:
	//
	// example
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	// The specifications of the ELB instance.
	//
	// example:
	//
	// elb.s2.medium
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	// The status of the ELB instance. Valid values:
	//
	// 	- **Active*	- (default): The listener for the instance can forward the received traffic based on the rule.
	//
	// 	- **InActive**: The listener for the instance does not forward the received traffic.
	//
	// example:
	//
	// Active
	LoadBalancerStatus *string `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	LoadBalancerType   *string `json:"LoadBalancerType,omitempty" xml:"LoadBalancerType,omitempty"`
	// The ID of the network.
	//
	// example:
	//
	// n-5rwbi3g9zvgxcsiufwhw8****
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The billing method. Valid values:
	//
	// 	- **PrePaid**: subscription.
	//
	// 	- **PostPaid**: pay-as-you-go. Only this billing method is supported.
	//
	// example:
	//
	// PostPaid
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-5s78m2pdr9osa0j64bn78****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetAddress() *string {
	return s.Address
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetAddressIPVersion() *string {
	return s.AddressIPVersion
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetBackendServers() []*DescribeLoadBalancerAttributeResponseBodyBackendServers {
	return s.BackendServers
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetListenerPorts() []*string {
	return s.ListenerPorts
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetListenerPortsAndProtocols() []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols {
	return s.ListenerPortsAndProtocols
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetLoadBalancerStatus() *string {
	return s.LoadBalancerStatus
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetNetworkId() *string {
	return s.NetworkId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddress(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddressIPVersion(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.AddressIPVersion = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddressType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetBackendServers(v []*DescribeLoadBalancerAttributeResponseBodyBackendServers) *DescribeLoadBalancerAttributeResponseBody {
	s.BackendServers = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetBandwidth(v int32) *DescribeLoadBalancerAttributeResponseBody {
	s.Bandwidth = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetCreateTime(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetEndTime(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetEnsRegionId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.EnsRegionId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPorts(v []*string) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPorts = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPortsAndProtocols(v []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPortsAndProtocols = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerName(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerName = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerSpec(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerSpec = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerStatus(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerStatus = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetLoadBalancerType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.LoadBalancerType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetNetworkId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.NetworkId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetPayType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetVSwitchId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) Validate() error {
	if s.BackendServers != nil {
		for _, item := range s.BackendServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ListenerPortsAndProtocols != nil {
		for _, item := range s.ListenerPortsAndProtocols {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyBackendServers struct {
	// The IP address of the backend server.
	//
	// example:
	//
	// 192.168.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The port that is used by the backend server.
	//
	// example:
	//
	// 0
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the backend server.
	//
	// example:
	//
	// i-5vb5h5njxiuhn48a****
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	// The type of backend server.
	//
	// example:
	//
	// ens
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The weight of the backend server.
	//
	// example:
	//
	// 100
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) GetIp() *string {
	return s.Ip
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) GetPort() *string {
	return s.Port
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) GetType() *string {
	return s.Type
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) SetIp(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	s.Ip = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) SetPort(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	s.Port = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) SetServerId(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	s.ServerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) SetType(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	s.Type = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) SetWeight(v int32) *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	s.Weight = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols struct {
	// The backend port that is used by the ELB instance. Valid values: **1*	- to **65535**.
	//
	// example:
	//
	// 6000
	BackendServerPort *int32 `json:"BackendServerPort,omitempty" xml:"BackendServerPort,omitempty"`
	// The description of the listener.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination listening port to which requests are forwarded.
	//
	// example:
	//
	// 0
	ForwardPort *int32 `json:"ForwardPort,omitempty" xml:"ForwardPort,omitempty"`
	// Indicates whether the listener is enabled.
	//
	// example:
	//
	// off
	ListenerForward *string `json:"ListenerForward,omitempty" xml:"ListenerForward,omitempty"`
	// The listener port of the instance.
	//
	// example:
	//
	// 8080
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The listener protocol of the instance.
	//
	// example:
	//
	// tcp
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) GetBackendServerPort() *int32 {
	return s.BackendServerPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) GetDescription() *string {
	return s.Description
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) GetForwardPort() *int32 {
	return s.ForwardPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) GetListenerForward() *string {
	return s.ListenerForward
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) SetBackendServerPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols {
	s.BackendServerPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) SetDescription(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols {
	s.Description = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) SetForwardPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols {
	s.ForwardPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) SetListenerForward(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols {
	s.ListenerForward = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) SetListenerPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) SetListenerProtocol(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocols) Validate() error {
	return dara.Validate(s)
}
