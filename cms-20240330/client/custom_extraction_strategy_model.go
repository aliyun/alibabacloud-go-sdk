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
	// Description
	//
	// example:
	//
	// test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Extraction prompt
	//
	// example:
	//
	// test
	ExtractionPrompt *string `json:"extractionPrompt,omitempty" xml:"extractionPrompt,omitempty"`
	// Policy Name
	//
	// example:
	//
	// test1
	StrategyName *string `json:"strategyName,omitempty" xml:"strategyName,omitempty"`
	// Policy type. Supports "Episodic", "Summary", and "Fact".
	//
	// example:
	//
	// Fact
	StrategyType *string `json:"strategyType,omitempty" xml:"strategyType,omitempty"`
	// Update prompt
	//
	// example:
	//
	// test
	UpdatePrompt *string `json:"updatePrompt,omitempty" xml:"updatePrompt,omitempty"`
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
