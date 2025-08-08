// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMcubeMiniPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UploadMcubeMiniPackageRequest
	GetAppId() *string
	SetAutoInstall(v int64) *UploadMcubeMiniPackageRequest
	GetAutoInstall() *int64
	SetClientVersionMax(v string) *UploadMcubeMiniPackageRequest
	GetClientVersionMax() *string
	SetClientVersionMin(v string) *UploadMcubeMiniPackageRequest
	GetClientVersionMin() *string
	SetEnableKeepAlive(v string) *UploadMcubeMiniPackageRequest
	GetEnableKeepAlive() *string
	SetEnableOptionMenu(v string) *UploadMcubeMiniPackageRequest
	GetEnableOptionMenu() *string
	SetEnableTabBar(v int64) *UploadMcubeMiniPackageRequest
	GetEnableTabBar() *int64
	SetExtendInfo(v string) *UploadMcubeMiniPackageRequest
	GetExtendInfo() *string
	SetH5Id(v string) *UploadMcubeMiniPackageRequest
	GetH5Id() *string
	SetH5Name(v string) *UploadMcubeMiniPackageRequest
	GetH5Name() *string
	SetH5Version(v string) *UploadMcubeMiniPackageRequest
	GetH5Version() *string
	SetIconFileUrl(v string) *UploadMcubeMiniPackageRequest
	GetIconFileUrl() *string
	SetIconUrl(v string) *UploadMcubeMiniPackageRequest
	GetIconUrl() *string
	SetInstallType(v int64) *UploadMcubeMiniPackageRequest
	GetInstallType() *int64
	SetMainUrl(v string) *UploadMcubeMiniPackageRequest
	GetMainUrl() *string
	SetOnexFlag(v bool) *UploadMcubeMiniPackageRequest
	GetOnexFlag() *bool
	SetPackageType(v int64) *UploadMcubeMiniPackageRequest
	GetPackageType() *int64
	SetPlatform(v string) *UploadMcubeMiniPackageRequest
	GetPlatform() *string
	SetResourceFileUrl(v string) *UploadMcubeMiniPackageRequest
	GetResourceFileUrl() *string
	SetResourceType(v int64) *UploadMcubeMiniPackageRequest
	GetResourceType() *int64
	SetTenantId(v string) *UploadMcubeMiniPackageRequest
	GetTenantId() *string
	SetUserId(v string) *UploadMcubeMiniPackageRequest
	GetUserId() *string
	SetUuid(v string) *UploadMcubeMiniPackageRequest
	GetUuid() *string
	SetVhost(v string) *UploadMcubeMiniPackageRequest
	GetVhost() *string
	SetWorkspaceId(v string) *UploadMcubeMiniPackageRequest
	GetWorkspaceId() *string
}

type UploadMcubeMiniPackageRequest struct {
	// This parameter is required.
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	AutoInstall      *int64  `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	ClientVersionMax *string `json:"ClientVersionMax,omitempty" xml:"ClientVersionMax,omitempty"`
	// This parameter is required.
	ClientVersionMin *string `json:"ClientVersionMin,omitempty" xml:"ClientVersionMin,omitempty"`
	// This parameter is required.
	EnableKeepAlive *string `json:"EnableKeepAlive,omitempty" xml:"EnableKeepAlive,omitempty"`
	// This parameter is required.
	EnableOptionMenu *string `json:"EnableOptionMenu,omitempty" xml:"EnableOptionMenu,omitempty"`
	// This parameter is required.
	EnableTabBar *int64  `json:"EnableTabBar,omitempty" xml:"EnableTabBar,omitempty"`
	ExtendInfo   *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// This parameter is required.
	H5Id *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	// This parameter is required.
	H5Name *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	// This parameter is required.
	H5Version   *string `json:"H5Version,omitempty" xml:"H5Version,omitempty"`
	IconFileUrl *string `json:"IconFileUrl,omitempty" xml:"IconFileUrl,omitempty"`
	IconUrl     *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// This parameter is required.
	InstallType *int64 `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	// This parameter is required.
	MainUrl *string `json:"MainUrl,omitempty" xml:"MainUrl,omitempty"`
	// This parameter is required.
	OnexFlag *bool `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	// This parameter is required.
	PackageType *int64 `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// This parameter is required.
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// This parameter is required.
	ResourceFileUrl *string `json:"ResourceFileUrl,omitempty" xml:"ResourceFileUrl,omitempty"`
	// This parameter is required.
	ResourceType *int64 `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	Uuid   *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// This parameter is required.
	Vhost *string `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UploadMcubeMiniPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadMcubeMiniPackageRequest) GoString() string {
	return s.String()
}

func (s *UploadMcubeMiniPackageRequest) GetAppId() *string {
	return s.AppId
}

func (s *UploadMcubeMiniPackageRequest) GetAutoInstall() *int64 {
	return s.AutoInstall
}

