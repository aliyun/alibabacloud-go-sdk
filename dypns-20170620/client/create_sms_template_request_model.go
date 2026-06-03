// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizUrl(v string) *CreateSmsTemplateRequest
	GetBizUrl() *string
	SetBusinessType(v string) *CreateSmsTemplateRequest
	GetBusinessType() *string
	SetOwnerId(v int64) *CreateSmsTemplateRequest
	GetOwnerId() *int64
	SetProdCode(v string) *CreateSmsTemplateRequest
	GetProdCode() *string
	SetRemark(v string) *CreateSmsTemplateRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateSmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsTemplateRequest
	GetResourceOwnerId() *int64
	SetSmsTemplateContent(v string) *CreateSmsTemplateRequest
	GetSmsTemplateContent() *string
	SetSmsTemplateName(v string) *CreateSmsTemplateRequest
	GetSmsTemplateName() *string
	SetSmsTemplateRule(v string) *CreateSmsTemplateRequest
	GetSmsTemplateRule() *string
}

type CreateSmsTemplateRequest struct {
	// This parameter is required.
	BizUrl *string `json:"BizUrl,omitempty" xml:"BizUrl,omitempty"`
	// This parameter is required.
	BusinessType         *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ProdCode             *string `json:"ProdCode,omitempty" xml:"ProdCode,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	SmsTemplateContent *string `json:"SmsTemplateContent,omitempty" xml:"SmsTemplateContent,omitempty"`
	// This parameter is required.
	SmsTemplateName *string `json:"SmsTemplateName,omitempty" xml:"SmsTemplateName,omitempty"`
	// This parameter is required.
	SmsTemplateRule *string `json:"SmsTemplateRule,omitempty" xml:"SmsTemplateRule,omitempty"`
}

func (s CreateSmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateRequest) GetBizUrl() *string {
	return s.BizUrl
}

func (s *CreateSmsTemplateRequest) GetBusinessType() *string {
	return s.BusinessType
}

func (s *CreateSmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsTemplateRequest) GetProdCode() *string {
	return s.ProdCode
}

func (s *CreateSmsTemplateRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateSmsTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsTemplateRequest) GetSmsTemplateContent() *string {
	return s.SmsTemplateContent
}

func (s *CreateSmsTemplateRequest) GetSmsTemplateName() *string {
	return s.SmsTemplateName
}

func (s *CreateSmsTemplateRequest) GetSmsTemplateRule() *string {
	return s.SmsTemplateRule
}

func (s *CreateSmsTemplateRequest) SetBizUrl(v string) *CreateSmsTemplateRequest {
	s.BizUrl = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetBusinessType(v string) *CreateSmsTemplateRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetOwnerId(v int64) *CreateSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetProdCode(v string) *CreateSmsTemplateRequest {
	s.ProdCode = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetRemark(v string) *CreateSmsTemplateRequest {
	s.Remark = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetResourceOwnerAccount(v string) *CreateSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetResourceOwnerId(v int64) *CreateSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetSmsTemplateContent(v string) *CreateSmsTemplateRequest {
	s.SmsTemplateContent = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetSmsTemplateName(v string) *CreateSmsTemplateRequest {
	s.SmsTemplateName = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetSmsTemplateRule(v string) *CreateSmsTemplateRequest {
	s.SmsTemplateRule = &v
	return s
}

func (s *CreateSmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
