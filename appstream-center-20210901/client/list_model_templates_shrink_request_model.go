// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *ListModelTemplatesShrinkRequest
	GetAgentPlatform() *string
	SetAgentPlatformList(v []*string) *ListModelTemplatesShrinkRequest
	GetAgentPlatformList() []*string
	SetAgentProvider(v string) *ListModelTemplatesShrinkRequest
	GetAgentProvider() *string
	SetAgentProviderList(v []*string) *ListModelTemplatesShrinkRequest
	GetAgentProviderList() []*string
	SetBizType(v int32) *ListModelTemplatesShrinkRequest
	GetBizType() *int32
	SetHasModel(v bool) *ListModelTemplatesShrinkRequest
	GetHasModel() *bool
	SetModelTemplateIdListShrink(v string) *ListModelTemplatesShrinkRequest
	GetModelTemplateIdListShrink() *string
	SetName(v string) *ListModelTemplatesShrinkRequest
	GetName() *string
	SetPageNumber(v int32) *ListModelTemplatesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelTemplatesShrinkRequest
	GetPageSize() *int32
	SetRefScope(v string) *ListModelTemplatesShrinkRequest
	GetRefScope() *string
	SetSource(v string) *ListModelTemplatesShrinkRequest
	GetSource() *string
}

type ListModelTemplatesShrinkRequest struct {
	// The Agent platform.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The Agent platform list. Supports COMMON. If specified together with AgentPlatform, AgentPlatform takes precedence and this list is ignored. Defaults to ENTERPRISE if no platform filter is specified. To query Common model groups, explicitly include COMMON. If filtering by Provider simultaneously, set the value to Common.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatformList []*string `json:"AgentPlatformList,omitempty" xml:"AgentPlatformList,omitempty" type:"Repeated"`
	// The Agent provider name.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The Agent provider list. Supports Common. If specified together with AgentProvider, AgentProvider takes precedence and this list is ignored. To query Common model groups, explicitly include COMMON in the platform filter.
	//
	// example:
	//
	// OpenClaw
	AgentProviderList []*string `json:"AgentProviderList,omitempty" xml:"AgentProviderList,omitempty" type:"Repeated"`
	// The business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Specifies whether models are configured in the group.
	//
	// example:
	//
	// true
	HasModel *bool `json:"HasModel,omitempty" xml:"HasModel,omitempty"`
	// The list of template group IDs to filter by.
	ModelTemplateIdListShrink *string `json:"ModelTemplateIdList,omitempty" xml:"ModelTemplateIdList,omitempty"`
	// The model group name. Fuzzy match is supported.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// The authorization scope filter. Valid values: ALL_USER, USER_MIXED, or RESOURCE_MIXED (strictly uppercase. Case variants or unknown values return InvalidParameter). If not specified, no filtering is applied. Unlike create/update operations, the filter scenario allows RESOURCE_MIXED (to filter non-Common model groups).
	//
	// example:
	//
	// ALL_USER
	RefScope *string `json:"RefScope,omitempty" xml:"RefScope,omitempty"`
	// The template source filter. Valid values:
	//
	// - User: tenant-created (default if not specified).
	//
	// - System: system preset.
	//
	// example:
	//
	// User
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s ListModelTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModelTemplatesShrinkRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListModelTemplatesShrinkRequest) GetAgentPlatformList() []*string {
	return s.AgentPlatformList
}

func (s *ListModelTemplatesShrinkRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListModelTemplatesShrinkRequest) GetAgentProviderList() []*string {
	return s.AgentProviderList
}

func (s *ListModelTemplatesShrinkRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListModelTemplatesShrinkRequest) GetHasModel() *bool {
	return s.HasModel
}

func (s *ListModelTemplatesShrinkRequest) GetModelTemplateIdListShrink() *string {
	return s.ModelTemplateIdListShrink
}

func (s *ListModelTemplatesShrinkRequest) GetName() *string {
	return s.Name
}

func (s *ListModelTemplatesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelTemplatesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelTemplatesShrinkRequest) GetRefScope() *string {
	return s.RefScope
}

func (s *ListModelTemplatesShrinkRequest) GetSource() *string {
	return s.Source
}

func (s *ListModelTemplatesShrinkRequest) SetAgentPlatform(v string) *ListModelTemplatesShrinkRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetAgentPlatformList(v []*string) *ListModelTemplatesShrinkRequest {
	s.AgentPlatformList = v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetAgentProvider(v string) *ListModelTemplatesShrinkRequest {
	s.AgentProvider = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetAgentProviderList(v []*string) *ListModelTemplatesShrinkRequest {
	s.AgentProviderList = v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetBizType(v int32) *ListModelTemplatesShrinkRequest {
	s.BizType = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetHasModel(v bool) *ListModelTemplatesShrinkRequest {
	s.HasModel = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetModelTemplateIdListShrink(v string) *ListModelTemplatesShrinkRequest {
	s.ModelTemplateIdListShrink = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetName(v string) *ListModelTemplatesShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetPageNumber(v int32) *ListModelTemplatesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetPageSize(v int32) *ListModelTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetRefScope(v string) *ListModelTemplatesShrinkRequest {
	s.RefScope = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetSource(v string) *ListModelTemplatesShrinkRequest {
	s.Source = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
