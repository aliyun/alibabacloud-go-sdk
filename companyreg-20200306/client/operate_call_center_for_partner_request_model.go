// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCallCenterForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *OperateCallCenterForPartnerRequest
	GetBizType() *string
	SetCallAction(v string) *OperateCallCenterForPartnerRequest
	GetCallAction() *string
	SetEmployeeCode(v string) *OperateCallCenterForPartnerRequest
	GetEmployeeCode() *string
	SetRequest(v string) *OperateCallCenterForPartnerRequest
	GetRequest() *string
	SetTenantId(v string) *OperateCallCenterForPartnerRequest
	GetTenantId() *string
}

type OperateCallCenterForPartnerRequest struct {
	// example:
	//
	// esp.beian_assist
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// MakeCall
	CallAction *string `json:"CallAction,omitempty" xml:"CallAction,omitempty"`
	// example:
	//
	// 12323213
	EmployeeCode *string `json:"EmployeeCode,omitempty" xml:"EmployeeCode,omitempty"`
	// example:
	//
	// {}
	Request *string `json:"Request,omitempty" xml:"Request,omitempty"`
	// example:
	//
	// t4uor8evmq9nk
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s OperateCallCenterForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateCallCenterForPartnerRequest) GoString() string {
	return s.String()
}

func (s *OperateCallCenterForPartnerRequest) GetBizType() *string {
	return s.BizType
}

func (s *OperateCallCenterForPartnerRequest) GetCallAction() *string {
	return s.CallAction
}

func (s *OperateCallCenterForPartnerRequest) GetEmployeeCode() *string {
	return s.EmployeeCode
}

func (s *OperateCallCenterForPartnerRequest) GetRequest() *string {
	return s.Request
}

func (s *OperateCallCenterForPartnerRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *OperateCallCenterForPartnerRequest) SetBizType(v string) *OperateCallCenterForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *OperateCallCenterForPartnerRequest) SetCallAction(v string) *OperateCallCenterForPartnerRequest {
	s.CallAction = &v
	return s
}

func (s *OperateCallCenterForPartnerRequest) SetEmployeeCode(v string) *OperateCallCenterForPartnerRequest {
	s.EmployeeCode = &v
	return s
}

func (s *OperateCallCenterForPartnerRequest) SetRequest(v string) *OperateCallCenterForPartnerRequest {
	s.Request = &v
	return s
}

func (s *OperateCallCenterForPartnerRequest) SetTenantId(v string) *OperateCallCenterForPartnerRequest {
	s.TenantId = &v
	return s
}

func (s *OperateCallCenterForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
