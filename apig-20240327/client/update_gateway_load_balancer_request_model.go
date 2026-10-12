// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerDTO(v *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) *UpdateGatewayLoadBalancerRequest
	GetLoadBalancerDTO() *UpdateGatewayLoadBalancerRequestLoadBalancerDTO
	SetOption(v string) *UpdateGatewayLoadBalancerRequest
	GetOption() *string
	SetPorts(v []*UpdateGatewayLoadBalancerRequestPorts) *UpdateGatewayLoadBalancerRequest
	GetPorts() []*UpdateGatewayLoadBalancerRequestPorts
}

type UpdateGatewayLoadBalancerRequest struct {
	// The load balancer operation DTO.
	LoadBalancerDTO *UpdateGatewayLoadBalancerRequestLoadBalancerDTO `json:"loadBalancerDTO,omitempty" xml:"loadBalancerDTO,omitempty" type:"Struct"`
	// The operation type. If this parameter is left empty or not specified, only the listening ports are updated (only the ports field is used).
	//
	// example:
	//
	// Add
	Option *string `json:"option,omitempty" xml:"option,omitempty"`
	// The port list.
	Ports []*UpdateGatewayLoadBalancerRequestPorts `json:"ports,omitempty" xml:"ports,omitempty" type:"Repeated"`
}

func (s UpdateGatewayLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLoadBalancerRequest) GetLoadBalancerDTO() *UpdateGatewayLoadBalancerRequestLoadBalancerDTO {
	return s.LoadBalancerDTO
}

func (s *UpdateGatewayLoadBalancerRequest) GetOption() *string {
	return s.Option
}

func (s *UpdateGatewayLoadBalancerRequest) GetPorts() []*UpdateGatewayLoadBalancerRequestPorts {
	return s.Ports
}

func (s *UpdateGatewayLoadBalancerRequest) SetLoadBalancerDTO(v *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) *UpdateGatewayLoadBalancerRequest {
	s.LoadBalancerDTO = v
	return s
}

func (s *UpdateGatewayLoadBalancerRequest) SetOption(v string) *UpdateGatewayLoadBalancerRequest {
	s.Option = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequest) SetPorts(v []*UpdateGatewayLoadBalancerRequestPorts) *UpdateGatewayLoadBalancerRequest {
	s.Ports = v
	return s
}

