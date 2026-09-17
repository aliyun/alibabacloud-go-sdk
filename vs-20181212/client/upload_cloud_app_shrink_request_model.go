// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCloudAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UploadCloudAppShrinkRequest
	GetAppName() *string
	SetAppVersion(v string) *UploadCloudAppShrinkRequest
	GetAppVersion() *string
	SetDescription(v string) *UploadCloudAppShrinkRequest
	GetDescription() *string
	SetDownloadUrl(v string) *UploadCloudAppShrinkRequest
	GetDownloadUrl() *string
	SetMd5(v string) *UploadCloudAppShrinkRequest
	GetMd5() *string
	SetPkgFormat(v string) *UploadCloudAppShrinkRequest
	GetPkgFormat() *string
	SetPkgLabelsShrink(v string) *UploadCloudAppShrinkRequest
	GetPkgLabelsShrink() *string
	SetPkgType(v string) *UploadCloudAppShrinkRequest
	GetPkgType() *string
	SetPostCommandPath(v string) *UploadCloudAppShrinkRequest
	GetPostCommandPath() *string
	SetPostCommandTimeoutSec(v int32) *UploadCloudAppShrinkRequest
	GetPostCommandTimeoutSec() *int32
}

type UploadCloudAppShrinkRequest struct {
	// The application name. For Android applications, use the package name, such as com.aaa.bbb.
	//
	// Value rules:
	//
	// 1. Length: 4 to 50 characters.
	//
	// 2. Lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 3. The first and last characters must be letters or digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// com.aaa.bbb
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The application version. For Android applications, use the VersionName, such as 1.1.1.
	//
	// Value rules:
	//
	// 1. Length: 1 to 50 characters.
	//
	// 2. Lowercase letters, digits, underscores (_), hyphens (-), and periods (.).
	//
	// 3. The first and last characters must be letters or digits.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// Test application package
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The download URL of the application package.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.xxx.xxx.apk
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The MD5 checksum of the application package, used to verify package integrity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0CFBB7BD10CDD7279642ADAB8FEF3DEE
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The package format. The default value is the file extension of the download URL. Valid values:
	//
	// 1. apk
	//
	// 2. tar.gz
	//
	// 3. tar
	//
	// 4. zip
	//
	// 5. rar
	//
	// example:
	//
	// apk
	PkgFormat *string `json:"PkgFormat,omitempty" xml:"PkgFormat,omitempty"`
	// The cloud application labels. You can select multiple values. Valid values:
	//
	// 1. hot
	//
	// 2. game
	//
	// 3. app
	PkgLabelsShrink *string `json:"PkgLabels,omitempty" xml:"PkgLabels,omitempty"`
	// The package type.
	//
	// ## Valid values:
	//
	// 1. android
	//
	// 2. win
	//
	// 3. android_appmarket: corresponds to the Android app marketplace scenario. In this scenario, the actual APK PackageName is restricted:
	//
	// a. Different AppName values cannot share the same PackageName.
	//
	// b. The same AppName with different AppVersion values can be associated with different PackageName values.
	//
	// ## Default value:
	//
	// If not specified, the package type is automatically mapped based on PkgFormat (or the file extension of DownloadUrl). Default mappings between PkgFormat and package type:
	//
	// 1. android: apk (the apk format is mapped to android by default).
	//
	// 2. win: tar.gz, tar, zip, rar.
	//
	// 3. android_appmarket: apk.
	//
	// example:
	//
	// android
	PkgType *string `json:"PkgType,omitempty" xml:"PkgType,omitempty"`
	// The relative path of the post-installation command within the application package. Only supported for win type applications.
	//
	// example:
	//
	// install.ps1
	PostCommandPath *string `json:"PostCommandPath,omitempty" xml:"PostCommandPath,omitempty"`
	// The timeout period (in seconds) for the post-installation command. Only supported for win type applications.
	//
	// example:
	//
	// 10
	PostCommandTimeoutSec *int32 `json:"PostCommandTimeoutSec,omitempty" xml:"PostCommandTimeoutSec,omitempty"`
}

func (s UploadCloudAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadCloudAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadCloudAppShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *UploadCloudAppShrinkRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *UploadCloudAppShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UploadCloudAppShrinkRequest) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *UploadCloudAppShrinkRequest) GetMd5() *string {
	return s.Md5
}

func (s *UploadCloudAppShrinkRequest) GetPkgFormat() *string {
	return s.PkgFormat
}

func (s *UploadCloudAppShrinkRequest) GetPkgLabelsShrink() *string {
	return s.PkgLabelsShrink
}

func (s *UploadCloudAppShrinkRequest) GetPkgType() *string {
	return s.PkgType
}

func (s *UploadCloudAppShrinkRequest) GetPostCommandPath() *string {
	return s.PostCommandPath
}

func (s *UploadCloudAppShrinkRequest) GetPostCommandTimeoutSec() *int32 {
	return s.PostCommandTimeoutSec
}

func (s *UploadCloudAppShrinkRequest) SetAppName(v string) *UploadCloudAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) SetAppVersion(v string) *UploadCloudAppShrinkRequest {
	s.AppVersion = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) SetDescription(v string) *UploadCloudAppShrinkRequest {
	s.Description = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) SetDownloadUrl(v string) *UploadCloudAppShrinkRequest {
	s.DownloadUrl = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) SetMd5(v string) *UploadCloudAppShrinkRequest {
	s.Md5 = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) SetPkgFormat(v string) *UploadCloudAppShrinkRequest {
	s.PkgFormat = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) SetPkgLabelsShrink(v string) *UploadCloudAppShrinkRequest {
	s.PkgLabelsShrink = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) SetPkgType(v string) *UploadCloudAppShrinkRequest {
	s.PkgType = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) SetPostCommandPath(v string) *UploadCloudAppShrinkRequest {
	s.PostCommandPath = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) SetPostCommandTimeoutSec(v int32) *UploadCloudAppShrinkRequest {
	s.PostCommandTimeoutSec = &v
	return s
}

func (s *UploadCloudAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
