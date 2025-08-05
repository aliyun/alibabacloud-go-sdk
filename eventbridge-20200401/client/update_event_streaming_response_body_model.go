// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStreamingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateEventStreamingResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateEventStreamingResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEventStreamingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEventStreamingResponseBody
	GetSuccess() *bool
}

type UpdateEventStreamingResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// 0FDD73AA-7A2D-5BD4-B4C0-88AFEBF5F0F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventStreamingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEventStreamingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEventStreamingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventStreamingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEventStreamingResponseBody) SetCode(v string) *UpdateEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventStreamingResponseBody) SetMessage(v string) *UpdateEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEventStreamingResponseBody) SetRequestId(v string) *UpdateEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventStreamingResponseBody) SetSuccess(v bool) *UpdateEventStreamingResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEventStreamingResponseBody) Validate() error {
	return dara.Validate(s)
}
