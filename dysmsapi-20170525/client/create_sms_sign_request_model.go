// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIcpRecordId(v int64) *CreateSmsSignRequest
	GetAppIcpRecordId() *int64
	SetApplySceneContent(v string) *CreateSmsSignRequest
	GetApplySceneContent() *string
	SetAuthorizationLetterId(v int64) *CreateSmsSignRequest
	GetAuthorizationLetterId() *int64
	SetMoreData(v []*string) *CreateSmsSignRequest
	GetMoreData() []*string
	SetOwnerId(v int64) *CreateSmsSignRequest
	GetOwnerId() *int64
	SetQualificationId(v int64) *CreateSmsSignRequest
	GetQualificationId() *int64
	SetRemark(v string) *CreateSmsSignRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *CreateSmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSmsSignRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *CreateSmsSignRequest
	GetSignName() *string
	SetSignSource(v int32) *CreateSmsSignRequest
	GetSignSource() *int32
	SetSignType(v int32) *CreateSmsSignRequest
	GetSignType() *int32
	SetThirdParty(v bool) *CreateSmsSignRequest
	GetThirdParty() *bool
	SetTrademarkId(v int64) *CreateSmsSignRequest
	GetTrademarkId() *int64
}

type CreateSmsSignRequest struct {
	// The APP-ICP filing entity ID.
	//
	// > - This parameter is required when SignSource is set to 2.
	//
	// > - You can obtain the filing entity ID by calling the [CreateSmsAppIcpRecord](~~CreateSmsAppIcpRecord~~) operation.
	//
	// example:
	//
	// 10000****029
	AppIcpRecordId *int64 `json:"AppIcpRecordId,omitempty" xml:"AppIcpRecordId,omitempty"`
	// 	Notice:  The signature source of launched apps is no longer supported.
	//
	// The app store link. If the signature source is a launched app, that is, SignSource is set to 2, specify a link that starts with http:// or https:// and make sure the app is already launched.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	// The ID of the power of attorney. When the signature is for third-party use, this parameter is required. Otherwise, the signature review will not pass. The unified social credit code in the power of attorney must match the unified social credit code in the qualification information bound to the signature. Otherwise, the signature creation fails.
	//
	// example:
	//
	// 1000********1234
	AuthorizationLetterId *int64 `json:"AuthorizationLetterId,omitempty" xml:"AuthorizationLetterId,omitempty"`
	// The supplementary materials. Upload business proof files or business screenshots to help reviewers understand your business details. See [Signature application materials](~~108076#section-xup-k46-yi4~~) and upload the relevant materials.
	MoreData []*string `json:"MoreData,omitempty" xml:"MoreData,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the approved qualification.
	//
	// > - Before applying for an SMS signature, [apply for a qualification](https://help.aliyun.com/document_detail/2539801.html).
	//
	// > - You can view the qualification ID on the [Qualification Management](https://dysms.console.aliyun.com/domestic/text/qualification) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8563**
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// The description of the SMS signature scenario. This is one of the reference materials for signature review. The description can be up to 200 characters in length.
	//
	// >  - You can describe the scenarios of your online service and provide links to the actual business website or marketplace download page.
	//
	// >  - You can provide a complete SMS example that reflects your business scenario.
	//
	// >  - You can provide the pass parameter content of variables and describe in detail the business scenario and the reason for selecting the variable property.
	//
	// >  - If the signature involves a government or public institution, specify the landline phone number of the institution.
	//
	// A well-documented application description improves the review efficiency for signatures and templates. Failure to follow the specifications or leaving this field empty may affect the approval of your signature.
	//
	// example:
	//
	// 登录场景使用验证码
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The signature name. The signature name must comply with the [signature specifications](~~108076#section-0p8-qn8-mmy~~):
	//
	// - The name must be 2 to 12 characters in length and cannot contain words such as "test".
	//
	// - The name cannot contain symbols such as 【】, (), or []. Special characters such as commas, periods, and spaces are not supported.
	//
	// > - Signature names are case-sensitive. For example, 【Aliyun通信】 and 【aliyun通信】 are treated as two different signatures.
	//
	// > - If your verification code signature and general-purpose signature have the same name, the system uses the general-purpose signature to send SMS messages by default.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云验证码
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The signature source. Valid values:
	//
	// -  **0**: full name or abbreviation of an enterprise or public institution. **(Recommended)**
	//
	// -  **5**: full name or abbreviation of a trademark.
	//
	// -  **2**: full name or abbreviation of an app. **(Not recommended)**
	//
	// For more information about signature sources, see [Signature sources](~~108076#section-fow-bfu-wo9~~).
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// The signature type. Valid values:
	//
	// - **0**: verification code.
	//
	// - **1**: general-purpose (default).
	//
	// We recommend that you use the default value: **general-purpose**.
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
	// The signature purpose. Valid values:
	//
	// - false: for personal use (default). The signature is the enterprise name, website, or product name verified under this account.
	//
	// - true: for third-party use. The signature is the enterprise name, website, or product name not verified under this account.
	//
	// 	Notice: If the signature is for personal use, select a qualification ID for personal use. If the signature is for third-party use, select a qualification ID for third-party use..
	//
	// example:
	//
	// false
	ThirdParty *bool `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
	// The trademark entity ID.
	//
	// > - 1. This parameter is required when SignSource is set to 5.
	//
	// > - 2. You can obtain the trademark entity ID by calling the [CreateSmsTrademark](~~CreateSmsTrademark~~) operation.
	//
	// > - 3. Based on carrier real-name registration requirements, provide the relevant field information. Otherwise, the probability of review rejection or carrier registration failure increases significantly.
	//
	// example:
	//
	// 1000009081***
	TrademarkId *int64 `json:"TrademarkId,omitempty" xml:"TrademarkId,omitempty"`
}

