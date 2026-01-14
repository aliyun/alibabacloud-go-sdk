// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSmsTemplateCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SmsTemplateCreateRequest
	GetContent() *string
	SetOwnerId(v int64) *SmsTemplateCreateRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SmsTemplateCreateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SmsTemplateCreateRequest
	GetResourceOwnerId() *int64
	SetSign(v string) *SmsTemplateCreateRequest
	GetSign() *string
	SetSmsType(v int64) *SmsTemplateCreateRequest
	GetSmsType() *int64
	SetTemplateName(v string) *SmsTemplateCreateRequest
	GetTemplateName() *string
	SetTemplateType(v int64) *SmsTemplateCreateRequest
	GetTemplateType() *int64
}

type SmsTemplateCreateRequest struct {
	// 短信内容
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值示例值
	Content              *string `json:"Content,omitempty" xml:"Content,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// 短信签名
	//
	// This parameter is required.
	//
	// example:
	//
	// ef2i29fsljf
	Sign *string `json:"Sign,omitempty" xml:"Sign,omitempty"`
	// 短信类型
	//
	// example:
	//
	// 73
	SmsType *int64 `json:"SmsType,omitempty" xml:"SmsType,omitempty"`
	// 模板名称
	//
	// This parameter is required.
	//
	// example:
	//
	// 示例值示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板类型
	//
	// example:
	//
	// 56
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s SmsTemplateCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s SmsTemplateCreateRequest) GoString() string {
	return s.String()
}

func (s *SmsTemplateCreateRequest) GetContent() *string {
	return s.Content
}

func (s *SmsTemplateCreateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SmsTemplateCreateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SmsTemplateCreateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SmsTemplateCreateRequest) GetSign() *string {
	return s.Sign
}

func (s *SmsTemplateCreateRequest) GetSmsType() *int64 {
	return s.SmsType
}

func (s *SmsTemplateCreateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SmsTemplateCreateRequest) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *SmsTemplateCreateRequest) SetContent(v string) *SmsTemplateCreateRequest {
	s.Content = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetOwnerId(v int64) *SmsTemplateCreateRequest {
	s.OwnerId = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetResourceOwnerAccount(v string) *SmsTemplateCreateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetResourceOwnerId(v int64) *SmsTemplateCreateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetSign(v string) *SmsTemplateCreateRequest {
	s.Sign = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetSmsType(v int64) *SmsTemplateCreateRequest {
	s.SmsType = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetTemplateName(v string) *SmsTemplateCreateRequest {
	s.TemplateName = &v
	return s
}

func (s *SmsTemplateCreateRequest) SetTemplateType(v int64) *SmsTemplateCreateRequest {
	s.TemplateType = &v
	return s
}

func (s *SmsTemplateCreateRequest) Validate() error {
	return dara.Validate(s)
}
