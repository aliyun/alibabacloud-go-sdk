// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadCloudAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UploadCloudAppRequest
	GetAppName() *string
	SetAppVersion(v string) *UploadCloudAppRequest
	GetAppVersion() *string
	SetDescription(v string) *UploadCloudAppRequest
	GetDescription() *string
	SetDownloadUrl(v string) *UploadCloudAppRequest
	GetDownloadUrl() *string
	SetMd5(v string) *UploadCloudAppRequest
	GetMd5() *string
	SetPkgFormat(v string) *UploadCloudAppRequest
	GetPkgFormat() *string
	SetPkgLabels(v []*string) *UploadCloudAppRequest
	GetPkgLabels() []*string
	SetPkgType(v string) *UploadCloudAppRequest
	GetPkgType() *string
}

type UploadCloudAppRequest struct {
	// The application name. For Android apps, use the package name, such as com.aaa.bbb.
	//
	// Value requirements:
	//
	// 1. Length: 4–50 characters
	//
	// 2. Allowed characters: lowercase letters, digits, underscores (_), hyphens (-), and dots (.)
	//
	// 3. The first and last characters must be a letter or digit
	//
	// This parameter is required.
	//
	// example:
	//
	// com.aaa.bbb
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// Value requirements:
	//
	// 1. Length: 1–50 characters
	//
	// 2. Allowed characters: lowercase letters, digits, underscores (_), hyphens (-), and dots (.)
	//
	// 3. The first and last characters must be a letter or digit
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// A description of the application.
	//
	// example:
	//
	// 测试应用包
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The download URL of the application package.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.xxx.xxx.apk
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The MD5 hash of the application package, used to verify package integrity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0CFBB7BD10CDD7279642ADAB8FEF3DEE
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The package format. By default, this is inferred from the file extension in the DownloadUrl. Valid values:
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
	// Cloud application labels. You can select multiple. Valid values:
	//
	// 1. hot
	//
	// 2. game
	//
	// 3. app
	PkgLabels []*string `json:"PkgLabels,omitempty" xml:"PkgLabels,omitempty" type:"Repeated"`
	// The package type.
	//
	// ## Valid values:
	//
	// 1. android
	//
	// 2. win
	//
	// 3. android_appmarket: for Android app marketplace scenarios. This scenario enforces real APK PackageName restrictions:
	//
	//    a. PackageNames must be unique across different AppNames.
	//
	//    b. The same AppName with different AppVersions can map to different PackageNames.
	//
	// ## Default behavior:
	//
	// If not specified, the system automatically maps the package type based on PkgFormat (or infers PkgFormat from the DownloadUrl file extension). The default mapping is:
	//
	// 1. android: apk
	//
	// 2. win: tar.gz, tar, zip, rar
	//
	// 3. android_appmarket: apk
	//
	// example:
	//
	// android
	PkgType *string `json:"PkgType,omitempty" xml:"PkgType,omitempty"`
}

func (s UploadCloudAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadCloudAppRequest) GoString() string {
	return s.String()
}

func (s *UploadCloudAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *UploadCloudAppRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *UploadCloudAppRequest) GetDescription() *string {
	return s.Description
}

func (s *UploadCloudAppRequest) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *UploadCloudAppRequest) GetMd5() *string {
	return s.Md5
}

func (s *UploadCloudAppRequest) GetPkgFormat() *string {
	return s.PkgFormat
}

func (s *UploadCloudAppRequest) GetPkgLabels() []*string {
	return s.PkgLabels
}

func (s *UploadCloudAppRequest) GetPkgType() *string {
	return s.PkgType
}

func (s *UploadCloudAppRequest) SetAppName(v string) *UploadCloudAppRequest {
	s.AppName = &v
	return s
}

func (s *UploadCloudAppRequest) SetAppVersion(v string) *UploadCloudAppRequest {
	s.AppVersion = &v
	return s
}

func (s *UploadCloudAppRequest) SetDescription(v string) *UploadCloudAppRequest {
	s.Description = &v
	return s
}

func (s *UploadCloudAppRequest) SetDownloadUrl(v string) *UploadCloudAppRequest {
	s.DownloadUrl = &v
	return s
}

func (s *UploadCloudAppRequest) SetMd5(v string) *UploadCloudAppRequest {
	s.Md5 = &v
	return s
}

func (s *UploadCloudAppRequest) SetPkgFormat(v string) *UploadCloudAppRequest {
	s.PkgFormat = &v
	return s
}

func (s *UploadCloudAppRequest) SetPkgLabels(v []*string) *UploadCloudAppRequest {
	s.PkgLabels = v
	return s
}

func (s *UploadCloudAppRequest) SetPkgType(v string) *UploadCloudAppRequest {
	s.PkgType = &v
	return s
}

func (s *UploadCloudAppRequest) Validate() error {
	return dara.Validate(s)
}
