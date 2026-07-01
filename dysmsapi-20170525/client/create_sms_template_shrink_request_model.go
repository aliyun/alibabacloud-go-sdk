// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplySceneContent(v string) *CreateSmsTemplateShrinkRequest
	GetApplySceneContent() *string
	SetIntlType(v int32) *CreateSmsTemplateShrinkRequest
	GetIntlType() *int32
	SetMoreDataShrink(v string) *CreateSmsTemplateShrinkRequest
	GetMoreDataShrink() *string
	SetOwnerId(v int64) *CreateSmsTemplateShrinkRequest
	GetOwnerId() *int64
	SetRelatedSignName(v string) *CreateSmsTemplateShrinkRequest
	GetRelatedSignName() *string
	SetRemark(v string) *CreateSmsTemplateShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateSmsTemplateShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsTemplateShrinkRequest
	GetResourceOwnerId() *int64
	SetTemplateContent(v string) *CreateSmsTemplateShrinkRequest
	GetTemplateContent() *string
	SetTemplateName(v string) *CreateSmsTemplateShrinkRequest
	GetTemplateName() *string
	SetTemplateRule(v string) *CreateSmsTemplateShrinkRequest
	GetTemplateRule() *string
	SetTemplateType(v int32) *CreateSmsTemplateShrinkRequest
	GetTemplateType() *int32
	SetTrafficDriving(v string) *CreateSmsTemplateShrinkRequest
	GetTrafficDriving() *string
}

