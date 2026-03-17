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
	// The list of the port information.
	Ports []*DescribeSagPortListResponseBodyPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// CE6642D4-21EB-4168-9BF9-F217953F9892
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the query task.
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
	// The Mac address of the port.
	//
	// example:
	//
	// c4:00:ad:a2:f5:****
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// The name of the port.
	//
	// example:
	//
	// 5
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// Port role:
	//
	// 	- **NONE**: No role is assigned to the port.
	//
	// 	- **WAN**: The port is used as a WAN port. The WAN port supports a Dynamic Host Configuration Protocol (DHCP) client, PPPoE, or a static IP address to access the Internet.
	//
	// 	- **LAN**: The port is used as a LAN port. The LAN port supports a DHCP server or a static IP address to connect to a local terminal or switch.
	//
	// 	- **ECC**: The port is used as a leased line port to connect to a leased line.
	//
	// 	- **MGT**: The port is used as the management port.
	//
	// example:
	//
	// NONE
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// Port states:
	//
	// 	- **Up**: The port is enabled.
	//
	// 	- **Down**: The port is disabled.
	//
	// 	- **Unavailable**: The SAG device is disconnected from Alibaba Cloud.
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
	// The error code. 200 indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. Successful indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Asynchronous task states:
	//
	// 	- **Initialized**: The query task is initialized.
	//
	// 	- **Offline**: The SAG device is disconnected from Alibaba Cloud and Alibaba Cloud has not assigned the query task to the SAG device. When the SAG device is connected to Alibaba Cloud, Alibaba Cloud continues to assign the query task to the SAG device.
	//
	// 	- **Succeed**: Alibaba Cloud has assigned the query task to the SAG device.
	//
	// 	- **Processing**: Alibaba Cloud is assigning the query task to the SAG device.
	//
	// 	- **VersionNotSupport**: not supported by the current version of the SAG device.
	//
	// 	- **BuildRequestError**: not supported by the control and management center in the cloud.
	//
	// 	- **HardwareError**: Alibaba Cloud failed to assign the query task to the SAG device because the SAG device is faulty.
	//
	// 	- **TaskNotExist**: The query task does not exist.
	//
	// 	- **OfflineNotConfiged**: The SAG device is disconnected from Alibaba Cloud and Alibaba Cloud has not assigned the query task to the SAG device. When the SAG device is connected to Alibaba Cloud, Alibaba Cloud does not assign the query task to the SAG device.
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
