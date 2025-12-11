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
	AppIcpRecordId *int64 `json:"AppIcpRecordId,omitempty" xml:"AppIcpRecordId,omitempty"`
	// Application scenarios, instructions as follows:
	//
	// - For registered websites, enter the domain name with HTTP or HTTPS that has been registered with the MIIT.
	//
	// - For launched apps, provide a display link from the app store with HTTP or HTTPS, ensuring the app is online.
	//
	// - For public accounts or mini-programs, input the full name, ensuring they are online.
	//
	// - For e-commerce platform store names, applicable only to enterprise users, provide a display link with HTTP or HTTPS for the store.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent     *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	AuthorizationLetterId *int64  `json:"AuthorizationLetterId,omitempty" xml:"AuthorizationLetterId,omitempty"`
	// Additional information to supplement uploaded business proof documents or screenshots, which helps reviewers understand your business details.
	//
	// This parameter is optional; please fill it out based on your actual needs.
	MoreData []*string `json:"MoreData,omitempty" xml:"MoreData,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Approved or under-review qualification ID.
	//
	// > - Before applying for an SMS signature, please first [Apply for Qualification](https://help.aliyun.com/zh/sms/user-guide/new-qualification?spm=a2c4g.11186623.0.0.718d187bbkpMRK).
	//
	// > - You can view the qualification ID on the [Qualification Management](https://dysms.console.aliyun.com/domestic/text/qualification) page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8563**
	QualificationId *int64 `json:"QualificationId,omitempty" xml:"QualificationId,omitempty"`
	// Explanation of the SMS signature scenario, with a maximum length of 200 characters.
	//
	// > The scenario explanation is one of the reference materials for signature review. Please provide a detailed description of the usage scenarios for your live services, along with links to verify these services such as website URLs with MIIT备案, app store display links, full names of public accounts or mini-programs, etc. For login scenarios, test account credentials are also required. A comprehensive application explanation enhances the efficiency of signature and template reviews. Refer to the **Application Scenario*	- column in the [Signature Source](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.i2#section-xup-k46-yi4) table for filling in SMS scenarios.
	//
	// example:
	//
	// SMS signature for the login scenario using verification code.
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Signature name. Please adhere to the [Signature Specifications](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.0.4f9710fder2gR7#section-0p8-qn8-mmy).
	//
	// > - Signature names are case-insensitive; e.g., 【Aliyun Communication】 and 【aliyun communication】 are considered identical.
	//
	// > - If your verification code signature and general signature names are the same, the system defaults to using the general signature for sending SMS messages.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// Signature source. Values:
	//
	// - **0**: Full name or abbreviation of an enterprise or institution.
	//
	// - **1**: Full name or abbreviation of a MIIT-registered website.
	//
	// - **2**: Full name or abbreviation of an App.
	//
	// - **3**: Full name or abbreviation of an official account or mini-program.
	//
	// - **4**: Full name or abbreviation of an e-commerce platform store.
	//
	// - **5**: Full name or abbreviation of a trademark.
	//
	// For detailed information on signature sources, refer to [Signature Source](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.0.4f9710fder2gR7#section-xup-k46-yi4).
	//
	// > This interface does not support applying for signatures with sources as **Test or Learning*	- and **Trial Use**. If you need to apply for signatures with these sources, please go to the [SMS Service Console](https://dysms.console.aliyun.com/domestic/text/sign/add/qualification).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// Signature type. Values:
	//
	// - **0**: Verification Code
	//
	// - **1**: General (Default)
	//
	// > It is recommended to use the default value: **General**.
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
	// Choose whether the applied signature is for self-use or third-party use.
	//
	// - false: Self-use (default)
	//
	// - true: Third-party use
	//
	// 	Notice: Please select self-use qualification ID when the signature is for self-use; choose third-party use qualification ID when it\\"s for third-party use.
	//
	// example:
	//
	// false
	ThirdParty  *bool  `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
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
