// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeUpgradePackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateMcubeUpgradePackageRequest
	GetAppId() *string
	SetAppVersion(v string) *CreateMcubeUpgradePackageRequest
	GetAppVersion() *string
	SetAppstoreUrl(v string) *CreateMcubeUpgradePackageRequest
	GetAppstoreUrl() *string
	SetBundleId(v string) *CreateMcubeUpgradePackageRequest
	GetBundleId() *string
	SetCustomDomainName(v string) *CreateMcubeUpgradePackageRequest
	GetCustomDomainName() *string
	SetDesc(v string) *CreateMcubeUpgradePackageRequest
	GetDesc() *string
	SetDownloadUrl(v string) *CreateMcubeUpgradePackageRequest
	GetDownloadUrl() *string
	SetFileUrl(v string) *CreateMcubeUpgradePackageRequest
	GetFileUrl() *string
	SetHarmonyLabel(v string) *CreateMcubeUpgradePackageRequest
	GetHarmonyLabel() *string
	SetIconFileUrl(v string) *CreateMcubeUpgradePackageRequest
	GetIconFileUrl() *string
	SetInstallAmount(v int32) *CreateMcubeUpgradePackageRequest
	GetInstallAmount() *int32
	SetIosSymbolfileUrl(v string) *CreateMcubeUpgradePackageRequest
	GetIosSymbolfileUrl() *string
	SetIsEnterprise(v int32) *CreateMcubeUpgradePackageRequest
	GetIsEnterprise() *int32
	SetLargeIconUrl(v string) *CreateMcubeUpgradePackageRequest
	GetLargeIconUrl() *string
	SetNeedCheck(v int32) *CreateMcubeUpgradePackageRequest
	GetNeedCheck() *int32
	SetOnexFlag(v bool) *CreateMcubeUpgradePackageRequest
	GetOnexFlag() *bool
	SetPlatform(v string) *CreateMcubeUpgradePackageRequest
	GetPlatform() *string
	SetTenantId(v string) *CreateMcubeUpgradePackageRequest
	GetTenantId() *string
	SetValidDays(v int32) *CreateMcubeUpgradePackageRequest
	GetValidDays() *int32
	SetWorkspaceId(v string) *CreateMcubeUpgradePackageRequest
	GetWorkspaceId() *string
}

type CreateMcubeUpgradePackageRequest struct {
	AppId            *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppVersion       *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	AppstoreUrl      *string `json:"AppstoreUrl,omitempty" xml:"AppstoreUrl,omitempty"`
	BundleId         *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	CustomDomainName *string `json:"CustomDomainName,omitempty" xml:"CustomDomainName,omitempty"`
	Desc             *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	DownloadUrl      *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	FileUrl          *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	HarmonyLabel     *string `json:"HarmonyLabel,omitempty" xml:"HarmonyLabel,omitempty"`
	IconFileUrl      *string `json:"IconFileUrl,omitempty" xml:"IconFileUrl,omitempty"`
	InstallAmount    *int32  `json:"InstallAmount,omitempty" xml:"InstallAmount,omitempty"`
	IosSymbolfileUrl *string `json:"IosSymbolfileUrl,omitempty" xml:"IosSymbolfileUrl,omitempty"`
	IsEnterprise     *int32  `json:"IsEnterprise,omitempty" xml:"IsEnterprise,omitempty"`
	LargeIconUrl     *string `json:"LargeIconUrl,omitempty" xml:"LargeIconUrl,omitempty"`
	NeedCheck        *int32  `json:"NeedCheck,omitempty" xml:"NeedCheck,omitempty"`
	OnexFlag         *bool   `json:"OnexFlag,omitempty" xml:"OnexFlag,omitempty"`
	Platform         *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	TenantId         *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ValidDays        *int32  `json:"ValidDays,omitempty" xml:"ValidDays,omitempty"`
	WorkspaceId      *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateMcubeUpgradePackageRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeUpgradePackageRequest) GoString() string {
	return s.String()
}

func (s *CreateMcubeUpgradePackageRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateMcubeUpgradePackageRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *CreateMcubeUpgradePackageRequest) GetAppstoreUrl() *string {
	return s.AppstoreUrl
}

func (s *CreateMcubeUpgradePackageRequest) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateMcubeUpgradePackageRequest) GetCustomDomainName() *string {
	return s.CustomDomainName
}

func (s *CreateMcubeUpgradePackageRequest) GetDesc() *string {
	return s.Desc
}

func (s *CreateMcubeUpgradePackageRequest) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateMcubeUpgradePackageRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateMcubeUpgradePackageRequest) GetHarmonyLabel() *string {
	return s.HarmonyLabel
}

func (s *CreateMcubeUpgradePackageRequest) GetIconFileUrl() *string {
	return s.IconFileUrl
}

func (s *CreateMcubeUpgradePackageRequest) GetInstallAmount() *int32 {
	return s.InstallAmount
}

func (s *CreateMcubeUpgradePackageRequest) GetIosSymbolfileUrl() *string {
	return s.IosSymbolfileUrl
}

func (s *CreateMcubeUpgradePackageRequest) GetIsEnterprise() *int32 {
	return s.IsEnterprise
}

func (s *CreateMcubeUpgradePackageRequest) GetLargeIconUrl() *string {
	return s.LargeIconUrl
}

func (s *CreateMcubeUpgradePackageRequest) GetNeedCheck() *int32 {
	return s.NeedCheck
}

func (s *CreateMcubeUpgradePackageRequest) GetOnexFlag() *bool {
	return s.OnexFlag
}

func (s *CreateMcubeUpgradePackageRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateMcubeUpgradePackageRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateMcubeUpgradePackageRequest) GetValidDays() *int32 {
	return s.ValidDays
}

func (s *CreateMcubeUpgradePackageRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateMcubeUpgradePackageRequest) SetAppId(v string) *CreateMcubeUpgradePackageRequest {
	s.AppId = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetAppVersion(v string) *CreateMcubeUpgradePackageRequest {
	s.AppVersion = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetAppstoreUrl(v string) *CreateMcubeUpgradePackageRequest {
	s.AppstoreUrl = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetBundleId(v string) *CreateMcubeUpgradePackageRequest {
	s.BundleId = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetCustomDomainName(v string) *CreateMcubeUpgradePackageRequest {
	s.CustomDomainName = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetDesc(v string) *CreateMcubeUpgradePackageRequest {
	s.Desc = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetDownloadUrl(v string) *CreateMcubeUpgradePackageRequest {
	s.DownloadUrl = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetFileUrl(v string) *CreateMcubeUpgradePackageRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetHarmonyLabel(v string) *CreateMcubeUpgradePackageRequest {
	s.HarmonyLabel = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetIconFileUrl(v string) *CreateMcubeUpgradePackageRequest {
	s.IconFileUrl = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetInstallAmount(v int32) *CreateMcubeUpgradePackageRequest {
	s.InstallAmount = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetIosSymbolfileUrl(v string) *CreateMcubeUpgradePackageRequest {
	s.IosSymbolfileUrl = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetIsEnterprise(v int32) *CreateMcubeUpgradePackageRequest {
	s.IsEnterprise = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetLargeIconUrl(v string) *CreateMcubeUpgradePackageRequest {
	s.LargeIconUrl = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetNeedCheck(v int32) *CreateMcubeUpgradePackageRequest {
	s.NeedCheck = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetOnexFlag(v bool) *CreateMcubeUpgradePackageRequest {
	s.OnexFlag = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetPlatform(v string) *CreateMcubeUpgradePackageRequest {
	s.Platform = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetTenantId(v string) *CreateMcubeUpgradePackageRequest {
	s.TenantId = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetValidDays(v int32) *CreateMcubeUpgradePackageRequest {
	s.ValidDays = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) SetWorkspaceId(v string) *CreateMcubeUpgradePackageRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateMcubeUpgradePackageRequest) Validate() error {
	return dara.Validate(s)
}
