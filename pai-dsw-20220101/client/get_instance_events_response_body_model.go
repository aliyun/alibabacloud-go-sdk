// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceEventsResponseBody
	GetCode() *string
	SetEvents(v []*string) *GetInstanceEventsResponseBody
	GetEvents() []*string
	SetHttpStatusCode(v int32) *GetInstanceEventsResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *GetInstanceEventsResponseBody
	GetInstanceId() *string
	SetMessage(v string) *GetInstanceEventsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceEventsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceEventsResponseBody
	GetSuccess() *bool
}

type GetInstanceEventsResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: an internal error. All errors, except for parameter validation errors, are classified as internal errors.
	//
	// 	- ValidationError: a parameter validation error.
	//
	// example:
	//
	// None
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The events.
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// The HTTP status code. Valid values:
	//
	// 	- 400: One or more parameters are invalid.
	//
	// 	- 404: The instance does not exist.
	//
	// 	- 200: The request is normal.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response message.
	//
	// example:
	//
	// XXX
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceEventsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceEventsResponseBody) GetEvents() []*string {
	return s.Events
}

func (s *GetInstanceEventsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceEventsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceEventsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceEventsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceEventsResponseBody) SetCode(v string) *GetInstanceEventsResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceEventsResponseBody) SetEvents(v []*string) *GetInstanceEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetInstanceEventsResponseBody) SetHttpStatusCode(v int32) *GetInstanceEventsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceEventsResponseBody) SetInstanceId(v string) *GetInstanceEventsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceEventsResponseBody) SetMessage(v string) *GetInstanceEventsResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceEventsResponseBody) SetRequestId(v string) *GetInstanceEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceEventsResponseBody) SetSuccess(v bool) *GetInstanceEventsResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceEventsResponseBody) Validate() error {
	return dara.Validate(s)
}
