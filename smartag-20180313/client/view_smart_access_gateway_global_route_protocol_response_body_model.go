// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewSmartAccessGatewayGlobalRouteProtocolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody
	GetRequestId() *string
	SetRouteProtocol(v string) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody
	GetRouteProtocol() *string
	SetTaskStates(v []*ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody
	GetTaskStates() []*ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates
}

type ViewSmartAccessGatewayGlobalRouteProtocolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D10FFDB2-AF7D-530A-A2AC-EBDC16500399
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routing protocol. Valid values:
	//
	// 	- **STATIC**
	//
	// 	- **OSPF**
	//
	// 	- **BGP**
	//
	// example:
	//
	// STATIC
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
	// The status of the task.
	TaskStates []*ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) GetRouteProtocol() *string {
	return s.RouteProtocol
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) GetTaskStates() []*ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates {
	return s.TaskStates
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) SetRequestId(v string) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) SetRouteProtocol(v string) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody {
	s.RouteProtocol = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) SetTaskStates(v []*ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody {
	s.TaskStates = v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBody) Validate() error {
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

type ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates struct {
	// The timestamp when the task was created. Unit: milliseconds.
	//
	// The value is a UNIX timestamp representing the number of milliseconds that have elapsed since the epoch time January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1586843621000
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

func (s ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) SetCreateTime(v string) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) SetErrorCode(v string) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) SetErrorMessage(v string) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) SetState(v string) *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *ViewSmartAccessGatewayGlobalRouteProtocolResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
