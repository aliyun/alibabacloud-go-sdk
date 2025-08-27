// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterModifyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CostCenterModifyResponseBody
	GetCode() *string
	SetMessage(v string) *CostCenterModifyResponseBody
	GetMessage() *string
	SetRequestId(v string) *CostCenterModifyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CostCenterModifyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CostCenterModifyResponseBody
	GetTraceId() *string
}

type CostCenterModifyResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CostCenterModifyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CostCenterModifyResponseBody) GoString() string {
	return s.String()
}

func (s *CostCenterModifyResponseBody) GetCode() *string {
	return s.Code
}

func (s *CostCenterModifyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CostCenterModifyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CostCenterModifyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CostCenterModifyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CostCenterModifyResponseBody) SetCode(v string) *CostCenterModifyResponseBody {
	s.Code = &v
	return s
}

func (s *CostCenterModifyResponseBody) SetMessage(v string) *CostCenterModifyResponseBody {
	s.Message = &v
	return s
}

func (s *CostCenterModifyResponseBody) SetRequestId(v string) *CostCenterModifyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CostCenterModifyResponseBody) SetSuccess(v bool) *CostCenterModifyResponseBody {
	s.Success = &v
	return s
}

func (s *CostCenterModifyResponseBody) SetTraceId(v string) *CostCenterModifyResponseBody {
	s.TraceId = &v
	return s
}

func (s *CostCenterModifyResponseBody) Validate() error {
	return dara.Validate(s)
}
