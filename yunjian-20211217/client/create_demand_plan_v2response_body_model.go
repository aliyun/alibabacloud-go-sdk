// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemandPlanV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateDemandPlanV2ResponseBody
	GetCode() *int64
	SetData(v int64) *CreateDemandPlanV2ResponseBody
	GetData() *int64
	SetMessage(v string) *CreateDemandPlanV2ResponseBody
	GetMessage() *string
	SetSuccess(v bool) *CreateDemandPlanV2ResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateDemandPlanV2ResponseBody
	GetTraceId() *string
}

type CreateDemandPlanV2ResponseBody struct {
	Code    *int64  `json:"code,omitempty" xml:"code,omitempty"`
	Data    *int64  `json:"data,omitempty" xml:"data,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CreateDemandPlanV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDemandPlanV2ResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanV2ResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateDemandPlanV2ResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateDemandPlanV2ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDemandPlanV2ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDemandPlanV2ResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateDemandPlanV2ResponseBody) SetCode(v int64) *CreateDemandPlanV2ResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDemandPlanV2ResponseBody) SetData(v int64) *CreateDemandPlanV2ResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDemandPlanV2ResponseBody) SetMessage(v string) *CreateDemandPlanV2ResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDemandPlanV2ResponseBody) SetSuccess(v bool) *CreateDemandPlanV2ResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDemandPlanV2ResponseBody) SetTraceId(v string) *CreateDemandPlanV2ResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateDemandPlanV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
