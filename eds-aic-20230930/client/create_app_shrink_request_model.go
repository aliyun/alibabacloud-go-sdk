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
	CustomAppInfoShrink *string `json:"CustomAppInfo,omitempty" xml:"CustomAppInfo,omitempty"`
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
