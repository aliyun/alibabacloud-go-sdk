// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsSignRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIcpRecordId(v int64) *UpdateSmsSignRequest
	GetAppIcpRecordId() *int64
	SetApplySceneContent(v string) *UpdateSmsSignRequest
	GetApplySceneContent() *string
	SetAuthorizationLetterId(v int64) *UpdateSmsSignRequest
	GetAuthorizationLetterId() *int64
	SetMoreData(v []*string) *UpdateSmsSignRequest
	GetMoreData() []*string
	SetOwnerId(v int64) *UpdateSmsSignRequest
	GetOwnerId() *int64
	SetQualificationId(v int64) *UpdateSmsSignRequest
	GetQualificationId() *int64
	SetRemark(v string) *UpdateSmsSignRequest
	GetRemark() *string
	SetResourceOwnerAccount(v string) *UpdateSmsSignRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateSmsSignRequest
	GetResourceOwnerId() *int64
	SetSignName(v string) *UpdateSmsSignRequest
	GetSignName() *string
	SetSignSource(v int32) *UpdateSmsSignRequest
	GetSignSource() *int32
	SetSignType(v int32) *UpdateSmsSignRequest
	GetSignType() *int32
	SetThirdParty(v bool) *UpdateSmsSignRequest
	GetThirdParty() *bool
	SetTrademarkId(v int64) *UpdateSmsSignRequest
	GetTrademarkId() *int64
}

