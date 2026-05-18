// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromType(v int32) *CreateTemplateRequest
	GetFromType() *int32
	SetOwnerId(v int64) *CreateTemplateRequest
	GetOwnerId() *int64
	SetRemark(v string) *CreateTemplateRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateTemplateRequest
	GetResourceOwnerId() *int64
	SetSmsContent(v string) *CreateTemplateRequest
	GetSmsContent() *string
	SetSmsType(v int32) *CreateTemplateRequest
	GetSmsType() *int32
	SetTemplateName(v string) *CreateTemplateRequest
	GetTemplateName() *string
	SetTemplateNickName(v string) *CreateTemplateRequest
	GetTemplateNickName() *string
	SetTemplateSubject(v string) *CreateTemplateRequest
	GetTemplateSubject() *string
	SetTemplateText(v string) *CreateTemplateRequest
	GetTemplateText() *string
	SetTemplateType(v int32) *CreateTemplateRequest
	GetTemplateType() *int32
}

type CreateTemplateRequest struct {
	FromType             *int32  `json:"FromType,omitempty" xml:"FromType,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SmsContent           *string `json:"SmsContent,omitempty" xml:"SmsContent,omitempty"`
	SmsType              *int32  `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 001 Condom Manufacturer J48
	TemplateName     *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	TemplateNickName *string `json:"TemplateNickName,omitempty" xml:"TemplateNickName,omitempty"`
	TemplateSubject  *string `json:"TemplateSubject,omitempty" xml:"TemplateSubject,omitempty"`
	TemplateText     *string `json:"TemplateText,omitempty" xml:"TemplateText,omitempty"`
	// example:
	//
	// 0
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s CreateTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateTemplateRequest) GetFromType() *int32 {
	return s.FromType
}

func (s *CreateTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateTemplateRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateTemplateRequest) GetSmsContent() *string {
	return s.SmsContent
}

func (s *CreateTemplateRequest) GetSmsType() *int32 {
	return s.SmsType
}

func (s *CreateTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateTemplateRequest) GetTemplateNickName() *string {
	return s.TemplateNickName
}

func (s *CreateTemplateRequest) GetTemplateSubject() *string {
	return s.TemplateSubject
}

func (s *CreateTemplateRequest) GetTemplateText() *string {
	return s.TemplateText
}

func (s *CreateTemplateRequest) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *CreateTemplateRequest) SetFromType(v int32) *CreateTemplateRequest {
	s.FromType = &v
	return s
}

func (s *CreateTemplateRequest) SetOwnerId(v int64) *CreateTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateTemplateRequest) SetRemark(v string) *CreateTemplateRequest {
	s.Remark = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceOwnerAccount(v string) *CreateTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateTemplateRequest) SetResourceOwnerId(v int64) *CreateTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateTemplateRequest) SetSmsContent(v string) *CreateTemplateRequest {
	s.SmsContent = &v
	return s
}

func (s *CreateTemplateRequest) SetSmsType(v int32) *CreateTemplateRequest {
	s.SmsType = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateName(v string) *CreateTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateNickName(v string) *CreateTemplateRequest {
	s.TemplateNickName = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateSubject(v string) *CreateTemplateRequest {
	s.TemplateSubject = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateText(v string) *CreateTemplateRequest {
	s.TemplateText = &v
	return s
}

func (s *CreateTemplateRequest) SetTemplateType(v int32) *CreateTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateTemplateRequest) Validate() error {
	return dara.Validate(s)
}
