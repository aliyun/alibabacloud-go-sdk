// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmsSignShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
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
}

type UpdateSmsSignShrinkRequest struct {
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
	MoreDataShrink *string `json:"MoreData,omitempty" xml:"MoreData,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	ThirdParty *bool `json:"ThirdParty,omitempty" xml:"ThirdParty,omitempty"`
}

func (s UpdateSmsSignShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmsSignShrinkRequest) GoString() string {
	return s.String()
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

func (s *UpdateSmsSignShrinkRequest) Validate() error {
	return dara.Validate(s)
}
