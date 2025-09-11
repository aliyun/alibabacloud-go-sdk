// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *ModifySmsTemplateRequest
	GetOwnerId() *int64
	SetRemark(v string) *ModifySmsTemplateRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *ModifySmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySmsTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateCode(v string) *ModifySmsTemplateRequest
	GetTemplateCode() *string
	SetTemplateContent(v string) *ModifySmsTemplateRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *ModifySmsTemplateRequest
	GetTemplateName() *string
	SetTemplateType(v int32) *ModifySmsTemplateRequest
	GetTemplateType() *int32
}

type ModifySmsTemplateRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the message template. It is one of the reference information for template review. The description cannot exceed 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Modify the parameters of the template.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The code of the message template.
	//
	// You can log on to the [Short Message Service (SMS) console](https://dysms.console.aliyun.com/dysms.htm), click **Go China*	- or **Go Globe*	- in the left-side navigation pane, and then view the template code on the **Templates*	- tab. You can also call the [AddSmsTemplate](https://help.aliyun.com/document_detail/121208.html) operation to obtain the template code.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_15255****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The content of the template. The content must be 1 to 500 characters in length.
	//
	// > When you modify a template, design the template content based on the review comments.
	//
	// This parameter is required.
	//
	// example:
	//
	// You are applying for mobile registration. The verification code is: ${code}, valid for 5 minutes!
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the template. The name must be 1 to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun verification code
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The type of the message. Valid values:
	//
	// 	- **0**: verification code
	//
	// 	- **1**: text message
	//
	// 	- **2**: promotional message
	//
	// 	- **3**: message sent to countries or regions outside the Chinese mainland
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s ModifySmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *ModifySmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySmsTemplateRequest) GetRemark() *string {
	return s.Remark
}

func (s *ModifySmsTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySmsTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySmsTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ModifySmsTemplateRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *ModifySmsTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ModifySmsTemplateRequest) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *ModifySmsTemplateRequest) SetOwnerId(v int64) *ModifySmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetRemark(v string) *ModifySmsTemplateRequest {
	s.Remark = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetResourceOwnerAccount(v string) *ModifySmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetResourceOwnerId(v int64) *ModifySmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateCode(v string) *ModifySmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateContent(v string) *ModifySmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateName(v string) *ModifySmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *ModifySmsTemplateRequest) SetTemplateType(v int32) *ModifySmsTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *ModifySmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
