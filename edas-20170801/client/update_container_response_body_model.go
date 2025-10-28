// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContainerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChangeOrderId(v string) *UpdateContainerResponseBody
	GetChangeOrderId() *string
	SetCode(v int32) *UpdateContainerResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateContainerResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContainerResponseBody
	GetRequestId() *string
}

type UpdateContainerResponseBody struct {
	// The ID of the change process.
	//
	// example:
	//
	// eb1b9862-****-476f-9e78-d6aa0842835a
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// a5281053-08e4-47a5-b2ab-5c0323******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateContainerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContainerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContainerResponseBody) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *UpdateContainerResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateContainerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContainerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContainerResponseBody) SetChangeOrderId(v string) *UpdateContainerResponseBody {
	s.ChangeOrderId = &v
	return s
}

func (s *UpdateContainerResponseBody) SetCode(v int32) *UpdateContainerResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContainerResponseBody) SetMessage(v string) *UpdateContainerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContainerResponseBody) SetRequestId(v string) *UpdateContainerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContainerResponseBody) Validate() error {
	return dara.Validate(s)
}
