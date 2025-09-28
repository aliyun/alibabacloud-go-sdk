// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVerifySchemeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateVerifySchemeRequest
	GetAppName() *string
	SetAuthType(v string) *CreateVerifySchemeRequest
	GetAuthType() *string
	SetBundleId(v string) *CreateVerifySchemeRequest
	GetBundleId() *string
	SetCmApiCode(v int64) *CreateVerifySchemeRequest
	GetCmApiCode() *int64
	SetCtApiCode(v int64) *CreateVerifySchemeRequest
	GetCtApiCode() *int64
	SetCuApiCode(v int64) *CreateVerifySchemeRequest
	GetCuApiCode() *int64
	SetEmail(v string) *CreateVerifySchemeRequest
	GetEmail() *string
	SetHmAppIdentifier(v string) *CreateVerifySchemeRequest
	GetHmAppIdentifier() *string
	SetHmPackageName(v string) *CreateVerifySchemeRequest
	GetHmPackageName() *string
	SetHmSignName(v string) *CreateVerifySchemeRequest
	GetHmSignName() *string
	SetIpWhiteList(v string) *CreateVerifySchemeRequest
	GetIpWhiteList() *string
	SetOrigin(v string) *CreateVerifySchemeRequest
	GetOrigin() *string
	SetOsType(v string) *CreateVerifySchemeRequest
	GetOsType() *string
	SetOwnerId(v int64) *CreateVerifySchemeRequest
	GetOwnerId() *int64
	SetPackName(v string) *CreateVerifySchemeRequest
	GetPackName() *string
	SetPackSign(v string) *CreateVerifySchemeRequest
	GetPackSign() *string
	SetResourceOwnerAccount(v string) *CreateVerifySchemeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateVerifySchemeRequest
	GetResourceOwnerId() *int64
	SetSceneType(v string) *CreateVerifySchemeRequest
	GetSceneType() *string
	SetSchemeName(v string) *CreateVerifySchemeRequest
	GetSchemeName() *string
	SetSmsSignName(v string) *CreateVerifySchemeRequest
	GetSmsSignName() *string
	SetUrl(v string) *CreateVerifySchemeRequest
	GetUrl() *string
}

