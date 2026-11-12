// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagGlobalRouteProtocolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSagGlobalRouteProtocolResponseBody
	GetRequestId() *string
	SetRouteProtocol(v string) *DescribeSagGlobalRouteProtocolResponseBody
	GetRouteProtocol() *string
	SetTaskStates(v []*DescribeSagGlobalRouteProtocolResponseBodyTaskStates) *DescribeSagGlobalRouteProtocolResponseBody
	GetTaskStates() []*DescribeSagGlobalRouteProtocolResponseBodyTaskStates
}

type DescribeSagGlobalRouteProtocolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3712F0B2-721E-4FBF-BBEF-888E3BFE0A20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The routable protocol. Valid values:
	//
	// - **STATIC**: static routable protocol
	//
	// - **OSPF**: OSPF dynamic routable protocol
	//
	// - **BGP**: BGP dynamic routable protocol
	//
	// example:
	//
	// STATIC
	RouteProtocol *string `json:"RouteProtocol,omitempty" xml:"RouteProtocol,omitempty"`
	// The status of the query task.
	TaskStates []*DescribeSagGlobalRouteProtocolResponseBodyTaskStates `json:"TaskStates,omitempty" xml:"TaskStates,omitempty" type:"Repeated"`
}

func (s DescribeSagGlobalRouteProtocolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagGlobalRouteProtocolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagGlobalRouteProtocolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagGlobalRouteProtocolResponseBody) GetRouteProtocol() *string {
	return s.RouteProtocol
}

func (s *DescribeSagGlobalRouteProtocolResponseBody) GetTaskStates() []*DescribeSagGlobalRouteProtocolResponseBodyTaskStates {
	return s.TaskStates
}

func (s *DescribeSagGlobalRouteProtocolResponseBody) SetRequestId(v string) *DescribeSagGlobalRouteProtocolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponseBody) SetRouteProtocol(v string) *DescribeSagGlobalRouteProtocolResponseBody {
	s.RouteProtocol = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponseBody) SetTaskStates(v []*DescribeSagGlobalRouteProtocolResponseBodyTaskStates) *DescribeSagGlobalRouteProtocolResponseBody {
	s.TaskStates = v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponseBody) Validate() error {
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

type DescribeSagGlobalRouteProtocolResponseBodyTaskStates struct {
	// The time when the query task was created.
	//
	// example:
	//
	// 1586855592000
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
	// - **Offline**: The Smart Access Gateway device is offline and the query task is not delivered. The task will be delivered after the device comes online.
	//
	// - **Succeed**: The query task is delivered.
	//
	// - **Processing**: The query task is being delivered.
	//
	// - **VersionNotSupport**: The current version of the Smart Access Gateway device does not support the feature.
	//
	// - **BuildRequestError**: The management and control system does not support the operation.
	//
	// - **HardwareError**: The query task delivery failed due to a device error.
	//
	// - **TaskNotExist**: The query task does not exist.
	//
	// - **OfflineNotConfiged**: The Smart Access Gateway device is offline and the query task is not delivered. The task will not be delivered after the device comes online.
	//
	// example:
	//
	// Succeed
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeSagGlobalRouteProtocolResponseBodyTaskStates) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagGlobalRouteProtocolResponseBodyTaskStates) GoString() string {
	return s.String()
}

func (s *DescribeSagGlobalRouteProtocolResponseBodyTaskStates) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSagGlobalRouteProtocolResponseBodyTaskStates) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSagGlobalRouteProtocolResponseBodyTaskStates) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSagGlobalRouteProtocolResponseBodyTaskStates) GetState() *string {
	return s.State
}

func (s *DescribeSagGlobalRouteProtocolResponseBodyTaskStates) SetCreateTime(v string) *DescribeSagGlobalRouteProtocolResponseBodyTaskStates {
	s.CreateTime = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponseBodyTaskStates) SetErrorCode(v string) *DescribeSagGlobalRouteProtocolResponseBodyTaskStates {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponseBodyTaskStates) SetErrorMessage(v string) *DescribeSagGlobalRouteProtocolResponseBodyTaskStates {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponseBodyTaskStates) SetState(v string) *DescribeSagGlobalRouteProtocolResponseBodyTaskStates {
	s.State = &v
	return s
}

func (s *DescribeSagGlobalRouteProtocolResponseBodyTaskStates) Validate() error {
	return dara.Validate(s)
}
