// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderRefundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsureOrderRefundResponseBody
	GetCode() *string
	SetMessage(v string) *InsureOrderRefundResponseBody
	GetMessage() *string
	SetModule(v *InsureOrderRefundResponseBodyModule) *InsureOrderRefundResponseBody
	GetModule() *InsureOrderRefundResponseBodyModule
	SetRequestId(v string) *InsureOrderRefundResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsureOrderRefundResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InsureOrderRefundResponseBody
	GetTraceId() *string
}

type InsureOrderRefundResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                              `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                              `json:"message,omitempty" xml:"message,omitempty"`
	Module  *InsureOrderRefundResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 2103ad3116824902540648188de7ac
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210e846c16726306481681232d441f
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InsureOrderRefundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderRefundResponseBody) GoString() string {
	return s.String()
}

func (s *InsureOrderRefundResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsureOrderRefundResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsureOrderRefundResponseBody) GetModule() *InsureOrderRefundResponseBodyModule {
	return s.Module
}

func (s *InsureOrderRefundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsureOrderRefundResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsureOrderRefundResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InsureOrderRefundResponseBody) SetCode(v string) *InsureOrderRefundResponseBody {
	s.Code = &v
	return s
}

func (s *InsureOrderRefundResponseBody) SetMessage(v string) *InsureOrderRefundResponseBody {
	s.Message = &v
	return s
}

func (s *InsureOrderRefundResponseBody) SetModule(v *InsureOrderRefundResponseBodyModule) *InsureOrderRefundResponseBody {
	s.Module = v
	return s
}

func (s *InsureOrderRefundResponseBody) SetRequestId(v string) *InsureOrderRefundResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsureOrderRefundResponseBody) SetSuccess(v bool) *InsureOrderRefundResponseBody {
	s.Success = &v
	return s
}

func (s *InsureOrderRefundResponseBody) SetTraceId(v string) *InsureOrderRefundResponseBody {
	s.TraceId = &v
	return s
}

func (s *InsureOrderRefundResponseBody) Validate() error {
	return dara.Validate(s)
}

type InsureOrderRefundResponseBodyModule struct {
	// example:
	//
	// 118526587
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 1423041410342678003
	InsOrderId    *string                                             `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	InsRefundList []*InsureOrderRefundResponseBodyModuleInsRefundList `json:"ins_refund_list,omitempty" xml:"ins_refund_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1423041410342678022
	OutApplyId *string `json:"out_apply_id,omitempty" xml:"out_apply_id,omitempty"`
}

func (s InsureOrderRefundResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderRefundResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InsureOrderRefundResponseBodyModule) GetApplyId() *string {
	return s.ApplyId
}

func (s *InsureOrderRefundResponseBodyModule) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *InsureOrderRefundResponseBodyModule) GetInsRefundList() []*InsureOrderRefundResponseBodyModuleInsRefundList {
	return s.InsRefundList
}

func (s *InsureOrderRefundResponseBodyModule) GetOutApplyId() *string {
	return s.OutApplyId
}

func (s *InsureOrderRefundResponseBodyModule) SetApplyId(v string) *InsureOrderRefundResponseBodyModule {
	s.ApplyId = &v
	return s
}

func (s *InsureOrderRefundResponseBodyModule) SetInsOrderId(v string) *InsureOrderRefundResponseBodyModule {
	s.InsOrderId = &v
	return s
}

func (s *InsureOrderRefundResponseBodyModule) SetInsRefundList(v []*InsureOrderRefundResponseBodyModuleInsRefundList) *InsureOrderRefundResponseBodyModule {
	s.InsRefundList = v
	return s
}

func (s *InsureOrderRefundResponseBodyModule) SetOutApplyId(v string) *InsureOrderRefundResponseBodyModule {
	s.OutApplyId = &v
	return s
}

func (s *InsureOrderRefundResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type InsureOrderRefundResponseBodyModuleInsRefundList struct {
	// example:
	//
	// po102000399221
	PolicyRefundNo *string `json:"policy_refund_no,omitempty" xml:"policy_refund_no,omitempty"`
	// example:
	//
	// REFUND_SUCCESS
	RefundStatus *string `json:"refund_status,omitempty" xml:"refund_status,omitempty"`
	// example:
	//
	// 1230012499921
	SubInsOrderId *string `json:"sub_ins_order_id,omitempty" xml:"sub_ins_order_id,omitempty"`
}

func (s InsureOrderRefundResponseBodyModuleInsRefundList) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderRefundResponseBodyModuleInsRefundList) GoString() string {
	return s.String()
}

func (s *InsureOrderRefundResponseBodyModuleInsRefundList) GetPolicyRefundNo() *string {
	return s.PolicyRefundNo
}

func (s *InsureOrderRefundResponseBodyModuleInsRefundList) GetRefundStatus() *string {
	return s.RefundStatus
}

func (s *InsureOrderRefundResponseBodyModuleInsRefundList) GetSubInsOrderId() *string {
	return s.SubInsOrderId
}

func (s *InsureOrderRefundResponseBodyModuleInsRefundList) SetPolicyRefundNo(v string) *InsureOrderRefundResponseBodyModuleInsRefundList {
	s.PolicyRefundNo = &v
	return s
}

func (s *InsureOrderRefundResponseBodyModuleInsRefundList) SetRefundStatus(v string) *InsureOrderRefundResponseBodyModuleInsRefundList {
	s.RefundStatus = &v
	return s
}

func (s *InsureOrderRefundResponseBodyModuleInsRefundList) SetSubInsOrderId(v string) *InsureOrderRefundResponseBodyModuleInsRefundList {
	s.SubInsOrderId = &v
	return s
}

func (s *InsureOrderRefundResponseBodyModuleInsRefundList) Validate() error {
	return dara.Validate(s)
}
