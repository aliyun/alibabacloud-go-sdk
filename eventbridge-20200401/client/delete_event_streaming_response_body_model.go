// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventStreamingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEventStreamingResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteEventStreamingResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEventStreamingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEventStreamingResponseBody
	GetSuccess() *bool
}

type DeleteEventStreamingResponseBody struct {
	// The returned response code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// The event streaming [xxxx] not existed!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 499A9ACF-70CD-5D43-87F3-1B60529EE446
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventStreamingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventStreamingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventStreamingResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEventStreamingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEventStreamingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventStreamingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventStreamingResponseBody) SetCode(v string) *DeleteEventStreamingResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventStreamingResponseBody) SetMessage(v string) *DeleteEventStreamingResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventStreamingResponseBody) SetRequestId(v string) *DeleteEventStreamingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventStreamingResponseBody) SetSuccess(v bool) *DeleteEventStreamingResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventStreamingResponseBody) Validate() error {
	return dara.Validate(s)
}
