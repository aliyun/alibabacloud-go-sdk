// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplySceneContent(v string) *UpdateSmsTemplateRequest
	GetApplySceneContent() *string
	SetIntlType(v int32) *UpdateSmsTemplateRequest
	GetIntlType() *int32
	SetMoreData(v []*string) *UpdateSmsTemplateRequest
	GetMoreData() []*string
	SetOwnerId(v int64) *UpdateSmsTemplateRequest
	GetOwnerId() *int64
	SetRelatedSignName(v string) *UpdateSmsTemplateRequest
	GetRelatedSignName() *string
	SetRemark(v string) *UpdateSmsTemplateRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateSmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmsTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateCode(v string) *UpdateSmsTemplateRequest
	GetTemplateCode() *string
	SetTemplateContent(v string) *UpdateSmsTemplateRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *UpdateSmsTemplateRequest
	GetTemplateName() *string
	SetTemplateRule(v string) *UpdateSmsTemplateRequest
	GetTemplateRule() *string
	SetTemplateType(v int32) *UpdateSmsTemplateRequest
	GetTemplateType() *int32
	SetTrafficDriving(v string) *UpdateSmsTemplateRequest
	GetTrafficDriving() *string
}

type UpdateSmsTemplateRequest struct {
	// Application scenarios, instructions as follows:
	//
	// - For registered websites, please enter the MIIT-registered domain with HTTP or HTTPS.
	//
	// - For launched apps, provide the app store display link with HTTP or HTTPS, ensuring the app is online.
	//
	// - For public accounts or mini-programs, enter the full name of the public account or mini-program, ensuring they are online.
	//
	// - For e-commerce platform stores, applicable only to enterprise users, enter the display link of the e-commerce store with HTTP or HTTPS.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// International/Hong Kong, Macao, and Taiwan template type. When the **TemplateType*	- parameter is **3**, this parameter is required for international/Hong Kong, Macao, and Taiwan templates, with values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional SMS.
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// Additional information, such as uploading business proof documents or screenshots, to help reviewers understand your business details. Optional and can be left unset.
	MoreData []*string `json:"MoreData,omitempty" xml:"MoreData,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// SMS signature associated with the template during the application.
	//
	// example:
	//
	// 阿里云
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// Explanation for the SMS template application, which serves as a reference for template review.
	//
	// example:
	//
	// 登录场景使用验证码
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Template Code of an unapproved template.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_152550****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// Template content, up to 500 characters in length.
	//
	// Both the template content and variable content must comply with SMS regulations; otherwise, the template will fail the review. You can also view common template examples on the template application page. Using sample templates can enhance review efficiency and success rates. Variable specifications can be found in [TemplateContent Parameter Variable Specifications](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example?spm).
	//
	// This parameter is required.
	//
	// example:
	//
	// 您正在申请手机注册，验证码为：${code}，5分钟内有效！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// Template name, up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// 验证码
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template variable rules.
	//
	// For guidance on filling variable rules, refer to the [Sample Documentation](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example?spm).
	//
	// example:
	//
	// {"code":"characterWithNumber"}
	TemplateRule *string `json:"TemplateRule,omitempty" xml:"TemplateRule,omitempty"`
	// SMS type. Values:
	//
	// - **0**: Verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: Promotional SMS.
	//
	// - **3**: International/Hong Kong, Macao, and Taiwan messages.
	//
	// > Only enterprise-certified users can apply for promotional SMS and international/Hong Kong, Macao, and Taiwan messages. Details on differences between personal and enterprise user rights are available in [Usage Guidelines](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm=a2c4g.11186623.0.0.67447f576NJnE8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TemplateType   *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TrafficDriving *string `json:"TrafficDriving,omitempty" xml:"TrafficDriving,omitempty"`
}

func (s UpdateSmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateRequest) GetApplySceneContent() *string {
	return s.ApplySceneContent
}

func (s *UpdateSmsTemplateRequest) GetIntlType() *int32 {
	return s.IntlType
}

func (s *UpdateSmsTemplateRequest) GetMoreData() []*string {
	return s.MoreData
}

func (s *UpdateSmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmsTemplateRequest) GetRelatedSignName() *string {
	return s.RelatedSignName
}

func (s *UpdateSmsTemplateRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateSmsTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmsTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmsTemplateRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateSmsTemplateRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *UpdateSmsTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateSmsTemplateRequest) GetTemplateRule() *string {
	return s.TemplateRule
}

func (s *UpdateSmsTemplateRequest) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *UpdateSmsTemplateRequest) GetTrafficDriving() *string {
	return s.TrafficDriving
}

func (s *UpdateSmsTemplateRequest) SetApplySceneContent(v string) *UpdateSmsTemplateRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetIntlType(v int32) *UpdateSmsTemplateRequest {
	s.IntlType = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetMoreData(v []*string) *UpdateSmsTemplateRequest {
	s.MoreData = v
	return s
}

func (s *UpdateSmsTemplateRequest) SetOwnerId(v int64) *UpdateSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetRelatedSignName(v string) *UpdateSmsTemplateRequest {
	s.RelatedSignName = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetRemark(v string) *UpdateSmsTemplateRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetResourceOwnerAccount(v string) *UpdateSmsTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetResourceOwnerId(v int64) *UpdateSmsTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateCode(v string) *UpdateSmsTemplateRequest {
	s.TemplateCode = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateContent(v string) *UpdateSmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateName(v string) *UpdateSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateRule(v string) *UpdateSmsTemplateRequest {
	s.TemplateRule = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTemplateType(v int32) *UpdateSmsTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *UpdateSmsTemplateRequest) SetTrafficDriving(v string) *UpdateSmsTemplateRequest {
	s.TrafficDriving = &v
	return s
}

func (s *UpdateSmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
