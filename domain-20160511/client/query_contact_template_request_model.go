// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContactTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuditStatus(v string) *QueryContactTemplateRequest
	GetAuditStatus() *string
	SetCCompany(v string) *QueryContactTemplateRequest
	GetCCompany() *string
	SetContactTemplateId(v int64) *QueryContactTemplateRequest
	GetContactTemplateId() *int64
	SetDefaultTemplate(v bool) *QueryContactTemplateRequest
	GetDefaultTemplate() *bool
	SetECompany(v string) *QueryContactTemplateRequest
	GetECompany() *string
	SetLang(v string) *QueryContactTemplateRequest
	GetLang() *string
	SetPageNum(v int32) *QueryContactTemplateRequest
	GetPageNum() *int32
	SetPageSize(v int32) *QueryContactTemplateRequest
	GetPageSize() *int32
	SetRegType(v string) *QueryContactTemplateRequest
	GetRegType() *string
	SetUserClientIp(v string) *QueryContactTemplateRequest
	GetUserClientIp() *string
}

type QueryContactTemplateRequest struct {
	AuditStatus       *string `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	CCompany          *string `json:"CCompany,omitempty" xml:"CCompany,omitempty"`
	ContactTemplateId *int64  `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	DefaultTemplate   *bool   `json:"DefaultTemplate,omitempty" xml:"DefaultTemplate,omitempty"`
	ECompany          *string `json:"ECompany,omitempty" xml:"ECompany,omitempty"`
	Lang              *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	PageNum           *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize          *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegType           *string `json:"RegType,omitempty" xml:"RegType,omitempty"`
	UserClientIp      *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryContactTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryContactTemplateRequest) GoString() string {
	return s.String()
}

func (s *QueryContactTemplateRequest) GetAuditStatus() *string {
	return s.AuditStatus
}

func (s *QueryContactTemplateRequest) GetCCompany() *string {
	return s.CCompany
}

func (s *QueryContactTemplateRequest) GetContactTemplateId() *int64 {
	return s.ContactTemplateId
}

func (s *QueryContactTemplateRequest) GetDefaultTemplate() *bool {
	return s.DefaultTemplate
}

func (s *QueryContactTemplateRequest) GetECompany() *string {
	return s.ECompany
}

func (s *QueryContactTemplateRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryContactTemplateRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *QueryContactTemplateRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryContactTemplateRequest) GetRegType() *string {
	return s.RegType
}

func (s *QueryContactTemplateRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryContactTemplateRequest) SetAuditStatus(v string) *QueryContactTemplateRequest {
	s.AuditStatus = &v
	return s
}

func (s *QueryContactTemplateRequest) SetCCompany(v string) *QueryContactTemplateRequest {
	s.CCompany = &v
	return s
}

func (s *QueryContactTemplateRequest) SetContactTemplateId(v int64) *QueryContactTemplateRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *QueryContactTemplateRequest) SetDefaultTemplate(v bool) *QueryContactTemplateRequest {
	s.DefaultTemplate = &v
	return s
}

func (s *QueryContactTemplateRequest) SetECompany(v string) *QueryContactTemplateRequest {
	s.ECompany = &v
	return s
}

func (s *QueryContactTemplateRequest) SetLang(v string) *QueryContactTemplateRequest {
	s.Lang = &v
	return s
}

func (s *QueryContactTemplateRequest) SetPageNum(v int32) *QueryContactTemplateRequest {
	s.PageNum = &v
	return s
}

func (s *QueryContactTemplateRequest) SetPageSize(v int32) *QueryContactTemplateRequest {
	s.PageSize = &v
	return s
}

func (s *QueryContactTemplateRequest) SetRegType(v string) *QueryContactTemplateRequest {
	s.RegType = &v
	return s
}

func (s *QueryContactTemplateRequest) SetUserClientIp(v string) *QueryContactTemplateRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryContactTemplateRequest) Validate() error {
	return dara.Validate(s)
}
