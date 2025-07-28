// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitUrgentDemandPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SubmitUrgentDemandPlanResponseBody
	GetCode() *int64
	SetData(v bool) *SubmitUrgentDemandPlanResponseBody
	GetData() *bool
	SetMessage(v string) *SubmitUrgentDemandPlanResponseBody
	GetMessage() *string
	SetSuccess(v bool) *SubmitUrgentDemandPlanResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *SubmitUrgentDemandPlanResponseBody
	GetTraceId() *string
}

type SubmitUrgentDemandPlanResponseBody struct {
	// code
	//
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// success
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 1e2b798516402440016572132e1459
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s SubmitUrgentDemandPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitUrgentDemandPlanResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitUrgentDemandPlanResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SubmitUrgentDemandPlanResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitUrgentDemandPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitUrgentDemandPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitUrgentDemandPlanResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *SubmitUrgentDemandPlanResponseBody) SetCode(v int64) *SubmitUrgentDemandPlanResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponseBody) SetData(v bool) *SubmitUrgentDemandPlanResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponseBody) SetMessage(v string) *SubmitUrgentDemandPlanResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponseBody) SetSuccess(v bool) *SubmitUrgentDemandPlanResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponseBody) SetTraceId(v string) *SubmitUrgentDemandPlanResponseBody {
	s.TraceId = &v
	return s
}

func (s *SubmitUrgentDemandPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
