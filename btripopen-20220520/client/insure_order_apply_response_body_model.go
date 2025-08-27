// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsureOrderApplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsureOrderApplyResponseBody
	GetCode() *string
	SetMessage(v string) *InsureOrderApplyResponseBody
	GetMessage() *string
	SetModule(v *InsureOrderApplyResponseBodyModule) *InsureOrderApplyResponseBody
	GetModule() *InsureOrderApplyResponseBodyModule
	SetRequestId(v string) *InsureOrderApplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsureOrderApplyResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InsureOrderApplyResponseBody
	GetTraceId() *string
}

type InsureOrderApplyResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                             `json:"message,omitempty" xml:"message,omitempty"`
	Module  *InsureOrderApplyResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210e800316781571635951548d1e9d
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InsureOrderApplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderApplyResponseBody) GoString() string {
	return s.String()
}

func (s *InsureOrderApplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsureOrderApplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsureOrderApplyResponseBody) GetModule() *InsureOrderApplyResponseBodyModule {
	return s.Module
}

func (s *InsureOrderApplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsureOrderApplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsureOrderApplyResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InsureOrderApplyResponseBody) SetCode(v string) *InsureOrderApplyResponseBody {
	s.Code = &v
	return s
}

func (s *InsureOrderApplyResponseBody) SetMessage(v string) *InsureOrderApplyResponseBody {
	s.Message = &v
	return s
}

func (s *InsureOrderApplyResponseBody) SetModule(v *InsureOrderApplyResponseBodyModule) *InsureOrderApplyResponseBody {
	s.Module = v
	return s
}

func (s *InsureOrderApplyResponseBody) SetRequestId(v string) *InsureOrderApplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsureOrderApplyResponseBody) SetSuccess(v bool) *InsureOrderApplyResponseBody {
	s.Success = &v
	return s
}

func (s *InsureOrderApplyResponseBody) SetTraceId(v string) *InsureOrderApplyResponseBody {
	s.TraceId = &v
	return s
}

func (s *InsureOrderApplyResponseBody) Validate() error {
	return dara.Validate(s)
}

type InsureOrderApplyResponseBodyModule struct {
	// example:
	//
	// 1021000196500370003
	InsOrderId         *string                                                 `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	InsOrderPolicyList []*InsureOrderApplyResponseBodyModuleInsOrderPolicyList `json:"ins_order_policy_list,omitempty" xml:"ins_order_policy_list,omitempty" type:"Repeated"`
}

func (s InsureOrderApplyResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderApplyResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InsureOrderApplyResponseBodyModule) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *InsureOrderApplyResponseBodyModule) GetInsOrderPolicyList() []*InsureOrderApplyResponseBodyModuleInsOrderPolicyList {
	return s.InsOrderPolicyList
}

func (s *InsureOrderApplyResponseBodyModule) SetInsOrderId(v string) *InsureOrderApplyResponseBodyModule {
	s.InsOrderId = &v
	return s
}

func (s *InsureOrderApplyResponseBodyModule) SetInsOrderPolicyList(v []*InsureOrderApplyResponseBodyModuleInsOrderPolicyList) *InsureOrderApplyResponseBodyModule {
	s.InsOrderPolicyList = v
	return s
}

func (s *InsureOrderApplyResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type InsureOrderApplyResponseBodyModuleInsOrderPolicyList struct {
	// example:
	//
	// 1022196500378006
	OutSubInsOrderId *string `json:"out_sub_ins_order_id,omitempty" xml:"out_sub_ins_order_id,omitempty"`
	// example:
	//
	// po10002300201
	PolicyNo *string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// example:
	//
	// INIT
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1022196500378003
	SubInsOrderId *string `json:"sub_ins_order_id,omitempty" xml:"sub_ins_order_id,omitempty"`
}

func (s InsureOrderApplyResponseBodyModuleInsOrderPolicyList) String() string {
	return dara.Prettify(s)
}

func (s InsureOrderApplyResponseBodyModuleInsOrderPolicyList) GoString() string {
	return s.String()
}

func (s *InsureOrderApplyResponseBodyModuleInsOrderPolicyList) GetOutSubInsOrderId() *string {
	return s.OutSubInsOrderId
}

func (s *InsureOrderApplyResponseBodyModuleInsOrderPolicyList) GetPolicyNo() *string {
	return s.PolicyNo
}

func (s *InsureOrderApplyResponseBodyModuleInsOrderPolicyList) GetStatus() *string {
	return s.Status
}

func (s *InsureOrderApplyResponseBodyModuleInsOrderPolicyList) GetSubInsOrderId() *string {
	return s.SubInsOrderId
}

func (s *InsureOrderApplyResponseBodyModuleInsOrderPolicyList) SetOutSubInsOrderId(v string) *InsureOrderApplyResponseBodyModuleInsOrderPolicyList {
	s.OutSubInsOrderId = &v
	return s
}

func (s *InsureOrderApplyResponseBodyModuleInsOrderPolicyList) SetPolicyNo(v string) *InsureOrderApplyResponseBodyModuleInsOrderPolicyList {
	s.PolicyNo = &v
	return s
}

func (s *InsureOrderApplyResponseBodyModuleInsOrderPolicyList) SetStatus(v string) *InsureOrderApplyResponseBodyModuleInsOrderPolicyList {
	s.Status = &v
	return s
}

func (s *InsureOrderApplyResponseBodyModuleInsOrderPolicyList) SetSubInsOrderId(v string) *InsureOrderApplyResponseBodyModuleInsOrderPolicyList {
	s.SubInsOrderId = &v
	return s
}

func (s *InsureOrderApplyResponseBodyModuleInsOrderPolicyList) Validate() error {
	return dara.Validate(s)
}
