// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTopicResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateTopicResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateTopicResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateTopicResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateTopicResponseBody
	GetSuccess() *bool
}

type CreateTopicResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9C0F207C-77A6-43E5-991C-9D98510A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTopicResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTopicResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTopicResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateTopicResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateTopicResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTopicResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateTopicResponseBody) SetCode(v int32) *CreateTopicResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTopicResponseBody) SetMessage(v string) *CreateTopicResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTopicResponseBody) SetRequestId(v string) *CreateTopicResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTopicResponseBody) SetSuccess(v bool) *CreateTopicResponseBody {
	s.Success = &v
	return s
}

func (s *CreateTopicResponseBody) Validate() error {
	return dara.Validate(s)
}
