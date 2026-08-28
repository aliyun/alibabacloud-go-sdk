// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerPatrolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *TriggerPatrolResponseBodyData) *TriggerPatrolResponseBody
	GetData() *TriggerPatrolResponseBodyData
	SetErrorCode(v string) *TriggerPatrolResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *TriggerPatrolResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *TriggerPatrolResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *TriggerPatrolResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TriggerPatrolResponseBody
	GetSuccess() *bool
}

type TriggerPatrolResponseBody struct {
	// The response data of the triggered inspection.
	Data *TriggerPatrolResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// When success is false, this value is not empty and indicates the business error code. When success is true, this value is empty.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// When success is false, this value is not empty and indicates the business error message. When success is true, this value is empty.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The business status code, which is uniformly 200. Use success to determine whether the business request is successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the business request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TriggerPatrolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerPatrolResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerPatrolResponseBody) GetData() *TriggerPatrolResponseBodyData {
	return s.Data
}

func (s *TriggerPatrolResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *TriggerPatrolResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *TriggerPatrolResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *TriggerPatrolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerPatrolResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TriggerPatrolResponseBody) SetData(v *TriggerPatrolResponseBodyData) *TriggerPatrolResponseBody {
	s.Data = v
	return s
}

func (s *TriggerPatrolResponseBody) SetErrorCode(v string) *TriggerPatrolResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *TriggerPatrolResponseBody) SetErrorMessage(v string) *TriggerPatrolResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *TriggerPatrolResponseBody) SetHttpCode(v int32) *TriggerPatrolResponseBody {
	s.HttpCode = &v
	return s
}

func (s *TriggerPatrolResponseBody) SetRequestId(v string) *TriggerPatrolResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerPatrolResponseBody) SetSuccess(v bool) *TriggerPatrolResponseBody {
	s.Success = &v
	return s
}

func (s *TriggerPatrolResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TriggerPatrolResponseBodyData struct {
	// The generated report ID.
	//
	// example:
	//
	// inspection-cf8f8843-64e4-4b45-9500-06790107130f
	ReportId *string `json:"reportId,omitempty" xml:"reportId,omitempty"`
	// The report status.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TriggerPatrolResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TriggerPatrolResponseBodyData) GoString() string {
	return s.String()
}

func (s *TriggerPatrolResponseBodyData) GetReportId() *string {
	return s.ReportId
}

func (s *TriggerPatrolResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *TriggerPatrolResponseBodyData) SetReportId(v string) *TriggerPatrolResponseBodyData {
	s.ReportId = &v
	return s
}

func (s *TriggerPatrolResponseBodyData) SetStatus(v string) *TriggerPatrolResponseBodyData {
	s.Status = &v
	return s
}

func (s *TriggerPatrolResponseBodyData) Validate() error {
	return dara.Validate(s)
}
