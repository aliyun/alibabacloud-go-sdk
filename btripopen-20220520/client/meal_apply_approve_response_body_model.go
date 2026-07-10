// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealApplyApproveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MealApplyApproveResponseBody
	GetCode() *string
	SetMessage(v string) *MealApplyApproveResponseBody
	GetMessage() *string
	SetRequestId(v string) *MealApplyApproveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MealApplyApproveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MealApplyApproveResponseBody
	GetTraceId() *string
}

type MealApplyApproveResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId   *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MealApplyApproveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MealApplyApproveResponseBody) GoString() string {
	return s.String()
}

func (s *MealApplyApproveResponseBody) GetCode() *string {
	return s.Code
}

func (s *MealApplyApproveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MealApplyApproveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MealApplyApproveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MealApplyApproveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MealApplyApproveResponseBody) SetCode(v string) *MealApplyApproveResponseBody {
	s.Code = &v
	return s
}

func (s *MealApplyApproveResponseBody) SetMessage(v string) *MealApplyApproveResponseBody {
	s.Message = &v
	return s
}

func (s *MealApplyApproveResponseBody) SetRequestId(v string) *MealApplyApproveResponseBody {
	s.RequestId = &v
	return s
}

func (s *MealApplyApproveResponseBody) SetSuccess(v bool) *MealApplyApproveResponseBody {
	s.Success = &v
	return s
}

func (s *MealApplyApproveResponseBody) SetTraceId(v string) *MealApplyApproveResponseBody {
	s.TraceId = &v
	return s
}

func (s *MealApplyApproveResponseBody) Validate() error {
	return dara.Validate(s)
}
