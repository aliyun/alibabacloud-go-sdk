// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplySceneContent(v string) *UpdateSmsTemplateShrinkRequest
	GetApplySceneContent() *string
	SetIntlType(v int32) *UpdateSmsTemplateShrinkRequest
	GetIntlType() *int32
	SetMoreDataShrink(v string) *UpdateSmsTemplateShrinkRequest
	GetMoreDataShrink() *string
	SetOwnerId(v int64) *UpdateSmsTemplateShrinkRequest
	GetOwnerId() *int64
	SetRelatedSignName(v string) *UpdateSmsTemplateShrinkRequest
	GetRelatedSignName() *string
	SetRemark(v string) *UpdateSmsTemplateShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateSmsTemplateShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmsTemplateShrinkRequest
	GetResourceOwnerId() *int64
	SetTemplateCode(v string) *UpdateSmsTemplateShrinkRequest
	GetTemplateCode() *string
	SetTemplateContent(v string) *UpdateSmsTemplateShrinkRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *UpdateSmsTemplateShrinkRequest
	GetTemplateName() *string
	SetTemplateRule(v string) *UpdateSmsTemplateShrinkRequest
	GetTemplateRule() *string
	SetTemplateType(v int32) *UpdateSmsTemplateShrinkRequest
	GetTemplateType() *int32
	SetTrafficDriving(v string) *UpdateSmsTemplateShrinkRequest
	GetTrafficDriving() *string
}

type UpdateSmsTemplateShrinkRequest struct {
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
	MoreDataShrink *string `json:"MoreData,omitempty" xml:"MoreData,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s UpdateSmsTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsTemplateShrinkRequest) GetApplySceneContent() *string {
	return s.ApplySceneContent
}

func (s *UpdateSmsTemplateShrinkRequest) GetIntlType() *int32 {
	return s.IntlType
}

func (s *UpdateSmsTemplateShrinkRequest) GetMoreDataShrink() *string {
	return s.MoreDataShrink
}

func (s *UpdateSmsTemplateShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmsTemplateShrinkRequest) GetRelatedSignName() *string {
	return s.RelatedSignName
}

func (s *UpdateSmsTemplateShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateSmsTemplateShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmsTemplateShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmsTemplateShrinkRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *UpdateSmsTemplateShrinkRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *UpdateSmsTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateSmsTemplateShrinkRequest) GetTemplateRule() *string {
	return s.TemplateRule
}

func (s *UpdateSmsTemplateShrinkRequest) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *UpdateSmsTemplateShrinkRequest) GetTrafficDriving() *string {
	return s.TrafficDriving
}

func (s *UpdateSmsTemplateShrinkRequest) SetApplySceneContent(v string) *UpdateSmsTemplateShrinkRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetIntlType(v int32) *UpdateSmsTemplateShrinkRequest {
	s.IntlType = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetMoreDataShrink(v string) *UpdateSmsTemplateShrinkRequest {
	s.MoreDataShrink = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetOwnerId(v int64) *UpdateSmsTemplateShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetRelatedSignName(v string) *UpdateSmsTemplateShrinkRequest {
	s.RelatedSignName = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetRemark(v string) *UpdateSmsTemplateShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetResourceOwnerAccount(v string) *UpdateSmsTemplateShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetResourceOwnerId(v int64) *UpdateSmsTemplateShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateCode(v string) *UpdateSmsTemplateShrinkRequest {
	s.TemplateCode = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateContent(v string) *UpdateSmsTemplateShrinkRequest {
	s.TemplateContent = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateName(v string) *UpdateSmsTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateRule(v string) *UpdateSmsTemplateShrinkRequest {
	s.TemplateRule = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTemplateType(v int32) *UpdateSmsTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) SetTrafficDriving(v string) *UpdateSmsTemplateShrinkRequest {
	s.TrafficDriving = &v
	return s
}

func (s *UpdateSmsTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
