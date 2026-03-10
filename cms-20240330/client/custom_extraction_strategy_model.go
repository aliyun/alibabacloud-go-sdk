// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomExtractionStrategy interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CustomExtractionStrategy
	GetDescription() *string
	SetExtractionPrompt(v string) *CustomExtractionStrategy
	GetExtractionPrompt() *string
	SetStrategyName(v string) *CustomExtractionStrategy
	GetStrategyName() *string
	SetStrategyType(v string) *CustomExtractionStrategy
	GetStrategyType() *string
	SetUpdatePrompt(v string) *CustomExtractionStrategy
	GetUpdatePrompt() *string
}

type CustomExtractionStrategy struct {
	Description      *string `json:"description,omitempty" xml:"description,omitempty"`
	ExtractionPrompt *string `json:"extractionPrompt,omitempty" xml:"extractionPrompt,omitempty"`
	StrategyName     *string `json:"strategyName,omitempty" xml:"strategyName,omitempty"`
	StrategyType     *string `json:"strategyType,omitempty" xml:"strategyType,omitempty"`
	UpdatePrompt     *string `json:"updatePrompt,omitempty" xml:"updatePrompt,omitempty"`
}

func (s CustomExtractionStrategy) String() string {
	return dara.Prettify(s)
}

func (s CustomExtractionStrategy) GoString() string {
	return s.String()
}

func (s *CustomExtractionStrategy) GetDescription() *string {
	return s.Description
}

func (s *CustomExtractionStrategy) GetExtractionPrompt() *string {
	return s.ExtractionPrompt
}

func (s *CustomExtractionStrategy) GetStrategyName() *string {
	return s.StrategyName
}

func (s *CustomExtractionStrategy) GetStrategyType() *string {
	return s.StrategyType
}

func (s *CustomExtractionStrategy) GetUpdatePrompt() *string {
	return s.UpdatePrompt
}

func (s *CustomExtractionStrategy) SetDescription(v string) *CustomExtractionStrategy {
	s.Description = &v
	return s
}

func (s *CustomExtractionStrategy) SetExtractionPrompt(v string) *CustomExtractionStrategy {
	s.ExtractionPrompt = &v
	return s
}

func (s *CustomExtractionStrategy) SetStrategyName(v string) *CustomExtractionStrategy {
	s.StrategyName = &v
	return s
}

func (s *CustomExtractionStrategy) SetStrategyType(v string) *CustomExtractionStrategy {
	s.StrategyType = &v
	return s
}

func (s *CustomExtractionStrategy) SetUpdatePrompt(v string) *CustomExtractionStrategy {
	s.UpdatePrompt = &v
	return s
}

func (s *CustomExtractionStrategy) Validate() error {
	return dara.Validate(s)
}
