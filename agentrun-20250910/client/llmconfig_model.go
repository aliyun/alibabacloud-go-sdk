// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLLMConfig interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *LLMConfigConfig) *LLMConfig
	GetConfig() *LLMConfigConfig
	SetModelServiceName(v string) *LLMConfig
	GetModelServiceName() *string
}

type LLMConfig struct {
	Config           *LLMConfigConfig `json:"config,omitempty" xml:"config,omitempty" type:"Struct"`
	ModelServiceName *string          `json:"modelServiceName,omitempty" xml:"modelServiceName,omitempty"`
}

func (s LLMConfig) String() string {
	return dara.Prettify(s)
}

func (s LLMConfig) GoString() string {
	return s.String()
}

func (s *LLMConfig) GetConfig() *LLMConfigConfig {
	return s.Config
}

func (s *LLMConfig) GetModelServiceName() *string {
	return s.ModelServiceName
}

func (s *LLMConfig) SetConfig(v *LLMConfigConfig) *LLMConfig {
	s.Config = v
	return s
}

func (s *LLMConfig) SetModelServiceName(v string) *LLMConfig {
	s.ModelServiceName = &v
	return s
}

func (s *LLMConfig) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type LLMConfigConfig struct {
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
}

func (s LLMConfigConfig) String() string {
	return dara.Prettify(s)
}

func (s LLMConfigConfig) GoString() string {
	return s.String()
}

func (s *LLMConfigConfig) GetModel() *string {
	return s.Model
}

func (s *LLMConfigConfig) SetModel(v string) *LLMConfigConfig {
	s.Model = &v
	return s
}

func (s *LLMConfigConfig) Validate() error {
	return dara.Validate(s)
}
