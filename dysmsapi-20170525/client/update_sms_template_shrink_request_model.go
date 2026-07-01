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
	// The business scenario.
	//
	// - If the associated SMS signature\\"s business scenario is "Live App", the `ApplySceneContent` parameter must be an app URL starting with `http://` or `https://`.
	//
	// - The `ApplySceneContent` parameter is required if the associated SMS signature\\"s business scenario is "Registered Trademark" or "Name of Enterprise or Public Institution".
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// The type of the international/regional template. This parameter is required when the **TemplateType*	- parameter is set to **3**. Valid values:
	//
	// - **0**: SMS notification.
	//
	// - **1**: promotional SMS.
	//
	// - **2**: verification code.
	//
	// example:
	//
	// 0
	IntlType *int32 `json:"IntlType,omitempty" xml:"IntlType,omitempty"`
	// Additional materials, such as supporting documents or business screenshots, to help reviewers understand your business. If `TemplateType` is set to `2` (promotional SMS), you must upload proof of user authorization. For more information, see [Upload specifications for user authorization materials](https://help.aliyun.com/document_detail/312341.html).
	MoreDataShrink *string `json:"MoreData,omitempty" xml:"MoreData,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The SMS signature associated with the template.
	//
	// example:
	//
	// 阿里云
	RelatedSignName *string `json:"RelatedSignName,omitempty" xml:"RelatedSignName,omitempty"`
	// Describe your business scenario, including a URL if applicable. You must also provide a complete SMS message example with populated variables. Providing this information as required is critical for template approval.
	//
	// example:
	//
	// 登录场景使用验证码
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The code of the rejected SMS template. You can find the template code on the [Messages in Chinese Mainland > Template Management](https://dysms.console.aliyun.com/domestic/text/template) tab in the console or by calling the [QuerySmsTemplateList](~~QuerySmsTemplateList~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_152550****
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// The new template content, up to 500 characters long.
	//
	// The template content and its variables must comply with [SMS template specifications](https://help.aliyun.com/document_detail/463161.html) to be approved. To increase the approval rate and efficiency, refer to the common examples on the [Apply for Template](https://dysms.console.aliyun.com/domestic/text/template/add) page. For more information about variable specifications, see [TemplateContent parameter variable specifications](https://help.aliyun.com/document_detail/2806243.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 您正在申请手机注册，验证码为：${code}，5分钟内有效！
	TemplateContent *string `json:"TemplateContent,omitempty" xml:"TemplateContent,omitempty"`
	// The name of the SMS template, up to 30 characters long. You can find the names of rejected templates on the [Messages in Chinese Mainland > Template Management](https://dysms.console.aliyun.com/domestic/text/template) tab in the console or by calling the [QuerySmsTemplateList](~~QuerySmsTemplateList~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 验证码模板
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// The rules for the variables in the template. For details on how to define these rules, see the [example document](https://help.aliyun.com/document_detail/2806243.html).
	//
	// example:
	//
	// {"code":"characterWithNumber"}
	TemplateRule *string `json:"TemplateRule,omitempty" xml:"TemplateRule,omitempty"`
	// The SMS type. Valid values:
	//
	// - **0**: verification code.
	//
	// - **1**: SMS notification.
	//
	// - **2**: promotional SMS.
	//
	// - **3**: international/regional message.
	//
	// > Only enterprise-verified users can apply for promotional SMS and international/regional messages. For more information about the differences between personal and enterprise accounts, see [Usage notes](https://help.aliyun.com/zh/sms/user-guide/usage-notes?spm=a2c4g.11186623.0.0.67447f576NJnE8).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	TemplateType *int32 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// 	Warning:
	//
	// To manage SMS content security, messages that contain traffic-driving information such as phone numbers and URLs may be blocked by carriers, which can cause delivery failures. We recommend that you avoid including such information in your SMS templates to prevent delivery failures.
	//
	//
	//
	// A JSON string that contains a list of traffic-driving information.
	//
	// 	Notice: The value must be in the JSON format. Convert the value to a string before you pass it in.
	//
	// ### 1. Fields
	//
	// {
	//
	// "trafficDrivingType":"Traffic-driving type",
	//
	// "trafficDrivingContent":"Traffic-driving content",
	//
	// "variableName":"variable name",
	//
	// "companyName":"Name of the enterprise or public institution",
	//
	// "organizationCode":"Unified Social Credit Code",
	//
	// "icpNo":"ICP filing/permit number",
	//
	// "icpPicOssKey":"OSS key of the ICP filing screenshot",
	//
	// "companyDifferentFromSignQuaReason":"The reason why the name of the enterprise or public institution is different from that in the SMS signature qualification"
	//
	// }
	//
	// ### 2. Notes
	//
	// - If the content is not a variable, do not pass the `variableName` field.
	//
	// - If the name of the enterprise or public institution is different from that in the SMS signature qualification, provide the `companyDifferentFromSignQuaReason` field.
	//
	// - If `trafficDrivingType` is set to `DOMAIN`, you must provide all the fields.
	//
	// - For `trafficDrivingType` values other than `DOMAIN`, the `trafficDrivingType`, `trafficDrivingContent`, `companyName`, and `organizationCode` fields are required. The `variableName` and `companyDifferentFromSignQuaReason` fields are optional.
	//
	// ### 3. TrafficDrivingType enumeration
	//
	// 	Warning:
	//
	// Due to regulatory requirements, mobile numbers are not supported.
	//
	//
	//
	// - `DOMAIN`: A domain name.
	//
	// - `FIXED_PHONE`: A fixed-line phone number.
	//
	// - `400_PHONE`: A phone number that starts with 400.
	//
	// - `800_PHONE`: A phone number that starts with 800.
	//
	// - `95_PHONE`: A phone number that starts with 95.
	//
	// - `96_PHONE`: A phone number that starts with 96.
	//
	// - `1_PHONE`: A 3-digit to 8-digit phone number that starts with 1.
	//
	// - `OTHER_PHONE`: Another type of phone number.
	//
	// example:
	//
	// [{"trafficDrivingType":"DOMAIN","trafficDrivingContent":"aliyun.com","companyName":"阿里云计算有限公司","organizationCode":"91330****73959654P","icpNo":"浙B2-20****01-4","icpPicOssKey":"db7784d8-cb0c-498f-****-295f1ad6d665_mf29l7nf.png","companyDifferentFromSignQuaReason":"这是一段说明文字"},{"trafficDrivingType":"1_PHONE","trafficDrivingContent":"1**86","variableName":"my1Phone","companyName":"阿里云计算有限公司","organizationCode":"91330****73959654P","companyDifferentFromSignQuaReason":"这是一段说明文字"}]
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
