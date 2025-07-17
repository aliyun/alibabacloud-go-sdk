// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndPulishAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigShrink(v string) *CreateAndPulishAgentShrinkRequest
	GetApplicationConfigShrink() *string
	SetInstructions(v string) *CreateAndPulishAgentShrinkRequest
	GetInstructions() *string
	SetModelId(v string) *CreateAndPulishAgentShrinkRequest
	GetModelId() *string
	SetName(v string) *CreateAndPulishAgentShrinkRequest
	GetName() *string
	SetSampleLibraryShrink(v string) *CreateAndPulishAgentShrinkRequest
	GetSampleLibraryShrink() *string
}

type CreateAndPulishAgentShrinkRequest struct {
	ApplicationConfigShrink *string `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty"`
	Instructions            *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId                 *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name                    *string `json:"name,omitempty" xml:"name,omitempty"`
	SampleLibraryShrink     *string `json:"sampleLibrary,omitempty" xml:"sampleLibrary,omitempty"`
}

func (s CreateAndPulishAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentShrinkRequest) GetApplicationConfigShrink() *string {
	return s.ApplicationConfigShrink
}

func (s *CreateAndPulishAgentShrinkRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *CreateAndPulishAgentShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *CreateAndPulishAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateAndPulishAgentShrinkRequest) GetSampleLibraryShrink() *string {
	return s.SampleLibraryShrink
}

func (s *CreateAndPulishAgentShrinkRequest) SetApplicationConfigShrink(v string) *CreateAndPulishAgentShrinkRequest {
	s.ApplicationConfigShrink = &v
	return s
}

func (s *CreateAndPulishAgentShrinkRequest) SetInstructions(v string) *CreateAndPulishAgentShrinkRequest {
	s.Instructions = &v
	return s
}

func (s *CreateAndPulishAgentShrinkRequest) SetModelId(v string) *CreateAndPulishAgentShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *CreateAndPulishAgentShrinkRequest) SetName(v string) *CreateAndPulishAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateAndPulishAgentShrinkRequest) SetSampleLibraryShrink(v string) *CreateAndPulishAgentShrinkRequest {
	s.SampleLibraryShrink = &v
	return s
}

func (s *CreateAndPulishAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
