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
	// The description of the SMS template application. The description cannot exceed 100 characters in length.
	//
	// This information helps reviewers understand your business scenarios and improves review efficiency. Guidelines:
	//
	// - Provide the use case of your live business.
	//
	// - Provide SMS examples for real scenarios to reflect your business scenarios.
	//
	// - Provide variable values and describe in detail the business use case and the reason for choosing the variable attributes.
	//
	// - Provide the website URL, registered domain name, or application marketplace download link of the actual business.
	//
	// - For logon scenarios, provide the test account and password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 手机注册登录
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The code of the SMS template that failed the review.
	//
	// - Call the [QuerySmsTemplateList](https://help.aliyun.com/document_detail/419288.html) operation to obtain the code of the SMS template that failed the review.
	//
	// - View the code of the SMS template that failed the review on the [Templates](https://dysms.console.aliyun.com/domestic/text/template) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_15255****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The template content. The content cannot exceed 500 characters in length.
	//
	// The template content and variable content must comply with the [SMS template specifications](https://help.aliyun.com/document_detail/463161.html). Otherwise, the template fails the review. You can view common template examples on the [Apply for Template](https://dysms.console.aliyun.com/domestic/text/template/add) page. Using sample templates can improve review efficiency and success rate. For variable specifications, see [TemplateContent variable specifications](https://help.aliyun.com/document_detail/2806243.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 您正在申请手机注册，验证码为：${code}，5分钟内有效！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The template name. The name must be 1 to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 验证码模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The SMS type. Valid values:
	//
	// - **0**: verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: promotional message.
	//
	// - **3**: international or Hong Kong, Macao, and Taiwan message.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
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
