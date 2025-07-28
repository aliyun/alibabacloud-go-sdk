// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushResourcePlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *PushResourcePlanResponseBody
	GetCode() *int64
	SetData(v bool) *PushResourcePlanResponseBody
	GetData() *bool
	SetMessage(v string) *PushResourcePlanResponseBody
	GetMessage() *string
	SetSuccess(v bool) *PushResourcePlanResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *PushResourcePlanResponseBody
	GetTraceId() *string
}

type PushResourcePlanResponseBody struct {
	Code    *int64  `json:"code,omitempty" xml:"code,omitempty"`
	Data    *bool   `json:"data,omitempty" xml:"data,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Success *bool   `json:"success,omitempty" xml:"success,omitempty"`
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s PushResourcePlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushResourcePlanResponseBody) GoString() string {
	return s.String()
}

func (s *PushResourcePlanResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *PushResourcePlanResponseBody) GetData() *bool {
	return s.Data
}

func (s *PushResourcePlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PushResourcePlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PushResourcePlanResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *PushResourcePlanResponseBody) SetCode(v int64) *PushResourcePlanResponseBody {
	s.Code = &v
	return s
}

func (s *PushResourcePlanResponseBody) SetData(v bool) *PushResourcePlanResponseBody {
	s.Data = &v
	return s
}

func (s *PushResourcePlanResponseBody) SetMessage(v string) *PushResourcePlanResponseBody {
	s.Message = &v
	return s
}

func (s *PushResourcePlanResponseBody) SetSuccess(v bool) *PushResourcePlanResponseBody {
	s.Success = &v
	return s
}

func (s *PushResourcePlanResponseBody) SetTraceId(v string) *PushResourcePlanResponseBody {
	s.TraceId = &v
	return s
}

func (s *PushResourcePlanResponseBody) Validate() error {
	return dara.Validate(s)
}
