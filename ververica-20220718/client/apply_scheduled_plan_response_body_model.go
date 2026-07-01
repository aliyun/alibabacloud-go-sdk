// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyScheduledPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScheduledPlanAppliedInfo) *ApplyScheduledPlanResponseBody
	GetData() *ScheduledPlanAppliedInfo
	SetErrorCode(v string) *ApplyScheduledPlanResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ApplyScheduledPlanResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ApplyScheduledPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ApplyScheduledPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ApplyScheduledPlanResponseBody
	GetSuccess() *bool
}

type ApplyScheduledPlanResponseBody struct {
	// The data structure of the applied scheduled plan.
	Data *ScheduledPlanAppliedInfo `json:"data,omitempty" xml:"data,omitempty"`
	// The error code returned if the request fails. This parameter is empty if the request is successful.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message returned if the request fails. This parameter is empty if the request is successful.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code. A value of 200 is returned. Use the \\`success\\` parameter to determine if the request was successful.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The ID of the request.
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

func (s ApplyScheduledPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyScheduledPlanResponseBody) GetData() *ScheduledPlanAppliedInfo {
	return s.Data
}

func (s *ApplyScheduledPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ApplyScheduledPlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ApplyScheduledPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ApplyScheduledPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyScheduledPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ApplyScheduledPlanResponseBody) SetData(v *ScheduledPlanAppliedInfo) *ApplyScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetErrorCode(v string) *ApplyScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetErrorMessage(v string) *ApplyScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetHttpCode(v int32) *ApplyScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetRequestId(v string) *ApplyScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyScheduledPlanResponseBody) SetSuccess(v bool) *ApplyScheduledPlanResponseBody {
	s.Success = &v
	return s
}

func (s *ApplyScheduledPlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
