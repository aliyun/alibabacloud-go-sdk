// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEventSourceResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteEventSourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEventSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEventSourceResponseBody
	GetSuccess() *bool
}

type DeleteEventSourceResponseBody struct {
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
	// Remote error. requestId: [78B66E68-E778-1F33-84BD-xxxx], error code: [EventSourceNotExist], message: [The event source in request is not exist! ]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5f80e9b3-98d5-4f51-8412-c758818a03e4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventSourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEventSourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEventSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventSourceResponseBody) SetCode(v string) *DeleteEventSourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventSourceResponseBody) SetMessage(v string) *DeleteEventSourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventSourceResponseBody) SetRequestId(v string) *DeleteEventSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventSourceResponseBody) SetSuccess(v bool) *DeleteEventSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
