// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagPortListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPorts(v []*DescribeSagPortListResponseBodyPorts) *DescribeSagPortListResponseBody
	GetPorts() []*DescribeSagPortListResponseBodyPorts
	SetRequestId(v string) *DescribeSagPortListResponseBody
	GetRequestId() *string
	SetTaskStates(v []*DescribeSagPortListResponseBodyTaskStates) *DescribeSagPortListResponseBody
	GetTaskStates() []*DescribeSagPortListResponseBodyTaskStates
}

type DescribeSagPortListResponseBody struct {
	// The list of port information.
	Ports []*DescribeSagPortListResponseBodyPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the query task.
	TaskStates []*DescribeSagPortListResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagPortListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagPortListResponseBody) GetPorts() []*DescribeSagPortListResponseBodyPorts {
	return s.Ports
}

func (s *DescribeSagPortListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagPortListResponseBody) GetTaskStates() []*DescribeSagPortListResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagPortListResponseBody) SetPorts(v []*DescribeSagPortListResponseBodyPorts) *DescribeSagPortListResponseBody {
	s.Ports = v
	return s
}

func (s *DescribeSagPortListResponseBody) SetRequestId(v string) *DescribeSagPortListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagPortListResponseBody) SetTaskStates(v []*DescribeSagPortListResponseBodyTaskStates) *DescribeSagPortListResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagPortListResponseBody) Validate() error {
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

type DescribeSagPortListResponseBodyPorts struct {
	// The MAC address of the port.
	//
	// example:
	//
	// c4:00:ad:a2:f5:****
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The port name.
	//
	// example:
	//
	// 5
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The port role. Valid values:
	//
	// - **NONE**: The port role is unassigned.
	//
	// - **WAN**: The port is a WAN port. A WAN port is used to connect to the Internet and supports DHCP client, PPPoE, and static IP access methods.
	//
	// - **LAN**: The port is a LAN port. A LAN port is used to connect to local clients or switches and supports DHCP server and static IP access methods.
	//
	// - **ECC**: The port is a dedicated connection port used to connect to a dedicated connection.
	//
	// - **MGT**: The port is a management port.
	//
	// - **CELLULAR**: The port is a 4G port.
	//
	// example:
	//
	// NONE
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// The port status. Valid values:
	//
	// - **Up**: The port is enabled.
	//
	// - **Down**: The port is disabled.
	//
	// - **Unavailable**: The device is offline.
	//
	// example:
	//
	// Down
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSagPortListResponseBodyPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortListResponseBodyPorts) GoString() string {
	return s.String()
}

func (s *DescribeSagPortListResponseBodyPorts) GetMac() *string {
	return s.Mac
}

func (s *DescribeSagPortListResponseBodyPorts) GetPortName() *string {
	return s.PortName
}

func (s *DescribeSagPortListResponseBodyPorts) GetRole() *string {
	return s.Role
}

func (s *DescribeSagPortListResponseBodyPorts) GetStatus() *string {
	return s.Status
}

func (s *DescribeSagPortListResponseBodyPorts) SetMac(v string) *DescribeSagPortListResponseBodyPorts {
	s.Mac = &v
	return s
}

func (s *DescribeSagPortListResponseBodyPorts) SetPortName(v string) *DescribeSagPortListResponseBodyPorts {
	s.PortName = &v
	return s
}

func (s *DescribeSagPortListResponseBodyPorts) SetRole(v string) *DescribeSagPortListResponseBodyPorts {
	s.Role = &v
	return s
}

func (s *DescribeSagPortListResponseBodyPorts) SetStatus(v string) *DescribeSagPortListResponseBodyPorts {
	s.Status = &v
	return s
}

func (s *DescribeSagPortListResponseBodyPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeSagPortListResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586762479000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code. A value of 200 indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. A value of Successful indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the asynchronous task. Valid values:
	//
	// - **Initialized**: The query task is initialized.
	//
	// - **Offline**: The Smart Access Gateway device is offline and the query task is not delivered. The task will be delivered after the device goes online.
	//
	// - **Succeed**: The query task is delivered.
	//
	// - **Processing**: The query task is being delivered.
	//
	// - **VersionNotSupport**: The current version of the Smart Access Gateway device does not support this feature.
	//
	// - **BuildRequestError**: The management and control system does not support this operation.
	//
	// - **HardwareError**: The query task failed to be delivered due to a device error.
	//
	// - **TaskNotExist**: The query task does not exist.
	//
	// - **OfflineNotConfiged**: The Smart Access Gateway device is offline and the query task is not delivered. The task will not be delivered after the device goes online.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSagPortListResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagPortListResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagPortListResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagPortListResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagPortListResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagPortListResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagPortListResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagPortListResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagPortListResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagPortListResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagPortListResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagPortListResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagPortListResponseBodyTaskStates) SetState(v string) *DescribeSagPortListResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagPortListResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
