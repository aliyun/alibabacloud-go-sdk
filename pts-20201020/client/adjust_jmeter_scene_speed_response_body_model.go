// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustJMeterSceneSpeedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AdjustJMeterSceneSpeedResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AdjustJMeterSceneSpeedResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AdjustJMeterSceneSpeedResponseBody
	GetMessage() *string
	SetReportId(v string) *AdjustJMeterSceneSpeedResponseBody
	GetReportId() *string
	SetRequestId(v string) *AdjustJMeterSceneSpeedResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AdjustJMeterSceneSpeedResponseBody
	GetSuccess() *bool
}

type AdjustJMeterSceneSpeedResponseBody struct {
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
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the report.
	//
	// example:
	//
	// DYYPZIH
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DC4E31DDA77-6745-4925-B423-4E89VV34221A
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

func (s AdjustJMeterSceneSpeedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AdjustJMeterSceneSpeedResponseBody) GoString() string {
	return s.String()
}

func (s *AdjustJMeterSceneSpeedResponseBody) GetCode() *string {
	return s.Code
}

func (s *AdjustJMeterSceneSpeedResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AdjustJMeterSceneSpeedResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AdjustJMeterSceneSpeedResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *AdjustJMeterSceneSpeedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AdjustJMeterSceneSpeedResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AdjustJMeterSceneSpeedResponseBody) SetCode(v string) *AdjustJMeterSceneSpeedResponseBody {
	s.Code = &v
	return s
}

func (s *AdjustJMeterSceneSpeedResponseBody) SetHttpStatusCode(v int32) *AdjustJMeterSceneSpeedResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AdjustJMeterSceneSpeedResponseBody) SetMessage(v string) *AdjustJMeterSceneSpeedResponseBody {
	s.Message = &v
	return s
}

func (s *AdjustJMeterSceneSpeedResponseBody) SetReportId(v string) *AdjustJMeterSceneSpeedResponseBody {
	s.ReportId = &v
	return s
}

func (s *AdjustJMeterSceneSpeedResponseBody) SetRequestId(v string) *AdjustJMeterSceneSpeedResponseBody {
	s.RequestId = &v
	return s
}

func (s *AdjustJMeterSceneSpeedResponseBody) SetSuccess(v bool) *AdjustJMeterSceneSpeedResponseBody {
	s.Success = &v
	return s
}

func (s *AdjustJMeterSceneSpeedResponseBody) Validate() error {
	return dara.Validate(s)
}
