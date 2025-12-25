// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateCalendarResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateCalendarResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateCalendarResponseBody
	GetSuccess() *bool
}

type UpdateCalendarResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D0DE9C33-992A-580B-89C4-B609A292748D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCalendarResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateCalendarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateCalendarResponseBody) SetCode(v int32) *UpdateCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCalendarResponseBody) SetMessage(v string) *UpdateCalendarResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCalendarResponseBody) SetRequestId(v string) *UpdateCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCalendarResponseBody) SetSuccess(v bool) *UpdateCalendarResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
