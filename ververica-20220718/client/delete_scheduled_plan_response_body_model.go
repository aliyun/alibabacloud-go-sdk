// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ScheduledPlan) *DeleteScheduledPlanResponseBody
	GetData() *ScheduledPlan
	SetErrorCode(v string) *DeleteScheduledPlanResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteScheduledPlanResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *DeleteScheduledPlanResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *DeleteScheduledPlanResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteScheduledPlanResponseBody
	GetSuccess() *bool
}

type DeleteScheduledPlanResponseBody struct {
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

func (s DeleteScheduledPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPlanResponseBody) GetData() *ScheduledPlan {
	return s.Data
}

func (s *DeleteScheduledPlanResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteScheduledPlanResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteScheduledPlanResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *DeleteScheduledPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScheduledPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteScheduledPlanResponseBody) SetData(v *ScheduledPlan) *DeleteScheduledPlanResponseBody {
	s.Data = v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetErrorCode(v string) *DeleteScheduledPlanResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetErrorMessage(v string) *DeleteScheduledPlanResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetHttpCode(v int32) *DeleteScheduledPlanResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetRequestId(v string) *DeleteScheduledPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScheduledPlanResponseBody) SetSuccess(v bool) *DeleteScheduledPlanResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteScheduledPlanResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
