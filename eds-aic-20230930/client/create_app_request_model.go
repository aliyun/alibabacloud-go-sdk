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
	// The name of the application.
	//
	// example:
	//
	// Application Name 1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The information about the custom app.
	//
	// >
	//
	// 	- If you want to pass in a custom app, configure the `CustomAppInfo` parameter. Take note that the six fields within it are mandatory.
	//
	// 	- A custom app has a higher priority than an app from the Alibaba Cloud Workspace Application Center. If you configure the `CustomAppInfo` parameter, the `FileName` and `FilePath` pair or the `OssAppUrl` will not take effect.
	CustomAppInfo *CreateAppRequestCustomAppInfo `json:"CustomAppInfo,omitempty" xml:"CustomAppInfo,omitempty" type:"Struct"`
	// The description of the application.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name used by the app file in Object Storage Service (OSS). This parameter, combined with `FilePath`, uniquely identifies the OSS path of the app file.
	//
	// >
	//
	// 	- If you want to pass in an app from the Alibaba Cloud Workspace Application Center, configure the `FileName` and `FilePath` parameters. Alternatively, configure the `OssAppUrl` parameter. The FileName and FilePath parameters takes precedence over the OssAppUrl parameter.
	//
	// 	- Log on to the [Elastic Desktop Service (EDS) Enterprise](https://eds.console.aliyun.com/osshelp) console, upload the app file to the Application Center according to the on-screen instructions, and then retrieve the parameter value.
	//
	// example:
	//
	// testApp.apk
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The OSS bucket path to the app file. This parameter, combined with `FileName`, uniquely identifies the OSS path of the app file.
	//
	// >
	//
	// 	- If you want to pass in an app from the Alibaba Cloud Workspace Application Center, configure the `FileName` and `FilePath` parameters. Alternatively, configure the `OssAppUrl` parameter. The FileName and FilePath parameters takes precedence over the OssAppUrl parameter.
	//
	// 	- Log on to the [EDS Enterprise](https://eds.console.aliyun.com/osshelp) console, upload the app file to the Application Center according to the on-screen instructions, and then retrieve the parameter value.
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The icon URL of the application.
	//
	// example:
	//
	// https://www.example.com/icon.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The parameters used for installing the application. By default, the `-r` parameter is included when you install an application.
	//
	// example:
	//
	// -d
	InstallParam *string `json:"InstallParam,omitempty" xml:"InstallParam,omitempty"`
	// The OSS bucket endpoint of the app file.
	//
	// >
	//
	// 	- If you want to pass in an app from the Alibaba Cloud Workspace Application Center, configure the `FileName` and `FilePath` parameters. Alternatively, configure the `OssAppUrl` parameter. The FileName and FilePath parameters takes precedence over the OssAppUrl parameter.
	//
	// 	- Log on to the [EDS Enterprise](https://eds.console.aliyun.com/osshelp) console, upload the app file to the Application Center according to the on-screen instructions, and then retrieve the parameter value.
	//
	// example:
	//
	// http://testApp.apk
	OssAppUrl *string `json:"OssAppUrl,omitempty" xml:"OssAppUrl,omitempty"`
	SignApk   *string `json:"SignApk,omitempty" xml:"SignApk,omitempty"`
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
	return dara.Validate(s)
}

type CreateAppRequestCustomAppInfo struct {
	// The size of the .apk file. Unit: MB.
	//
	// example:
	//
	// 10
	ApkSize *string `json:"ApkSize,omitempty" xml:"ApkSize,omitempty"`
	// The download URL of the app.
	//
	// example:
	//
	// http://testApp.apk
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The MD5 value of the .apk file.
	//
	// example:
	//
	// df3f46ce5844ddb278f14c5a9cd2****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the app package.
	//
	// example:
	//
	// com.example.demo
	PackageName *string `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	// The version of the app.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// The code of the app version.
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
