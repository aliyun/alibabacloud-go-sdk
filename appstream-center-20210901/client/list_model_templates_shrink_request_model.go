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
	SetAgentProvider(v string) *ListModelTemplatesShrinkRequest
	GetAgentProvider() *string
	SetBizType(v int32) *ListModelTemplatesShrinkRequest
	GetBizType() *int32
	SetHasModel(v bool) *ListModelTemplatesShrinkRequest
	GetHasModel() *bool
	SetModelTemplateIdListShrink(v string) *ListModelTemplatesShrinkRequest
	GetModelTemplateIdListShrink() *string
	SetPageNumber(v int32) *ListModelTemplatesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelTemplatesShrinkRequest
	GetPageSize() *int32
}

type ListModelTemplatesShrinkRequest struct {
	// The Agent platform.
	//
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// The name of the Agent provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// The business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// Specifies whether models have been configured in the group.
	//
	// example:
	//
	// true
	HasModel *bool `json:"HasModel,omitempty" xml:"HasModel,omitempty"`
	// The list of template group IDs used for filtering.
	ModelTemplateIdListShrink *string `json:"ModelTemplateIdList,omitempty" xml:"ModelTemplateIdList,omitempty"`
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

func (s *ListModelTemplatesShrinkRequest) GetAgentProvider() *string {
	return s.AgentProvider
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

func (s *ListModelTemplatesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelTemplatesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelTemplatesShrinkRequest) SetAgentPlatform(v string) *ListModelTemplatesShrinkRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetAgentProvider(v string) *ListModelTemplatesShrinkRequest {
	s.AgentProvider = &v
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

func (s *ListModelTemplatesShrinkRequest) SetPageNumber(v int32) *ListModelTemplatesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) SetPageSize(v int32) *ListModelTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
