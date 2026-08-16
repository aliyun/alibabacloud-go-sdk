// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *ListModelTemplatesRequest
	GetAgentPlatform() *string
	SetAgentPlatformList(v []*string) *ListModelTemplatesRequest
	GetAgentPlatformList() []*string
	SetAgentProvider(v string) *ListModelTemplatesRequest
	GetAgentProvider() *string
	SetAgentProviderList(v []*string) *ListModelTemplatesRequest
	GetAgentProviderList() []*string
	SetBizType(v int32) *ListModelTemplatesRequest
	GetBizType() *int32
	SetHasModel(v bool) *ListModelTemplatesRequest
	GetHasModel() *bool
	SetModelTemplateIdList(v []*string) *ListModelTemplatesRequest
	GetModelTemplateIdList() []*string
	SetName(v string) *ListModelTemplatesRequest
	GetName() *string
	SetPageNumber(v int32) *ListModelTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelTemplatesRequest
	GetPageSize() *int32
	SetRefScope(v string) *ListModelTemplatesRequest
	GetRefScope() *string
	SetSource(v string) *ListModelTemplatesRequest
	GetSource() *string
}

type ListModelTemplatesRequest struct {
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
	ModelTemplateIdList []*string `json:"ModelTemplateIdList,omitempty" xml:"ModelTemplateIdList,omitempty" type:"Repeated"`
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

func (s ListModelTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListModelTemplatesRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListModelTemplatesRequest) GetAgentPlatformList() []*string {
	return s.AgentPlatformList
}

func (s *ListModelTemplatesRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListModelTemplatesRequest) GetAgentProviderList() []*string {
	return s.AgentProviderList
}

func (s *ListModelTemplatesRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListModelTemplatesRequest) GetHasModel() *bool {
	return s.HasModel
}

func (s *ListModelTemplatesRequest) GetModelTemplateIdList() []*string {
	return s.ModelTemplateIdList
}

func (s *ListModelTemplatesRequest) GetName() *string {
	return s.Name
}

func (s *ListModelTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelTemplatesRequest) GetRefScope() *string {
	return s.RefScope
}

func (s *ListModelTemplatesRequest) GetSource() *string {
	return s.Source
}

func (s *ListModelTemplatesRequest) SetAgentPlatform(v string) *ListModelTemplatesRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListModelTemplatesRequest) SetAgentPlatformList(v []*string) *ListModelTemplatesRequest {
	s.AgentPlatformList = v
	return s
}

func (s *ListModelTemplatesRequest) SetAgentProvider(v string) *ListModelTemplatesRequest {
	s.AgentProvider = &v
	return s
}

func (s *ListModelTemplatesRequest) SetAgentProviderList(v []*string) *ListModelTemplatesRequest {
	s.AgentProviderList = v
	return s
}

func (s *ListModelTemplatesRequest) SetBizType(v int32) *ListModelTemplatesRequest {
	s.BizType = &v
	return s
}

func (s *ListModelTemplatesRequest) SetHasModel(v bool) *ListModelTemplatesRequest {
	s.HasModel = &v
	return s
}

func (s *ListModelTemplatesRequest) SetModelTemplateIdList(v []*string) *ListModelTemplatesRequest {
	s.ModelTemplateIdList = v
	return s
}

func (s *ListModelTemplatesRequest) SetName(v string) *ListModelTemplatesRequest {
	s.Name = &v
	return s
}

func (s *ListModelTemplatesRequest) SetPageNumber(v int32) *ListModelTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelTemplatesRequest) SetPageSize(v int32) *ListModelTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelTemplatesRequest) SetRefScope(v string) *ListModelTemplatesRequest {
	s.RefScope = &v
	return s
}

func (s *ListModelTemplatesRequest) SetSource(v string) *ListModelTemplatesRequest {
	s.Source = &v
	return s
}

func (s *ListModelTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
