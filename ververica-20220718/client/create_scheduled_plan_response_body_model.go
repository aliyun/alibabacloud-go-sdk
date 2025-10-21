// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScheduledPlan) *CreateScheduledPlanResponseBody
	GetData() *ScheduledPlan
	SetErrorCode(v string) *CreateScheduledPlanResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateScheduledPlanResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *CreateScheduledPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *CreateScheduledPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateScheduledPlanResponseBody
	GetSuccess() *bool
}

type CreateScheduledPlanResponseBody struct {
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
	// CBC799F0-ABCD-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateScheduledPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledPlanResponseBody) GetData() *ScheduledPlan {
	return s.Data
}

func (s *CreateScheduledPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateScheduledPlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateScheduledPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *CreateScheduledPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduledPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateScheduledPlanResponseBody) SetData(v *ScheduledPlan) *CreateScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetErrorCode(v string) *CreateScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetErrorMessage(v string) *CreateScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetHttpCode(v int32) *CreateScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetRequestId(v string) *CreateScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledPlanResponseBody) SetSuccess(v bool) *CreateScheduledPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateScheduledPlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
