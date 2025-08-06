// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventStreamingBusinessOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateEventStreamingBusinessOptionResponseBody
	GetCode() *string
	SetData(v bool) *UpdateEventStreamingBusinessOptionResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateEventStreamingBusinessOptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEventStreamingBusinessOptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateEventStreamingBusinessOptionResponseBody
	GetSuccess() *bool
}

type UpdateEventStreamingBusinessOptionResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 34AD682D-5B91-5773-8132-AA38C130****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateEventStreamingBusinessOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventStreamingBusinessOptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) SetCode(v string) *UpdateEventStreamingBusinessOptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) SetData(v bool) *UpdateEventStreamingBusinessOptionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) SetMessage(v string) *UpdateEventStreamingBusinessOptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) SetRequestId(v string) *UpdateEventStreamingBusinessOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) SetSuccess(v bool) *UpdateEventStreamingBusinessOptionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateEventStreamingBusinessOptionResponseBody) Validate() error {
	return dara.Validate(s)
}
