// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDebuggingJMeterSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartDebuggingJMeterSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StartDebuggingJMeterSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartDebuggingJMeterSceneResponseBody
	GetMessage() *string
	SetReportId(v string) *StartDebuggingJMeterSceneResponseBody
	GetReportId() *string
	SetRequestId(v string) *StartDebuggingJMeterSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartDebuggingJMeterSceneResponseBody
	GetSuccess() *bool
}

type StartDebuggingJMeterSceneResponseBody struct {
	// The system status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the debugging report.
	//
	// example:
	//
	// MH0SU1I
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartDebuggingJMeterSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDebuggingJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartDebuggingJMeterSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartDebuggingJMeterSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartDebuggingJMeterSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartDebuggingJMeterSceneResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *StartDebuggingJMeterSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDebuggingJMeterSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartDebuggingJMeterSceneResponseBody) SetCode(v string) *StartDebuggingJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetHttpStatusCode(v int32) *StartDebuggingJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetMessage(v string) *StartDebuggingJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetReportId(v string) *StartDebuggingJMeterSceneResponseBody {
	s.ReportId = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetRequestId(v string) *StartDebuggingJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) SetSuccess(v bool) *StartDebuggingJMeterSceneResponseBody {
	s.Success = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
