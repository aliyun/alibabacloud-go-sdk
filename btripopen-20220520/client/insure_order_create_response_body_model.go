// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsureOrderCreateResponseBody
	GetCode() *string
	SetMessage(v string) *InsureOrderCreateResponseBody
	GetMessage() *string
	SetModule(v *InsureOrderCreateResponseBodyModule) *InsureOrderCreateResponseBody
	GetModule() *InsureOrderCreateResponseBodyModule
	SetRequestId(v string) *InsureOrderCreateResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsureOrderCreateResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InsureOrderCreateResponseBody
	GetTraceId() *string
}

type InsureOrderCreateResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  *InsureOrderCreateResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// C6055EA5-C566-511D-9FC4-5E4D45182711
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210f079916782711059363565d6be1
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InsureOrderCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateResponseBody) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsureOrderCreateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsureOrderCreateResponseBody) GetModule() *InsureOrderCreateResponseBodyModule {
	return s.Module
}

func (s *InsureOrderCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsureOrderCreateResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsureOrderCreateResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InsureOrderCreateResponseBody) SetCode(v string) *InsureOrderCreateResponseBody {
	s.Code = &v
	return s
}

func (s *InsureOrderCreateResponseBody) SetMessage(v string) *InsureOrderCreateResponseBody {
	s.Message = &v
	return s
}

func (s *InsureOrderCreateResponseBody) SetModule(v *InsureOrderCreateResponseBodyModule) *InsureOrderCreateResponseBody {
	s.Module = v
	return s
}

func (s *InsureOrderCreateResponseBody) SetRequestId(v string) *InsureOrderCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsureOrderCreateResponseBody) SetSuccess(v bool) *InsureOrderCreateResponseBody {
	s.Success = &v
	return s
}

func (s *InsureOrderCreateResponseBody) SetTraceId(v string) *InsureOrderCreateResponseBody {
	s.TraceId = &v
	return s
}

func (s *InsureOrderCreateResponseBody) Validate() error {
	return dara.Validate(s)
}

type InsureOrderCreateResponseBodyModule struct {
	// example:
	//
	// 1
	Copies                *int32                                                      `json:"copies,omitempty" xml:"copies,omitempty"`
	InsureOrderDetailList []*InsureOrderCreateResponseBodyModuleInsureOrderDetailList `json:"insure_order_detail_list,omitempty" xml:"insure_order_detail_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1021000196500370001
	InsureOrderId *string `json:"insure_order_id,omitempty" xml:"insure_order_id,omitempty"`
	// example:
	//
	// 3000
	Premium *int64 `json:"premium,omitempty" xml:"premium,omitempty"`
}

func (s InsureOrderCreateResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateResponseBodyModule) GetCopies() *int32 {
	return s.Copies
}

func (s *InsureOrderCreateResponseBodyModule) GetInsureOrderDetailList() []*InsureOrderCreateResponseBodyModuleInsureOrderDetailList {
	return s.InsureOrderDetailList
}

func (s *InsureOrderCreateResponseBodyModule) GetInsureOrderId() *string {
	return s.InsureOrderId
}

func (s *InsureOrderCreateResponseBodyModule) GetPremium() *int64 {
	return s.Premium
}

func (s *InsureOrderCreateResponseBodyModule) SetCopies(v int32) *InsureOrderCreateResponseBodyModule {
	s.Copies = &v
	return s
}

func (s *InsureOrderCreateResponseBodyModule) SetInsureOrderDetailList(v []*InsureOrderCreateResponseBodyModuleInsureOrderDetailList) *InsureOrderCreateResponseBodyModule {
	s.InsureOrderDetailList = v
	return s
}

func (s *InsureOrderCreateResponseBodyModule) SetInsureOrderId(v string) *InsureOrderCreateResponseBodyModule {
	s.InsureOrderId = &v
	return s
}

func (s *InsureOrderCreateResponseBodyModule) SetPremium(v int64) *InsureOrderCreateResponseBodyModule {
	s.Premium = &v
	return s
}

func (s *InsureOrderCreateResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type InsureOrderCreateResponseBodyModuleInsureOrderDetailList struct {
	// example:
	//
	// 1992939412431231
	OutSubInsOrderId *string `json:"out_sub_ins_order_id,omitempty" xml:"out_sub_ins_order_id,omitempty"`
	// example:
	//
	// 121234444
	PolicyNo *string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// example:
	//
	// INIT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1022196500378001
	SubInsOrderId *string `json:"sub_ins_order_id,omitempty" xml:"sub_ins_order_id,omitempty"`
}

func (s InsureOrderCreateResponseBodyModuleInsureOrderDetailList) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderCreateResponseBodyModuleInsureOrderDetailList) GoString() string {
	return s.String()
}

func (s *InsureOrderCreateResponseBodyModuleInsureOrderDetailList) GetOutSubInsOrderId() *string {
	return s.OutSubInsOrderId
}

func (s *InsureOrderCreateResponseBodyModuleInsureOrderDetailList) GetPolicyNo() *string {
	return s.PolicyNo
}

func (s *InsureOrderCreateResponseBodyModuleInsureOrderDetailList) GetStatus() *string {
	return s.Status
}

func (s *InsureOrderCreateResponseBodyModuleInsureOrderDetailList) GetSubInsOrderId() *string {
	return s.SubInsOrderId
}

func (s *InsureOrderCreateResponseBodyModuleInsureOrderDetailList) SetOutSubInsOrderId(v string) *InsureOrderCreateResponseBodyModuleInsureOrderDetailList {
	s.OutSubInsOrderId = &v
	return s
}

func (s *InsureOrderCreateResponseBodyModuleInsureOrderDetailList) SetPolicyNo(v string) *InsureOrderCreateResponseBodyModuleInsureOrderDetailList {
	s.PolicyNo = &v
	return s
}

func (s *InsureOrderCreateResponseBodyModuleInsureOrderDetailList) SetStatus(v string) *InsureOrderCreateResponseBodyModuleInsureOrderDetailList {
	s.Status = &v
	return s
}

func (s *InsureOrderCreateResponseBodyModuleInsureOrderDetailList) SetSubInsOrderId(v string) *InsureOrderCreateResponseBodyModuleInsureOrderDetailList {
	s.SubInsOrderId = &v
	return s
}

func (s *InsureOrderCreateResponseBodyModuleInsureOrderDetailList) Validate() error {
	return dara.Validate(s)
}
