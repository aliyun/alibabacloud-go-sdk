// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlidingAssistantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *CreateAlidingAssistantRequest
	GetAppCode() *string
	SetDescription(v string) *CreateAlidingAssistantRequest
	GetDescription() *string
	SetExt(v map[string]*string) *CreateAlidingAssistantRequest
	GetExt() map[string]*string
	SetIcon(v string) *CreateAlidingAssistantRequest
	GetIcon() *string
	SetInstructions(v string) *CreateAlidingAssistantRequest
	GetInstructions() *string
	SetName(v string) *CreateAlidingAssistantRequest
	GetName() *string
	SetRecommendPrompts(v []*string) *CreateAlidingAssistantRequest
	GetRecommendPrompts() []*string
	SetSource(v int32) *CreateAlidingAssistantRequest
	GetSource() *int32
	SetSourceIdentityId(v string) *CreateAlidingAssistantRequest
	GetSourceIdentityId() *string
	SetTenantContext(v *CreateAlidingAssistantRequestTenantContext) *CreateAlidingAssistantRequest
	GetTenantContext() *CreateAlidingAssistantRequestTenantContext
	SetWelcomeContent(v string) *CreateAlidingAssistantRequest
	GetWelcomeContent() *string
}

type CreateAlidingAssistantRequest struct {
	// example:
	//
	// f5cb37a0fb44441ab7b74c6f4a679dd3
	AppCode *string `json:"AppCode,omitempty" xml:"AppCode,omitempty"`
	// This parameter is required.
	Description *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Ext         map[string]*string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// @lADPDetfgMsFFUvNAkjNAkg
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// This parameter is required.
	Instructions *string `json:"Instructions,omitempty" xml:"Instructions,omitempty"`
	// This parameter is required.
	Name             *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	RecommendPrompts []*string `json:"RecommendPrompts,omitempty" xml:"RecommendPrompts,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 11
	Source *int32 `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// chatBot-123
	SourceIdentityId *string                                     `json:"SourceIdentityId,omitempty" xml:"SourceIdentityId,omitempty"`
	TenantContext    *CreateAlidingAssistantRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
	// This parameter is required.
	WelcomeContent *string `json:"WelcomeContent,omitempty" xml:"WelcomeContent,omitempty"`
}

func (s CreateAlidingAssistantRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAlidingAssistantRequest) GoString() string {
	return s.String()
}

func (s *CreateAlidingAssistantRequest) GetAppCode() *string {
	return s.AppCode
}

func (s *CreateAlidingAssistantRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAlidingAssistantRequest) GetExt() map[string]*string {
	return s.Ext
}

func (s *CreateAlidingAssistantRequest) GetIcon() *string {
	return s.Icon
}

func (s *CreateAlidingAssistantRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *CreateAlidingAssistantRequest) GetName() *string {
	return s.Name
}

func (s *CreateAlidingAssistantRequest) GetRecommendPrompts() []*string {
	return s.RecommendPrompts
}

func (s *CreateAlidingAssistantRequest) GetSource() *int32 {
	return s.Source
}

func (s *CreateAlidingAssistantRequest) GetSourceIdentityId() *string {
	return s.SourceIdentityId
}

func (s *CreateAlidingAssistantRequest) GetTenantContext() *CreateAlidingAssistantRequestTenantContext {
	return s.TenantContext
}

func (s *CreateAlidingAssistantRequest) GetWelcomeContent() *string {
	return s.WelcomeContent
}

func (s *CreateAlidingAssistantRequest) SetAppCode(v string) *CreateAlidingAssistantRequest {
	s.AppCode = &v
	return s
}

func (s *CreateAlidingAssistantRequest) SetDescription(v string) *CreateAlidingAssistantRequest {
	s.Description = &v
	return s
}

func (s *CreateAlidingAssistantRequest) SetExt(v map[string]*string) *CreateAlidingAssistantRequest {
	s.Ext = v
	return s
}

func (s *CreateAlidingAssistantRequest) SetIcon(v string) *CreateAlidingAssistantRequest {
	s.Icon = &v
	return s
}

func (s *CreateAlidingAssistantRequest) SetInstructions(v string) *CreateAlidingAssistantRequest {
	s.Instructions = &v
	return s
}

func (s *CreateAlidingAssistantRequest) SetName(v string) *CreateAlidingAssistantRequest {
	s.Name = &v
	return s
}

func (s *CreateAlidingAssistantRequest) SetRecommendPrompts(v []*string) *CreateAlidingAssistantRequest {
	s.RecommendPrompts = v
	return s
}

func (s *CreateAlidingAssistantRequest) SetSource(v int32) *CreateAlidingAssistantRequest {
	s.Source = &v
	return s
}

func (s *CreateAlidingAssistantRequest) SetSourceIdentityId(v string) *CreateAlidingAssistantRequest {
	s.SourceIdentityId = &v
	return s
}

func (s *CreateAlidingAssistantRequest) SetTenantContext(v *CreateAlidingAssistantRequestTenantContext) *CreateAlidingAssistantRequest {
	s.TenantContext = v
	return s
}

func (s *CreateAlidingAssistantRequest) SetWelcomeContent(v string) *CreateAlidingAssistantRequest {
	s.WelcomeContent = &v
	return s
}

func (s *CreateAlidingAssistantRequest) Validate() error {
	return dara.Validate(s)
}

type CreateAlidingAssistantRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateAlidingAssistantRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s CreateAlidingAssistantRequestTenantContext) GoString() string {
	return s.String()
}

func (s *CreateAlidingAssistantRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateAlidingAssistantRequestTenantContext) SetTenantId(v string) *CreateAlidingAssistantRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *CreateAlidingAssistantRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
