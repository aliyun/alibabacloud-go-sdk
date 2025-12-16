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
	Md5       *string   `json:"Md5,omitempty" xml:"Md5,omitempty"`
	PkgFormat *string   `json:"PkgFormat,omitempty" xml:"PkgFormat,omitempty"`
	PkgLabels []*string `json:"PkgLabels,omitempty" xml:"PkgLabels,omitempty" type:"Repeated"`
	PkgType   *string   `json:"PkgType,omitempty" xml:"PkgType,omitempty"`
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
