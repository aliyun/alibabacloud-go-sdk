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
	// The high-availability (HA) pattern. Valid values:
	//
	// - **NONE**: The HA feature is not enabled.
	//
	// - **STATIC**: static pattern.
	//
	// - **DYNAMIC**: dynamic schema.
	//
	// example:
	//
	// NONE
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The list of port information.
	Ports []*DescribeSagHaResponseBodyPorts `json:"Ports,omitempty" xml:"Ports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 273D62FD-E346-4959-AA18-D79B9276FEFB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of query task information.
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
	// The port name.
	//
	// example:
	//
	// 5
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The virtual IP address of the Smart Access Gateway.
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
	// The error code. 200 indicates that the query task succeeded.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. "Successful" indicates that the query task succeeded.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the asynchronous task. Valid values:
	//
	// - **Initialized**: The query task is initialized.
	//
	// - **Offline**: The Smart Access Gateway device is offline and the query task is not delivered. The task will be delivered after the device comes online.
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
	// - **OfflineNotConfiged**: The Smart Access Gateway device is offline and the query task is not delivered. The task will not be delivered even after the device comes online.
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
