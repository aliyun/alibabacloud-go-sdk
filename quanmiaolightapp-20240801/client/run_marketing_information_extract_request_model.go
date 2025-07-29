// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMarketingInformationExtractRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomPrompt(v string) *RunMarketingInformationExtractRequest
	GetCustomPrompt() *string
	SetExtractType(v string) *RunMarketingInformationExtractRequest
	GetExtractType() *string
	SetModelId(v string) *RunMarketingInformationExtractRequest
	GetModelId() *string
	SetSourceMaterials(v []*string) *RunMarketingInformationExtractRequest
	GetSourceMaterials() []*string
}

type RunMarketingInformationExtractRequest struct {
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
	ExtractType  *string `json:"extractType,omitempty" xml:"extractType,omitempty"`
	// example:
	//
	// qwen-max
	//
	// qwen-plus
	ModelId         *string   `json:"modelId,omitempty" xml:"modelId,omitempty"`
	SourceMaterials []*string `json:"sourceMaterials,omitempty" xml:"sourceMaterials,omitempty" type:"Repeated"`
}

func (s RunMarketingInformationExtractRequest) String() string {
	return dara.Prettify(s)
}

func (s RunMarketingInformationExtractRequest) GoString() string {
	return s.String()
}

func (s *RunMarketingInformationExtractRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *RunMarketingInformationExtractRequest) GetExtractType() *string {
	return s.ExtractType
}

func (s *RunMarketingInformationExtractRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunMarketingInformationExtractRequest) GetSourceMaterials() []*string {
	return s.SourceMaterials
}

func (s *RunMarketingInformationExtractRequest) SetCustomPrompt(v string) *RunMarketingInformationExtractRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RunMarketingInformationExtractRequest) SetExtractType(v string) *RunMarketingInformationExtractRequest {
	s.ExtractType = &v
	return s
}

func (s *RunMarketingInformationExtractRequest) SetModelId(v string) *RunMarketingInformationExtractRequest {
	s.ModelId = &v
	return s
}

func (s *RunMarketingInformationExtractRequest) SetSourceMaterials(v []*string) *RunMarketingInformationExtractRequest {
	s.SourceMaterials = v
	return s
}

func (s *RunMarketingInformationExtractRequest) Validate() error {
	return dara.Validate(s)
}
