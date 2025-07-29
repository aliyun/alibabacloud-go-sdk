// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMarketingInformationExtractShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrompt(v string) *RunMarketingInformationExtractShrinkRequest
	GetCustomPrompt() *string
	SetExtractType(v string) *RunMarketingInformationExtractShrinkRequest
	GetExtractType() *string
	SetModelId(v string) *RunMarketingInformationExtractShrinkRequest
	GetModelId() *string
	SetSourceMaterialsShrink(v string) *RunMarketingInformationExtractShrinkRequest
	GetSourceMaterialsShrink() *string
}

type RunMarketingInformationExtractShrinkRequest struct {
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	ExtractType  *string `json:"extractType,omitempty" xml:"extractType,omitempty"`
	// example:
	//
	// qwen-max
	//
	// qwen-plus
	ModelId               *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	SourceMaterialsShrink *string `json:"sourceMaterials,omitempty" xml:"sourceMaterials,omitempty"`
}

func (s RunMarketingInformationExtractShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationExtractShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractShrinkRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *RunMarketingInformationExtractShrinkRequest) GetExtractType() *string {
	return s.ExtractType
}

func (s *RunMarketingInformationExtractShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunMarketingInformationExtractShrinkRequest) GetSourceMaterialsShrink() *string {
	return s.SourceMaterialsShrink
}

func (s *RunMarketingInformationExtractShrinkRequest) SetCustomPrompt(v string) *RunMarketingInformationExtractShrinkRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RunMarketingInformationExtractShrinkRequest) SetExtractType(v string) *RunMarketingInformationExtractShrinkRequest {
	s.ExtractType = &v
	return s
}

func (s *RunMarketingInformationExtractShrinkRequest) SetModelId(v string) *RunMarketingInformationExtractShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunMarketingInformationExtractShrinkRequest) SetSourceMaterialsShrink(v string) *RunMarketingInformationExtractShrinkRequest {
	s.SourceMaterialsShrink = &v
	return s
}

func (s *RunMarketingInformationExtractShrinkRequest) Validate() error {
	return dara.Validate(s)
}