type CreateSmsTemplateShrinkRequest struct {
	// The business scenario.
	//
	// - If the associated signature\\"s use case is "Live App", `ApplySceneContent` must be an app URL that starts with `http://` or `https://`.
	//
	// - This parameter is required if the associated signature\\"s use case is "Registered Trademark Name" or "organization name".
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// The type of the template for international/Hong Kong, Macao, and Taiwan messages. This parameter is required when **TemplateType*	- is set to **3**. Valid values:
	//
	// - **0**: notification message.
	//
	// - **1**: promotional message.
	//
	// - **2**: verification code.
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// Additional information. You can upload supporting documents or business screenshots to help reviewers better understand your business scenario. If you are applying for a promotional message template (where `TemplateType` is `2`), you must upload user authorization materials. For more information, see [Specifications for Uploading User Authorization Materials](https://help.aliyun.com/document_detail/312341.html).
	MoreDataShrink *string `json:"MoreData,omitempty" xml:"MoreData,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the signature to associate with the template. The signature must be an approved signature.
	//
	// 	Notice:
	//
	// - This parameter is required if **TemplateType*	- is set to **0**, **1**, or **2**.
	//
	// - Associating a signature can expedite the review process. The signature associated here is unrelated to the one you select when sending SMS messages.
	//
	// example:
	//
	// 验证码签名
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// Describe the business scenario for the SMS messages, or provide a URL for online scenarios. You must also provide a complete SMS example with actual values for any variables. Complete information increases the chance of template approval. Templates that do not provide this information as specified may be rejected.
	//
	// example:
	//
	// 申请验证码短信
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The template content. The content must be 500 characters or less.
	//
	// The template content and variables must comply with the [SMS Template Specifications](https://help.aliyun.com/document_detail/463161.html). Templates that do not comply may be rejected. You can find common template examples on the [Apply for Template](https://dysms.console.aliyun.com/domestic/text/template/add) page. Using these examples can speed up the review process and increase the approval rate. For variable specifications, see [Variable Specifications for the TemplateContent Parameter](https://help.aliyun.com/document_detail/2806243.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 您正在申请手机注册，验证码为：${code}，5分钟内有效！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The template name. The name must be 30 characters or less.
	//
	// This parameter is required.
	//
	// example:
	//
	// 验证码
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The rules for variables in the template. For instructions on how to define these rules, see [Sample Document](https://help.aliyun.com/document_detail/2806243.html).
	//
	// > - This parameter is required if the message template contains variables.
	//
	// example:
	//
	// {"code":"characterWithNumber2"}
	TemplateRule *string `json:"TemplateRule,omitempty" xml:"TemplateRule,omitempty"`
	// The SMS type. Valid values:
	//
	// - **0**: verification code.
	//
	// - **1**: notification message.
	//
	// - **2**: promotional message.
	//
	// - **3**: international/Hong Kong, Macao, and Taiwan messages.
	//
	// > Only enterprise-verified users can apply for promotional messages or international/Hong Kong, Macao, and Taiwan messages. For more information about the differences in privileges between individual and enterprise users, see [Usage Notes](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm=a2c4g.11186623.0.0.67447f576NJnE8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// 	Warning:
	//
	// To control the security of SMS content, messages that contain traffic-driving information, such as phone numbers and links, may be blocked by carriers, which can lead to delivery failures. To reduce this risk, we recommend that you avoid including such information in message templates.
	//
	//
	//
	// A JSON string that contains a list of traffic-driving information.
	//
	// 	Notice: The value must be a JSON array serialized into a string.
	//
	// ### 1. Fields
	//
	// {
	//
	// "trafficDrivingType":"traffic driving type",
	//
	// "trafficDrivingContent":"traffic driving content",
	//
	// "variableName":"variable name",
	//
	// "companyName":"organization name",
	//
	// "organizationCode":"unified social credit code",
	//
	// "icpNo":"ICP filing or license number",
	//
	// "icpPicOssKey":"OSS key of the ICP filing screenshot",
	//
	// "companyDifferentFromSignQuaReason":"Reason for the discrepancy between the organization name and the signature qualification"
	//
	// }
	//
	// ### 2. Notes
	//
	// - If the content is not a variable, do not pass the `variableName` parameter.
	//
	// - If the organization name is different from the one in the signature qualification, pass the `companyDifferentFromSignQuaReason` parameter.
	//
	// - If `trafficDrivingType` is set to `DOMAIN`, all parameters in this object are required.
	//
	// - If `trafficDrivingType` is set to another value, pass the `trafficDrivingType`, `trafficDrivingContent`, `variableName` (if applicable), `companyName`, `organizationCode`, and `companyDifferentFromSignQuaReason` (if applicable) parameters.
	//
	// ### 3. trafficDrivingType enum values
	//
	// 	Warning:
	//
	// Due to regulatory requirements, mobile phone numbers are not supported.
	//
	//
	//
	// - DOMAIN: A domain link.
	//
	// - FIXED_PHONE: Fixed-line phone.
	//
	// - 400_PHONE: Phone number prefixed with 400.
	//
	// - 800_PHONE: Phone number prefixed with 800.
	//
	// - 95_PHONE: Phone number prefixed with 95.
	//
	// - 96_PHONE: Phone number prefixed with 96.
	//
	// - 1_PHONE: A 3- to 8-digit phone number that starts with 1.
	//
	// - OTHER_PHONE: Other phone number.
	//
	// example:
	//
	// [{"trafficDrivingType":"DOMAIN","trafficDrivingContent":"aliyun.com","companyName":"阿里云计算有限公司","organizationCode":"91330****73959654P","icpNo":"浙B2-20****01-4","icpPicOssKey":"db7784d8-cb0c-498f-****-295f1ad6d665_mf29l7nf.png", "companyDifferentFromSignQuaReason":"这是一段说明文字"},{"trafficDrivingType":"1_PHONE","trafficDrivingContent":"1**86","variableName":"my1Phone","companyName":"阿里云计算有限公司","organizationCode":"91330****73959654P","companyDifferentFromSignQuaReason":"这是一段说明文字"}]
	TrafficDriving *string `json:"TrafficDriving,omitempty" xml:"TrafficDriving,omitempty"`
}

func (s CreateSmsTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsTemplateShrinkRequest) GetApplySceneContent() *string {
	return s.ApplySceneContent
}

func (s *CreateSmsTemplateShrinkRequest) GetIntlType() *int32 {
	return s.IntlType
}

func (s *CreateSmsTemplateShrinkRequest) GetMoreDataShrink() *string {
	return s.MoreDataShrink
}

func (s *CreateSmsTemplateShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsTemplateShrinkRequest) GetRelatedSignName() *string {
	return s.RelatedSignName
}

func (s *CreateSmsTemplateShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateSmsTemplateShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsTemplateShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsTemplateShrinkRequest) GetTemplateContent() *string {
	return s.TemplateContent
}

func (s *CreateSmsTemplateShrinkRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateSmsTemplateShrinkRequest) GetTemplateRule() *string {
	return s.TemplateRule
}

func (s *CreateSmsTemplateShrinkRequest) GetTemplateType() *int32 {
	return s.TemplateType
}

func (s *CreateSmsTemplateShrinkRequest) GetTrafficDriving() *string {
	return s.TrafficDriving
}

func (s *CreateSmsTemplateShrinkRequest) SetApplySceneContent(v string) *CreateSmsTemplateShrinkRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetIntlType(v int32) *CreateSmsTemplateShrinkRequest {
	s.IntlType = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetMoreDataShrink(v string) *CreateSmsTemplateShrinkRequest {
	s.MoreDataShrink = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetOwnerId(v int64) *CreateSmsTemplateShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetRelatedSignName(v string) *CreateSmsTemplateShrinkRequest {
	s.RelatedSignName = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetRemark(v string) *CreateSmsTemplateShrinkRequest {
	s.Remark = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetResourceOwnerAccount(v string) *CreateSmsTemplateShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetResourceOwnerId(v int64) *CreateSmsTemplateShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetTemplateContent(v string) *CreateSmsTemplateShrinkRequest {
	s.TemplateContent = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetTemplateName(v string) *CreateSmsTemplateShrinkRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetTemplateRule(v string) *CreateSmsTemplateShrinkRequest {
	s.TemplateRule = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetTemplateType(v int32) *CreateSmsTemplateShrinkRequest {
	s.TemplateType = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) SetTrafficDriving(v string) *CreateSmsTemplateShrinkRequest {
	s.TrafficDriving = &v
	return s
}

func (s *CreateSmsTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
