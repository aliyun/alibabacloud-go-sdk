// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayWanSnatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ViewSmartAccessGatewayWanSnatResponseBody
	GetRequestId() *string
	SetSnat(v string) *ViewSmartAccessGatewayWanSnatResponseBody
	GetSnat() *string
	SetTaskStates(v []*ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) *ViewSmartAccessGatewayWanSnatResponseBody
	GetTaskStates() []*ViewSmartAccessGatewayWanSnatResponseBodyTaskStates
}

type ViewSmartAccessGatewayWanSnatResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0CFC6919-8F3D-524F-A7A6-E5FADCD36A20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether to enable SNAT. Valid values: Valid values:
	//
	// 	- **1**: enables SNAT
	//
	// 	- **0**: disables SNAT
	//
	// example:
	//
	// 1
	Snat *string `json:"Snat,omitempty" xml:"Snat,omitempty"`
	// The status of the task.
	TaskStates []*ViewSmartAccessGatewayWanSnatResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s ViewSmartAccessGatewayWanSnatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayWanSnatResponseBody) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayWanSnatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewSmartAccessGatewayWanSnatResponseBody) GetSnat() *string {
	return s.Snat
}

func (s *ViewSmartAccessGatewayWanSnatResponseBody) GetTaskStates() []*ViewSmartAccessGatewayWanSnatResponseBodyTaskStates {
	return s.TaskStates
}

func (s *ViewSmartAccessGatewayWanSnatResponseBody) SetRequestId(v string) *ViewSmartAccessGatewayWanSnatResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponseBody) SetSnat(v string) *ViewSmartAccessGatewayWanSnatResponseBody {
	s.Snat = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponseBody) SetTaskStates(v []*ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) *ViewSmartAccessGatewayWanSnatResponseBody {
	s.TaskStates = v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponseBody) Validate() error {
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

type ViewSmartAccessGatewayWanSnatResponseBodyTaskStates struct {
	// The timestamp when the task was created. Unit: milliseconds.
	//
	// The value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1586852928000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code. A value of 200 indicates that the task is successful.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. A value of Successful indicates that the task is successful.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the asynchronous task. Valid values:
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

func (s ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) SetCreateTime(v string) *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) SetErrorCode(v string) *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) SetErrorMessage(v string) *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) SetState(v string) *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *ViewSmartAccessGatewayWanSnatResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
