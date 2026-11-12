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
	// The ID of the request.
	//
	// example:
	//
	// D10FFDB2-AF7D-530A-A2AC-EBDC16500399
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routing protocol. Valid values:
	//
	// - **STATIC**: Static routing protocol.
	//
	// - **OSPF**: OSPF dynamic routing protocol.
	//
	// - **BGP**: BGP dynamic routing protocol.
	//
	// example:
	//
	// STATIC
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
	// The status of the query task.
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
	// The timestamp when the query task was created. Unit: milliseconds.
	//
	// The timestamp is in the UNIX format, which represents the total time elapsed from 00:00:00 on January 1, 1970 (UTC) to the time when the query task was created.
	//
	// example:
	//
	// 1586843621000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The error code. A value of 200 indicates that the query task succeeded.
	//
	// example:
	//
	// 200
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message. A value of Successful indicates that the query task succeeded.
	//
	// example:
	//
	// Successful
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The status of the asynchronous task. Valid values:
	//
	// - **Initialized**: The query task is initialized.
	//
	// - **Offline**: The SAG device is offline and the query task is not delivered. The task will be delivered Smart Access Gateway device goes online.
	//
	// - **Succeed**: The query task is delivered.
	//
	// - **Processing**: The query task is being delivered.
	//
	// - **VersionNotSupport**: The current version of the SAG device does not support this operation.
	//
	// - **BuildRequestError**: The management and control system does not support this operation.
	//
	// - **HardwareError**: The query task failed to be delivered due to a device error.
	//
	// - **TaskNotExist**: The query task does not exist.
	//
	// - **OfflineNotConfiged**: The SAG device is offline and the query task is not delivered. The task will not be delivered Smart Access Gateway device goes online.
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
