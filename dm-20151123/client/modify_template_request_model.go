// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromType(v int32) *ModifyTemplateRequest
	GetFromType() *int32
	SetOwnerId(v int64) *ModifyTemplateRequest
	GetOwnerId() *int64
	SetRemark(v string) *ModifyTemplateRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *ModifyTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyTemplateRequest
	GetResourceOwnerId() *int64
	SetSmsContent(v string) *ModifyTemplateRequest
	GetSmsContent() *string
	SetSmsType(v int32) *ModifyTemplateRequest
	GetSmsType() *int32
	SetTemplateId(v int32) *ModifyTemplateRequest
	GetTemplateId() *int32
	SetTemplateName(v string) *ModifyTemplateRequest
	GetTemplateName() *string
	SetTemplateNickName(v string) *ModifyTemplateRequest
	GetTemplateNickName() *string
	SetTemplateSubject(v string) *ModifyTemplateRequest
	GetTemplateSubject() *string
	SetTemplateText(v string) *ModifyTemplateRequest
	GetTemplateText() *string
}

type ModifyTemplateRequest struct {
	// The source channel through which the user connects. Default value: 1.
	//
	// example:
	//
	// 1
	FromType *int32 `json:"FromType,omitempty" xml:"FromType,omitempty"`
	OwnerId  *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The remarks or application description for the SMS template. This parameter is required only when the templatetype is SMS. Maximum length: 100 characters.
	//
	// example:
	//
	// for verification
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The body content of the SMS template. This parameter is required only when the templatetype is SMS. Length: 2 to 400 characters.
	//
	// example:
	//
	// <p>hello {name}</p>
	SmsContent *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	// The business type of the SMS template. This parameter is required only when the templatetype is SMS.
	//
	// example:
	//
	// 0
	SmsType *int32 `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 409481
	TemplateId *int32 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The template name. Maximum length: 30 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// verification code
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The nickname of the template or the alias of the sender. This parameter is required only when the templatetype is email. Maximum length: 30 characters.
	//
	// example:
	//
	// Lisa Gao
	TemplateNickName *string `json:"TemplateNickName,omitempty" xml:"TemplateNickName,omitempty"`
	// The subject of the email template. This parameter is required only when the templatetype is email. Maximum length: 256 characters.
	//
	// example:
	//
	// REAL\\"EN OPEN TONNAGE
	TemplateSubject *string `json:"TemplateSubject,omitempty" xml:"TemplateSubject,omitempty"`
	// The body content of the email. This parameter is required only when the templatetype is email. Maximum size: 1 MB.
	//
	// example:
	//
	// <p>hello {name}</p>
	TemplateText *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
}

func (s ModifyTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifyTemplateRequest) GetFromType() *int32 {
	return s.FromType
}

func (s *ModifyTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyTemplateRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifyTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyTemplateRequest) GetSmsContent() *string {
	return s.SmsContent
}

func (s *ModifyTemplateRequest) GetSmsType() *int32 {
	return s.SmsType
}

func (s *ModifyTemplateRequest) GetTemplateId() *int32 {
	return s.TemplateId
}

func (s *ModifyTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifyTemplateRequest) GetTemplateNickName() *string {
	return s.TemplateNickName
}

func (s *ModifyTemplateRequest) GetTemplateSubject() *string {
	return s.TemplateSubject
}

func (s *ModifyTemplateRequest) GetTemplateText() *string {
	return s.TemplateText
}

func (s *ModifyTemplateRequest) SetFromType(v int32) *ModifyTemplateRequest {
	s.FromType = &v
	return s
}

func (s *ModifyTemplateRequest) SetOwnerId(v int64) *ModifyTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyTemplateRequest) SetRemark(v string) *ModifyTemplateRequest {
	s.Remark = &v
	return s
}

func (s *ModifyTemplateRequest) SetResourceOwnerAccount(v string) *ModifyTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyTemplateRequest) SetResourceOwnerId(v int64) *ModifyTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyTemplateRequest) SetSmsContent(v string) *ModifyTemplateRequest {
	s.SmsContent = &v
	return s
}

func (s *ModifyTemplateRequest) SetSmsType(v int32) *ModifyTemplateRequest {
	s.SmsType = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateId(v int32) *ModifyTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateName(v string) *ModifyTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateNickName(v string) *ModifyTemplateRequest {
	s.TemplateNickName = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateSubject(v string) *ModifyTemplateRequest {
	s.TemplateSubject = &v
	return s
}

func (s *ModifyTemplateRequest) SetTemplateText(v string) *ModifyTemplateRequest {
	s.TemplateText = &v
	return s
}

func (s *ModifyTemplateRequest) Validate() error {
	return dara.Validate(s)
}
