// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsSignShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIcpRecordId(v int64) *UpdateSmsSignShrinkRequest
	GetAppIcpRecordId() *int64
	SetApplySceneContent(v string) *UpdateSmsSignShrinkRequest
	GetApplySceneContent() *string
	SetAuthorizationLetterId(v int64) *UpdateSmsSignShrinkRequest
	GetAuthorizationLetterId() *int64
	SetMoreDataShrink(v string) *UpdateSmsSignShrinkRequest
	GetMoreDataShrink() *string
	SetOwnerId(v int64) *UpdateSmsSignShrinkRequest
	GetOwnerId() *int64
	SetQualificationId(v int64) *UpdateSmsSignShrinkRequest
	GetQualificationId() *int64
	SetRemark(v string) *UpdateSmsSignShrinkRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateSmsSignShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmsSignShrinkRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *UpdateSmsSignShrinkRequest
	GetSignName() *string
	SetSignSource(v int32) *UpdateSmsSignShrinkRequest
	GetSignSource() *int32
	SetSignType(v int32) *UpdateSmsSignShrinkRequest
	GetSignType() *int32
	SetThirdParty(v bool) *UpdateSmsSignShrinkRequest
	GetThirdParty() *bool
	SetTrademarkId(v int64) *UpdateSmsSignShrinkRequest
	GetTrademarkId() *int64
}

type UpdateSmsSignShrinkRequest struct {
	// The ID of the app\\"s ICP filing entity.
	//
	// > - This parameter is required if `SignSource` is set to 2.
	//
	// >
	//
	// > - You can obtain the filing entity ID by calling the [Create ICP Filing Entity](~~CreateSmsAppIcpRecord~~) operation.
	//
	// example:
	//
	// 100001***1234
	AppIcpRecordId *int64 `json:"AppIcpRecordId,omitempty" xml:"AppIcpRecordId,omitempty"`
	// The app store link. This parameter is required if the signature source (`SignSource`) is an app (the value is 2). The link must start with `http://` or `https://`, and the app must be published in the app store.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// The authorization letter ID. This parameter is required if the signature is for third-party use (`ThirdParty` is set to `true`). The Unified Social Credit Code on the authorization letter must match the code in the selected qualification\\"s information.
	//
	// example:
	//
	// 1000********1234
	AuthorizationLetterId *int64 `json:"AuthorizationLetterId,omitempty" xml:"AuthorizationLetterId,omitempty"`
	// Additional supporting materials. You can upload supporting business documents or business screenshots to help with the review. For details on what to upload, see [Signature application materials](~~108076#section-xup-k46-yi4~~).
	MoreDataShrink *string `json:"MoreData,omitempty" xml:"MoreData,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the approved qualification.
	//
	// > - You must [apply for a qualification](https://help.aliyun.com/zh/sms/user-guide/new-qualification?spm=a2c4g.11186623.0.0.718d187bbkpMRK) before applying for an SMS signature.
	//
	// >
	//
	// > - You can find the qualification ID on the [qualification management](https://dysms.console.aliyun.com/domestic/text/qualification) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8563**
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// A description of the SMS signature\\"s use case. This information is used during the review and must be 200 characters or less.
	//
	// > - Describe the use case for your live service. Include relevant links, such as a website link or an app store link.
	//
	// >
	//
	// > - Provide a complete example of an SMS message that reflects your use case.
	//
	// >
	//
	// > - Provide the values for any variables. Describe the use case in detail and explain why the variables are necessary.
	//
	// >
	//
	// > - If the signature involves a government agency or public institution, provide its official landline number.
	//
	// Providing complete and accurate information accelerates the review process. If you do not provide the required information, your signature application may be rejected.
	//
	// example:
	//
	// 登录场景申请验证码
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The name of the rejected SMS signature. You can find rejected SMS signatures on the [Domestic Messages > Signature Management](https://dysms.console.aliyun.com/domestic/text/sign) page in the console, or by calling the [QuerySmsSignList](~~QuerySmsSignList~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云验证码
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The signature source. Valid values:
	//
	// - **0**: The full or abbreviated name of an enterprise or public institution. **(Recommended)**
	//
	// - **5**: The full or abbreviated trademark name.
	//
	// - **2**: The full or abbreviated name of an app. **(Not recommended)**
	//
	// For more information, see [signature source](~~108076#section-fow-bfu-wo9~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// The signature type. Valid values:
	//
	// - **0**: verification code.
	//
	// - **1**: general (default).
	//
	// We recommend that you use the default value, **general**.
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
	// The signature purpose. Valid values:
	//
	// - false: for own use (default). The signature is for a business, website, or product owned by your account\\"s verified entity.
	//
	// - true: for third-party use. The signature is for a business, website, or product not owned by your account\\"s verified entity.
	//
	//   	Notice: Ensure the selected qualification ID matches the signature purpose (for own use or for third-party use).
	//
	// example:
	//
	// false
	ThirdParty *bool `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
	// The trademark entity ID.
	//
	// > - This parameter is required if `SignSource` is set to 5.
	//
	// >
	//
	// > - You can obtain the trademark entity ID by calling the [Create Trademark Entity](~~CreateSmsTrademark~~) operation.
	//
	// example:
	//
	// 1000009081***
	TrademarkId *int64 `json:"TrademarkId,omitempty" xml:"TrademarkId,omitempty"`
}

