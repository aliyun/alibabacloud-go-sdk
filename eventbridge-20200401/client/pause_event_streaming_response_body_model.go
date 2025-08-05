// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseEventStreamingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PauseEventStreamingResponseBody
	GetCode() *string
	SetMessage(v string) *PauseEventStreamingResponseBody
	GetMessage() *string
	SetRequestId(v string) *PauseEventStreamingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PauseEventStreamingResponseBody
	GetSuccess() *bool
}

type PauseEventStreamingResponseBody struct {
	// The response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// The event streaming [xxxx] not existed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 72ce027c-546a-4231-9cf6-ec58766027f9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PauseEventStreamingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *PauseEventStreamingResponseBody) GetCode() *string {
	return s.Code
}

func (s *PauseEventStreamingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PauseEventStreamingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseEventStreamingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PauseEventStreamingResponseBody) SetCode(v string) *PauseEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *PauseEventStreamingResponseBody) SetMessage(v string) *PauseEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *PauseEventStreamingResponseBody) SetRequestId(v string) *PauseEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseEventStreamingResponseBody) SetSuccess(v bool) *PauseEventStreamingResponseBody {
	s.Success = &v
	return s
}

func (s *PauseEventStreamingResponseBody) Validate() error {
	return dara.Validate(s)
}
