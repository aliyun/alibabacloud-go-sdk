// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplySceneContent(v string) *CreateSmsTemplateRequest
	GetApplySceneContent() *string
	SetIntlType(v int32) *CreateSmsTemplateRequest
	GetIntlType() *int32
	SetMoreData(v []*string) *CreateSmsTemplateRequest
	GetMoreData() []*string
	SetOwnerId(v int64) *CreateSmsTemplateRequest
	GetOwnerId() *int64
	SetRelatedSignName(v string) *CreateSmsTemplateRequest
	GetRelatedSignName() *string
	SetRemark(v string) *CreateSmsTemplateRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateSmsTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsTemplateRequest
	GetResourceOwnerId() *int64
	SetTemplateContent(v string) *CreateSmsTemplateRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *CreateSmsTemplateRequest
	GetTemplateName() *string
	SetTemplateRule(v string) *CreateSmsTemplateRequest
	GetTemplateRule() *string
	SetTemplateType(v int32) *CreateSmsTemplateRequest
	GetTemplateType() *int32
	SetTrafficDriving(v string) *CreateSmsTemplateRequest
	GetTrafficDriving() *string
}

type CreateSmsTemplateRequest struct {
	// If there is an applicable scenario, you can fill it in.
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
	// - **2**: Promotional message.
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// Additional materials you can upload, such as business proof documents or screenshots, to help reviewers understand your business details.
	//
	// This parameter is optional; please fill it in according to actual needs.
	MoreData []*string `json:"MoreData,omitempty" xml:"MoreData,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The signature name that the template needs to be associated with. The associated SMS signature must have passed the review.
	//
	// This parameter is mandatory when the TemplateType parameter is **0**, **1**, or **2**.
	//
	// <notice>Associating a signature can expedite the review process. Note that this associated signature is unrelated to the signature selected when sending SMS messages.</notice>
	//
	// example:
	//
	// Aliyun
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// Please describe the business scenario where you use SMS or provide an online link to the scenario, along with a complete example of the SMS (with variable contents filled), as complete information helps increase the template approval rate. Failure to follow guidelines or leaving this field blank may affect the approval of your template.
	//
	// example:
	//
	// Request verification code SMS.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Template content, up to 500 characters in length.
	//
	// Both the template content and variable content must comply with SMS specifications; otherwise, the template will fail the review. You can also view common template examples on the template application page. Using sample templates can enhance review efficiency and success rates. For variable specifications, see [TemplateContent Variable Parameter Filling Specifications](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example).
	//
	// This parameter is required.
	//
	// example:
	//
	// You are applying for mobile registration. The verification code is: ${code}. It is valid for 5 minutes!
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// Template name, up to 30 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyunCode
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// Template variable rules.
	//
	// For filling in variable rules, refer to the [Sample Documentation](https://help.aliyun.com/zh/sms/templaterule-template-variable-parameter-filling-example).
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
	// - **2**: Promotional message.
	//
	// - **3**: International/Hong Kong, Macao, and Taiwan messages.
	//
	// > Only enterprise-verified users can apply for promotional messages and international/Hong Kong, Macao, and Taiwan messages. For details on the differences between personal and enterprise user rights, please refer to [Usage Instructions](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm=a2c4g.11186623.0.0.67447f576NJnE8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TemplateType   *int32  `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	TrafficDriving *string `json:"TrafficDriving,omitempty" xml:"TrafficDriving,omitempty"`
}

func (s CreateSmsTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateRequest) GetApplySceneContent() *string {
	return s.ApplySceneContent
}

func (s *CreateSmsTemplateRequest) GetIntlType() *int32 {
	return s.IntlType
}

func (s *CreateSmsTemplateRequest) GetMoreData() []*string {
	return s.MoreData
}

func (s *CreateSmsTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsTemplateRequest) GetRelatedSignName() *string {
	return s.RelatedSignName
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

func (s *CreateSmsTemplateRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *CreateSmsTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateSmsTemplateRequest) GetTemplateRule() *string {
	return s.TemplateRule
}

func (s *CreateSmsTemplateRequest) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *CreateSmsTemplateRequest) GetTrafficDriving() *string {
	return s.TrafficDriving
}

func (s *CreateSmsTemplateRequest) SetApplySceneContent(v string) *CreateSmsTemplateRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetIntlType(v int32) *CreateSmsTemplateRequest {
	s.IntlType = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetMoreData(v []*string) *CreateSmsTemplateRequest {
	s.MoreData = v
	return s
}

func (s *CreateSmsTemplateRequest) SetOwnerId(v int64) *CreateSmsTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetRelatedSignName(v string) *CreateSmsTemplateRequest {
	s.RelatedSignName = &v
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

func (s *CreateSmsTemplateRequest) SetTemplateContent(v string) *CreateSmsTemplateRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateName(v string) *CreateSmsTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateRule(v string) *CreateSmsTemplateRequest {
	s.TemplateRule = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTemplateType(v int32) *CreateSmsTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateSmsTemplateRequest) SetTrafficDriving(v string) *CreateSmsTemplateRequest {
	s.TrafficDriving = &v
	return s
}

func (s *CreateSmsTemplateRequest) Validate() error {
	return dara.Validate(s)
}
