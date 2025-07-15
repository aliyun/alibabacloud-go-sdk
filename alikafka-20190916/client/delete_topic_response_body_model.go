// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteTopicResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteTopicResponseBody
	GetSuccess() *bool
}

type DeleteTopicResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06084011-E093-46F3-A51F-4B19A8AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTopicResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTopicResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTopicResponseBody) SetCode(v int32) *DeleteTopicResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTopicResponseBody) SetMessage(v string) *DeleteTopicResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTopicResponseBody) SetRequestId(v string) *DeleteTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTopicResponseBody) SetSuccess(v bool) *DeleteTopicResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
