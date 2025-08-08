// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeNebulaResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeNebulaResourceRequest
	GetAppId() *string
	SetAutoInstall(v int32) *CreateMcubeNebulaResourceRequest
	GetAutoInstall() *int32
	SetClientVersionMax(v string) *CreateMcubeNebulaResourceRequest
	GetClientVersionMax() *string
	SetClientVersionMin(v string) *CreateMcubeNebulaResourceRequest
	GetClientVersionMin() *string
	SetCustomDomainName(v string) *CreateMcubeNebulaResourceRequest
	GetCustomDomainName() *string
	SetExtendInfo(v string) *CreateMcubeNebulaResourceRequest
	GetExtendInfo() *string
	SetFileUrl(v string) *CreateMcubeNebulaResourceRequest
	GetFileUrl() *string
	SetH5Id(v string) *CreateMcubeNebulaResourceRequest
	GetH5Id() *string
	SetH5Name(v string) *CreateMcubeNebulaResourceRequest
	GetH5Name() *string
	SetH5Version(v string) *CreateMcubeNebulaResourceRequest
	GetH5Version() *string
	SetInstallType(v int32) *CreateMcubeNebulaResourceRequest
	GetInstallType() *int32
	SetMainUrl(v string) *CreateMcubeNebulaResourceRequest
	GetMainUrl() *string
	SetOnexFlag(v bool) *CreateMcubeNebulaResourceRequest
	GetOnexFlag() *bool
	SetPlatform(v string) *CreateMcubeNebulaResourceRequest
	GetPlatform() *string
	SetRepeatNebula(v int32) *CreateMcubeNebulaResourceRequest
	GetRepeatNebula() *int32
	SetResourceType(v int32) *CreateMcubeNebulaResourceRequest
	GetResourceType() *int32
	SetSubUrl(v string) *CreateMcubeNebulaResourceRequest
	GetSubUrl() *string
	SetTenantId(v string) *CreateMcubeNebulaResourceRequest
	GetTenantId() *string
	SetVhost(v string) *CreateMcubeNebulaResourceRequest
	GetVhost() *string
	SetWorkspaceId(v string) *CreateMcubeNebulaResourceRequest
	GetWorkspaceId() *string
}

type CreateMcubeNebulaResourceRequest struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AutoInstall      *int32  `json:"AutoInstall,omitempty" xml:"AutoInstall,omitempty"`
	ClientVersionMax *string `json:"ClientVersionMax,omitempty" xml:"ClientVersionMax,omitempty"`
	ClientVersionMin *string `json:"ClientVersionMin,omitempty" xml:"ClientVersionMin,omitempty"`
	CustomDomainName *string `json:"CustomDomainName,omitempty" xml:"CustomDomainName,omitempty"`
	ExtendInfo       *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	FileUrl          *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	H5Id             *string `json:"H5Id,omitempty" xml:"H5Id,omitempty"`
	H5Name           *string `json:"H5Name,omitempty" xml:"H5Name,omitempty"`
	H5Version        *string `json:"H5Version,omitempty" xml:"H5Version,omitempty"`
	InstallType      *int32  `json:"InstallType,omitempty" xml:"InstallType,omitempty"`
	MainUrl          *string `json:"MainUrl,omitempty" xml:"MainUrl,omitempty"`
	OnexFlag         *bool   `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	Platform         *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RepeatNebula     *int32  `json:"RepeatNebula,omitempty" xml:"RepeatNebula,omitempty"`
	ResourceType     *int32  `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	SubUrl           *string `json:"SubUrl,omitempty" xml:"SubUrl,omitempty"`
	TenantId         *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Vhost            *string `json:"Vhost,omitempty" xml:"Vhost,omitempty"`
	WorkspaceId      *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeNebulaResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeNebulaResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeNebulaResourceRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeNebulaResourceRequest) GetAutoInstall() *int32 {
	return s.AutoInstall
}

func (s *CreateMcubeNebulaResourceRequest) GetClientVersionMax() *string {
	return s.ClientVersionMax
}

func (s *CreateMcubeNebulaResourceRequest) GetClientVersionMin() *string {
	return s.ClientVersionMin
}

func (s *CreateMcubeNebulaResourceRequest) GetCustomDomainName() *string {
	return s.CustomDomainName
}

func (s *CreateMcubeNebulaResourceRequest) GetExtendInfo() *string {
	return s.ExtendInfo
}

func (s *CreateMcubeNebulaResourceRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateMcubeNebulaResourceRequest) GetH5Id() *string {
	return s.H5Id
}

func (s *CreateMcubeNebulaResourceRequest) GetH5Name() *string {
	return s.H5Name
}

func (s *CreateMcubeNebulaResourceRequest) GetH5Version() *string {
	return s.H5Version
}

func (s *CreateMcubeNebulaResourceRequest) GetInstallType() *int32 {
	return s.InstallType
}

func (s *CreateMcubeNebulaResourceRequest) GetMainUrl() *string {
	return s.MainUrl
}

func (s *CreateMcubeNebulaResourceRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *CreateMcubeNebulaResourceRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateMcubeNebulaResourceRequest) GetRepeatNebula() *int32 {
	return s.RepeatNebula
}

func (s *CreateMcubeNebulaResourceRequest) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *CreateMcubeNebulaResourceRequest) GetSubUrl() *string {
	return s.SubUrl
}

func (s *CreateMcubeNebulaResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeNebulaResourceRequest) GetVhost() *string {
	return s.Vhost
}

func (s *CreateMcubeNebulaResourceRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeNebulaResourceRequest) SetAppId(v string) *CreateMcubeNebulaResourceRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetAutoInstall(v int32) *CreateMcubeNebulaResourceRequest {
	s.AutoInstall = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetClientVersionMax(v string) *CreateMcubeNebulaResourceRequest {
	s.ClientVersionMax = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetClientVersionMin(v string) *CreateMcubeNebulaResourceRequest {
	s.ClientVersionMin = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetCustomDomainName(v string) *CreateMcubeNebulaResourceRequest {
	s.CustomDomainName = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetExtendInfo(v string) *CreateMcubeNebulaResourceRequest {
	s.ExtendInfo = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetFileUrl(v string) *CreateMcubeNebulaResourceRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetH5Id(v string) *CreateMcubeNebulaResourceRequest {
	s.H5Id = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetH5Name(v string) *CreateMcubeNebulaResourceRequest {
	s.H5Name = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetH5Version(v string) *CreateMcubeNebulaResourceRequest {
	s.H5Version = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetInstallType(v int32) *CreateMcubeNebulaResourceRequest {
	s.InstallType = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetMainUrl(v string) *CreateMcubeNebulaResourceRequest {
	s.MainUrl = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetOnexFlag(v bool) *CreateMcubeNebulaResourceRequest {
	s.OnexFlag = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetPlatform(v string) *CreateMcubeNebulaResourceRequest {
	s.Platform = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetRepeatNebula(v int32) *CreateMcubeNebulaResourceRequest {
	s.RepeatNebula = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetResourceType(v int32) *CreateMcubeNebulaResourceRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetSubUrl(v string) *CreateMcubeNebulaResourceRequest {
	s.SubUrl = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetTenantId(v string) *CreateMcubeNebulaResourceRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetVhost(v string) *CreateMcubeNebulaResourceRequest {
	s.Vhost = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) SetWorkspaceId(v string) *CreateMcubeNebulaResourceRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeNebulaResourceRequest) Validate() error {
	return dara.Validate(s)
}
