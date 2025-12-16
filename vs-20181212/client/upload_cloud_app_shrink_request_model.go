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
}

type UploadCloudAppShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// com.aaa.bbb
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1.1.1
	AppVersion  *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://xxx.xxx.xxx.apk
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0CFBB7BD10CDD7279642ADAB8FEF3DEE
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	PkgFormat       *string `json:"PkgFormat,omitempty" xml:"PkgFormat,omitempty"`
	PkgLabelsShrink *string `json:"PkgLabels,omitempty" xml:"PkgLabels,omitempty"`
	PkgType         *string `json:"PkgType,omitempty" xml:"PkgType,omitempty"`
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

func (s *UploadCloudAppShrinkRequest) Validate() error {
	return dara.Validate(s)
}
