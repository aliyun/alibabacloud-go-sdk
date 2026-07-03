// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayPortRouteProtocolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPorts(v []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) *ViewSmartAccessGatewayPortRouteProtocolResponseBody
	GetPorts() []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts
	SetRequestId(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBody
	GetRequestId() *string
	SetTaskStates(v []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) *ViewSmartAccessGatewayPortRouteProtocolResponseBody
	GetTaskStates() []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates
}

type ViewSmartAccessGatewayPortRouteProtocolResponseBody struct {
	// The list of port information.
	Ports []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 877F5673-FFD1-5168-99D1-1E8009FBFF7B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the query task.
	TaskStates []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s ViewSmartAccessGatewayPortRouteProtocolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayPortRouteProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBody) GetPorts() []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts {
	return s.Ports
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBody) GetTaskStates() []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates {
	return s.TaskStates
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBody) SetPorts(v []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) *ViewSmartAccessGatewayPortRouteProtocolResponseBody {
	s.Ports = v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBody) SetRequestId(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBody) SetTaskStates(v []*ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) *ViewSmartAccessGatewayPortRouteProtocolResponseBody {
	s.TaskStates = v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBody) Validate() error {
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

type ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts struct {
	// The IP address of the neighbor.
	//
	// example:
	//
	// 192.XX.XX.2
	NeighborIp *string `json:"NeighborIp,omitempty" xml:"NeighborIp,omitempty"`
	// The name of the port.
	//
	// example:
	//
	// 5
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The autonomous system (AS) number of the BGP peer.
	//
	// example:
	//
	// 65535
	RemoteAs *string `json:"RemoteAs,omitempty" xml:"RemoteAs,omitempty"`
	// The IP address of the peer.
	//
	// example:
	//
	// 192.XX.XX.1
	RemoteIp *string `json:"RemoteIp,omitempty" xml:"RemoteIp,omitempty"`
	// The routable protocol of the port. Valid values:
	//
	// - **STATIC**: static routing protocol.
	//
	// - **OSPF**: Open Shortest Path First (OSPF) dynamic routing protocol.
	//
	// - **BGP**: Border Gateway Protocol (BGP) dynamic routing protocol.
	//
	// example:
	//
	// BGP
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
	// The status of the port. Valid values:
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

func (s ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) GetNeighborIp() *string {
	return s.NeighborIp
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) GetPortName() *string {
	return s.PortName
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) GetRemoteAs() *string {
	return s.RemoteAs
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) GetRemoteIp() *string {
	return s.RemoteIp
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) GetRouteProtocol() *string {
	return s.RouteProtocol
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) GetStatus() *string {
	return s.Status
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) GetVlan() *string {
	return s.Vlan
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) SetNeighborIp(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts {
	s.NeighborIp = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) SetPortName(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts {
	s.PortName = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) SetRemoteAs(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts {
	s.RemoteAs = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) SetRemoteIp(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts {
	s.RemoteIp = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) SetRouteProtocol(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts {
	s.RouteProtocol = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) SetStatus(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts {
	s.Status = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) SetVlan(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts {
	s.Vlan = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyPorts) Validate() error {
	return dara.Validate(s)
}

type ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// This is a UNIX timestamp that represents the number of milliseconds that have elapsed since 00:00:00 UTC on January 1, 1970.
	//
	// example:
	//
	// 1586765938000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code. \\`200\\` indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. \\`Successful\\` indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the asynchronous task:
	//
	// - **Initialized**: The query task is being initialized.
	//
	// - **Offline**: The SAG device is offline and the query task is not sent. The task will be sent after the device goes online.
	//
	// - **Succeed**: The query task is sent.
	//
	// - **Processing**: The query task is being sent.
	//
	// - **VersionNotSupport**: The current version of the SAG device is not supported.
	//
	// - **BuildRequestError**: The management plane does not support the operation.
	//
	// - **HardwareError**: The query task failed to be sent due to a device error.
	//
	// - **TaskNotExist**: The query task does not exist.
	//
	// - **OfflineNotConfiged**: The SAG device is offline and the query task is not sent. The task will not be sent even after the device goes online.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) SetCreateTime(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) SetErrorCode(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) SetErrorMessage(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) SetState(v string) *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *ViewSmartAccessGatewayPortRouteProtocolResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
