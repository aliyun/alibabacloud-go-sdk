// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagHaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *DescribeSagHaResponseBody
	GetMode() *string
	SetPorts(v []*DescribeSagHaResponseBodyPorts) *DescribeSagHaResponseBody
	GetPorts() []*DescribeSagHaResponseBodyPorts
	SetRequestId(v string) *DescribeSagHaResponseBody
	GetRequestId() *string
	SetTaskStates(v []*DescribeSagHaResponseBodyTaskStates) *DescribeSagHaResponseBody
	GetTaskStates() []*DescribeSagHaResponseBodyTaskStates
}

type DescribeSagHaResponseBody struct {
	// The HA mode. Valid values:
	//
	// 	- **NONE**: HA is disabled.
	//
	// 	- **STATIC**: static HA is enabled.
	//
	// 	- **DYNAMIC**: dynamic HA is enabled.
	//
	// example:
	//
	// NONE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The information about the port.
	Ports []*DescribeSagHaResponseBodyPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 273D62FD-E346-4959-AA18-D79B9276FEFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the query task.
	TaskStates []*DescribeSagHaResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagHaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagHaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagHaResponseBody) GetMode() *string {
	return s.Mode
}

func (s *DescribeSagHaResponseBody) GetPorts() []*DescribeSagHaResponseBodyPorts {
	return s.Ports
}

func (s *DescribeSagHaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagHaResponseBody) GetTaskStates() []*DescribeSagHaResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagHaResponseBody) SetMode(v string) *DescribeSagHaResponseBody {
	s.Mode = &v
	return s
}

func (s *DescribeSagHaResponseBody) SetPorts(v []*DescribeSagHaResponseBodyPorts) *DescribeSagHaResponseBody {
	s.Ports = v
	return s
}

func (s *DescribeSagHaResponseBody) SetRequestId(v string) *DescribeSagHaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagHaResponseBody) SetTaskStates(v []*DescribeSagHaResponseBodyTaskStates) *DescribeSagHaResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagHaResponseBody) Validate() error {
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

type DescribeSagHaResponseBodyPorts struct {
	// The name of the port.
	//
	// example:
	//
	// 5
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The virtual IP address of the SAG device.
	//
	// example:
	//
	// 192.XX.XX.1
	VirtualIp *string `json:"VirtualIp,omitempty" xml:"VirtualIp,omitempty"`
}

func (s DescribeSagHaResponseBodyPorts) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagHaResponseBodyPorts) GoString() string {
	return s.String()
}

func (s *DescribeSagHaResponseBodyPorts) GetPortName() *string {
	return s.PortName
}

func (s *DescribeSagHaResponseBodyPorts) GetVirtualIp() *string {
	return s.VirtualIp
}

func (s *DescribeSagHaResponseBodyPorts) SetPortName(v string) *DescribeSagHaResponseBodyPorts {
	s.PortName = &v
	return s
}

func (s *DescribeSagHaResponseBodyPorts) SetVirtualIp(v string) *DescribeSagHaResponseBodyPorts {
	s.VirtualIp = &v
	return s
}

func (s *DescribeSagHaResponseBodyPorts) Validate() error {
	return dara.Validate(s)
}

type DescribeSagHaResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586836343000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code returned for a query task. The 200 error code indicates that the query task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned for a query task. The Successful error message indicates that the query task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the query task. Valid values:
	//
	// 	- **Initialized**: The query task has been initialized.
	//
	// 	- **Offline**: The query task is not dispatched because the SAG device is disconnected from Alibaba Cloud. The task will be dispatched after the SAG device is connected to Alibaba Cloud.
	//
	// 	- **Succeed**: The query task has been dispatched.
	//
	// 	- **Processing**: The query task is being dispatched.
	//
	// 	- **VersionNotSupport**: The current version of the SAG device does not support query tasks.
	//
	// 	- **BuildRequestError**: The SAG control system does not support query tasks.
	//
	// 	- **HardwareError**: The query task failed to be dispatched due to device errors.
	//
	// 	- **TaskNotExist**: The query task does not exist.
	//
	// 	- **OfflineNotConfiged**: The query task is not dispatched because the SAG device is disconnected from Alibaba Cloud. The task will not be dispatched after the device is connected to Alibaba Cloud.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSagHaResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagHaResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagHaResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagHaResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagHaResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagHaResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagHaResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagHaResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagHaResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagHaResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagHaResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagHaResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagHaResponseBodyTaskStates) SetState(v string) *DescribeSagHaResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagHaResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
