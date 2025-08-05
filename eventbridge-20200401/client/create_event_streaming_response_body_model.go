// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventStreamingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEventStreamingResponseBody
	GetCode() *string
	SetData(v *CreateEventStreamingResponseBodyData) *CreateEventStreamingResponseBody
	GetData() *CreateEventStreamingResponseBodyData
	SetMessage(v string) *CreateEventStreamingResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEventStreamingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateEventStreamingResponseBody
	GetSuccess() *bool
}

type CreateEventStreamingResponseBody struct {
	// The response code. Valid values:
	//
	// 	- Success: The request is successful.
	//
	// 	- Other codes: The request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateEventStreamingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// The name [xxxx] of event streaming in request is already exist!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B896B484-F16D-59DE-9E23-DD0E5C361108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateEventStreamingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEventStreamingResponseBody) GetData() *CreateEventStreamingResponseBodyData {
	return s.Data
}

func (s *CreateEventStreamingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEventStreamingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEventStreamingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateEventStreamingResponseBody) SetCode(v string) *CreateEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEventStreamingResponseBody) SetData(v *CreateEventStreamingResponseBodyData) *CreateEventStreamingResponseBody {
	s.Data = v
	return s
}

func (s *CreateEventStreamingResponseBody) SetMessage(v string) *CreateEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEventStreamingResponseBody) SetRequestId(v string) *CreateEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEventStreamingResponseBody) SetSuccess(v bool) *CreateEventStreamingResponseBody {
	s.Success = &v
	return s
}

func (s *CreateEventStreamingResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateEventStreamingResponseBodyData struct {
	// The ARN of the event stream.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:164901546557****:eventstreaming/myeventstreaming
	EventStreamingARN *string `json:"EventStreamingARN,omitempty" xml:"EventStreamingARN,omitempty"`
}

func (s CreateEventStreamingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEventStreamingResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEventStreamingResponseBodyData) GetEventStreamingARN() *string {
	return s.EventStreamingARN
}

func (s *CreateEventStreamingResponseBodyData) SetEventStreamingARN(v string) *CreateEventStreamingResponseBodyData {
	s.EventStreamingARN = &v
	return s
}

func (s *CreateEventStreamingResponseBodyData) Validate() error {
	return dara.Validate(s)
}
