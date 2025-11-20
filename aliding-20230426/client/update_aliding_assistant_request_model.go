// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlidingAssistantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *UpdateAlidingAssistantRequest
	GetAssistantId() *string
	SetDescription(v string) *UpdateAlidingAssistantRequest
	GetDescription() *string
	SetExt(v map[string]*string) *UpdateAlidingAssistantRequest
	GetExt() map[string]*string
	SetFallbackContent(v string) *UpdateAlidingAssistantRequest
	GetFallbackContent() *string
	SetFeature(v map[string]*string) *UpdateAlidingAssistantRequest
	GetFeature() map[string]*string
	SetIcon(v string) *UpdateAlidingAssistantRequest
	GetIcon() *string
	SetInstructions(v string) *UpdateAlidingAssistantRequest
	GetInstructions() *string
	SetName(v string) *UpdateAlidingAssistantRequest
	GetName() *string
	SetRecommendPrompts(v []*string) *UpdateAlidingAssistantRequest
	GetRecommendPrompts() []*string
	SetTenantContext(v *UpdateAlidingAssistantRequestTenantContext) *UpdateAlidingAssistantRequest
	GetTenantContext() *UpdateAlidingAssistantRequestTenantContext
	SetWelcomeContent(v string) *UpdateAlidingAssistantRequest
	GetWelcomeContent() *string
}

type UpdateAlidingAssistantRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234
	AssistantId *string            `json:"AssistantId,omitempty" xml:"AssistantId,omitempty"`
	Description *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Ext         map[string]*string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// fallbackContent
	FallbackContent *string            `json:"FallbackContent,omitempty" xml:"FallbackContent,omitempty"`
	Feature         map[string]*string `json:"Feature,omitempty" xml:"Feature,omitempty"`
	// example:
	//
	// @lADPDetfgMsFFUvNAkjNAkg
	Icon             *string                                     `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Instructions     *string                                     `json:"Instructions,omitempty" xml:"Instructions,omitempty"`
	Name             *string                                     `json:"Name,omitempty" xml:"Name,omitempty"`
	RecommendPrompts []*string                                   `json:"RecommendPrompts,omitempty" xml:"RecommendPrompts,omitempty" type:"Repeated"`
	TenantContext    *UpdateAlidingAssistantRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	WelcomeContent   *string                                     `json:"WelcomeContent,omitempty" xml:"WelcomeContent,omitempty"`
}

func (s UpdateAlidingAssistantRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlidingAssistantRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlidingAssistantRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *UpdateAlidingAssistantRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAlidingAssistantRequest) GetExt() map[string]*string {
	return s.Ext
}

func (s *UpdateAlidingAssistantRequest) GetFallbackContent() *string {
	return s.FallbackContent
}

func (s *UpdateAlidingAssistantRequest) GetFeature() map[string]*string {
	return s.Feature
}

func (s *UpdateAlidingAssistantRequest) GetIcon() *string {
	return s.Icon
}

func (s *UpdateAlidingAssistantRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *UpdateAlidingAssistantRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAlidingAssistantRequest) GetRecommendPrompts() []*string {
	return s.RecommendPrompts
}

func (s *UpdateAlidingAssistantRequest) GetTenantContext() *UpdateAlidingAssistantRequestTenantContext {
	return s.TenantContext
}

func (s *UpdateAlidingAssistantRequest) GetWelcomeContent() *string {
	return s.WelcomeContent
}

func (s *UpdateAlidingAssistantRequest) SetAssistantId(v string) *UpdateAlidingAssistantRequest {
	s.AssistantId = &v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetDescription(v string) *UpdateAlidingAssistantRequest {
	s.Description = &v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetExt(v map[string]*string) *UpdateAlidingAssistantRequest {
	s.Ext = v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetFallbackContent(v string) *UpdateAlidingAssistantRequest {
	s.FallbackContent = &v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetFeature(v map[string]*string) *UpdateAlidingAssistantRequest {
	s.Feature = v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetIcon(v string) *UpdateAlidingAssistantRequest {
	s.Icon = &v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetInstructions(v string) *UpdateAlidingAssistantRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetName(v string) *UpdateAlidingAssistantRequest {
	s.Name = &v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetRecommendPrompts(v []*string) *UpdateAlidingAssistantRequest {
	s.RecommendPrompts = v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetTenantContext(v *UpdateAlidingAssistantRequestTenantContext) *UpdateAlidingAssistantRequest {
	s.TenantContext = v
	return s
}

func (s *UpdateAlidingAssistantRequest) SetWelcomeContent(v string) *UpdateAlidingAssistantRequest {
	s.WelcomeContent = &v
	return s
}

func (s *UpdateAlidingAssistantRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAlidingAssistantRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateAlidingAssistantRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlidingAssistantRequestTenantContext) GoString() string {
	return s.String()
}

func (s *UpdateAlidingAssistantRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateAlidingAssistantRequestTenantContext) SetTenantId(v string) *UpdateAlidingAssistantRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *UpdateAlidingAssistantRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
