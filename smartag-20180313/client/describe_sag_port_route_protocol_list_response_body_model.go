// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagPortRouteProtocolListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPorts(v []*DescribeSagPortRouteProtocolListResponseBodyPorts) *DescribeSagPortRouteProtocolListResponseBody
	GetPorts() []*DescribeSagPortRouteProtocolListResponseBodyPorts
	SetRequestId(v string) *DescribeSagPortRouteProtocolListResponseBody
	GetRequestId() *string
	SetTaskStates(v []*DescribeSagPortRouteProtocolListResponseBodyTaskStates) *DescribeSagPortRouteProtocolListResponseBody
	GetTaskStates() []*DescribeSagPortRouteProtocolListResponseBodyTaskStates
}

type DescribeSagPortRouteProtocolListResponseBody struct {
	// The list of port information.
	Ports []*DescribeSagPortRouteProtocolListResponseBodyPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The query task status.
	TaskStates []*DescribeSagPortRouteProtocolListResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagPortRouteProtocolListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortRouteProtocolListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagPortRouteProtocolListResponseBody) GetPorts() []*DescribeSagPortRouteProtocolListResponseBodyPorts {
	return s.Ports
}

func (s *DescribeSagPortRouteProtocolListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagPortRouteProtocolListResponseBody) GetTaskStates() []*DescribeSagPortRouteProtocolListResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagPortRouteProtocolListResponseBody) SetPorts(v []*DescribeSagPortRouteProtocolListResponseBodyPorts) *DescribeSagPortRouteProtocolListResponseBody {
	s.Ports = v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBody) SetRequestId(v string) *DescribeSagPortRouteProtocolListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBody) SetTaskStates(v []*DescribeSagPortRouteProtocolListResponseBodyTaskStates) *DescribeSagPortRouteProtocolListResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBody) Validate() error {
	if s.Ports != nil {
		for _, item := range s.Ports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TaskStates != nil {
		for _, item := range s.TaskStates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSagPortRouteProtocolListResponseBodyPorts struct {
	// The neighbor IP address.
	//
	// example:
	//
	// 192.XX.XX.1
	NeighborIp *string `json:"NeighborIp,omitempty" xml:"NeighborIp,omitempty"`
	// The port name.
	//
	// example:
	//
	// 3
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The autonomous system number of the peer BGP network.
	//
	// example:
	//
	// 12345
	RemoteAs *string `json:"RemoteAs,omitempty" xml:"RemoteAs,omitempty"`
	// The IP address of the peer.
	//
	// example:
	//
	// 192.XX.XX.1
	RemoteIp *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	// The routable protocol of the port. Valid values:
	//
	// - **STATIC**: static routable protocol.
	//
	// - **OSPF**: OSPF dynamic routable protocol.
	//
	// - **BGP**: BGP dynamic routable protocol.
	//
	// example:
	//
	// BGP
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
	// The port status. Valid values:
	//
	// - **UP**: The port is enabled.
	//
	// - **DOWN**: The port is disabled.
	//
	// example:
	//
	// UP
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VLAN ID.
	//
	// example:
	//
	// 2
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s DescribeSagPortRouteProtocolListResponseBodyPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortRouteProtocolListResponseBodyPorts) GoString() string {
	return s.String()
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) GetNeighborIp() *string {
	return s.NeighborIp
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) GetPortName() *string {
	return s.PortName
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) GetRemoteAs() *string {
	return s.RemoteAs
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) GetRemoteIp() *string {
	return s.RemoteIp
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) GetRouteProtocol() *string {
	return s.RouteProtocol
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) GetStatus() *string {
	return s.Status
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) GetVlan() *string {
	return s.Vlan
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) SetNeighborIp(v string) *DescribeSagPortRouteProtocolListResponseBodyPorts {
	s.NeighborIp = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) SetPortName(v string) *DescribeSagPortRouteProtocolListResponseBodyPorts {
	s.PortName = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) SetRemoteAs(v string) *DescribeSagPortRouteProtocolListResponseBodyPorts {
	s.RemoteAs = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) SetRemoteIp(v string) *DescribeSagPortRouteProtocolListResponseBodyPorts {
	s.RemoteIp = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) SetRouteProtocol(v string) *DescribeSagPortRouteProtocolListResponseBodyPorts {
	s.RouteProtocol = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) SetStatus(v string) *DescribeSagPortRouteProtocolListResponseBodyPorts {
	s.Status = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) SetVlan(v string) *DescribeSagPortRouteProtocolListResponseBodyPorts {
	s.Vlan = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeSagPortRouteProtocolListResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586843621000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code. A value of 200 indicates that the query task succeeded.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. A value of Successful indicates that the query task succeeded.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the asynchronous task. Valid values:
	//
	// - **Initialized**: The query task is initialized.
	//
	// - **Offline**: The Smart Access Gateway device is offline and the query task has not been delivered. The task will be delivered after the device comes online.
	//
	// - **Succeed**: The query task is delivered.
	//
	// - **Processing**: The query task is being delivered.
	//
	// - **VersionNotSupport**: The current version of the Smart Access Gateway device does not support this operation.
	//
	// - **BuildRequestError**: The China Cloud Management Platform does not support this operation.
	//
	// - **HardwareError**: The query task failed to be delivered due to a device error.
	//
	// - **TaskNotExist**: The query task does not exist.
	//
	// - **OfflineNotConfiged**: The Smart Access Gateway device is offline and the query task has not been delivered. The task will not be delivered even after the device comes online.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSagPortRouteProtocolListResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortRouteProtocolListResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagPortRouteProtocolListResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagPortRouteProtocolListResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagPortRouteProtocolListResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagPortRouteProtocolListResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagPortRouteProtocolListResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagPortRouteProtocolListResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagPortRouteProtocolListResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagPortRouteProtocolListResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyTaskStates) SetState(v string) *DescribeSagPortRouteProtocolListResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagPortRouteProtocolListResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
