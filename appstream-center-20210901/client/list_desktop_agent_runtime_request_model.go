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
	AgentInstanceStatuses []*string `json:"AgentInstanceStatuses,omitempty" xml:"AgentInstanceStatuses,omitempty" type:"Repeated"`
	AgentInstanceVersions []*string `json:"AgentInstanceVersions,omitempty" xml:"AgentInstanceVersions,omitempty" type:"Repeated"`
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string   `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	AuthUsers     []*string `json:"AuthUsers,omitempty" xml:"AuthUsers,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// true
	ChannelConfigure *bool `json:"ChannelConfigure,omitempty" xml:"ChannelConfigure,omitempty"`
	// example:
	//
	// Admin
	DeploymentSource *string   `json:"DeploymentSource,omitempty" xml:"DeploymentSource,omitempty"`
	DesktopIds       []*string `json:"DesktopIds,omitempty" xml:"DesktopIds,omitempty" type:"Repeated"`
	DesktopNames     []*string `json:"DesktopNames,omitempty" xml:"DesktopNames,omitempty" type:"Repeated"`
	DesktopStatuses  []*string `json:"DesktopStatuses,omitempty" xml:"DesktopStatuses,omitempty" type:"Repeated"`
	// example:
	//
	// true
	HasAuthUser *bool `json:"HasAuthUser,omitempty" xml:"HasAuthUser,omitempty"`
	// example:
	//
	// true
	HasRisk *bool `json:"HasRisk,omitempty" xml:"HasRisk,omitempty"`
	// example:
	//
	// true
	IncludeRiskInfo *bool `json:"IncludeRiskInfo,omitempty" xml:"IncludeRiskInfo,omitempty"`
	// example:
	//
	// true
	ModelConfigure *bool `json:"ModelConfigure,omitempty" xml:"ModelConfigure,omitempty"`
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceIds     []*string `json:"ResourceIds,omitempty" xml:"ResourceIds,omitempty" type:"Repeated"`
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
