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
	SetAddressType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetAddressType() *string
	SetAutoReleaseTime(v int64) *DescribeLoadBalancerAttributeResponseBody
	GetAutoReleaseTime() *int64
	SetBackendServers(v *DescribeLoadBalancerAttributeResponseBodyBackendServers) *DescribeLoadBalancerAttributeResponseBody
	GetBackendServers() *DescribeLoadBalancerAttributeResponseBodyBackendServers
	SetBandwidth(v int32) *DescribeLoadBalancerAttributeResponseBody
	GetBandwidth() *int32
	SetCreateTime(v string) *DescribeLoadBalancerAttributeResponseBody
	GetCreateTime() *string
	SetCreateTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody
	GetCreateTimeStamp() *int64
	SetEndTime(v string) *DescribeLoadBalancerAttributeResponseBody
	GetEndTime() *string
	SetEndTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody
	GetEndTimeStamp() *int64
	SetInternetChargeType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetInternetChargeType() *string
	SetIsPublicAddress(v string) *DescribeLoadBalancerAttributeResponseBody
	GetIsPublicAddress() *string
	SetListenerPorts(v *DescribeLoadBalancerAttributeResponseBodyListenerPorts) *DescribeLoadBalancerAttributeResponseBody
	GetListenerPorts() *DescribeLoadBalancerAttributeResponseBodyListenerPorts
	SetListenerPortsAndProtocal(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) *DescribeLoadBalancerAttributeResponseBody
	GetListenerPortsAndProtocal() *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal
	SetListenerPortsAndProtocol(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) *DescribeLoadBalancerAttributeResponseBody
	GetListenerPortsAndProtocol() *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol
	SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerName() *string
	SetLoadBalancerSpec(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerSpec() *string
	SetLoadBalancerStatus(v string) *DescribeLoadBalancerAttributeResponseBody
	GetLoadBalancerStatus() *string
	SetMasterZoneId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetMasterZoneId() *string
	SetNetworkType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetNetworkType() *string
	SetPayType(v string) *DescribeLoadBalancerAttributeResponseBody
	GetPayType() *string
	SetRegionId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetRegionId() *string
	SetRegionIdAlias(v string) *DescribeLoadBalancerAttributeResponseBody
	GetRegionIdAlias() *string
	SetRequestId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetRequestId() *string
	SetSlaveZoneId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetSlaveZoneId() *string
	SetVSwitchId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetVSwitchId() *string
	SetVpcId(v string) *DescribeLoadBalancerAttributeResponseBody
	GetVpcId() *string
}

