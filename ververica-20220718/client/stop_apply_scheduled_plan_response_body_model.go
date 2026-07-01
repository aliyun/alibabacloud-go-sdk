// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopApplyScheduledPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScheduledPlanAppliedInfo) *StopApplyScheduledPlanResponseBody
	GetData() *ScheduledPlanAppliedInfo
	SetErrorCode(v string) *StopApplyScheduledPlanResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StopApplyScheduledPlanResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *StopApplyScheduledPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *StopApplyScheduledPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopApplyScheduledPlanResponseBody
	GetSuccess() *bool
}

type StopApplyScheduledPlanResponseBody struct {
	// The application information for the scheduled plan.
	Data *ScheduledPlanAppliedInfo `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. This parameter is not empty if the request fails. This parameter is empty if the request is successful.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message. This parameter is not empty if the request fails. This parameter is empty if the request is successful.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code. A value of 200 is always returned. Use the \\`success\\` parameter to check if the request was successful.
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
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s StopApplyScheduledPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopApplyScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *StopApplyScheduledPlanResponseBody) GetData() *ScheduledPlanAppliedInfo {
	return s.Data
}

func (s *StopApplyScheduledPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StopApplyScheduledPlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopApplyScheduledPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *StopApplyScheduledPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopApplyScheduledPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopApplyScheduledPlanResponseBody) SetData(v *ScheduledPlanAppliedInfo) *StopApplyScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetErrorCode(v string) *StopApplyScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetErrorMessage(v string) *StopApplyScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetHttpCode(v int32) *StopApplyScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetRequestId(v string) *StopApplyScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) SetSuccess(v bool) *StopApplyScheduledPlanResponseBody {
	s.Success = &v
	return s
}

func (s *StopApplyScheduledPlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
