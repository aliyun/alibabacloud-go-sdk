// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemandPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateDemandPlanResponseBody
	GetCode() *int64
	SetData(v int64) *CreateDemandPlanResponseBody
	GetData() *int64
	SetMessage(v string) *CreateDemandPlanResponseBody
	GetMessage() *string
	SetSuccess(v bool) *CreateDemandPlanResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateDemandPlanResponseBody
	GetTraceId() *string
}

type CreateDemandPlanResponseBody struct {
	// code
	//
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 111223
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
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

func (s CreateDemandPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDemandPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateDemandPlanResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateDemandPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDemandPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDemandPlanResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateDemandPlanResponseBody) SetCode(v int64) *CreateDemandPlanResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDemandPlanResponseBody) SetData(v int64) *CreateDemandPlanResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDemandPlanResponseBody) SetMessage(v string) *CreateDemandPlanResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDemandPlanResponseBody) SetSuccess(v bool) *CreateDemandPlanResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDemandPlanResponseBody) SetTraceId(v string) *CreateDemandPlanResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateDemandPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
