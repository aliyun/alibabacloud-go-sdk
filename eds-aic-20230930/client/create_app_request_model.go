// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateAppRequest
	GetAppName() *string
	SetBizRegionId(v string) *CreateAppRequest
	GetBizRegionId() *string
	SetCustomAppInfo(v *CreateAppRequestCustomAppInfo) *CreateAppRequest
	GetCustomAppInfo() *CreateAppRequestCustomAppInfo
	SetDescription(v string) *CreateAppRequest
	GetDescription() *string
	SetFileName(v string) *CreateAppRequest
	GetFileName() *string
	SetFilePath(v string) *CreateAppRequest
	GetFilePath() *string
	SetIconUrl(v string) *CreateAppRequest
	GetIconUrl() *string
	SetInstallParam(v string) *CreateAppRequest
	GetInstallParam() *string
	SetOssAppUrl(v string) *CreateAppRequest
	GetOssAppUrl() *string
	SetSignApk(v string) *CreateAppRequest
	GetSignApk() *string
}

type CreateAppRequest struct {
	// The application name.
	//
	// example:
	//
	// Application name 1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The custom application information.
	//
	// > - If you pass a custom application, pass the `CustomAppInfo` parameter. All six fields in this object parameter are required.
	//
	// >
	//
	// > - Custom applications have a higher priority than applications from the WUYING Workspace app center. If you pass the `CustomAppInfo` parameter, `FileName` and `FilePath`, or `OssAppUrl` will be invalid.
	CustomAppInfo *CreateAppRequestCustomAppInfo `json:"CustomAppInfo,omitempty" xml:"CustomAppInfo,omitempty" type:"Struct"`
	// The application description.
	//
	// example:
	//
	// Application description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the application file stored in Object Storage Service (OSS). This parameter and `FilePath` together determine the unique OSS address.
	//
	// > - If you pass an application from the WUYING Workspace app center, you must pass `FileName` and `FilePath`, or `OssAppUrl`. The former takes precedence.
	//
	// >
	//
	// > - Log on to the [WUYING Workspace console](https://eds.console.aliyun.com/osshelp). Follow the on-screen instructions to upload your application file to the WUYING Workspace app center to obtain this parameter.
	//
	// example:
	//
	// testApp.apk
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The storage address of the application file in an OSS bucket. This parameter and `FileName` together determine the unique OSS address.
	//
	// > - If you pass an application from the WUYING Workspace app center, you must pass `FileName` and `FilePath`, or `OssAppUrl`. The former takes precedence.
	//
	// >
	//
	// > - Log on to the [WUYING Workspace console](https://eds.console.aliyun.com/osshelp). Follow the on-screen instructions to upload your application file to the WUYING Workspace app center to obtain this parameter.
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The URL of the application icon.
	//
	// example:
	//
	// https://www.example.com/icon.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The installation parameters. The `-r` installation parameter is included by default when you install the application.
	//
	// example:
	//
	// -d
	InstallParam *string `json:"InstallParam,omitempty" xml:"InstallParam,omitempty"`
	// The OSS address of the application.
	//
	// > - If you pass an application from the WUYING Workspace app center, you must pass `FileName` and `FilePath`, or `OssAppUrl`. The former takes precedence.
	//
	// >
	//
	// > - Log on to the [WUYING Workspace console](https://eds.console.aliyun.com/osshelp). Follow the on-screen instructions to upload your application file to the WUYING Workspace app center to obtain this parameter.
	//
	// example:
	//
	// http://testApp.apk
	OssAppUrl *string `json:"OssAppUrl,omitempty" xml:"OssAppUrl,omitempty"`
	// Specifies whether to perform a system signature.
	//
	// example:
	//
	// false
	SignApk *string `json:"SignApk,omitempty" xml:"SignApk,omitempty"`
}

func (s CreateAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateAppRequest) GetCustomAppInfo() *CreateAppRequestCustomAppInfo {
	return s.CustomAppInfo
}

func (s *CreateAppRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateAppRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateAppRequest) GetIconUrl() *string {
	return s.IconUrl
}

func (s *CreateAppRequest) GetInstallParam() *string {
	return s.InstallParam
}

func (s *CreateAppRequest) GetOssAppUrl() *string {
	return s.OssAppUrl
}

func (s *CreateAppRequest) GetSignApk() *string {
	return s.SignApk
}

func (s *CreateAppRequest) SetAppName(v string) *CreateAppRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppRequest) SetBizRegionId(v string) *CreateAppRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppRequest) SetCustomAppInfo(v *CreateAppRequestCustomAppInfo) *CreateAppRequest {
	s.CustomAppInfo = v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetFileName(v string) *CreateAppRequest {
	s.FileName = &v
	return s
}

func (s *CreateAppRequest) SetFilePath(v string) *CreateAppRequest {
	s.FilePath = &v
	return s
}

func (s *CreateAppRequest) SetIconUrl(v string) *CreateAppRequest {
	s.IconUrl = &v
	return s
}

func (s *CreateAppRequest) SetInstallParam(v string) *CreateAppRequest {
	s.InstallParam = &v
	return s
}

func (s *CreateAppRequest) SetOssAppUrl(v string) *CreateAppRequest {
	s.OssAppUrl = &v
	return s
}

func (s *CreateAppRequest) SetSignApk(v string) *CreateAppRequest {
	s.SignApk = &v
	return s
}

func (s *CreateAppRequest) Validate() error {
	if s.CustomAppInfo != nil {
		if err := s.CustomAppInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAppRequestCustomAppInfo struct {
	// The size of the .apk file. Unit: MB.
	//
	// example:
	//
	// 10
	ApkSize *string `json:"ApkSize,omitempty" xml:"ApkSize,omitempty"`
	// The download URL of the application.
	//
	// example:
	//
	// http://testApp.apk
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The MD5 value of the .apk package.
	//
	// example:
	//
	// df3f46ce5844ddb278f14c5a9cd2****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The application package name.
	//
	// example:
	//
	// com.example.demo
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The application version.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The application version code.
	//
	// example:
	//
	// 10000
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s CreateAppRequestCustomAppInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateAppRequestCustomAppInfo) GoString() string {
	return s.String()
}

func (s *CreateAppRequestCustomAppInfo) GetApkSize() *string {
	return s.ApkSize
}

func (s *CreateAppRequestCustomAppInfo) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateAppRequestCustomAppInfo) GetMd5() *string {
	return s.Md5
}

func (s *CreateAppRequestCustomAppInfo) GetPackageName() *string {
	return s.PackageName
}

func (s *CreateAppRequestCustomAppInfo) GetVersion() *string {
	return s.Version
}

func (s *CreateAppRequestCustomAppInfo) GetVersionCode() *string {
	return s.VersionCode
}

func (s *CreateAppRequestCustomAppInfo) SetApkSize(v string) *CreateAppRequestCustomAppInfo {
	s.ApkSize = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetDownloadUrl(v string) *CreateAppRequestCustomAppInfo {
	s.DownloadUrl = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetMd5(v string) *CreateAppRequestCustomAppInfo {
	s.Md5 = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetPackageName(v string) *CreateAppRequestCustomAppInfo {
	s.PackageName = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetVersion(v string) *CreateAppRequestCustomAppInfo {
	s.Version = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) SetVersionCode(v string) *CreateAppRequestCustomAppInfo {
	s.VersionCode = &v
	return s
}

func (s *CreateAppRequestCustomAppInfo) Validate() error {
	return dara.Validate(s)
}
