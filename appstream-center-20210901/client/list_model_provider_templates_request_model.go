// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelProviderTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentPlatform(v string) *ListModelProviderTemplatesRequest
	GetAgentPlatform() *string
	SetAgentProvider(v string) *ListModelProviderTemplatesRequest
	GetAgentProvider() *string
	SetBizType(v int32) *ListModelProviderTemplatesRequest
	GetBizType() *int32
	SetModelTemplateId(v string) *ListModelProviderTemplatesRequest
	GetModelTemplateId() *string
	SetPageNumber(v int32) *ListModelProviderTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelProviderTemplatesRequest
	GetPageSize() *int32
	SetProviderName(v string) *ListModelProviderTemplatesRequest
	GetProviderName() *string
	SetProviderTemplateIds(v []*string) *ListModelProviderTemplatesRequest
	GetProviderTemplateIds() []*string
}

type ListModelProviderTemplatesRequest struct {
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
	ProviderName        *string   `json:"ProviderName,omitempty" xml:"ProviderName,omitempty"`
	ProviderTemplateIds []*string `json:"ProviderTemplateIds,omitempty" xml:"ProviderTemplateIds,omitempty" type:"Repeated"`
}

func (s ListModelProviderTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelProviderTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListModelProviderTemplatesRequest) GetAgentPlatform() *string {
	return s.AgentPlatform
}

func (s *ListModelProviderTemplatesRequest) GetAgentProvider() *string {
	return s.AgentProvider
}

func (s *ListModelProviderTemplatesRequest) GetBizType() *int32 {
	return s.BizType
}

func (s *ListModelProviderTemplatesRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ListModelProviderTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelProviderTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelProviderTemplatesRequest) GetProviderName() *string {
	return s.ProviderName
}

func (s *ListModelProviderTemplatesRequest) GetProviderTemplateIds() []*string {
	return s.ProviderTemplateIds
}

func (s *ListModelProviderTemplatesRequest) SetAgentPlatform(v string) *ListModelProviderTemplatesRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListModelProviderTemplatesRequest) SetAgentProvider(v string) *ListModelProviderTemplatesRequest {
	s.AgentProvider = &v
	return s
}

func (s *ListModelProviderTemplatesRequest) SetBizType(v int32) *ListModelProviderTemplatesRequest {
	s.BizType = &v
	return s
}

func (s *ListModelProviderTemplatesRequest) SetModelTemplateId(v string) *ListModelProviderTemplatesRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *ListModelProviderTemplatesRequest) SetPageNumber(v int32) *ListModelProviderTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelProviderTemplatesRequest) SetPageSize(v int32) *ListModelProviderTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelProviderTemplatesRequest) SetProviderName(v string) *ListModelProviderTemplatesRequest {
	s.ProviderName = &v
	return s
}

func (s *ListModelProviderTemplatesRequest) SetProviderTemplateIds(v []*string) *ListModelProviderTemplatesRequest {
	s.ProviderTemplateIds = v
	return s
}

func (s *ListModelProviderTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
