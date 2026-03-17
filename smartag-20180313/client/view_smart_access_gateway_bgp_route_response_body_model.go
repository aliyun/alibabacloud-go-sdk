// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayBgpRouteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHoldTime(v int32) *ViewSmartAccessGatewayBgpRouteResponseBody
	GetHoldTime() *int32
	SetKeepAlive(v int32) *ViewSmartAccessGatewayBgpRouteResponseBody
	GetKeepAlive() *int32
	SetLocalAs(v int32) *ViewSmartAccessGatewayBgpRouteResponseBody
	GetLocalAs() *int32
	SetRequestId(v string) *ViewSmartAccessGatewayBgpRouteResponseBody
	GetRequestId() *string
	SetRouterId(v string) *ViewSmartAccessGatewayBgpRouteResponseBody
	GetRouterId() *string
	SetTaskStates(v []*ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) *ViewSmartAccessGatewayBgpRouteResponseBody
	GetTaskStates() []*ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates
}

type ViewSmartAccessGatewayBgpRouteResponseBody struct {
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
	// The autonomous system number (ASN) to which the SAG device belongs.
	//
	// example:
	//
	// 12****
	LocalAs *int32 `json:"LocalAs,omitempty" xml:"LocalAs,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F1FEABC0-F7B7-53EA-83EE-AA470ABACE60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the BGP router.
	//
	// example:
	//
	// 192.XX.XX.1
	RouterId *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	// The status of the task.
	TaskStates []*ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s ViewSmartAccessGatewayBgpRouteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayBgpRouteResponseBody) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) GetHoldTime() *int32 {
	return s.HoldTime
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) GetKeepAlive() *int32 {
	return s.KeepAlive
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) GetLocalAs() *int32 {
	return s.LocalAs
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) GetRouterId() *string {
	return s.RouterId
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) GetTaskStates() []*ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates {
	return s.TaskStates
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) SetHoldTime(v int32) *ViewSmartAccessGatewayBgpRouteResponseBody {
	s.HoldTime = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) SetKeepAlive(v int32) *ViewSmartAccessGatewayBgpRouteResponseBody {
	s.KeepAlive = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) SetLocalAs(v int32) *ViewSmartAccessGatewayBgpRouteResponseBody {
	s.LocalAs = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) SetRequestId(v string) *ViewSmartAccessGatewayBgpRouteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) SetRouterId(v string) *ViewSmartAccessGatewayBgpRouteResponseBody {
	s.RouterId = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) SetTaskStates(v []*ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) *ViewSmartAccessGatewayBgpRouteResponseBody {
	s.TaskStates = v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBody) Validate() error {
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

type ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates struct {
	// The timestamp when the task was created. Unit: milliseconds.
	//
	// The value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1586855592000
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

func (s ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) SetCreateTime(v string) *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) SetErrorCode(v string) *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) SetErrorMessage(v string) *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) SetState(v string) *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *ViewSmartAccessGatewayBgpRouteResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
