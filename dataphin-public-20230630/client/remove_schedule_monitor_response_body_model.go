// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveScheduleMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveScheduleMonitorResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveScheduleMonitorResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveScheduleMonitorResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveScheduleMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveScheduleMonitorResponseBody
	GetSuccess() *bool
}

type RemoveScheduleMonitorResponseBody struct {
	// The error code. A value of OK indicates that the request is successful.
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

func (s RemoveScheduleMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveScheduleMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveScheduleMonitorResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveScheduleMonitorResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveScheduleMonitorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveScheduleMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveScheduleMonitorResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveScheduleMonitorResponseBody) SetCode(v string) *RemoveScheduleMonitorResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveScheduleMonitorResponseBody) SetHttpStatusCode(v int32) *RemoveScheduleMonitorResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveScheduleMonitorResponseBody) SetMessage(v string) *RemoveScheduleMonitorResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveScheduleMonitorResponseBody) SetRequestId(v string) *RemoveScheduleMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveScheduleMonitorResponseBody) SetSuccess(v bool) *RemoveScheduleMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveScheduleMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
