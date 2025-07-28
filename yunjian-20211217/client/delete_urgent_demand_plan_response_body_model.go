// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUrgentDemandPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DeleteUrgentDemandPlanResponseBody
	GetCode() *int64
	SetData(v int64) *DeleteUrgentDemandPlanResponseBody
	GetData() *int64
	SetMessage(v string) *DeleteUrgentDemandPlanResponseBody
	GetMessage() *string
	SetSuccess(v bool) *DeleteUrgentDemandPlanResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteUrgentDemandPlanResponseBody
	GetTraceId() *string
}

type DeleteUrgentDemandPlanResponseBody struct {
	// example:
	//
	// 0
	Code *int64 `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// msg
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 212cf01016405759151137225e83cd
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s DeleteUrgentDemandPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteUrgentDemandPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteUrgentDemandPlanResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DeleteUrgentDemandPlanResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeleteUrgentDemandPlanResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteUrgentDemandPlanResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteUrgentDemandPlanResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteUrgentDemandPlanResponseBody) SetCode(v int64) *DeleteUrgentDemandPlanResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponseBody) SetData(v int64) *DeleteUrgentDemandPlanResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponseBody) SetMessage(v string) *DeleteUrgentDemandPlanResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponseBody) SetSuccess(v bool) *DeleteUrgentDemandPlanResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponseBody) SetTraceId(v string) *DeleteUrgentDemandPlanResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteUrgentDemandPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
