// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderTemplatesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *ListModelProviderTemplatesShrinkRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *ListModelProviderTemplatesShrinkRequest
	GetAgentProvider() *string
	SetBizType(v int32) *ListModelProviderTemplatesShrinkRequest
	GetBizType() *int32
	SetModelTemplateId(v string) *ListModelProviderTemplatesShrinkRequest
	GetModelTemplateId() *string
	SetPageNumber(v int32) *ListModelProviderTemplatesShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelProviderTemplatesShrinkRequest
	GetPageSize() *int32
	SetProviderName(v string) *ListModelProviderTemplatesShrinkRequest
	GetProviderName() *string
	SetProviderTemplateIdsShrink(v string) *ListModelProviderTemplatesShrinkRequest
	GetProviderTemplateIdsShrink() *string
}

type ListModelProviderTemplatesShrinkRequest struct {
	// example:
	//
	// ENTERPRISE
	AgentPlatform *string `json:"AgentPlatform,omitempty" xml:"AgentPlatform,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// OpenClaw
	AgentProvider *string `json:"AgentProvider,omitempty" xml:"AgentProvider,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	BizType *int32 `json:"BizType,omitempty" xml:"BizType,omitempty"`
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
	// bailian
	ProviderName              *string `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	ProviderTemplateIdsShrink *string `json:"ProviderTemplateIds,omitempty" xml:"ProviderTemplateIds,omitempty"`
}

func (s ListModelProviderTemplatesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderTemplatesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListModelProviderTemplatesShrinkRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListModelProviderTemplatesShrinkRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListModelProviderTemplatesShrinkRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListModelProviderTemplatesShrinkRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListModelProviderTemplatesShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelProviderTemplatesShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelProviderTemplatesShrinkRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListModelProviderTemplatesShrinkRequest) GetProviderTemplateIdsShrink() *string {
	return s.ProviderTemplateIdsShrink
}

func (s *ListModelProviderTemplatesShrinkRequest) SetAgentPlatform(v string) *ListModelProviderTemplatesShrinkRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListModelProviderTemplatesShrinkRequest) SetAgentProvider(v string) *ListModelProviderTemplatesShrinkRequest {
	s.AgentProvider = &v
	return s
}

func (s *ListModelProviderTemplatesShrinkRequest) SetBizType(v int32) *ListModelProviderTemplatesShrinkRequest {
	s.BizType = &v
	return s
}

func (s *ListModelProviderTemplatesShrinkRequest) SetModelTemplateId(v string) *ListModelProviderTemplatesShrinkRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *ListModelProviderTemplatesShrinkRequest) SetPageNumber(v int32) *ListModelProviderTemplatesShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelProviderTemplatesShrinkRequest) SetPageSize(v int32) *ListModelProviderTemplatesShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelProviderTemplatesShrinkRequest) SetProviderName(v string) *ListModelProviderTemplatesShrinkRequest {
	s.ProviderName = &v
	return s
}

func (s *ListModelProviderTemplatesShrinkRequest) SetProviderTemplateIdsShrink(v string) *ListModelProviderTemplatesShrinkRequest {
	s.ProviderTemplateIdsShrink = &v
	return s
}

func (s *ListModelProviderTemplatesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
