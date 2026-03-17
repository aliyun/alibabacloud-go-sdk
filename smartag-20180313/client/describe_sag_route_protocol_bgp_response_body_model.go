// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRouteProtocolBgpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHoldTime(v int32) *DescribeSagRouteProtocolBgpResponseBody
	GetHoldTime() *int32
	SetKeepAlive(v int32) *DescribeSagRouteProtocolBgpResponseBody
	GetKeepAlive() *int32
	SetLocalAs(v int32) *DescribeSagRouteProtocolBgpResponseBody
	GetLocalAs() *int32
	SetRequestId(v string) *DescribeSagRouteProtocolBgpResponseBody
	GetRequestId() *string
	SetRouterId(v string) *DescribeSagRouteProtocolBgpResponseBody
	GetRouterId() *string
	SetTaskStates(v []*DescribeSagRouteProtocolBgpResponseBodyTaskStates) *DescribeSagRouteProtocolBgpResponseBody
	GetTaskStates() []*DescribeSagRouteProtocolBgpResponseBodyTaskStates
}

type DescribeSagRouteProtocolBgpResponseBody struct {
	// The hold time. Unit: seconds.
	//
	// example:
	//
	// 9
	HoldTime *int32 `json:"HoldTime,omitempty" xml:"HoldTime,omitempty"`
	// The time interval at which keep-alive packets are sent. Unit: seconds.
	//
	// example:
	//
	// 3
	KeepAlive *int32 `json:"KeepAlive,omitempty" xml:"KeepAlive,omitempty"`
	// The autonomous system (AS) number to which the SAG device belongs.
	//
	// example:
	//
	// 65536
	LocalAs *int32 `json:"LocalAs,omitempty" xml:"LocalAs,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2F39E4FE-B45C-47FF-9921-95780486F52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the BGP router.
	//
	// example:
	//
	// 192.XX.XX.1
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The details about the query task.
	TaskStates []*DescribeSagRouteProtocolBgpResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagRouteProtocolBgpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteProtocolBgpResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteProtocolBgpResponseBody) GetHoldTime() *int32 {
	return s.HoldTime
}

func (s *DescribeSagRouteProtocolBgpResponseBody) GetKeepAlive() *int32 {
	return s.KeepAlive
}

func (s *DescribeSagRouteProtocolBgpResponseBody) GetLocalAs() *int32 {
	return s.LocalAs
}

func (s *DescribeSagRouteProtocolBgpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagRouteProtocolBgpResponseBody) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeSagRouteProtocolBgpResponseBody) GetTaskStates() []*DescribeSagRouteProtocolBgpResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagRouteProtocolBgpResponseBody) SetHoldTime(v int32) *DescribeSagRouteProtocolBgpResponseBody {
	s.HoldTime = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBody) SetKeepAlive(v int32) *DescribeSagRouteProtocolBgpResponseBody {
	s.KeepAlive = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBody) SetLocalAs(v int32) *DescribeSagRouteProtocolBgpResponseBody {
	s.LocalAs = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBody) SetRequestId(v string) *DescribeSagRouteProtocolBgpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBody) SetRouterId(v string) *DescribeSagRouteProtocolBgpResponseBody {
	s.RouterId = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBody) SetTaskStates(v []*DescribeSagRouteProtocolBgpResponseBodyTaskStates) *DescribeSagRouteProtocolBgpResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBody) Validate() error {
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

type DescribeSagRouteProtocolBgpResponseBodyTaskStates struct {
	// The timestamp when the task was created.
	//
	// example:
	//
	// 1586765938000
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
	// The state of the query task. Valid values:
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

func (s DescribeSagRouteProtocolBgpResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRouteProtocolBgpResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagRouteProtocolBgpResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagRouteProtocolBgpResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagRouteProtocolBgpResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagRouteProtocolBgpResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagRouteProtocolBgpResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagRouteProtocolBgpResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagRouteProtocolBgpResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagRouteProtocolBgpResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBodyTaskStates) SetState(v string) *DescribeSagRouteProtocolBgpResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagRouteProtocolBgpResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
