// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidPackageName(v string) *CreateSchemeConfigRequest
	GetAndroidPackageName() *string
	SetAndroidPackageSign(v string) *CreateSchemeConfigRequest
	GetAndroidPackageSign() *string
	SetAppName(v string) *CreateSchemeConfigRequest
	GetAppName() *string
	SetH5Origin(v string) *CreateSchemeConfigRequest
	GetH5Origin() *string
	SetH5Url(v string) *CreateSchemeConfigRequest
	GetH5Url() *string
	SetIosBundleId(v string) *CreateSchemeConfigRequest
	GetIosBundleId() *string
	SetOwnerId(v int64) *CreateSchemeConfigRequest
	GetOwnerId() *int64
	SetPlatform(v string) *CreateSchemeConfigRequest
	GetPlatform() *string
	SetResourceOwnerAccount(v string) *CreateSchemeConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSchemeConfigRequest
	GetResourceOwnerId() *int64
	SetSchemeName(v string) *CreateSchemeConfigRequest
	GetSchemeName() *string
}

type CreateSchemeConfigRequest struct {
	// The package name. This parameter is required when Platform is set to Android. The name must be 1 to 128 characters in length and can contain digits, letters, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// com.aliyun.android
	AndroidPackageName *string `json:"AndroidPackageName,omitempty" xml:"AndroidPackageName,omitempty"`
	// The package signature. This parameter is required when Platform is set to Android. The signature must be 32 characters in length and can contain digits and letters.
	//
	// example:
	//
	// dfsfaawklll1****olkweklk***
	AndroidPackageSign *string `json:"AndroidPackageSign,omitempty" xml:"AndroidPackageSign,omitempty"`
	// The app name, which can be up to 20 characters in length and can contain letters.
	//
	// example:
	//
	// Alibaba Cloud Communications
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The reserved field. HTML5 apps are not supported.
	//
	// example:
	//
	// -
	H5Origin *string `json:"H5Origin,omitempty" xml:"H5Origin,omitempty"`
	// The reserved field. HTML5 apps are not supported.
	//
	// example:
	//
	// -
	H5Url *string `json:"H5Url,omitempty" xml:"H5Url,omitempty"`
	// The bundle ID. This parameter is required when OsType is set to iOS. The bundle ID must be 1 to 128 characters in length and can contain digits, letters, hyphens (-), underscores (_), and periods (.).
	//
	// example:
	//
	// com.aliyun.ios
	IosBundleId *string `json:"IosBundleId,omitempty" xml:"IosBundleId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The app platform.
	//
	// Valid values:
	//
	// 	- Android
	//
	// 	- iOS
	//
	// This parameter is required.
	//
	// example:
	//
	// Android
	Platform             *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service name, which can be up to 10 characters in length and can contain letters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Aliyun
	SchemeName *string `json:"SchemeName,omitempty" xml:"SchemeName,omitempty"`
}

func (s CreateSchemeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemeConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateSchemeConfigRequest) GetAndroidPackageName() *string {
	return s.AndroidPackageName
}

func (s *CreateSchemeConfigRequest) GetAndroidPackageSign() *string {
	return s.AndroidPackageSign
}

func (s *CreateSchemeConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateSchemeConfigRequest) GetH5Origin() *string {
	return s.H5Origin
}

func (s *CreateSchemeConfigRequest) GetH5Url() *string {
	return s.H5Url
}

func (s *CreateSchemeConfigRequest) GetIosBundleId() *string {
	return s.IosBundleId
}

func (s *CreateSchemeConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSchemeConfigRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateSchemeConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSchemeConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSchemeConfigRequest) GetSchemeName() *string {
	return s.SchemeName
}

func (s *CreateSchemeConfigRequest) SetAndroidPackageName(v string) *CreateSchemeConfigRequest {
	s.AndroidPackageName = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetAndroidPackageSign(v string) *CreateSchemeConfigRequest {
	s.AndroidPackageSign = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetAppName(v string) *CreateSchemeConfigRequest {
	s.AppName = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetH5Origin(v string) *CreateSchemeConfigRequest {
	s.H5Origin = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetH5Url(v string) *CreateSchemeConfigRequest {
	s.H5Url = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetIosBundleId(v string) *CreateSchemeConfigRequest {
	s.IosBundleId = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetOwnerId(v int64) *CreateSchemeConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetPlatform(v string) *CreateSchemeConfigRequest {
	s.Platform = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetResourceOwnerAccount(v string) *CreateSchemeConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetResourceOwnerId(v int64) *CreateSchemeConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSchemeConfigRequest) SetSchemeName(v string) *CreateSchemeConfigRequest {
	s.SchemeName = &v
	return s
}

func (s *CreateSchemeConfigRequest) Validate() error {
	return dara.Validate(s)
}
