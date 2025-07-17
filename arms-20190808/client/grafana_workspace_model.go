// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrafanaWorkspace interface {
	dara.Model
	String() string
	GoString() string
	SetCommercial(v bool) *GrafanaWorkspace
	GetCommercial() *bool
	SetDeployType(v string) *GrafanaWorkspace
	GetDeployType() *string
	SetDescription(v string) *GrafanaWorkspace
	GetDescription() *string
	SetEndTime(v float32) *GrafanaWorkspace
	GetEndTime() *float32
	SetGmtCreate(v float32) *GrafanaWorkspace
	GetGmtCreate() *float32
	SetGrafanaVersion(v string) *GrafanaWorkspace
	GetGrafanaVersion() *string
	SetGrafanaWorkspaceDomain(v string) *GrafanaWorkspace
	GetGrafanaWorkspaceDomain() *string
	SetGrafanaWorkspaceDomainStatus(v string) *GrafanaWorkspace
	GetGrafanaWorkspaceDomainStatus() *string
	SetGrafanaWorkspaceEdition(v string) *GrafanaWorkspace
	GetGrafanaWorkspaceEdition() *string
	SetGrafanaWorkspaceId(v string) *GrafanaWorkspace
	GetGrafanaWorkspaceId() *string
	SetGrafanaWorkspaceIp(v string) *GrafanaWorkspace
	GetGrafanaWorkspaceIp() *string
	SetGrafanaWorkspaceName(v string) *GrafanaWorkspace
	GetGrafanaWorkspaceName() *string
	SetMaxAccount(v string) *GrafanaWorkspace
	GetMaxAccount() *string
	SetNtmId(v string) *GrafanaWorkspace
	GetNtmId() *string
	SetPersonalDomain(v string) *GrafanaWorkspace
	GetPersonalDomain() *string
	SetPersonalDomainPrefix(v string) *GrafanaWorkspace
	GetPersonalDomainPrefix() *string
	SetPrivateDomain(v string) *GrafanaWorkspace
	GetPrivateDomain() *string
	SetPrivateIp(v string) *GrafanaWorkspace
	GetPrivateIp() *string
	SetProtocol(v string) *GrafanaWorkspace
	GetProtocol() *string
	SetRegionId(v string) *GrafanaWorkspace
	GetRegionId() *string
	SetResourceGroupId(v string) *GrafanaWorkspace
	GetResourceGroupId() *string
	SetShareSynced(v bool) *GrafanaWorkspace
	GetShareSynced() *bool
	SetSnatIp(v string) *GrafanaWorkspace
	GetSnatIp() *string
	SetStatus(v string) *GrafanaWorkspace
	GetStatus() *string
	SetTags(v []*GrafanaWorkspaceTags) *GrafanaWorkspace
	GetTags() []*GrafanaWorkspaceTags
	SetUpgradeVersion(v []*string) *GrafanaWorkspace
	GetUpgradeVersion() []*string
	SetUserId(v string) *GrafanaWorkspace
	GetUserId() *string
}

