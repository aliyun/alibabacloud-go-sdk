// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateScheduleMonitorResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateScheduleMonitorResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateScheduleMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateScheduleMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateScheduleMonitorResponseBody
	GetSuccess() *bool
}

type UpdateScheduleMonitorResponseBody struct {
	// The error code. A value of OK indicates that the request is normal.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateScheduleMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduleMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateScheduleMonitorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateScheduleMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateScheduleMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScheduleMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateScheduleMonitorResponseBody) SetCode(v string) *UpdateScheduleMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateScheduleMonitorResponseBody) SetHttpStatusCode(v int32) *UpdateScheduleMonitorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateScheduleMonitorResponseBody) SetMessage(v string) *UpdateScheduleMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateScheduleMonitorResponseBody) SetRequestId(v string) *UpdateScheduleMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduleMonitorResponseBody) SetSuccess(v bool) *UpdateScheduleMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateScheduleMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
