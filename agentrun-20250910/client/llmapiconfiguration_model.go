// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLLMAPIConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetAiProtocols(v []*string) *LLMAPIConfiguration
	GetAiProtocols() []*string
	SetAttachPolicyConfigs(v []*AttachPolicyConfig) *LLMAPIConfiguration
	GetAttachPolicyConfigs() []*AttachPolicyConfig
	SetBasePath(v string) *LLMAPIConfiguration
	GetBasePath() *string
	SetDeployConfigs(v []*LLMDeployConfig) *LLMAPIConfiguration
	GetDeployConfigs() []*LLMDeployConfig
	SetModelCategory(v string) *LLMAPIConfiguration
	GetModelCategory() *string
	SetRemoveBasePathOnForward(v bool) *LLMAPIConfiguration
	GetRemoveBasePathOnForward() *bool
	SetType(v string) *LLMAPIConfiguration
	GetType() *string
}

type LLMAPIConfiguration struct {
	AiProtocols             []*string             `json:"aiProtocols,omitempty" xml:"aiProtocols,omitempty" type:"Repeated"`
	AttachPolicyConfigs     []*AttachPolicyConfig `json:"attachPolicyConfigs,omitempty" xml:"attachPolicyConfigs,omitempty" type:"Repeated"`
	BasePath                *string               `json:"basePath,omitempty" xml:"basePath,omitempty"`
	DeployConfigs           []*LLMDeployConfig    `json:"deployConfigs,omitempty" xml:"deployConfigs,omitempty" type:"Repeated"`
	ModelCategory           *string               `json:"modelCategory,omitempty" xml:"modelCategory,omitempty"`
	RemoveBasePathOnForward *bool                 `json:"removeBasePathOnForward,omitempty" xml:"removeBasePathOnForward,omitempty"`
	Type                    *string               `json:"type,omitempty" xml:"type,omitempty"`
}

func (s LLMAPIConfiguration) String() string {
	return dara.Prettify(s)
}

func (s LLMAPIConfiguration) GoString() string {
	return s.String()
}

func (s *LLMAPIConfiguration) GetAiProtocols() []*string {
	return s.AiProtocols
}

func (s *LLMAPIConfiguration) GetAttachPolicyConfigs() []*AttachPolicyConfig {
	return s.AttachPolicyConfigs
}

func (s *LLMAPIConfiguration) GetBasePath() *string {
	return s.BasePath
}

func (s *LLMAPIConfiguration) GetDeployConfigs() []*LLMDeployConfig {
	return s.DeployConfigs
}

func (s *LLMAPIConfiguration) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *LLMAPIConfiguration) GetRemoveBasePathOnForward() *bool {
	return s.RemoveBasePathOnForward
}

func (s *LLMAPIConfiguration) GetType() *string {
	return s.Type
}

func (s *LLMAPIConfiguration) SetAiProtocols(v []*string) *LLMAPIConfiguration {
	s.AiProtocols = v
	return s
}

func (s *LLMAPIConfiguration) SetAttachPolicyConfigs(v []*AttachPolicyConfig) *LLMAPIConfiguration {
	s.AttachPolicyConfigs = v
	return s
}

func (s *LLMAPIConfiguration) SetBasePath(v string) *LLMAPIConfiguration {
	s.BasePath = &v
	return s
}

func (s *LLMAPIConfiguration) SetDeployConfigs(v []*LLMDeployConfig) *LLMAPIConfiguration {
	s.DeployConfigs = v
	return s
}

func (s *LLMAPIConfiguration) SetModelCategory(v string) *LLMAPIConfiguration {
	s.ModelCategory = &v
	return s
}

func (s *LLMAPIConfiguration) SetRemoveBasePathOnForward(v bool) *LLMAPIConfiguration {
	s.RemoveBasePathOnForward = &v
	return s
}

func (s *LLMAPIConfiguration) SetType(v string) *LLMAPIConfiguration {
	s.Type = &v
	return s
}

func (s *LLMAPIConfiguration) Validate() error {
	return dara.Validate(s)
}