type GrafanaWorkspace struct {
	Commercial  *bool   `json:"commercial,omitempty" xml:"commercial,omitempty"`
	DeployType  *string `json:"deployType,omitempty" xml:"deployType,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1652803200000
	EndTime *float32 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 创建时间
	GmtCreate              *float32 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	GrafanaVersion         *string  `json:"grafanaVersion,omitempty" xml:"grafanaVersion,omitempty"`
	GrafanaWorkspaceDomain *string  `json:"grafanaWorkspaceDomain,omitempty" xml:"grafanaWorkspaceDomain,omitempty"`
	// example:
	//
	// on
	GrafanaWorkspaceDomainStatus *string `json:"grafanaWorkspaceDomainStatus,omitempty" xml:"grafanaWorkspaceDomainStatus,omitempty"`
	// example:
	//
	// experts_edition/advanced_edition/standard
	GrafanaWorkspaceEdition *string `json:"grafanaWorkspaceEdition,omitempty" xml:"grafanaWorkspaceEdition,omitempty"`
	// example:
	//
	// g-thisisademo666
	GrafanaWorkspaceId *string `json:"grafanaWorkspaceId,omitempty" xml:"grafanaWorkspaceId,omitempty"`
	// example:
	//
	// 127.0.0.1:3000
	GrafanaWorkspaceIp   *string `json:"grafanaWorkspaceIp,omitempty" xml:"grafanaWorkspaceIp,omitempty"`
	GrafanaWorkspaceName *string `json:"grafanaWorkspaceName,omitempty" xml:"grafanaWorkspaceName,omitempty"`
	// example:
	//
	// 10
	MaxAccount           *string `json:"maxAccount,omitempty" xml:"maxAccount,omitempty"`
	NtmId                *string `json:"ntmId,omitempty" xml:"ntmId,omitempty"`
	PersonalDomain       *string `json:"personalDomain,omitempty" xml:"personalDomain,omitempty"`
	PersonalDomainPrefix *string `json:"personalDomainPrefix,omitempty" xml:"personalDomainPrefix,omitempty"`
	PrivateDomain        *string `json:"privateDomain,omitempty" xml:"privateDomain,omitempty"`
	PrivateIp            *string `json:"privateIp,omitempty" xml:"privateIp,omitempty"`
	// example:
	//
	// http/https
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	ShareSynced     *bool   `json:"shareSynced,omitempty" xml:"shareSynced,omitempty"`
	// example:
	//
	// 1.1.1.1
	SnatIp *string `json:"snatIp,omitempty" xml:"snatIp,omitempty"`
	// example:
	//
	// Starting/Running/Stop/DeleteSucceed
	Status         *string                 `json:"status,omitempty" xml:"status,omitempty"`
	Tags           []*GrafanaWorkspaceTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	UpgradeVersion []*string               `json:"upgradeVersion,omitempty" xml:"upgradeVersion,omitempty" type:"Repeated"`
	// example:
	//
	// 66666666
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s GrafanaWorkspace) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspace) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspace) GetCommercial() *bool {
	return s.Commercial
}

func (s *GrafanaWorkspace) GetDeployType() *string {
	return s.DeployType
}

func (s *GrafanaWorkspace) GetDescription() *string {
	return s.Description
}

func (s *GrafanaWorkspace) GetEndTime() *float32 {
	return s.EndTime
}

func (s *GrafanaWorkspace) GetGmtCreate() *float32 {
	return s.GmtCreate
}

func (s *GrafanaWorkspace) GetGrafanaVersion() *string {
	return s.GrafanaVersion
}

func (s *GrafanaWorkspace) GetGrafanaWorkspaceDomain() *string {
	return s.GrafanaWorkspaceDomain
}

func (s *GrafanaWorkspace) GetGrafanaWorkspaceDomainStatus() *string {
	return s.GrafanaWorkspaceDomainStatus
}

func (s *GrafanaWorkspace) GetGrafanaWorkspaceEdition() *string {
	return s.GrafanaWorkspaceEdition
}

func (s *GrafanaWorkspace) GetGrafanaWorkspaceId() *string {
	return s.GrafanaWorkspaceId
}

func (s *GrafanaWorkspace) GetGrafanaWorkspaceIp() *string {
	return s.GrafanaWorkspaceIp
}

func (s *GrafanaWorkspace) GetGrafanaWorkspaceName() *string {
	return s.GrafanaWorkspaceName
}

func (s *GrafanaWorkspace) GetMaxAccount() *string {
	return s.MaxAccount
}

func (s *GrafanaWorkspace) GetNtmId() *string {
	return s.NtmId
}

func (s *GrafanaWorkspace) GetPersonalDomain() *string {
	return s.PersonalDomain
}

func (s *GrafanaWorkspace) GetPersonalDomainPrefix() *string {
	return s.PersonalDomainPrefix
}

func (s *GrafanaWorkspace) GetPrivateDomain() *string {
	return s.PrivateDomain
}

func (s *GrafanaWorkspace) GetPrivateIp() *string {
	return s.PrivateIp
}

func (s *GrafanaWorkspace) GetProtocol() *string {
	return s.Protocol
}

func (s *GrafanaWorkspace) GetRegionId() *string {
	return s.RegionId
}

func (s *GrafanaWorkspace) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GrafanaWorkspace) GetShareSynced() *bool {
	return s.ShareSynced
}

func (s *GrafanaWorkspace) GetSnatIp() *string {
	return s.SnatIp
}

func (s *GrafanaWorkspace) GetStatus() *string {
	return s.Status
}

func (s *GrafanaWorkspace) GetTags() []*GrafanaWorkspaceTags {
	return s.Tags
}

func (s *GrafanaWorkspace) GetUpgradeVersion() []*string {
	return s.UpgradeVersion
}

func (s *GrafanaWorkspace) GetUserId() *string {
	return s.UserId
}

func (s *GrafanaWorkspace) SetCommercial(v bool) *GrafanaWorkspace {
	s.Commercial = &v
	return s
}

func (s *GrafanaWorkspace) SetDeployType(v string) *GrafanaWorkspace {
	s.DeployType = &v
	return s
}

func (s *GrafanaWorkspace) SetDescription(v string) *GrafanaWorkspace {
	s.Description = &v
	return s
}

func (s *GrafanaWorkspace) SetEndTime(v float32) *GrafanaWorkspace {
	s.EndTime = &v
	return s
}

func (s *GrafanaWorkspace) SetGmtCreate(v float32) *GrafanaWorkspace {
	s.GmtCreate = &v
	return s
}

func (s *GrafanaWorkspace) SetGrafanaVersion(v string) *GrafanaWorkspace {
	s.GrafanaVersion = &v
	return s
}

func (s *GrafanaWorkspace) SetGrafanaWorkspaceDomain(v string) *GrafanaWorkspace {
	s.GrafanaWorkspaceDomain = &v
	return s
}

func (s *GrafanaWorkspace) SetGrafanaWorkspaceDomainStatus(v string) *GrafanaWorkspace {
	s.GrafanaWorkspaceDomainStatus = &v
	return s
}

func (s *GrafanaWorkspace) SetGrafanaWorkspaceEdition(v string) *GrafanaWorkspace {
	s.GrafanaWorkspaceEdition = &v
	return s
}

func (s *GrafanaWorkspace) SetGrafanaWorkspaceId(v string) *GrafanaWorkspace {
	s.GrafanaWorkspaceId = &v
	return s
}

func (s *GrafanaWorkspace) SetGrafanaWorkspaceIp(v string) *GrafanaWorkspace {
	s.GrafanaWorkspaceIp = &v
	return s
}

func (s *GrafanaWorkspace) SetGrafanaWorkspaceName(v string) *GrafanaWorkspace {
	s.GrafanaWorkspaceName = &v
	return s
}

func (s *GrafanaWorkspace) SetMaxAccount(v string) *GrafanaWorkspace {
	s.MaxAccount = &v
	return s
}

func (s *GrafanaWorkspace) SetNtmId(v string) *GrafanaWorkspace {
	s.NtmId = &v
	return s
}

func (s *GrafanaWorkspace) SetPersonalDomain(v string) *GrafanaWorkspace {
	s.PersonalDomain = &v
	return s
}

func (s *GrafanaWorkspace) SetPersonalDomainPrefix(v string) *GrafanaWorkspace {
	s.PersonalDomainPrefix = &v
	return s
}

func (s *GrafanaWorkspace) SetPrivateDomain(v string) *GrafanaWorkspace {
	s.PrivateDomain = &v
	return s
}

func (s *GrafanaWorkspace) SetPrivateIp(v string) *GrafanaWorkspace {
	s.PrivateIp = &v
	return s
}

func (s *GrafanaWorkspace) SetProtocol(v string) *GrafanaWorkspace {
	s.Protocol = &v
	return s
}

func (s *GrafanaWorkspace) SetRegionId(v string) *GrafanaWorkspace {
	s.RegionId = &v
	return s
}

func (s *GrafanaWorkspace) SetResourceGroupId(v string) *GrafanaWorkspace {
	s.ResourceGroupId = &v
	return s
}

func (s *GrafanaWorkspace) SetShareSynced(v bool) *GrafanaWorkspace {
	s.ShareSynced = &v
	return s
}

func (s *GrafanaWorkspace) SetSnatIp(v string) *GrafanaWorkspace {
	s.SnatIp = &v
	return s
}

func (s *GrafanaWorkspace) SetStatus(v string) *GrafanaWorkspace {
	s.Status = &v
	return s
}

func (s *GrafanaWorkspace) SetTags(v []*GrafanaWorkspaceTags) *GrafanaWorkspace {
	s.Tags = v
	return s
}

func (s *GrafanaWorkspace) SetUpgradeVersion(v []*string) *GrafanaWorkspace {
	s.UpgradeVersion = v
	return s
}

func (s *GrafanaWorkspace) SetUserId(v string) *GrafanaWorkspace {
	s.UserId = &v
	return s
}

func (s *GrafanaWorkspace) Validate() error {
	return dara.Validate(s)
}

type GrafanaWorkspaceTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s GrafanaWorkspaceTags) String() string {
	return dara.Prettify(s)
}

func (s GrafanaWorkspaceTags) GoString() string {
	return s.String()
}

func (s *GrafanaWorkspaceTags) GetKey() *string {
	return s.Key
}

func (s *GrafanaWorkspaceTags) GetValue() *string {
	return s.Value
}

func (s *GrafanaWorkspaceTags) SetKey(v string) *GrafanaWorkspaceTags {
	s.Key = &v
	return s
}

func (s *GrafanaWorkspaceTags) SetValue(v string) *GrafanaWorkspaceTags {
	s.Value = &v
	return s
}

func (s *GrafanaWorkspaceTags) Validate() error {
	return dara.Validate(s)
}
