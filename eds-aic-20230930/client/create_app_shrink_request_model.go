// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateAppShrinkRequest
	GetAppName() *string
	SetBizRegionId(v string) *CreateAppShrinkRequest
	GetBizRegionId() *string
	SetCustomAppInfoShrink(v string) *CreateAppShrinkRequest
	GetCustomAppInfoShrink() *string
	SetDescription(v string) *CreateAppShrinkRequest
	GetDescription() *string
	SetFileName(v string) *CreateAppShrinkRequest
	GetFileName() *string
	SetFilePath(v string) *CreateAppShrinkRequest
	GetFilePath() *string
	SetIconUrl(v string) *CreateAppShrinkRequest
	GetIconUrl() *string
	SetInstallParam(v string) *CreateAppShrinkRequest
	GetInstallParam() *string
	SetOssAppUrl(v string) *CreateAppShrinkRequest
	GetOssAppUrl() *string
	SetSignApk(v string) *CreateAppShrinkRequest
	GetSignApk() *string
}

type CreateAppShrinkRequest struct {
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
	CustomAppInfoShrink *string `json:"CustomAppInfo,omitempty" xml:"CustomAppInfo,omitempty"`
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

func (s CreateAppShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAppShrinkRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateAppShrinkRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *CreateAppShrinkRequest) GetCustomAppInfoShrink() *string {
	return s.CustomAppInfoShrink
}

func (s *CreateAppShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateAppShrinkRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateAppShrinkRequest) GetIconUrl() *string {
	return s.IconUrl
}

func (s *CreateAppShrinkRequest) GetInstallParam() *string {
	return s.InstallParam
}

func (s *CreateAppShrinkRequest) GetOssAppUrl() *string {
	return s.OssAppUrl
}

func (s *CreateAppShrinkRequest) GetSignApk() *string {
	return s.SignApk
}

func (s *CreateAppShrinkRequest) SetAppName(v string) *CreateAppShrinkRequest {
	s.AppName = &v
	return s
}

func (s *CreateAppShrinkRequest) SetBizRegionId(v string) *CreateAppShrinkRequest {
	s.BizRegionId = &v
	return s
}

func (s *CreateAppShrinkRequest) SetCustomAppInfoShrink(v string) *CreateAppShrinkRequest {
	s.CustomAppInfoShrink = &v
	return s
}

func (s *CreateAppShrinkRequest) SetDescription(v string) *CreateAppShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateAppShrinkRequest) SetFileName(v string) *CreateAppShrinkRequest {
	s.FileName = &v
	return s
}

func (s *CreateAppShrinkRequest) SetFilePath(v string) *CreateAppShrinkRequest {
	s.FilePath = &v
	return s
}

func (s *CreateAppShrinkRequest) SetIconUrl(v string) *CreateAppShrinkRequest {
	s.IconUrl = &v
	return s
}

func (s *CreateAppShrinkRequest) SetInstallParam(v string) *CreateAppShrinkRequest {
	s.InstallParam = &v
	return s
}

func (s *CreateAppShrinkRequest) SetOssAppUrl(v string) *CreateAppShrinkRequest {
	s.OssAppUrl = &v
	return s
}

func (s *CreateAppShrinkRequest) SetSignApk(v string) *CreateAppShrinkRequest {
	s.SignApk = &v
	return s
}

func (s *CreateAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
