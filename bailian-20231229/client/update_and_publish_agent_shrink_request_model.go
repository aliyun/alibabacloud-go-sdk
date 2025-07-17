// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndPublishAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigShrink(v string) *UpdateAndPublishAgentShrinkRequest
	GetApplicationConfigShrink() *string
	SetInstructions(v string) *UpdateAndPublishAgentShrinkRequest
	GetInstructions() *string
	SetModelId(v string) *UpdateAndPublishAgentShrinkRequest
	GetModelId() *string
	SetName(v string) *UpdateAndPublishAgentShrinkRequest
	GetName() *string
	SetSampleLibraryShrink(v string) *UpdateAndPublishAgentShrinkRequest
	GetSampleLibraryShrink() *string
}

type UpdateAndPublishAgentShrinkRequest struct {
	ApplicationConfigShrink *string `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty"`
	Instructions            *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId                 *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name                    *string `json:"name,omitempty" xml:"name,omitempty"`
	SampleLibraryShrink     *string `json:"sampleLibrary,omitempty" xml:"sampleLibrary,omitempty"`
}

func (s UpdateAndPublishAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentShrinkRequest) GetApplicationConfigShrink() *string {
	return s.ApplicationConfigShrink
}

func (s *UpdateAndPublishAgentShrinkRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *UpdateAndPublishAgentShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *UpdateAndPublishAgentShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAndPublishAgentShrinkRequest) GetSampleLibraryShrink() *string {
	return s.SampleLibraryShrink
}

func (s *UpdateAndPublishAgentShrinkRequest) SetApplicationConfigShrink(v string) *UpdateAndPublishAgentShrinkRequest {
	s.ApplicationConfigShrink = &v
	return s
}

func (s *UpdateAndPublishAgentShrinkRequest) SetInstructions(v string) *UpdateAndPublishAgentShrinkRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateAndPublishAgentShrinkRequest) SetModelId(v string) *UpdateAndPublishAgentShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *UpdateAndPublishAgentShrinkRequest) SetName(v string) *UpdateAndPublishAgentShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateAndPublishAgentShrinkRequest) SetSampleLibraryShrink(v string) *UpdateAndPublishAgentShrinkRequest {
	s.SampleLibraryShrink = &v
	return s
}

func (s *UpdateAndPublishAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
