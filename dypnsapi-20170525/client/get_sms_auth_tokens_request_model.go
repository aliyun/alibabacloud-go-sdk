// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsAuthTokensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleId(v string) *GetSmsAuthTokensRequest
	GetBundleId() *string
	SetExpire(v int64) *GetSmsAuthTokensRequest
	GetExpire() *int64
	SetOsType(v string) *GetSmsAuthTokensRequest
	GetOsType() *string
	SetOwnerId(v int64) *GetSmsAuthTokensRequest
	GetOwnerId() *int64
	SetPackageName(v string) *GetSmsAuthTokensRequest
	GetPackageName() *string
	SetResourceOwnerAccount(v string) *GetSmsAuthTokensRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmsAuthTokensRequest
	GetResourceOwnerId() *int64
	SetSceneCode(v string) *GetSmsAuthTokensRequest
	GetSceneCode() *string
	SetSignName(v string) *GetSmsAuthTokensRequest
	GetSignName() *string
	SetSmsCodeExpire(v int32) *GetSmsAuthTokensRequest
	GetSmsCodeExpire() *int32
	SetSmsTemplateCode(v string) *GetSmsAuthTokensRequest
	GetSmsTemplateCode() *string
}

type GetSmsAuthTokensRequest struct {
	// The ID of the iOS application. This parameter is required if OsType is set to **iOS**.
	//
	// example:
	//
	// 12345****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The validity period of the token. Unit: seconds. Valid values: 900 to 43200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 900
	Expire *int64 `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The type of the operating system. Valid values: **Android*	- and **iOS**.
	//
	// This parameter is required.
	//
	// example:
	//
	// Android
	OsType  *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The package name. This parameter is required if OsType is set to **Android**.
	//
	// example:
	//
	// com.aliqin.mytel.test
	PackageName          *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FC100000134840112
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// The signature. This parameter is required if OsType is set to **Android**.
	//
	// example:
	//
	// 47fcc6615485e83b4100433****
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// The validity period of the SMS verification code. Unit: seconds. Default value: 180.
	//
	// example:
	//
	// 60
	SmsCodeExpire *int32 `json:"SmsCodeExpire,omitempty" xml:"SmsCodeExpire,omitempty"`
	// The code of the text message template.
	//
	// This parameter is required.
	//
	// example:
	//
	// SMS_13987****
	SmsTemplateCode *string `json:"SmsTemplateCode,omitempty" xml:"SmsTemplateCode,omitempty"`
}

func (s GetSmsAuthTokensRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmsAuthTokensRequest) GoString() string {
	return s.String()
}

func (s *GetSmsAuthTokensRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *GetSmsAuthTokensRequest) GetExpire() *int64 {
	return s.Expire
}

func (s *GetSmsAuthTokensRequest) GetOsType() *string {
	return s.OsType
}

func (s *GetSmsAuthTokensRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSmsAuthTokensRequest) GetPackageName() *string {
	return s.PackageName
}

func (s *GetSmsAuthTokensRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSmsAuthTokensRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSmsAuthTokensRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *GetSmsAuthTokensRequest) GetSignName() *string {
	return s.SignName
}

func (s *GetSmsAuthTokensRequest) GetSmsCodeExpire() *int32 {
	return s.SmsCodeExpire
}

func (s *GetSmsAuthTokensRequest) GetSmsTemplateCode() *string {
	return s.SmsTemplateCode
}

func (s *GetSmsAuthTokensRequest) SetBundleId(v string) *GetSmsAuthTokensRequest {
	s.BundleId = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetExpire(v int64) *GetSmsAuthTokensRequest {
	s.Expire = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetOsType(v string) *GetSmsAuthTokensRequest {
	s.OsType = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetOwnerId(v int64) *GetSmsAuthTokensRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetPackageName(v string) *GetSmsAuthTokensRequest {
	s.PackageName = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetResourceOwnerAccount(v string) *GetSmsAuthTokensRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetResourceOwnerId(v int64) *GetSmsAuthTokensRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSceneCode(v string) *GetSmsAuthTokensRequest {
	s.SceneCode = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSignName(v string) *GetSmsAuthTokensRequest {
	s.SignName = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSmsCodeExpire(v int32) *GetSmsAuthTokensRequest {
	s.SmsCodeExpire = &v
	return s
}

func (s *GetSmsAuthTokensRequest) SetSmsTemplateCode(v string) *GetSmsAuthTokensRequest {
	s.SmsTemplateCode = &v
	return s
}

func (s *GetSmsAuthTokensRequest) Validate() error {
	return dara.Validate(s)
}