type CreateVerifySchemeRequest struct {
	// The app name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Alibaba Cloud Communications
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The verification type. You can select multiple types only when the phone number verification is supported. Separate multiple types with commas (,).
	//
	// 	- **1**: phone number verification
	//
	// 	- **2**: SMS verification
	//
	// example:
	//
	// 1,2
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// The bundle ID. This parameter is required when OsType is set to iOS. The bundle ID must be 1 to 128 characters in length and can contain digits, letters, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// example.aliyundoc.com
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The channel code of China Mobile.
	//
	// example:
	//
	// 1
	CmApiCode *int64 `json:"CmApiCode,omitempty" xml:"CmApiCode,omitempty"`
	// The channel code of China Telecom.
	//
	// example:
	//
	// 3
	CtApiCode *int64 `json:"CtApiCode,omitempty" xml:"CtApiCode,omitempty"`
	// The channel code of China Unicom.
	//
	// example:
	//
	// 2
	CuApiCode *int64 `json:"CuApiCode,omitempty" xml:"CuApiCode,omitempty"`
	// The email address that receives the key.
	//
	// example:
	//
	// username@aliyundoc.com
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	HmAppIdentifier *string `json:"HmAppIdentifier,omitempty" xml:"HmAppIdentifier,omitempty"`
	HmPackageName   *string `json:"HmPackageName,omitempty" xml:"HmPackageName,omitempty"`
	HmSignName      *string `json:"HmSignName,omitempty" xml:"HmSignName,omitempty"`
	// The IP address whitelist.
	//
	// example:
	//
	// 139.9.167.181
	//
	// 122.112.210.205
	//
	// 139.9.172.0/24
	IpWhiteList *string `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty"`
	// The source URL of the HTML5 app page. We recommend that you specify this parameter as a domain name.
	//
	// example:
	//
	// https://h5.minexiot.com
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The type of the operating system for the terminal. Valid values: iOS and Android.
	//
	// This parameter is required.
	//
	// example:
	//
	// iOS
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The package name. This parameter is required when OsType is set to Android. The name must be 1 to 128 characters in length and can contain digits, letters, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// com.aliyun
	PackName *string `json:"PackName,omitempty" xml:"PackName,omitempty"`
	// The package signature. This parameter is required when OsType is set to Android. The signature must be 32 characters in length and can contain digits and letters.
	//
	// example:
	//
	// 123aliyun
	PackSign             *string `json:"PackSign,omitempty" xml:"PackSign,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service type.
	//
	// example:
	//
	// 0
	SceneType *string `json:"SceneType,omitempty" xml:"SceneType,omitempty"`
	// The service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
	// The bound SMS signature. This parameter is valid only when AuthType is set to 2. The signature must be approved.
	//
	// example:
	//
	// Aliyun Test
	SmsSignName *string `json:"SmsSignName,omitempty" xml:"SmsSignName,omitempty"`
	// The URL of the HTML5 app page.
	//
	// example:
	//
	// https://h5.minexiot.com/index.html
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateVerifySchemeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVerifySchemeRequest) GoString() string {
	return s.String()
}

func (s *CreateVerifySchemeRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateVerifySchemeRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *CreateVerifySchemeRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateVerifySchemeRequest) GetCmApiCode() *int64 {
	return s.CmApiCode
}

func (s *CreateVerifySchemeRequest) GetCtApiCode() *int64 {
	return s.CtApiCode
}

func (s *CreateVerifySchemeRequest) GetCuApiCode() *int64 {
	return s.CuApiCode
}

func (s *CreateVerifySchemeRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateVerifySchemeRequest) GetHmAppIdentifier() *string {
	return s.HmAppIdentifier
}

func (s *CreateVerifySchemeRequest) GetHmPackageName() *string {
	return s.HmPackageName
}

func (s *CreateVerifySchemeRequest) GetHmSignName() *string {
	return s.HmSignName
}

func (s *CreateVerifySchemeRequest) GetIpWhiteList() *string {
	return s.IpWhiteList
}

func (s *CreateVerifySchemeRequest) GetOrigin() *string {
	return s.Origin
}

func (s *CreateVerifySchemeRequest) GetOsType() *string {
	return s.OsType
}

func (s *CreateVerifySchemeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateVerifySchemeRequest) GetPackName() *string {
	return s.PackName
}

func (s *CreateVerifySchemeRequest) GetPackSign() *string {
	return s.PackSign
}

func (s *CreateVerifySchemeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateVerifySchemeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateVerifySchemeRequest) GetSceneType() *string {
	return s.SceneType
}

func (s *CreateVerifySchemeRequest) GetSchemeName() *string {
	return s.SchemeName
}

func (s *CreateVerifySchemeRequest) GetSmsSignName() *string {
	return s.SmsSignName
}

func (s *CreateVerifySchemeRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateVerifySchemeRequest) SetAppName(v string) *CreateVerifySchemeRequest {
	s.AppName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetAuthType(v string) *CreateVerifySchemeRequest {
	s.AuthType = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetBundleId(v string) *CreateVerifySchemeRequest {
	s.BundleId = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetCmApiCode(v int64) *CreateVerifySchemeRequest {
	s.CmApiCode = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetCtApiCode(v int64) *CreateVerifySchemeRequest {
	s.CtApiCode = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetCuApiCode(v int64) *CreateVerifySchemeRequest {
	s.CuApiCode = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetEmail(v string) *CreateVerifySchemeRequest {
	s.Email = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetHmAppIdentifier(v string) *CreateVerifySchemeRequest {
	s.HmAppIdentifier = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetHmPackageName(v string) *CreateVerifySchemeRequest {
	s.HmPackageName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetHmSignName(v string) *CreateVerifySchemeRequest {
	s.HmSignName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetIpWhiteList(v string) *CreateVerifySchemeRequest {
	s.IpWhiteList = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetOrigin(v string) *CreateVerifySchemeRequest {
	s.Origin = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetOsType(v string) *CreateVerifySchemeRequest {
	s.OsType = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetOwnerId(v int64) *CreateVerifySchemeRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetPackName(v string) *CreateVerifySchemeRequest {
	s.PackName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetPackSign(v string) *CreateVerifySchemeRequest {
	s.PackSign = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetResourceOwnerAccount(v string) *CreateVerifySchemeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetResourceOwnerId(v int64) *CreateVerifySchemeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetSceneType(v string) *CreateVerifySchemeRequest {
	s.SceneType = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetSchemeName(v string) *CreateVerifySchemeRequest {
	s.SchemeName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetSmsSignName(v string) *CreateVerifySchemeRequest {
	s.SmsSignName = &v
	return s
}

func (s *CreateVerifySchemeRequest) SetUrl(v string) *CreateVerifySchemeRequest {
	s.Url = &v
	return s
}

func (s *CreateVerifySchemeRequest) Validate() error {
	return dara.Validate(s)
}
