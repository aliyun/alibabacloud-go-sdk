// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEventStreamingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartEventStreamingResponseBody
	GetCode() *string
	SetMessage(v string) *StartEventStreamingResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartEventStreamingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartEventStreamingResponseBody
	GetSuccess() *bool
}

type StartEventStreamingResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful. Other values indicate that the request failed. For more information about error codes, see Error codes.
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
	// 8CEAD24D-328D-5539-9D30-FD2D33204FBB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. The value true indicates that the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartEventStreamingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *StartEventStreamingResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartEventStreamingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartEventStreamingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartEventStreamingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartEventStreamingResponseBody) SetCode(v string) *StartEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *StartEventStreamingResponseBody) SetMessage(v string) *StartEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *StartEventStreamingResponseBody) SetRequestId(v string) *StartEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartEventStreamingResponseBody) SetSuccess(v bool) *StartEventStreamingResponseBody {
	s.Success = &v
	return s
}

func (s *StartEventStreamingResponseBody) Validate() error {
	return dara.Validate(s)
}
