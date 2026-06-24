// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDesktopAgentRuntimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentInstanceStatuses(v []*string) *ListDesktopAgentRuntimeRequest
	GetAgentInstanceStatuses() []*string
	SetAgentInstanceVersions(v []*string) *ListDesktopAgentRuntimeRequest
	GetAgentInstanceVersions() []*string
	SetAgentPlatform(v string) *ListDesktopAgentRuntimeRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *ListDesktopAgentRuntimeRequest
	GetAgentProvider() *string
	SetAuthUsers(v []*string) *ListDesktopAgentRuntimeRequest
	GetAuthUsers() []*string
	SetBizType(v int32) *ListDesktopAgentRuntimeRequest
	GetBizType() *int32
	SetChannelConfigure(v bool) *ListDesktopAgentRuntimeRequest
	GetChannelConfigure() *bool
	SetDeploymentSource(v string) *ListDesktopAgentRuntimeRequest
	GetDeploymentSource() *string
	SetDesktopIds(v []*string) *ListDesktopAgentRuntimeRequest
	GetDesktopIds() []*string
	SetDesktopNames(v []*string) *ListDesktopAgentRuntimeRequest
	GetDesktopNames() []*string
	SetDesktopStatuses(v []*string) *ListDesktopAgentRuntimeRequest
	GetDesktopStatuses() []*string
	SetHasAuthUser(v bool) *ListDesktopAgentRuntimeRequest
	GetHasAuthUser() *bool
	SetHasRisk(v bool) *ListDesktopAgentRuntimeRequest
	GetHasRisk() *bool
	SetIncludeRiskInfo(v bool) *ListDesktopAgentRuntimeRequest
	GetIncludeRiskInfo() *bool
	SetManagementStatus(v string) *ListDesktopAgentRuntimeRequest
	GetManagementStatus() *string
	SetModelConfigure(v bool) *ListDesktopAgentRuntimeRequest
	GetModelConfigure() *bool
	SetModelTemplateId(v string) *ListDesktopAgentRuntimeRequest
	GetModelTemplateId() *string
	SetPageNumber(v int32) *ListDesktopAgentRuntimeRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDesktopAgentRuntimeRequest
	GetPageSize() *int32
	SetResourceGroupId(v string) *ListDesktopAgentRuntimeRequest
	GetResourceGroupId() *string
	SetResourceIds(v []*string) *ListDesktopAgentRuntimeRequest
	GetResourceIds() []*string
}

type ListDesktopAgentRuntimeRequest struct {
	// The list of agent instance statuses.
	//
	// example:
	//
	// Running
	AgentInstanceStatuses []*string `json:"AgentInstanceStatuses,omitempty" xml:"AgentInstanceStatuses,omitempty" type:"Repeated"`
	// The list of agent instance versions.
	//
	// example:
	//
	// OpenClaw 2026.4.15
	AgentInstanceVersions []*string `json:"AgentInstanceVersions,omitempty" xml:"AgentInstanceVersions,omitempty" type:"Repeated"`
	// The agent platform.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The name of the agent provider.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The list of authorized users.
	//
	// example:
	//
	// user001
	AuthUsers []*string `json:"AuthUsers,omitempty" xml:"AuthUsers,omitempty" type:"Repeated"`
	// The business type.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Specifies whether the third-party channel is configured.
	//
	// example:
	//
	// true
	ChannelConfigure *bool `json:"ChannelConfigure,omitempty" xml:"ChannelConfigure,omitempty"`
	// The deployment source.
	//
	// example:
	//
	// Admin
	DeploymentSource *string `json:"DeploymentSource,omitempty" xml:"DeploymentSource,omitempty"`
	// The list of agent runtime IDs.
	//
	// example:
	//
	// jvs-xxxxx
	DesktopIds []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	// The list of agent runtime names.
	//
	// example:
	//
	// Desktop001
	DesktopNames []*string `json:"DesktopNames,omitempty" xml:"DesktopNames,omitempty" type:"Repeated"`
	// The list of cloud computer statuses.
	//
	// example:
	//
	// Running
	DesktopStatuses []*string `json:"DesktopStatuses,omitempty" xml:"DesktopStatuses,omitempty" type:"Repeated"`
	// Specifies whether authorized users exist.
	//
	// example:
	//
	// true
	HasAuthUser *bool `json:"HasAuthUser,omitempty" xml:"HasAuthUser,omitempty"`
	// Specifies whether a risk exists. Used to filter cloud computers with or without risks. This parameter takes effect only when IncludeRiskInfo is set to true.
	//
	// Set to true to return only records with risks. Set to false to return only records without risks. If not specified, no filtering is applied.
	//
	// example:
	//
	// true
	HasRisk *bool `json:"HasRisk,omitempty" xml:"HasRisk,omitempty"`
	// Specifies whether to query and return risk information. Default value: false. When set to true, the response includes the RiskInfo field, and the HasRisk filter condition takes effect.
	//
	// example:
	//
	// true
	IncludeRiskInfo *bool `json:"IncludeRiskInfo,omitempty" xml:"IncludeRiskInfo,omitempty"`
	// example:
	//
	// Hibernated
	ManagementStatus *string `json:"ManagementStatus,omitempty" xml:"ManagementStatus,omitempty"`
	// Specifies whether the model is configured.
	//
	// example:
	//
	// true
	ModelConfigure *bool `json:"ModelConfigure,omitempty" xml:"ModelConfigure,omitempty"`
	// The model group ID.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// The page number, starting from 1. Values 0 and 1 return the same result.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The list of resource IDs (underlying real resource IDs).
	//
	// example:
	//
	// ecd-xxxxx
	ResourceIds []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
}

func (s ListDesktopAgentRuntimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDesktopAgentRuntimeRequest) GoString() string {
	return s.String()
}

func (s *ListDesktopAgentRuntimeRequest) GetAgentInstanceStatuses() []*string {
	return s.AgentInstanceStatuses
}

func (s *ListDesktopAgentRuntimeRequest) GetAgentInstanceVersions() []*string {
	return s.AgentInstanceVersions
}

func (s *ListDesktopAgentRuntimeRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListDesktopAgentRuntimeRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListDesktopAgentRuntimeRequest) GetAuthUsers() []*string {
	return s.AuthUsers
}

func (s *ListDesktopAgentRuntimeRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListDesktopAgentRuntimeRequest) GetChannelConfigure() *bool {
	return s.ChannelConfigure
}

func (s *ListDesktopAgentRuntimeRequest) GetDeploymentSource() *string {
	return s.DeploymentSource
}

func (s *ListDesktopAgentRuntimeRequest) GetDesktopIds() []*string {
	return s.DesktopIds
}

func (s *ListDesktopAgentRuntimeRequest) GetDesktopNames() []*string {
	return s.DesktopNames
}

func (s *ListDesktopAgentRuntimeRequest) GetDesktopStatuses() []*string {
	return s.DesktopStatuses
}

func (s *ListDesktopAgentRuntimeRequest) GetHasAuthUser() *bool {
	return s.HasAuthUser
}

func (s *ListDesktopAgentRuntimeRequest) GetHasRisk() *bool {
	return s.HasRisk
}

func (s *ListDesktopAgentRuntimeRequest) GetIncludeRiskInfo() *bool {
	return s.IncludeRiskInfo
}

func (s *ListDesktopAgentRuntimeRequest) GetManagementStatus() *string {
	return s.ManagementStatus
}

func (s *ListDesktopAgentRuntimeRequest) GetModelConfigure() *bool {
	return s.ModelConfigure
}

func (s *ListDesktopAgentRuntimeRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListDesktopAgentRuntimeRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDesktopAgentRuntimeRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDesktopAgentRuntimeRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListDesktopAgentRuntimeRequest) GetResourceIds() []*string {
	return s.ResourceIds
}

func (s *ListDesktopAgentRuntimeRequest) SetAgentInstanceStatuses(v []*string) *ListDesktopAgentRuntimeRequest {
	s.AgentInstanceStatuses = v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetAgentInstanceVersions(v []*string) *ListDesktopAgentRuntimeRequest {
	s.AgentInstanceVersions = v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetAgentPlatform(v string) *ListDesktopAgentRuntimeRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetAgentProvider(v string) *ListDesktopAgentRuntimeRequest {
	s.AgentProvider = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetAuthUsers(v []*string) *ListDesktopAgentRuntimeRequest {
	s.AuthUsers = v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetBizType(v int32) *ListDesktopAgentRuntimeRequest {
	s.BizType = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetChannelConfigure(v bool) *ListDesktopAgentRuntimeRequest {
	s.ChannelConfigure = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetDeploymentSource(v string) *ListDesktopAgentRuntimeRequest {
	s.DeploymentSource = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetDesktopIds(v []*string) *ListDesktopAgentRuntimeRequest {
	s.DesktopIds = v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetDesktopNames(v []*string) *ListDesktopAgentRuntimeRequest {
	s.DesktopNames = v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetDesktopStatuses(v []*string) *ListDesktopAgentRuntimeRequest {
	s.DesktopStatuses = v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetHasAuthUser(v bool) *ListDesktopAgentRuntimeRequest {
	s.HasAuthUser = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetHasRisk(v bool) *ListDesktopAgentRuntimeRequest {
	s.HasRisk = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetIncludeRiskInfo(v bool) *ListDesktopAgentRuntimeRequest {
	s.IncludeRiskInfo = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetManagementStatus(v string) *ListDesktopAgentRuntimeRequest {
	s.ManagementStatus = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetModelConfigure(v bool) *ListDesktopAgentRuntimeRequest {
	s.ModelConfigure = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetModelTemplateId(v string) *ListDesktopAgentRuntimeRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetPageNumber(v int32) *ListDesktopAgentRuntimeRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetPageSize(v int32) *ListDesktopAgentRuntimeRequest {
	s.PageSize = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetResourceGroupId(v string) *ListDesktopAgentRuntimeRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) SetResourceIds(v []*string) *ListDesktopAgentRuntimeRequest {
	s.ResourceIds = v
	return s
}

func (s *ListDesktopAgentRuntimeRequest) Validate() error {
	return dara.Validate(s)
}