func (s *UploadMcubeMiniPackageRequest) GetClientVersionMax() *string {
	return s.ClientVersionMax
}

func (s *UploadMcubeMiniPackageRequest) GetClientVersionMin() *string {
	return s.ClientVersionMin
}

func (s *UploadMcubeMiniPackageRequest) GetEnableKeepAlive() *string {
	return s.EnableKeepAlive
}

func (s *UploadMcubeMiniPackageRequest) GetEnableOptionMenu() *string {
	return s.EnableOptionMenu
}

func (s *UploadMcubeMiniPackageRequest) GetEnableTabBar() *int64 {
	return s.EnableTabBar
}

func (s *UploadMcubeMiniPackageRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *UploadMcubeMiniPackageRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *UploadMcubeMiniPackageRequest) GetH5Name() *string {
	return s.H5Name
}

func (s *UploadMcubeMiniPackageRequest) GetH5Version() *string {
	return s.H5Version
}

func (s *UploadMcubeMiniPackageRequest) GetIconFileUrl() *string {
	return s.IconFileUrl
}

func (s *UploadMcubeMiniPackageRequest) GetIconUrl() *string {
	return s.IconUrl
}

func (s *UploadMcubeMiniPackageRequest) GetInstallType() *int64 {
	return s.InstallType
}

func (s *UploadMcubeMiniPackageRequest) GetMainUrl() *string {
	return s.MainUrl
}

func (s *UploadMcubeMiniPackageRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *UploadMcubeMiniPackageRequest) GetPackageType() *int64 {
	return s.PackageType
}

func (s *UploadMcubeMiniPackageRequest) GetPlatform() *string {
	return s.Platform
}

func (s *UploadMcubeMiniPackageRequest) GetResourceFileUrl() *string {
	return s.ResourceFileUrl
}

func (s *UploadMcubeMiniPackageRequest) GetResourceType() *int64 {
	return s.ResourceType
}

func (s *UploadMcubeMiniPackageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UploadMcubeMiniPackageRequest) GetUserId() *string {
	return s.UserId
}

func (s *UploadMcubeMiniPackageRequest) GetUuid() *string {
	return s.Uuid
}

func (s *UploadMcubeMiniPackageRequest) GetVhost() *string {
	return s.Vhost
}

func (s *UploadMcubeMiniPackageRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UploadMcubeMiniPackageRequest) SetAppId(v string) *UploadMcubeMiniPackageRequest {
	s.AppId = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetAutoInstall(v int64) *UploadMcubeMiniPackageRequest {
	s.AutoInstall = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetClientVersionMax(v string) *UploadMcubeMiniPackageRequest {
	s.ClientVersionMax = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetClientVersionMin(v string) *UploadMcubeMiniPackageRequest {
	s.ClientVersionMin = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetEnableKeepAlive(v string) *UploadMcubeMiniPackageRequest {
	s.EnableKeepAlive = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetEnableOptionMenu(v string) *UploadMcubeMiniPackageRequest {
	s.EnableOptionMenu = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetEnableTabBar(v int64) *UploadMcubeMiniPackageRequest {
	s.EnableTabBar = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetExtendInfo(v string) *UploadMcubeMiniPackageRequest {
	s.ExtendInfo = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetH5Id(v string) *UploadMcubeMiniPackageRequest {
	s.H5Id = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetH5Name(v string) *UploadMcubeMiniPackageRequest {
	s.H5Name = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetH5Version(v string) *UploadMcubeMiniPackageRequest {
	s.H5Version = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetIconFileUrl(v string) *UploadMcubeMiniPackageRequest {
	s.IconFileUrl = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetIconUrl(v string) *UploadMcubeMiniPackageRequest {
	s.IconUrl = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetInstallType(v int64) *UploadMcubeMiniPackageRequest {
	s.InstallType = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetMainUrl(v string) *UploadMcubeMiniPackageRequest {
	s.MainUrl = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetOnexFlag(v bool) *UploadMcubeMiniPackageRequest {
	s.OnexFlag = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetPackageType(v int64) *UploadMcubeMiniPackageRequest {
	s.PackageType = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetPlatform(v string) *UploadMcubeMiniPackageRequest {
	s.Platform = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetResourceFileUrl(v string) *UploadMcubeMiniPackageRequest {
	s.ResourceFileUrl = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetResourceType(v int64) *UploadMcubeMiniPackageRequest {
	s.ResourceType = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetTenantId(v string) *UploadMcubeMiniPackageRequest {
	s.TenantId = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetUserId(v string) *UploadMcubeMiniPackageRequest {
	s.UserId = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetUuid(v string) *UploadMcubeMiniPackageRequest {
	s.Uuid = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetVhost(v string) *UploadMcubeMiniPackageRequest {
	s.Vhost = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) SetWorkspaceId(v string) *UploadMcubeMiniPackageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UploadMcubeMiniPackageRequest) Validate() error {
	return dara.Validate(s)
}
