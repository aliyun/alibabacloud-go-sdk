// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelSetSpec interface {
	dara.Model
	String() string
	GoString() string
	SetApiInvokeType(v string) *ModelSetSpec
	GetApiInvokeType() *string
	SetAuthConfig(v *Authorization) *ModelSetSpec
	GetAuthConfig() *Authorization
	SetBaseUrl(v string) *ModelSetSpec
	GetBaseUrl() *string
	SetFeatures(v *ModelSetSpecFeatures) *ModelSetSpec
	GetFeatures() *ModelSetSpecFeatures
	SetModelType(v string) *ModelSetSpec
	GetModelType() *string
	SetModels(v []*ModelSetModelProfile) *ModelSetSpec
	GetModels() []*ModelSetModelProfile
	SetProvider(v string) *ModelSetSpec
	GetProvider() *string
}

type ModelSetSpec struct {
	// example:
	//
	// openai
	ApiInvokeType *string               `json:"apiInvokeType,omitempty" xml:"apiInvokeType,omitempty"`
	AuthConfig    *Authorization        `json:"authConfig,omitempty" xml:"authConfig,omitempty"`
	BaseUrl       *string               `json:"baseUrl,omitempty" xml:"baseUrl,omitempty"`
	Features      *ModelSetSpecFeatures `json:"features,omitempty" xml:"features,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// llm
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// This parameter is required.
	Models []*ModelSetModelProfile `json:"models,omitempty" xml:"models,omitempty" type:"Repeated"`
	// example:
	//
	// FunctionAI
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ModelSetSpec) String() string {
	return dara.Prettify(s)
}

func (s ModelSetSpec) GoString() string {
	return s.String()
}

func (s *ModelSetSpec) GetApiInvokeType() *string {
	return s.ApiInvokeType
}

func (s *ModelSetSpec) GetAuthConfig() *Authorization {
	return s.AuthConfig
}

func (s *ModelSetSpec) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModelSetSpec) GetFeatures() *ModelSetSpecFeatures {
	return s.Features
}

func (s *ModelSetSpec) GetModelType() *string {
	return s.ModelType
}

func (s *ModelSetSpec) GetModels() []*ModelSetModelProfile {
	return s.Models
}

func (s *ModelSetSpec) GetProvider() *string {
	return s.Provider
}

func (s *ModelSetSpec) SetApiInvokeType(v string) *ModelSetSpec {
	s.ApiInvokeType = &v
	return s
}

func (s *ModelSetSpec) SetAuthConfig(v *Authorization) *ModelSetSpec {
	s.AuthConfig = v
	return s
}

func (s *ModelSetSpec) SetBaseUrl(v string) *ModelSetSpec {
	s.BaseUrl = &v
	return s
}

func (s *ModelSetSpec) SetFeatures(v *ModelSetSpecFeatures) *ModelSetSpec {
	s.Features = v
	return s
}

func (s *ModelSetSpec) SetModelType(v string) *ModelSetSpec {
	s.ModelType = &v
	return s
}

func (s *ModelSetSpec) SetModels(v []*ModelSetModelProfile) *ModelSetSpec {
	s.Models = v
	return s
}

func (s *ModelSetSpec) SetProvider(v string) *ModelSetSpec {
	s.Provider = &v
	return s
}

func (s *ModelSetSpec) Validate() error {
	return dara.Validate(s)
}

type ModelSetSpecFeatures struct {
	AgentThought *bool `json:"agentThought,omitempty" xml:"agentThought,omitempty"`
	ToolCall     *bool `json:"toolCall,omitempty" xml:"toolCall,omitempty"`
	Vision       *bool `json:"vision,omitempty" xml:"vision,omitempty"`
}

func (s ModelSetSpecFeatures) String() string {
	return dara.Prettify(s)
}

func (s ModelSetSpecFeatures) GoString() string {
	return s.String()
}

func (s *ModelSetSpecFeatures) GetAgentThought() *bool {
	return s.AgentThought
}

func (s *ModelSetSpecFeatures) GetToolCall() *bool {
	return s.ToolCall
}

func (s *ModelSetSpecFeatures) GetVision() *bool {
	return s.Vision
}

func (s *ModelSetSpecFeatures) SetAgentThought(v bool) *ModelSetSpecFeatures {
	s.AgentThought = &v
	return s
}

func (s *ModelSetSpecFeatures) SetToolCall(v bool) *ModelSetSpecFeatures {
	s.ToolCall = &v
	return s
}

func (s *ModelSetSpecFeatures) SetVision(v bool) *ModelSetSpecFeatures {
	s.Vision = &v
	return s
}

func (s *ModelSetSpecFeatures) Validate() error {
	return dara.Validate(s)
}
