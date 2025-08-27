// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCostCenterSaveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CostCenterSaveResponseBody
	GetCode() *string
	SetMessage(v string) *CostCenterSaveResponseBody
	GetMessage() *string
	SetModule(v *CostCenterSaveResponseBodyModule) *CostCenterSaveResponseBody
	GetModule() *CostCenterSaveResponseBodyModule
	SetRequestId(v string) *CostCenterSaveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CostCenterSaveResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CostCenterSaveResponseBody
	GetTraceId() *string
}

type CostCenterSaveResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                           `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                           `json:"message,omitempty" xml:"message,omitempty"`
	Module  *CostCenterSaveResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
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

func (s CostCenterSaveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CostCenterSaveResponseBody) GoString() string {
	return s.String()
}

func (s *CostCenterSaveResponseBody) GetCode() *string {
	return s.Code
}

func (s *CostCenterSaveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CostCenterSaveResponseBody) GetModule() *CostCenterSaveResponseBodyModule {
	return s.Module
}

func (s *CostCenterSaveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CostCenterSaveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CostCenterSaveResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CostCenterSaveResponseBody) SetCode(v string) *CostCenterSaveResponseBody {
	s.Code = &v
	return s
}

func (s *CostCenterSaveResponseBody) SetMessage(v string) *CostCenterSaveResponseBody {
	s.Message = &v
	return s
}

func (s *CostCenterSaveResponseBody) SetModule(v *CostCenterSaveResponseBodyModule) *CostCenterSaveResponseBody {
	s.Module = v
	return s
}

func (s *CostCenterSaveResponseBody) SetRequestId(v string) *CostCenterSaveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CostCenterSaveResponseBody) SetSuccess(v bool) *CostCenterSaveResponseBody {
	s.Success = &v
	return s
}

func (s *CostCenterSaveResponseBody) SetTraceId(v string) *CostCenterSaveResponseBody {
	s.TraceId = &v
	return s
}

func (s *CostCenterSaveResponseBody) Validate() error {
	return dara.Validate(s)
}

type CostCenterSaveResponseBodyModule struct {
	// example:
	//
	// 17690
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
}

func (s CostCenterSaveResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CostCenterSaveResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CostCenterSaveResponseBodyModule) GetId() *int64 {
	return s.Id
}

func (s *CostCenterSaveResponseBodyModule) SetId(v int64) *CostCenterSaveResponseBodyModule {
	s.Id = &v
	return s
}

func (s *CostCenterSaveResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