func (s UpdateSmsSignShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsSignShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignShrinkRequest) GetAppIcpRecordId() *int64 {
	return s.AppIcpRecordId
}

func (s *UpdateSmsSignShrinkRequest) GetApplySceneContent() *string {
	return s.ApplySceneContent
}

func (s *UpdateSmsSignShrinkRequest) GetAuthorizationLetterId() *int64 {
	return s.AuthorizationLetterId
}

func (s *UpdateSmsSignShrinkRequest) GetMoreDataShrink() *string {
	return s.MoreDataShrink
}

func (s *UpdateSmsSignShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmsSignShrinkRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *UpdateSmsSignShrinkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateSmsSignShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmsSignShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmsSignShrinkRequest) GetSignName() *string {
	return s.SignName
}

func (s *UpdateSmsSignShrinkRequest) GetSignSource() *int32 {
	return s.SignSource
}

func (s *UpdateSmsSignShrinkRequest) GetSignType() *int32 {
	return s.SignType
}

func (s *UpdateSmsSignShrinkRequest) GetThirdParty() *bool {
	return s.ThirdParty
}

func (s *UpdateSmsSignShrinkRequest) GetTrademarkId() *int64 {
	return s.TrademarkId
}

func (s *UpdateSmsSignShrinkRequest) SetAppIcpRecordId(v int64) *UpdateSmsSignShrinkRequest {
	s.AppIcpRecordId = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetApplySceneContent(v string) *UpdateSmsSignShrinkRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetAuthorizationLetterId(v int64) *UpdateSmsSignShrinkRequest {
	s.AuthorizationLetterId = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetMoreDataShrink(v string) *UpdateSmsSignShrinkRequest {
	s.MoreDataShrink = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetOwnerId(v int64) *UpdateSmsSignShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetQualificationId(v int64) *UpdateSmsSignShrinkRequest {
	s.QualificationId = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetRemark(v string) *UpdateSmsSignShrinkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetResourceOwnerAccount(v string) *UpdateSmsSignShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetResourceOwnerId(v int64) *UpdateSmsSignShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetSignName(v string) *UpdateSmsSignShrinkRequest {
	s.SignName = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetSignSource(v int32) *UpdateSmsSignShrinkRequest {
	s.SignSource = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetSignType(v int32) *UpdateSmsSignShrinkRequest {
	s.SignType = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetThirdParty(v bool) *UpdateSmsSignShrinkRequest {
	s.ThirdParty = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) SetTrademarkId(v int64) *UpdateSmsSignShrinkRequest {
	s.TrademarkId = &v
	return s
}

func (s *UpdateSmsSignShrinkRequest) Validate() error {
	return dara.Validate(s)
}