type DescribeLoadBalancerAttributeResponseBody struct {
	Address                  *string                                                            `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressType              *string                                                            `json:"AddressType,omitempty" xml:"AddressType,omitempty"`
	AutoReleaseTime          *int64                                                             `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	BackendServers           *DescribeLoadBalancerAttributeResponseBodyBackendServers           `json:"BackendServers,omitempty" xml:"BackendServers,omitempty" type:"Struct"`
	Bandwidth                *int32                                                             `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	CreateTime               *string                                                            `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimeStamp          *int64                                                             `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	EndTime                  *string                                                            `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeStamp             *int64                                                             `json:"EndTimeStamp,omitempty" xml:"EndTimeStamp,omitempty"`
	InternetChargeType       *string                                                            `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IsPublicAddress          *string                                                            `json:"IsPublicAddress,omitempty" xml:"IsPublicAddress,omitempty"`
	ListenerPorts            *DescribeLoadBalancerAttributeResponseBodyListenerPorts            `json:"ListenerPorts,omitempty" xml:"ListenerPorts,omitempty" type:"Struct"`
	ListenerPortsAndProtocal *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal `json:"ListenerPortsAndProtocal,omitempty" xml:"ListenerPortsAndProtocal,omitempty" type:"Struct"`
	ListenerPortsAndProtocol *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol `json:"ListenerPortsAndProtocol,omitempty" xml:"ListenerPortsAndProtocol,omitempty" type:"Struct"`
	LoadBalancerId           *string                                                            `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	LoadBalancerName         *string                                                            `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	LoadBalancerSpec         *string                                                            `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
	LoadBalancerStatus       *string                                                            `json:"LoadBalancerStatus,omitempty" xml:"LoadBalancerStatus,omitempty"`
	MasterZoneId             *string                                                            `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	NetworkType              *string                                                            `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	PayType                  *string                                                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId                 *string                                                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionIdAlias            *string                                                            `json:"RegionIdAlias,omitempty" xml:"RegionIdAlias,omitempty"`
	RequestId                *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlaveZoneId              *string                                                            `json:"SlaveZoneId,omitempty" xml:"SlaveZoneId,omitempty"`
	VSwitchId                *string                                                            `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                    *string                                                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *DescribeLoadBalancerAttributeResponseBody) GetAddressType() *string {
	return s.AddressType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetAutoReleaseTime() *int64 {
	return s.AutoReleaseTime
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetBackendServers() *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	return s.BackendServers
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetEndTimeStamp() *int64 {
	return s.EndTimeStamp
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetIsPublicAddress() *string {
	return s.IsPublicAddress
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetListenerPorts() *DescribeLoadBalancerAttributeResponseBodyListenerPorts {
	return s.ListenerPorts
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetListenerPortsAndProtocal() *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal {
	return s.ListenerPortsAndProtocal
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetListenerPortsAndProtocol() *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol {
	return s.ListenerPortsAndProtocol
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

func (s *DescribeLoadBalancerAttributeResponseBody) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRegionIdAlias() *string {
	return s.RegionIdAlias
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetSlaveZoneId() *string {
	return s.SlaveZoneId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeLoadBalancerAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddress(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAddressType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.AddressType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetAutoReleaseTime(v int64) *DescribeLoadBalancerAttributeResponseBody {
	s.AutoReleaseTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetBackendServers(v *DescribeLoadBalancerAttributeResponseBodyBackendServers) *DescribeLoadBalancerAttributeResponseBody {
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

func (s *DescribeLoadBalancerAttributeResponseBody) SetCreateTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody {
	s.CreateTimeStamp = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetEndTime(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetEndTimeStamp(v int64) *DescribeLoadBalancerAttributeResponseBody {
	s.EndTimeStamp = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetInternetChargeType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.InternetChargeType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetIsPublicAddress(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.IsPublicAddress = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPorts(v *DescribeLoadBalancerAttributeResponseBodyListenerPorts) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPorts = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPortsAndProtocal(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPortsAndProtocal = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetListenerPortsAndProtocol(v *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) *DescribeLoadBalancerAttributeResponseBody {
	s.ListenerPortsAndProtocol = v
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

func (s *DescribeLoadBalancerAttributeResponseBody) SetMasterZoneId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.MasterZoneId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetNetworkType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.NetworkType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetPayType(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRegionId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRegionIdAlias(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RegionIdAlias = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetRequestId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetSlaveZoneId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.SlaveZoneId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetVSwitchId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) SetVpcId(v string) *DescribeLoadBalancerAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBody) Validate() error {
	if s.BackendServers != nil {
		if err := s.BackendServers.Validate(); err != nil {
			return err
		}
	}
	if s.ListenerPorts != nil {
		if err := s.ListenerPorts.Validate(); err != nil {
			return err
		}
	}
	if s.ListenerPortsAndProtocal != nil {
		if err := s.ListenerPortsAndProtocal.Validate(); err != nil {
			return err
		}
	}
	if s.ListenerPortsAndProtocol != nil {
		if err := s.ListenerPortsAndProtocol.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyBackendServers struct {
	BackendServer []*DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer `json:"BackendServer,omitempty" xml:"BackendServer,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServers) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServers) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) GetBackendServer() []*DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	return s.BackendServer
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) SetBackendServer(v []*DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) *DescribeLoadBalancerAttributeResponseBodyBackendServers {
	s.BackendServer = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServers) Validate() error {
	if s.BackendServer != nil {
		for _, item := range s.BackendServer {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer struct {
	ServerId *string `json:"ServerId,omitempty" xml:"ServerId,omitempty"`
	Weight   *int32  `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GetServerId() *string {
	return s.ServerId
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) GetWeight() *int32 {
	return s.Weight
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetServerId(v string) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.ServerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) SetWeight(v int32) *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer {
	s.Weight = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyBackendServersBackendServer) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerAttributeResponseBodyListenerPorts struct {
	ListenerPort []*string `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPorts) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPorts) GetListenerPort() []*string {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPorts) SetListenerPort(v []*string) *DescribeLoadBalancerAttributeResponseBodyListenerPorts {
	s.ListenerPort = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal struct {
	ListenerPortAndProtocal []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal `json:"ListenerPortAndProtocal,omitempty" xml:"ListenerPortAndProtocal,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) GetListenerPortAndProtocal() []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal {
	return s.ListenerPortAndProtocal
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) SetListenerPortAndProtocal(v []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal {
	s.ListenerPortAndProtocal = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocal) Validate() error {
	if s.ListenerPortAndProtocal != nil {
		for _, item := range s.ListenerPortAndProtocal {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal struct {
	ListenerPort     *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocal *string `json:"ListenerProtocal,omitempty" xml:"ListenerProtocal,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) GetListenerProtocal() *string {
	return s.ListenerProtocal
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) SetListenerPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) SetListenerProtocal(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal {
	s.ListenerProtocal = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocalListenerPortAndProtocal) Validate() error {
	return dara.Validate(s)
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol struct {
	ListenerPortAndProtocol []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol `json:"ListenerPortAndProtocol,omitempty" xml:"ListenerPortAndProtocol,omitempty" type:"Repeated"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) GetListenerPortAndProtocol() []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	return s.ListenerPortAndProtocol
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) SetListenerPortAndProtocol(v []*DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol {
	s.ListenerPortAndProtocol = v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocol) Validate() error {
	if s.ListenerPortAndProtocol != nil {
		for _, item := range s.ListenerPortAndProtocol {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol struct {
	ListenerPort     *int32  `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	ListenerProtocol *string `json:"ListenerProtocol,omitempty" xml:"ListenerProtocol,omitempty"`
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) GetListenerProtocol() *string {
	return s.ListenerProtocol
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetListenerPort(v int32) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) SetListenerProtocol(v string) *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol {
	s.ListenerProtocol = &v
	return s
}

func (s *DescribeLoadBalancerAttributeResponseBodyListenerPortsAndProtocolListenerPortAndProtocol) Validate() error {
	return dara.Validate(s)
}