func (s CreateSmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsSignRequest) GoString() string {
	return s.String()
}

func (s *CreateSmsSignRequest) GetAppIcpRecordId() *int64 {
	return s.AppIcpRecordId
}

func (s *CreateSmsSignRequest) GetApplySceneContent() *string {
	return s.ApplySceneContent
}

func (s *CreateSmsSignRequest) GetAuthorizationLetterId() *int64 {
	return s.AuthorizationLetterId
}

func (s *CreateSmsSignRequest) GetMoreData() []*string {
	return s.MoreData
}

func (s *CreateSmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSmsSignRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *CreateSmsSignRequest) GetRemark() *string {
	return s.Remark
}

func (s *CreateSmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSmsSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *CreateSmsSignRequest) GetSignSource() *int32 {
	return s.SignSource
}

func (s *CreateSmsSignRequest) GetSignType() *int32 {
	return s.SignType
}

func (s *CreateSmsSignRequest) GetThirdParty() *bool {
	return s.ThirdParty
}

func (s *CreateSmsSignRequest) GetTrademarkId() *int64 {
	return s.TrademarkId
}

func (s *CreateSmsSignRequest) SetAppIcpRecordId(v int64) *CreateSmsSignRequest {
	s.AppIcpRecordId = &v
	return s
}

func (s *CreateSmsSignRequest) SetApplySceneContent(v string) *CreateSmsSignRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *CreateSmsSignRequest) SetAuthorizationLetterId(v int64) *CreateSmsSignRequest {
	s.AuthorizationLetterId = &v
	return s
}

func (s *CreateSmsSignRequest) SetMoreData(v []*string) *CreateSmsSignRequest {
	s.MoreData = v
	return s
}

func (s *CreateSmsSignRequest) SetOwnerId(v int64) *CreateSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSmsSignRequest) SetQualificationId(v int64) *CreateSmsSignRequest {
	s.QualificationId = &v
	return s
}

func (s *CreateSmsSignRequest) SetRemark(v string) *CreateSmsSignRequest {
	s.Remark = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerAccount(v string) *CreateSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSmsSignRequest) SetResourceOwnerId(v int64) *CreateSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSmsSignRequest) SetSignName(v string) *CreateSmsSignRequest {
	s.SignName = &v
	return s
}

func (s *CreateSmsSignRequest) SetSignSource(v int32) *CreateSmsSignRequest {
	s.SignSource = &v
	return s
}

func (s *CreateSmsSignRequest) SetSignType(v int32) *CreateSmsSignRequest {
	s.SignType = &v
	return s
}

func (s *CreateSmsSignRequest) SetThirdParty(v bool) *CreateSmsSignRequest {
	s.ThirdParty = &v
	return s
}

func (s *CreateSmsSignRequest) SetTrademarkId(v int64) *CreateSmsSignRequest {
	s.TrademarkId = &v
	return s
}

func (s *CreateSmsSignRequest) Validate() error {
	return dara.Validate(s)
}