func (s *UpdateGatewayLoadBalancerRequest) Validate() error {
	if s.LoadBalancerDTO != nil {
		if err := s.LoadBalancerDTO.Validate(); err != nil {
			return err
		}
	}
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateGatewayLoadBalancerRequestLoadBalancerDTO struct {
	// The load balancer instance ID.
	//
	// example:
	//
	// lb-bp1xxxx
	LoadBalancerId *string `json:"loadBalancerId,omitempty" xml:"loadBalancerId,omitempty"`
	// The load balancer type.
	//
	// example:
	//
	// CLB
	LoadBalancerType *string `json:"loadBalancerType,omitempty" xml:"loadBalancerType,omitempty"`
	// The network type.
	//
	// example:
	//
	// Internet
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
	// The service weight. Valid values: 0 to 100.
	//
	// example:
	//
	// 0
	ServiceWeight *int64 `json:"serviceWeight,omitempty" xml:"serviceWeight,omitempty"`
	// The backend server group bindings.
	VirtualServiceList []*UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList `json:"virtualServiceList,omitempty" xml:"virtualServiceList,omitempty" type:"Repeated"`
}

func (s UpdateGatewayLoadBalancerRequestLoadBalancerDTO) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLoadBalancerRequestLoadBalancerDTO) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) GetLoadBalancerType() *string {
	return s.LoadBalancerType
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) GetNetworkType() *string {
	return s.NetworkType
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) GetServiceWeight() *int64 {
	return s.ServiceWeight
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) GetVirtualServiceList() []*UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList {
	return s.VirtualServiceList
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) SetLoadBalancerId(v string) *UpdateGatewayLoadBalancerRequestLoadBalancerDTO {
	s.LoadBalancerId = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) SetLoadBalancerType(v string) *UpdateGatewayLoadBalancerRequestLoadBalancerDTO {
	s.LoadBalancerType = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) SetNetworkType(v string) *UpdateGatewayLoadBalancerRequestLoadBalancerDTO {
	s.NetworkType = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) SetServiceWeight(v int64) *UpdateGatewayLoadBalancerRequestLoadBalancerDTO {
	s.ServiceWeight = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) SetVirtualServiceList(v []*UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) *UpdateGatewayLoadBalancerRequestLoadBalancerDTO {
	s.VirtualServiceList = v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTO) Validate() error {
	if s.VirtualServiceList != nil {
		for _, item := range s.VirtualServiceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList struct {
	// The port number. Valid values: 1 to 65535. An error is returned if an invalid value is specified.
	//
	// example:
	//
	// 80
	Port *int64 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol. This parameter is required for CLB. If this parameter is not specified for CLB, the error message "protocol is missing for CLB" is returned. No backend enumeration validation is performed. Valid values: tcp, udp, http, https, and tls.
	//
	// example:
	//
	// http
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The CLB VServerGroupId or NLB ServerGroupId. This parameter is required. If a duplicate group ID is specified, an error is returned.
	//
	// example:
	//
	// rsp-xxxx
	VirtualServiceGroupId *string `json:"virtualServiceGroupId,omitempty" xml:"virtualServiceGroupId,omitempty"`
	// The group name. This parameter is required. If this parameter is not specified, the error message "virtual service parameter is missing" is returned.
	//
	// example:
	//
	// 80-tcp
	VirtualServiceGroupName *string `json:"virtualServiceGroupName,omitempty" xml:"virtualServiceGroupName,omitempty"`
}

func (s UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) GetPort() *int64 {
	return s.Port
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) GetVirtualServiceGroupId() *string {
	return s.VirtualServiceGroupId
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) GetVirtualServiceGroupName() *string {
	return s.VirtualServiceGroupName
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) SetPort(v int64) *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList {
	s.Port = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) SetProtocol(v string) *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList {
	s.Protocol = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) SetVirtualServiceGroupId(v string) *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList {
	s.VirtualServiceGroupId = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) SetVirtualServiceGroupName(v string) *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList {
	s.VirtualServiceGroupName = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestLoadBalancerDTOVirtualServiceList) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayLoadBalancerRequestPorts struct {
	// The listening ports.
	GatewayLoadBalancerPorts []*UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts `json:"gatewayLoadBalancerPorts,omitempty" xml:"gatewayLoadBalancerPorts,omitempty" type:"Repeated"`
	// The load balancer type.
	//
	// example:
	//
	// NLB
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s UpdateGatewayLoadBalancerRequestPorts) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLoadBalancerRequestPorts) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLoadBalancerRequestPorts) GetGatewayLoadBalancerPorts() []*UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts {
	return s.GatewayLoadBalancerPorts
}

func (s *UpdateGatewayLoadBalancerRequestPorts) GetType() *string {
	return s.Type
}

func (s *UpdateGatewayLoadBalancerRequestPorts) SetGatewayLoadBalancerPorts(v []*UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts) *UpdateGatewayLoadBalancerRequestPorts {
	s.GatewayLoadBalancerPorts = v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestPorts) SetType(v string) *UpdateGatewayLoadBalancerRequestPorts {
	s.Type = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestPorts) Validate() error {
	if s.GatewayLoadBalancerPorts != nil {
		for _, item := range s.GatewayLoadBalancerPorts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts struct {
	// The listening port. Valid values: 0 to 65535. A maximum of 50 custom ports can be configured in the console.
	//
	// example:
	//
	// 8080
	Port *int32 `json:"port,omitempty" xml:"port,omitempty"`
	// The protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts) GoString() string {
	return s.String()
}

func (s *UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts) GetPort() *int32 {
	return s.Port
}

func (s *UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts) SetPort(v int32) *UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts {
	s.Port = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts) SetProtocol(v string) *UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts {
	s.Protocol = &v
	return s
}

func (s *UpdateGatewayLoadBalancerRequestPortsGatewayLoadBalancerPorts) Validate() error {
	return dara.Validate(s)
}
