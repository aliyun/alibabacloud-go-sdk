// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateScheduleMonitorResponseBody
	GetCode() *string
	SetData(v *CreateScheduleMonitorResponseBodyData) *CreateScheduleMonitorResponseBody
	GetData() *CreateScheduleMonitorResponseBodyData
	SetHttpStatusCode(v int32) *CreateScheduleMonitorResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateScheduleMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateScheduleMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateScheduleMonitorResponseBody
	GetSuccess() *bool
}

type CreateScheduleMonitorResponseBody struct {
	// The error code. OK indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The monitoring creation result.
	Data *CreateScheduleMonitorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The ID of the request.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateScheduleMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateScheduleMonitorResponseBody) GetData() *CreateScheduleMonitorResponseBodyData {
	return s.Data
}

func (s *CreateScheduleMonitorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateScheduleMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateScheduleMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduleMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScheduleMonitorResponseBody) SetCode(v string) *CreateScheduleMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScheduleMonitorResponseBody) SetData(v *CreateScheduleMonitorResponseBodyData) *CreateScheduleMonitorResponseBody {
	s.Data = v
	return s
}

func (s *CreateScheduleMonitorResponseBody) SetHttpStatusCode(v int32) *CreateScheduleMonitorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateScheduleMonitorResponseBody) SetMessage(v string) *CreateScheduleMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateScheduleMonitorResponseBody) SetRequestId(v string) *CreateScheduleMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleMonitorResponseBody) SetSuccess(v bool) *CreateScheduleMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScheduleMonitorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateScheduleMonitorResponseBodyData struct {
	// The ID of the created monitored object.
	//
	// example:
	//
	// 6912345678901234
	MonitoredItemId *int64 `json:"MonitoredItemId,omitempty" xml:"MonitoredItemId,omitempty"`
}

func (s CreateScheduleMonitorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleMonitorResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateScheduleMonitorResponseBodyData) GetMonitoredItemId() *int64 {
	return s.MonitoredItemId
}

func (s *CreateScheduleMonitorResponseBodyData) SetMonitoredItemId(v int64) *CreateScheduleMonitorResponseBodyData {
	s.MonitoredItemId = &v
	return s
}

func (s *CreateScheduleMonitorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
