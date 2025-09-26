// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTargetServiceConfig interface {
	dara.Model
	String() string
	GoString() string
	SetModelId(v string) *TargetServiceConfig
	GetModelId() *string
	SetModelName(v string) *TargetServiceConfig
	GetModelName() *string
	SetModelNamePattern(v string) *TargetServiceConfig
	GetModelNamePattern() *string
	SetWeight(v int64) *TargetServiceConfig
	GetWeight() *int64
}

type TargetServiceConfig struct {
	ModelId          *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	ModelName        *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	ModelNamePattern *string `json:"modelNamePattern,omitempty" xml:"modelNamePattern,omitempty"`
	Weight           *int64  `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s TargetServiceConfig) String() string {
	return dara.Prettify(s)
}

func (s TargetServiceConfig) GoString() string {
	return s.String()
}

func (s *TargetServiceConfig) GetModelId() *string {
	return s.ModelId
}

func (s *TargetServiceConfig) GetModelName() *string {
	return s.ModelName
}

func (s *TargetServiceConfig) GetModelNamePattern() *string {
	return s.ModelNamePattern
}

func (s *TargetServiceConfig) GetWeight() *int64 {
	return s.Weight
}

func (s *TargetServiceConfig) SetModelId(v string) *TargetServiceConfig {
	s.ModelId = &v
	return s
}

func (s *TargetServiceConfig) SetModelName(v string) *TargetServiceConfig {
	s.ModelName = &v
	return s
}

func (s *TargetServiceConfig) SetModelNamePattern(v string) *TargetServiceConfig {
	s.ModelNamePattern = &v
	return s
}

func (s *TargetServiceConfig) SetWeight(v int64) *TargetServiceConfig {
	s.Weight = &v
	return s
}

func (s *TargetServiceConfig) Validate() error {
	return dara.Validate(s)
}
