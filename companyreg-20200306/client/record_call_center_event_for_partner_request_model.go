// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecordCallCenterEventForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *RecordCallCenterEventForPartnerRequest
	GetBizId() *string
	SetBizType(v string) *RecordCallCenterEventForPartnerRequest
	GetBizType() *string
	SetCallAction(v string) *RecordCallCenterEventForPartnerRequest
	GetCallAction() *string
	SetCallee(v string) *RecordCallCenterEventForPartnerRequest
	GetCallee() *string
	SetCaller(v string) *RecordCallCenterEventForPartnerRequest
	GetCaller() *string
	SetConnId(v string) *RecordCallCenterEventForPartnerRequest
	GetConnId() *string
	SetContactId(v string) *RecordCallCenterEventForPartnerRequest
	GetContactId() *string
	SetEmployeeCode(v string) *RecordCallCenterEventForPartnerRequest
	GetEmployeeCode() *string
	SetJobId(v string) *RecordCallCenterEventForPartnerRequest
	GetJobId() *string
	SetRelatedId(v int64) *RecordCallCenterEventForPartnerRequest
	GetRelatedId() *int64
	SetSecretMobile(v string) *RecordCallCenterEventForPartnerRequest
	GetSecretMobile() *string
	SetSkillType(v int32) *RecordCallCenterEventForPartnerRequest
	GetSkillType() *int32
	SetTenantId(v string) *RecordCallCenterEventForPartnerRequest
	GetTenantId() *string
}

type RecordCallCenterEventForPartnerRequest struct {
	// example:
	//
	// P20211117141528000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
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
	// 18578602216
	Callee *string `json:"Callee,omitempty" xml:"Callee,omitempty"`
	// example:
	//
	// 13162828888
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// example:
	//
	// 12313
	ConnId *string `json:"ConnId,omitempty" xml:"ConnId,omitempty"`
	// example:
	//
	// 897265
	ContactId    *string `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	EmployeeCode *string `json:"EmployeeCode,omitempty" xml:"EmployeeCode,omitempty"`
	// example:
	//
	// ufbo502ma94m480
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// RequestId
	//
	// example:
	//
	// 1000030990004
	RelatedId *int64 `json:"RelatedId,omitempty" xml:"RelatedId,omitempty"`
	// example:
	//
	// MTAFA/DF#
	SecretMobile *string `json:"SecretMobile,omitempty" xml:"SecretMobile,omitempty"`
	// example:
	//
	// 1
	SkillType *int32  `json:"SkillType,omitempty" xml:"SkillType,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s RecordCallCenterEventForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s RecordCallCenterEventForPartnerRequest) GoString() string {
	return s.String()
}

func (s *RecordCallCenterEventForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *RecordCallCenterEventForPartnerRequest) GetBizType() *string {
	return s.BizType
}

func (s *RecordCallCenterEventForPartnerRequest) GetCallAction() *string {
	return s.CallAction
}

func (s *RecordCallCenterEventForPartnerRequest) GetCallee() *string {
	return s.Callee
}

func (s *RecordCallCenterEventForPartnerRequest) GetCaller() *string {
	return s.Caller
}

func (s *RecordCallCenterEventForPartnerRequest) GetConnId() *string {
	return s.ConnId
}

func (s *RecordCallCenterEventForPartnerRequest) GetContactId() *string {
	return s.ContactId
}

func (s *RecordCallCenterEventForPartnerRequest) GetEmployeeCode() *string {
	return s.EmployeeCode
}

func (s *RecordCallCenterEventForPartnerRequest) GetJobId() *string {
	return s.JobId
}

func (s *RecordCallCenterEventForPartnerRequest) GetRelatedId() *int64 {
	return s.RelatedId
}

func (s *RecordCallCenterEventForPartnerRequest) GetSecretMobile() *string {
	return s.SecretMobile
}

func (s *RecordCallCenterEventForPartnerRequest) GetSkillType() *int32 {
	return s.SkillType
}

func (s *RecordCallCenterEventForPartnerRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RecordCallCenterEventForPartnerRequest) SetBizId(v string) *RecordCallCenterEventForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetBizType(v string) *RecordCallCenterEventForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetCallAction(v string) *RecordCallCenterEventForPartnerRequest {
	s.CallAction = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetCallee(v string) *RecordCallCenterEventForPartnerRequest {
	s.Callee = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetCaller(v string) *RecordCallCenterEventForPartnerRequest {
	s.Caller = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetConnId(v string) *RecordCallCenterEventForPartnerRequest {
	s.ConnId = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetContactId(v string) *RecordCallCenterEventForPartnerRequest {
	s.ContactId = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetEmployeeCode(v string) *RecordCallCenterEventForPartnerRequest {
	s.EmployeeCode = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetJobId(v string) *RecordCallCenterEventForPartnerRequest {
	s.JobId = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetRelatedId(v int64) *RecordCallCenterEventForPartnerRequest {
	s.RelatedId = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetSecretMobile(v string) *RecordCallCenterEventForPartnerRequest {
	s.SecretMobile = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetSkillType(v int32) *RecordCallCenterEventForPartnerRequest {
	s.SkillType = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) SetTenantId(v string) *RecordCallCenterEventForPartnerRequest {
	s.TenantId = &v
	return s
}

func (s *RecordCallCenterEventForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
