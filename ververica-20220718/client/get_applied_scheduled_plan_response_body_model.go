// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppliedScheduledPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScheduledPlanAppliedInfo) *GetAppliedScheduledPlanResponseBody
	GetData() *ScheduledPlanAppliedInfo
	SetErrorCode(v string) *GetAppliedScheduledPlanResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAppliedScheduledPlanResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *GetAppliedScheduledPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetAppliedScheduledPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAppliedScheduledPlanResponseBody
	GetSuccess() *bool
}

type GetAppliedScheduledPlanResponseBody struct {
	// The scheduled plan for the application.
	Data *ScheduledPlanAppliedInfo `json:"data,omitempty" xml:"data,omitempty"`
	// The error code. This parameter is returned if the request fails. If the request is successful, this parameter is empty.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message. This parameter is returned if the request fails. If the request is successful, this parameter is empty.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code. A value of 200 is always returned. Use the success parameter to determine whether the request was successful.
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

func (s GetAppliedScheduledPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAppliedScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *GetAppliedScheduledPlanResponseBody) GetData() *ScheduledPlanAppliedInfo {
	return s.Data
}

func (s *GetAppliedScheduledPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAppliedScheduledPlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAppliedScheduledPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetAppliedScheduledPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAppliedScheduledPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAppliedScheduledPlanResponseBody) SetData(v *ScheduledPlanAppliedInfo) *GetAppliedScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetErrorCode(v string) *GetAppliedScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetErrorMessage(v string) *GetAppliedScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetHttpCode(v int32) *GetAppliedScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetRequestId(v string) *GetAppliedScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) SetSuccess(v bool) *GetAppliedScheduledPlanResponseBody {
	s.Success = &v
	return s
}

func (s *GetAppliedScheduledPlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
