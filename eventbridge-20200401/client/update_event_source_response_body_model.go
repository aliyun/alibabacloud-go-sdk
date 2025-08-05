// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateEventSourceResponseBody
	GetCode() *string
	SetData(v bool) *UpdateEventSourceResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateEventSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEventSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEventSourceResponseBody
	GetSuccess() *bool
}

type UpdateEventSourceResponseBody struct {
	// The returned response code. Valid values:
	//
	// 	- Success: The request is successful.
	//
	// 	- Other codes: The request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// Remote error. requestId: [xxxx-9D10-65DFDFA3A75D], error code: [EventSourceNotExist], message: [The event source in request is not exist! ]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// c057d379-ea65-41ec-a8a8-90627a968204
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventSourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEventSourceResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateEventSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEventSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEventSourceResponseBody) SetCode(v string) *UpdateEventSourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventSourceResponseBody) SetData(v bool) *UpdateEventSourceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateEventSourceResponseBody) SetMessage(v string) *UpdateEventSourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEventSourceResponseBody) SetRequestId(v string) *UpdateEventSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventSourceResponseBody) SetSuccess(v bool) *UpdateEventSourceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEventSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
