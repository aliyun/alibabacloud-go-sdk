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
	// The description for the SMS template application. The length must not exceed 100 characters.
	//
	// This is reference information for template review. Providing a complete application description helps reviewers understand your business scenario and improves review efficiency. Guidelines:
	//
	// - You can provide the usage scenario of business that has been launched.
	//
	// - You can provide SMS examples for real scenarios to demonstrate your business scenario.
	//
	// - You can provide the variable parameter content, describing the business usage scenario in detail and the reason for choosing this variable attribute.
	//
	// - You can provide website links of actual business, registered domain addresses, app market download links, and so on.
	//
	// - For login scenarios, you can provide a test account and password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 申请验证码短信
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The template content. The length must not exceed 500 characters.
	//
	// The template content and variable content must comply with the [SMS template specifications](https://help.aliyun.com/document_detail/463161.html); otherwise, the template review will fail. You can also view common template examples on the [Apply for template](https://dysms.console.aliyun.com/domestic/text/template/add) page. Using example templates can improve review efficiency and success rate. For variable specifications, see [TemplateContent parameter variable specifications](https://help.aliyun.com/document_detail/2806243.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 您正在申请手机注册，验证码为：${code}，5分钟内有效！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The template name. The length must not exceed 30 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// 验证码
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The SMS type. Valid values:
	//
	// - **0**: verification code.
	//
	// - **1**: notification SMS.
	//
	// - **2**: promotional SMS.
	//
	// - **3**: international/Hong Kong, Macao, and Taiwan messages.
	//
	// > Only enterprise-verified users can apply for promotional SMS and international/Hong Kong, Macao, and Taiwan messages. For details about the differences between individual and enterprise user privileges, see [Usage notes](https://help.aliyun.com/document_detail/55324.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
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
