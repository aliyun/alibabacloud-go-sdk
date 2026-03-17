// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagExpressConnectInterfaceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInterfaces(v []*DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) *DescribeSagExpressConnectInterfaceListResponseBody
	GetInterfaces() []*DescribeSagExpressConnectInterfaceListResponseBodyInterfaces
	SetRequestId(v string) *DescribeSagExpressConnectInterfaceListResponseBody
	GetRequestId() *string
	SetTaskStates(v []*DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) *DescribeSagExpressConnectInterfaceListResponseBody
	GetTaskStates() []*DescribeSagExpressConnectInterfaceListResponseBodyTaskStates
}

type DescribeSagExpressConnectInterfaceListResponseBody struct {
	// The information about the port.
	Interfaces []*DescribeSagExpressConnectInterfaceListResponseBodyInterfaces `json:"Interfaces,omitempty" xml:"Interfaces,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 2718F7A6-EA67-41EF-BA39-E9F4A0F5D306
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the query task.
	TaskStates []*DescribeSagExpressConnectInterfaceListResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagExpressConnectInterfaceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagExpressConnectInterfaceListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagExpressConnectInterfaceListResponseBody) GetInterfaces() []*DescribeSagExpressConnectInterfaceListResponseBodyInterfaces {
	return s.Interfaces
}

func (s *DescribeSagExpressConnectInterfaceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagExpressConnectInterfaceListResponseBody) GetTaskStates() []*DescribeSagExpressConnectInterfaceListResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagExpressConnectInterfaceListResponseBody) SetInterfaces(v []*DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) *DescribeSagExpressConnectInterfaceListResponseBody {
	s.Interfaces = v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBody) SetRequestId(v string) *DescribeSagExpressConnectInterfaceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBody) SetTaskStates(v []*DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) *DescribeSagExpressConnectInterfaceListResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBody) Validate() error {
	if s.Interfaces != nil {
		for _, item := range s.Interfaces {
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

type DescribeSagExpressConnectInterfaceListResponseBodyInterfaces struct {
	// The IP address.
	//
	// example:
	//
	// 192.XX.XX.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The subnet mask of the IP address of the port.
	//
	// example:
	//
	// 255.255.255.0
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The VLAN ID.
	//
	// example:
	//
	// 2
	Vlan *string `json:"Vlan,omitempty" xml:"Vlan,omitempty"`
}

func (s DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) GoString() string {
	return s.String()
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) GetIP() *string {
	return s.IP
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) GetMask() *string {
	return s.Mask
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) GetVlan() *string {
	return s.Vlan
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) SetIP(v string) *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces {
	s.IP = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) SetMask(v string) *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces {
	s.Mask = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) SetVlan(v string) *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces {
	s.Vlan = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyInterfaces) Validate() error {
	return dara.Validate(s)
}

type DescribeSagExpressConnectInterfaceListResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586835287000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code returned. A value of 200 indicates that the query task is successful.
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
	// The state of the asynchronous query task. Valid values:
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

func (s DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) SetState(v string) *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
