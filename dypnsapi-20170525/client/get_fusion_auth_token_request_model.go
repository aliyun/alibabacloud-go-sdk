// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFusionAuthTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBundleId(v string) *GetFusionAuthTokenRequest
	GetBundleId() *string
	SetDurationSeconds(v int64) *GetFusionAuthTokenRequest
	GetDurationSeconds() *int64
	SetOwnerId(v int64) *GetFusionAuthTokenRequest
	GetOwnerId() *int64
	SetPackageName(v string) *GetFusionAuthTokenRequest
	GetPackageName() *string
	SetPackageSign(v string) *GetFusionAuthTokenRequest
	GetPackageSign() *string
	SetPlatform(v string) *GetFusionAuthTokenRequest
	GetPlatform() *string
	SetResourceOwnerAccount(v string) *GetFusionAuthTokenRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetFusionAuthTokenRequest
	GetResourceOwnerId() *int64
	SetSchemeCode(v string) *GetFusionAuthTokenRequest
	GetSchemeCode() *string
}

type GetFusionAuthTokenRequest struct {
	// The bundle ID of the app. This parameter is required when Platform is set to iOS.
	//
	// example:
	//
	// com.example.test
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The validity period of the token. Unit: seconds. Valid values: 900 to 43200.
	//
	// This parameter is required.
	//
	// example:
	//
	// 900
	DurationSeconds *int64 `json:"DurationSeconds,omitempty" xml:"DurationSeconds,omitempty"`
	OwnerId         *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The package name of the app. This parameter is required when Platform is set to Android.
	//
	// example:
	//
	// com.example.test
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The package signature of the app. This parameter is required when Platform is set to Android.
	//
	// example:
	//
	// 47fcc************************278
	PackageSign *string `json:"PackageSign,omitempty" xml:"PackageSign,omitempty"`
	// The platform type. Valid values: Android and iOS.
	//
	// This parameter is required.
	//
	// example:
	//
	// Android
	Platform             *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The service code.
	//
	// This parameter is required.
	//
	// example:
	//
	// FA1000*************201
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s GetFusionAuthTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFusionAuthTokenRequest) GoString() string {
	return s.String()
}

func (s *GetFusionAuthTokenRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *GetFusionAuthTokenRequest) GetDurationSeconds() *int64 {
	return s.DurationSeconds
}

func (s *GetFusionAuthTokenRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetFusionAuthTokenRequest) GetPackageName() *string {
	return s.PackageName
}

func (s *GetFusionAuthTokenRequest) GetPackageSign() *string {
	return s.PackageSign
}

func (s *GetFusionAuthTokenRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GetFusionAuthTokenRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetFusionAuthTokenRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetFusionAuthTokenRequest) GetSchemeCode() *string {
	return s.SchemeCode
}

func (s *GetFusionAuthTokenRequest) SetBundleId(v string) *GetFusionAuthTokenRequest {
	s.BundleId = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetDurationSeconds(v int64) *GetFusionAuthTokenRequest {
	s.DurationSeconds = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetOwnerId(v int64) *GetFusionAuthTokenRequest {
	s.OwnerId = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetPackageName(v string) *GetFusionAuthTokenRequest {
	s.PackageName = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetPackageSign(v string) *GetFusionAuthTokenRequest {
	s.PackageSign = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetPlatform(v string) *GetFusionAuthTokenRequest {
	s.Platform = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetResourceOwnerAccount(v string) *GetFusionAuthTokenRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetResourceOwnerId(v int64) *GetFusionAuthTokenRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetFusionAuthTokenRequest) SetSchemeCode(v string) *GetFusionAuthTokenRequest {
	s.SchemeCode = &v
	return s
}

func (s *GetFusionAuthTokenRequest) Validate() error {
	return dara.Validate(s)
}
