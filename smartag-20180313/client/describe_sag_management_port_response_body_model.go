// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagManagementPortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGateway(v string) *DescribeSagManagementPortResponseBody
	GetGateway() *string
	SetIP(v string) *DescribeSagManagementPortResponseBody
	GetIP() *string
	SetMask(v string) *DescribeSagManagementPortResponseBody
	GetMask() *string
	SetRequestId(v string) *DescribeSagManagementPortResponseBody
	GetRequestId() *string
	SetTaskStates(v []*DescribeSagManagementPortResponseBodyTaskStates) *DescribeSagManagementPortResponseBody
	GetTaskStates() []*DescribeSagManagementPortResponseBodyTaskStates
}

type DescribeSagManagementPortResponseBody struct {
	// The gateway IP address of the management port.
	//
	// example:
	//
	// 192.XX.XX.254
	Gateway *string `json:"Gateway,omitempty" xml:"Gateway,omitempty"`
	// The IP address of the management port.
	//
	// example:
	//
	// 192.XX.XX.10
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// The subnet mask of the management port.
	//
	// example:
	//
	// 255.255.255.0
	Mask *string `json:"Mask,omitempty" xml:"Mask,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3616AAA9-3A6F-4604-98AF-86753AB7F040
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the query task.
	TaskStates []*DescribeSagManagementPortResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagManagementPortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagManagementPortResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagManagementPortResponseBody) GetGateway() *string {
	return s.Gateway
}

func (s *DescribeSagManagementPortResponseBody) GetIP() *string {
	return s.IP
}

func (s *DescribeSagManagementPortResponseBody) GetMask() *string {
	return s.Mask
}

func (s *DescribeSagManagementPortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagManagementPortResponseBody) GetTaskStates() []*DescribeSagManagementPortResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagManagementPortResponseBody) SetGateway(v string) *DescribeSagManagementPortResponseBody {
	s.Gateway = &v
	return s
}

func (s *DescribeSagManagementPortResponseBody) SetIP(v string) *DescribeSagManagementPortResponseBody {
	s.IP = &v
	return s
}

func (s *DescribeSagManagementPortResponseBody) SetMask(v string) *DescribeSagManagementPortResponseBody {
	s.Mask = &v
	return s
}

func (s *DescribeSagManagementPortResponseBody) SetRequestId(v string) *DescribeSagManagementPortResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagManagementPortResponseBody) SetTaskStates(v []*DescribeSagManagementPortResponseBodyTaskStates) *DescribeSagManagementPortResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagManagementPortResponseBody) Validate() error {
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

type DescribeSagManagementPortResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586759657000
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
	// - **Succeed**: The query task is delivered successfully.
	//
	// - **Processing**: The query task is being delivered.
	//
	// - **VersionNotSupport**: The current version of the Smart Access Gateway device does not support the feature.
	//
	// - **BuildRequestError**: The control plane does not support the operation.
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

func (s DescribeSagManagementPortResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagManagementPortResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagManagementPortResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagManagementPortResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagManagementPortResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagManagementPortResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagManagementPortResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagManagementPortResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagManagementPortResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagManagementPortResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagManagementPortResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagManagementPortResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagManagementPortResponseBodyTaskStates) SetState(v string) *DescribeSagManagementPortResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagManagementPortResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
