// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventBusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEventBusResponseBody
	GetCode() *string
	SetData(v *GetEventBusResponseBodyData) *GetEventBusResponseBody
	GetData() *GetEventBusResponseBodyData
	SetMessage(v string) *GetEventBusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEventBusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetEventBusResponseBody
	GetSuccess() *bool
}

type GetEventBusResponseBody struct {
	// The response code. The value Success indicates that the request was successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetEventBusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// EventBusNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// d5bfc188-4452-4ba7-b73a-a9005e522439
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. If the operation was successful, the value true is returned.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetEventBusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventBusResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventBusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEventBusResponseBody) GetData() *GetEventBusResponseBodyData {
	return s.Data
}

func (s *GetEventBusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEventBusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEventBusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetEventBusResponseBody) SetCode(v string) *GetEventBusResponseBody {
	s.Code = &v
	return s
}

func (s *GetEventBusResponseBody) SetData(v *GetEventBusResponseBodyData) *GetEventBusResponseBody {
	s.Data = v
	return s
}

func (s *GetEventBusResponseBody) SetMessage(v string) *GetEventBusResponseBody {
	s.Message = &v
	return s
}

func (s *GetEventBusResponseBody) SetRequestId(v string) *GetEventBusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventBusResponseBody) SetSuccess(v bool) *GetEventBusResponseBody {
	s.Success = &v
	return s
}

func (s *GetEventBusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEventBusResponseBodyData struct {
	// The timestamp that indicates when the event bus was created.
	//
	// example:
	//
	// 1641781825000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The description of the event bus.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the event bus.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789098****:eventbus/MyEventBus
	EventBusARN *string `json:"EventBusARN,omitempty" xml:"EventBusARN,omitempty"`
	// The name of the event bus.
	//
	// example:
	//
	// MyEventBus
	EventBusName *string `json:"EventBusName,omitempty" xml:"EventBusName,omitempty"`
}

func (s GetEventBusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEventBusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEventBusResponseBodyData) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetEventBusResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetEventBusResponseBodyData) GetEventBusARN() *string {
	return s.EventBusARN
}

func (s *GetEventBusResponseBodyData) GetEventBusName() *string {
	return s.EventBusName
}

func (s *GetEventBusResponseBodyData) SetCreateTimestamp(v int64) *GetEventBusResponseBodyData {
	s.CreateTimestamp = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetDescription(v string) *GetEventBusResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetEventBusARN(v string) *GetEventBusResponseBodyData {
	s.EventBusARN = &v
	return s
}

func (s *GetEventBusResponseBodyData) SetEventBusName(v string) *GetEventBusResponseBodyData {
	s.EventBusName = &v
	return s
}

func (s *GetEventBusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
