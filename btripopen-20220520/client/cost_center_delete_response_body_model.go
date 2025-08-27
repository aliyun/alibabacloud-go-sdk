// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CostCenterDeleteResponseBody
	GetCode() *string
	SetMessage(v string) *CostCenterDeleteResponseBody
	GetMessage() *string
	SetRequestId(v string) *CostCenterDeleteResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CostCenterDeleteResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CostCenterDeleteResponseBody
	GetTraceId() *string
}

type CostCenterDeleteResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CostCenterDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CostCenterDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *CostCenterDeleteResponseBody) GetCode() *string {
	return s.Code
}

func (s *CostCenterDeleteResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CostCenterDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CostCenterDeleteResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CostCenterDeleteResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CostCenterDeleteResponseBody) SetCode(v string) *CostCenterDeleteResponseBody {
	s.Code = &v
	return s
}

func (s *CostCenterDeleteResponseBody) SetMessage(v string) *CostCenterDeleteResponseBody {
	s.Message = &v
	return s
}

func (s *CostCenterDeleteResponseBody) SetRequestId(v string) *CostCenterDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CostCenterDeleteResponseBody) SetSuccess(v bool) *CostCenterDeleteResponseBody {
	s.Success = &v
	return s
}

func (s *CostCenterDeleteResponseBody) SetTraceId(v string) *CostCenterDeleteResponseBody {
	s.TraceId = &v
	return s
}

func (s *CostCenterDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
