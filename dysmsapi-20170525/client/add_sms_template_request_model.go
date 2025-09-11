// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *AddSmsTemplateRequest
	GetOwnerId() *int64
	SetRemark(v string) *AddSmsTemplateRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *AddSmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddSmsTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateContent(v string) *AddSmsTemplateRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *AddSmsTemplateRequest
	GetTemplateName() *string
	SetTemplateType(v int32) *AddSmsTemplateRequest
	GetTemplateType() *int32
}

type AddSmsTemplateRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the message template. It is one of the reference information for template review. The description cannot exceed 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Apply for a template to send verification codes.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The content of the template. The content can be up to 500 characters in length. For more information, see [Message template specifications](https://help.aliyun.com/document_detail/108253.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// You are applying for mobile registration. The verification code is: ${code}, valid for 5 minutes!
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template. The name can be up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun Test
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: notification
	//
	// 	- **2**: promotional message
	//
	// 	- **3**: message sent to countries or regions outside the Chinese mainland
	//
	// > Only enterprise users can send promotional messages, or send messages to countries or regions outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s AddSmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s AddSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *AddSmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddSmsTemplateRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddSmsTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddSmsTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddSmsTemplateRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *AddSmsTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *AddSmsTemplateRequest) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *AddSmsTemplateRequest) SetOwnerId(v int64) *AddSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *AddSmsTemplateRequest) SetRemark(v string) *AddSmsTemplateRequest {
	s.Remark = &v
	return s
}

func (s *AddSmsTemplateRequest) SetResourceOwnerAccount(v string) *AddSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddSmsTemplateRequest) SetResourceOwnerId(v int64) *AddSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddSmsTemplateRequest) SetTemplateContent(v string) *AddSmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *AddSmsTemplateRequest) SetTemplateName(v string) *AddSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *AddSmsTemplateRequest) SetTemplateType(v int32) *AddSmsTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *AddSmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
