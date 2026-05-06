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
	SetAgentProvider(v string) *ListModelTemplatesRequest
	GetAgentProvider() *string
	SetBizType(v int32) *ListModelTemplatesRequest
	GetBizType() *int32
	SetHasModel(v bool) *ListModelTemplatesRequest
	GetHasModel() *bool
	SetModelTemplateIdList(v []*string) *ListModelTemplatesRequest
	GetModelTemplateIdList() []*string
	SetPageNumber(v int32) *ListModelTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelTemplatesRequest
	GetPageSize() *int32
}

type ListModelTemplatesRequest struct {
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
	// true
	HasModel            *bool     `json:"HasModel,omitempty" xml:"HasModel,omitempty"`
	ModelTemplateIdList []*string `json:"ModelTemplateIdList,omitempty" xml:"ModelTemplateIdList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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

func (s *ListModelTemplatesRequest) GetAgentProvider() *string {
	return s.AgentProvider
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

func (s *ListModelTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelTemplatesRequest) SetAgentPlatform(v string) *ListModelTemplatesRequest {
	s.AgentPlatform = &v
	return s
}

func (s *ListModelTemplatesRequest) SetAgentProvider(v string) *ListModelTemplatesRequest {
	s.AgentProvider = &v
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

func (s *ListModelTemplatesRequest) SetPageNumber(v int32) *ListModelTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelTemplatesRequest) SetPageSize(v int32) *ListModelTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
