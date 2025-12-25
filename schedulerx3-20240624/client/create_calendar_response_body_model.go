// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateCalendarResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateCalendarResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCalendarResponseBody
	GetSuccess() *bool
}

type CreateCalendarResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// CalendarName is already existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B880122A-B0E4-52E8-8F54-87DB7779EB74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCalendarResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateCalendarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCalendarResponseBody) SetCode(v int32) *CreateCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCalendarResponseBody) SetMessage(v string) *CreateCalendarResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCalendarResponseBody) SetRequestId(v string) *CreateCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCalendarResponseBody) SetSuccess(v bool) *CreateCalendarResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