type UpdateSmsSignRequest struct {
	AppIcpRecordId *int64 `json:"AppIcpRecordId,omitempty" xml:"AppIcpRecordId,omitempty"`
	// Application scenarios, instructions as follows:
	//
	// - For registered websites, please enter the domain name registered with MIIT, including HTTP or HTTPS.
	//
	// - For launched apps, provide the display link from the app store with HTTP or HTTPS, ensuring the app is online.
	//
	// - For public accounts or mini-programs, fill in the full name, ensuring they are online.
	//
	// - For e-commerce platform store names (for enterprise users only), provide the display link with HTTP or HTTPS.
	//
	// example:
	//
	// http://www.aliyun.com/
	ApplySceneContent     *string `json:"ApplySceneContent,omitempty" xml:"ApplySceneContent,omitempty"`
	AuthorizationLetterId *int64  `json:"AuthorizationLetterId,omitempty" xml:"AuthorizationLetterId,omitempty"`
	// Additional materials, such as uploading business proof documents or screenshots of business operations, to help reviewers understand your business details.
	MoreData []*string `json:"MoreData,omitempty" xml:"MoreData,omitempty" type:"Repeated"`
	OwnerId  *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Approved or under-review qualification ID.
	//
	// > - Before applying for an SMS signature, please first [apply for qualifications](https://help.aliyun.com/zh/sms/user-guide/new-qualification?spm=a2c4g.11186623.0.0.718d187bbkpMRK).
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
	// > The scenario explanation is one of the reference information for signature review. Please provide a detailed description of the usage scenarios of the launched business, along with verifiable information such as website links, registered domain addresses, app store download links, full names of public accounts or mini-programs, etc. For login scenarios, test account credentials are also required. A well-informed application explanation will enhance the efficiency of signature and template reviews. Refer to the **Application Scenarios*	- column in the [Signature Source](https://help.aliyun.com/zh/sms/user-guide/signature-specifications-1?spm=a2c4g.11186623.0.i2#section-xup-k46-yi4) table for filling in SMS scenarios.
	//
	// example:
	//
	// 登录场景申请验证码
	Remark               *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Signature not yet approved.
	//
	// This parameter is required.
	//
	// example:
	//
	// 阿里云验证码
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// Source of the signature. Values:
	//
	// - **0**: Full name or abbreviation of enterprises and institutions.
	//
	// - **1**: Full name or abbreviation of MIIT-registered websites.
	//
	// - **2**: Full name or abbreviation of app applications.
	//
	// - **3**: Full name or abbreviation of public accounts or mini-programs.
	//
	// - **4**: Full name or abbreviation of e-commerce platform store names.
	//
	// - **5**: Full name or abbreviation of trademarks.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	SignSource *int32 `json:"SignSource,omitempty" xml:"SignSource,omitempty"`
	// Signature type. It is recommended to use the default value.
	//
	// - **0**: Verification code
	//
	// - **1**: General (default)
	//
	// example:
	//
	// 1
	SignType *int32 `json:"SignType,omitempty" xml:"SignType,omitempty"`
	// Whether the signature is for self-use or others.
	//
	// - false: Self-use
	//
	// - true: Others
	//
	// 	Notice: When the signature is for self-use, select the self-use qualification ID; when it\\"s for others, choose the others\\" qualification ID.
	//
	// example:
	//
	// false
	ThirdParty  *bool  `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
	TrademarkId *int64 `json:"TrademarkId,omitempty" xml:"TrademarkId,omitempty"`
}

func (s UpdateSmsSignRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsSignRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmsSignRequest) GetAppIcpRecordId() *int64 {
	return s.AppIcpRecordId
}

func (s *UpdateSmsSignRequest) GetApplySceneContent() *string {
	return s.ApplySceneContent
}

func (s *UpdateSmsSignRequest) GetAuthorizationLetterId() *int64 {
	return s.AuthorizationLetterId
}

func (s *UpdateSmsSignRequest) GetMoreData() []*string {
	return s.MoreData
}

func (s *UpdateSmsSignRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateSmsSignRequest) GetQualificationId() *int64 {
	return s.QualificationId
}

func (s *UpdateSmsSignRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateSmsSignRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateSmsSignRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateSmsSignRequest) GetSignName() *string {
	return s.SignName
}

func (s *UpdateSmsSignRequest) GetSignSource() *int32 {
	return s.SignSource
}

func (s *UpdateSmsSignRequest) GetSignType() *int32 {
	return s.SignType
}

func (s *UpdateSmsSignRequest) GetThirdParty() *bool {
	return s.ThirdParty
}

func (s *UpdateSmsSignRequest) GetTrademarkId() *int64 {
	return s.TrademarkId
}

func (s *UpdateSmsSignRequest) SetAppIcpRecordId(v int64) *UpdateSmsSignRequest {
	s.AppIcpRecordId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetApplySceneContent(v string) *UpdateSmsSignRequest {
	s.ApplySceneContent = &v
	return s
}

func (s *UpdateSmsSignRequest) SetAuthorizationLetterId(v int64) *UpdateSmsSignRequest {
	s.AuthorizationLetterId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetMoreData(v []*string) *UpdateSmsSignRequest {
	s.MoreData = v
	return s
}

func (s *UpdateSmsSignRequest) SetOwnerId(v int64) *UpdateSmsSignRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetQualificationId(v int64) *UpdateSmsSignRequest {
	s.QualificationId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetRemark(v string) *UpdateSmsSignRequest {
	s.Remark = &v
	return s
}

func (s *UpdateSmsSignRequest) SetResourceOwnerAccount(v string) *UpdateSmsSignRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateSmsSignRequest) SetResourceOwnerId(v int64) *UpdateSmsSignRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateSmsSignRequest) SetSignName(v string) *UpdateSmsSignRequest {
	s.SignName = &v
	return s
}

func (s *UpdateSmsSignRequest) SetSignSource(v int32) *UpdateSmsSignRequest {
	s.SignSource = &v
	return s
}

func (s *UpdateSmsSignRequest) SetSignType(v int32) *UpdateSmsSignRequest {
	s.SignType = &v
	return s
}

func (s *UpdateSmsSignRequest) SetThirdParty(v bool) *UpdateSmsSignRequest {
	s.ThirdParty = &v
	return s
}

func (s *UpdateSmsSignRequest) SetTrademarkId(v int64) *UpdateSmsSignRequest {
	s.TrademarkId = &v
	return s
}

func (s *UpdateSmsSignRequest) Validate() error {
	return dara.Validate(s)
}
