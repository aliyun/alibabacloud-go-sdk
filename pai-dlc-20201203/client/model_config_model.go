// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelConfig interface {
	dara.Model
	String() string
	GoString() string
	SetModelName(v string) *ModelConfig
	GetModelName() *string
	SetModelTemplate(v *ModelTemplate) *ModelConfig
	GetModelTemplate() *ModelTemplate
}

type ModelConfig struct {
	ModelName     *string        `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ModelTemplate *ModelTemplate `json:"ModelTemplate,omitempty" xml:"ModelTemplate,omitempty"`
}

func (s ModelConfig) String() string {
	return dara.Prettify(s)
}

func (s ModelConfig) GoString() string {
	return s.String()
}

func (s *ModelConfig) GetModelName() *string {
	return s.ModelName
}

func (s *ModelConfig) GetModelTemplate() *ModelTemplate {
	return s.ModelTemplate
}

func (s *ModelConfig) SetModelName(v string) *ModelConfig {
	s.ModelName = &v
	return s
}

func (s *ModelConfig) SetModelTemplate(v *ModelTemplate) *ModelConfig {
	s.ModelTemplate = v
	return s
}

func (s *ModelConfig) Validate() error {
	if s.ModelTemplate != nil {
		if err := s.ModelTemplate.Validate(); err != nil {
			return err
		}
	}
	return nil
}
