// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTestingJMeterSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartTestingJMeterSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *StartTestingJMeterSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *StartTestingJMeterSceneResponseBody
	GetMessage() *string
	SetReportId(v string) *StartTestingJMeterSceneResponseBody
	GetReportId() *string
	SetRequestId(v string) *StartTestingJMeterSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartTestingJMeterSceneResponseBody
	GetSuccess() *bool
}

type StartTestingJMeterSceneResponseBody struct {
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
	// The ID of the report.
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

func (s StartTestingJMeterSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartTestingJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *StartTestingJMeterSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartTestingJMeterSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *StartTestingJMeterSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartTestingJMeterSceneResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *StartTestingJMeterSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartTestingJMeterSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartTestingJMeterSceneResponseBody) SetCode(v string) *StartTestingJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetHttpStatusCode(v int32) *StartTestingJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetMessage(v string) *StartTestingJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetReportId(v string) *StartTestingJMeterSceneResponseBody {
	s.ReportId = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetRequestId(v string) *StartTestingJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) SetSuccess(v bool) *StartTestingJMeterSceneResponseBody {
	s.Success = &v
	return s
}

func (s *StartTestingJMeterSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
