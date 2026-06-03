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
	// example:
	//
	// com.aliyun.android
	AndroidPackageName *string `json:"AndroidPackageName,omitempty" xml:"AndroidPackageName,omitempty"`
	// example:
	//
	// dfsfaawklll1****olkweklk***
	AndroidPackageSign *string `json:"AndroidPackageSign,omitempty" xml:"AndroidPackageSign,omitempty"`
	// example:
	//
	// 阿里云通信
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// example:
	//
	// -
	H5Origin *string `json:"H5Origin,omitempty" xml:"H5Origin,omitempty"`
	// example:
	//
	// -
	H5Url *string `json:"H5Url,omitempty" xml:"H5Url,omitempty"`
	// example:
	//
	// com.aliyun.ios
	IosBundleId *string `json:"IosBundleId,omitempty" xml:"IosBundleId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Android
	Platform             *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Android APP测试方案
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
