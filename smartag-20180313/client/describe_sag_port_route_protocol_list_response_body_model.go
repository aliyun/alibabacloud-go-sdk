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
	// An array that consists of the details of the port.
	Ports []*DescribeSagPortRouteProtocolListResponseBodyPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details about the status of the query task.
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
	// The IP address of the neighbor device.
	//
	// example:
	//
	// 192.XX.XX.1
	NeighborIp *string `json:"NeighborIp,omitempty" xml:"NeighborIp,omitempty"`
	// The name of the port.
	//
	// example:
	//
	// 3
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The number of the autonomous system (AS) to which the SAG device belongs.
	//
	// example:
	//
	// 12345
	RemoteAs *string `json:"RemoteAs,omitempty" xml:"RemoteAs,omitempty"`
	// The IP address of the peer device.
	//
	// example:
	//
	// 192.XX.XX.1
	RemoteIp *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	// The routing protocol. Valid values:
	//
	// 	- **STATIC**: static routing protocol
	//
	// 	- **OSPF**: Open Shortest Path First protocol (OSPF)
	//
	// 	- **BGP**: Border Gateway Protocol (BGP)
	//
	// example:
	//
	// BGP
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
	// The status of the port. Valid values:
	//
	// 	- **UP**: The port was enabled.
	//
	// 	- **DOWN**: The port was disabled.
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
	// The error code returned. A value of 200 indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned. A value of Successful indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the query task. Valid values:
	//
	// 	- **Initialized**: The query task is initialized.
	//
	// 	- **Offline**: The SAG device is disconnected from Alibaba Cloud and Alibaba Cloud has not assigned the query task to the SAG device. After the SAG device is connected to Alibaba Cloud, Alibaba Cloud assigns the query task to the SAG device.
	//
	// 	- **Succeed**: Alibaba Cloud has assigned the query task to the SAG device.
	//
	// 	- **Processing**: Alibaba Cloud is assigning the query task to the SAG device.
	//
	// 	- **VersionNotSupport**: The query task is not supported by the current version of the SAG device.
	//
	// 	- **BuildRequestError**: The query task is not supported by the controller of the SAG device.
	//
	// 	- **HardwareError**: Alibaba Cloud failed to assign the query task to the SAG device because the SAG device is faulty.
	//
	// 	- **TaskNotExist**: The query task does not exist.
	//
	// 	- **OfflineNotConfiged**: The SAG device is disconnected from Alibaba Cloud and Alibaba Cloud has not assigned the query task to the SAG device. Alibaba Cloud does not assign the query task to the SAG device even after the SAG device is connected to Alibaba Cloud.
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
