// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScheduledPlan) *UpdateScheduledPlanResponseBody
	GetData() *ScheduledPlan
	SetErrorCode(v string) *UpdateScheduledPlanResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateScheduledPlanResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *UpdateScheduledPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *UpdateScheduledPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateScheduledPlanResponseBody
	GetSuccess() *bool
}

type UpdateScheduledPlanResponseBody struct {
	Data *ScheduledPlan `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateScheduledPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPlanResponseBody) GetData() *ScheduledPlan {
	return s.Data
}

func (s *UpdateScheduledPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateScheduledPlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateScheduledPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *UpdateScheduledPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScheduledPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateScheduledPlanResponseBody) SetData(v *ScheduledPlan) *UpdateScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetErrorCode(v string) *UpdateScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetErrorMessage(v string) *UpdateScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetHttpCode(v int32) *UpdateScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetRequestId(v string) *UpdateScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduledPlanResponseBody) SetSuccess(v bool) *UpdateScheduledPlanResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateScheduledPlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
