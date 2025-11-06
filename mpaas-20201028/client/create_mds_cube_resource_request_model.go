// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMdsCubeResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidMaxVersion(v string) *CreateMdsCubeResourceRequest
	GetAndroidMaxVersion() *string
	SetAndroidMinVersion(v string) *CreateMdsCubeResourceRequest
	GetAndroidMinVersion() *string
	SetAppId(v string) *CreateMdsCubeResourceRequest
	GetAppId() *string
	SetExtendInfo(v string) *CreateMdsCubeResourceRequest
	GetExtendInfo() *string
	SetFileUrl(v string) *CreateMdsCubeResourceRequest
	GetFileUrl() *string
	SetIosMaxVersion(v string) *CreateMdsCubeResourceRequest
	GetIosMaxVersion() *string
	SetIosMinVersion(v string) *CreateMdsCubeResourceRequest
	GetIosMinVersion() *string
	SetMockDataUrl(v string) *CreateMdsCubeResourceRequest
	GetMockDataUrl() *string
	SetOnexFlag(v bool) *CreateMdsCubeResourceRequest
	GetOnexFlag() *bool
	SetPlatform(v string) *CreateMdsCubeResourceRequest
	GetPlatform() *string
	SetPreviewPictureUrl(v string) *CreateMdsCubeResourceRequest
	GetPreviewPictureUrl() *string
	SetTemplateId(v string) *CreateMdsCubeResourceRequest
	GetTemplateId() *string
	SetTemplateResourceDesc(v string) *CreateMdsCubeResourceRequest
	GetTemplateResourceDesc() *string
	SetTemplateResourceVersion(v string) *CreateMdsCubeResourceRequest
	GetTemplateResourceVersion() *string
	SetTenantId(v string) *CreateMdsCubeResourceRequest
	GetTenantId() *string
	SetWorkspaceId(v string) *CreateMdsCubeResourceRequest
	GetWorkspaceId() *string
}

type CreateMdsCubeResourceRequest struct {
	// example:
	//
	// 0.0.0.1
	AndroidMaxVersion *string `json:"AndroidMaxVersion,omitempty" xml:"AndroidMaxVersion,omitempty"`
	// example:
	//
	// 0.0.0.1
	AndroidMinVersion *string `json:"AndroidMinVersion,omitempty" xml:"AndroidMinVersion,omitempty"`
	// example:
	//
	// ALIPUBE5C3F6D091419
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// {}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// example:
	//
	// url
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// example:
	//
	// 0.0.0.1
	IosMaxVersion *string `json:"IosMaxVersion,omitempty" xml:"IosMaxVersion,omitempty"`
	// example:
	//
	// 0.0.0.1
	IosMinVersion *string `json:"IosMinVersion,omitempty" xml:"IosMinVersion,omitempty"`
	// example:
	//
	// url
	MockDataUrl *string `json:"MockDataUrl,omitempty" xml:"MockDataUrl,omitempty"`
	// example:
	//
	// true
	OnexFlag *bool `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	// example:
	//
	// Android,iOS
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// example:
	//
	// url
	PreviewPictureUrl *string `json:"PreviewPictureUrl,omitempty" xml:"PreviewPictureUrl,omitempty"`
	// example:
	//
	// tempalte_test
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// resource desc
	TemplateResourceDesc *string `json:"TemplateResourceDesc,omitempty" xml:"TemplateResourceDesc,omitempty"`
	// example:
	//
	// 0.0.0.1
	TemplateResourceVersion *string `json:"TemplateResourceVersion,omitempty" xml:"TemplateResourceVersion,omitempty"`
	// example:
	//
	// LZFPEFIM-zh_CN
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// dev
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMdsCubeResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMdsCubeResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateMdsCubeResourceRequest) GetAndroidMaxVersion() *string {
	return s.AndroidMaxVersion
}

func (s *CreateMdsCubeResourceRequest) GetAndroidMinVersion() *string {
	return s.AndroidMinVersion
}

func (s *CreateMdsCubeResourceRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMdsCubeResourceRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *CreateMdsCubeResourceRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateMdsCubeResourceRequest) GetIosMaxVersion() *string {
	return s.IosMaxVersion
}

func (s *CreateMdsCubeResourceRequest) GetIosMinVersion() *string {
	return s.IosMinVersion
}

func (s *CreateMdsCubeResourceRequest) GetMockDataUrl() *string {
	return s.MockDataUrl
}

func (s *CreateMdsCubeResourceRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *CreateMdsCubeResourceRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateMdsCubeResourceRequest) GetPreviewPictureUrl() *string {
	return s.PreviewPictureUrl
}

func (s *CreateMdsCubeResourceRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CreateMdsCubeResourceRequest) GetTemplateResourceDesc() *string {
	return s.TemplateResourceDesc
}

func (s *CreateMdsCubeResourceRequest) GetTemplateResourceVersion() *string {
	return s.TemplateResourceVersion
}

func (s *CreateMdsCubeResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMdsCubeResourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMdsCubeResourceRequest) SetAndroidMaxVersion(v string) *CreateMdsCubeResourceRequest {
	s.AndroidMaxVersion = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetAndroidMinVersion(v string) *CreateMdsCubeResourceRequest {
	s.AndroidMinVersion = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetAppId(v string) *CreateMdsCubeResourceRequest {
	s.AppId = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetExtendInfo(v string) *CreateMdsCubeResourceRequest {
	s.ExtendInfo = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetFileUrl(v string) *CreateMdsCubeResourceRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetIosMaxVersion(v string) *CreateMdsCubeResourceRequest {
	s.IosMaxVersion = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetIosMinVersion(v string) *CreateMdsCubeResourceRequest {
	s.IosMinVersion = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetMockDataUrl(v string) *CreateMdsCubeResourceRequest {
	s.MockDataUrl = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetOnexFlag(v bool) *CreateMdsCubeResourceRequest {
	s.OnexFlag = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetPlatform(v string) *CreateMdsCubeResourceRequest {
	s.Platform = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetPreviewPictureUrl(v string) *CreateMdsCubeResourceRequest {
	s.PreviewPictureUrl = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetTemplateId(v string) *CreateMdsCubeResourceRequest {
	s.TemplateId = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetTemplateResourceDesc(v string) *CreateMdsCubeResourceRequest {
	s.TemplateResourceDesc = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetTemplateResourceVersion(v string) *CreateMdsCubeResourceRequest {
	s.TemplateResourceVersion = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetTenantId(v string) *CreateMdsCubeResourceRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) SetWorkspaceId(v string) *CreateMdsCubeResourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMdsCubeResourceRequest) Validate() error {
	return dara.Validate(s)
}
