// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndPublishAgentSelectiveShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationConfigShrink(v string) *UpdateAndPublishAgentSelectiveShrinkRequest
	GetApplicationConfigShrink() *string
	SetInstructions(v string) *UpdateAndPublishAgentSelectiveShrinkRequest
	GetInstructions() *string
	SetModelId(v string) *UpdateAndPublishAgentSelectiveShrinkRequest
	GetModelId() *string
	SetName(v string) *UpdateAndPublishAgentSelectiveShrinkRequest
	GetName() *string
	SetSampleLibraryShrink(v string) *UpdateAndPublishAgentSelectiveShrinkRequest
	GetSampleLibraryShrink() *string
}

type UpdateAndPublishAgentSelectiveShrinkRequest struct {
	ApplicationConfigShrink *string `json:"applicationConfig,omitempty" xml:"applicationConfig,omitempty"`
	Instructions            *string `json:"instructions,omitempty" xml:"instructions,omitempty"`
	ModelId                 *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	Name                    *string `json:"name,omitempty" xml:"name,omitempty"`
	SampleLibraryShrink     *string `json:"sampleLibrary,omitempty" xml:"sampleLibrary,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) GetApplicationConfigShrink() *string {
	return s.ApplicationConfigShrink
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) GetSampleLibraryShrink() *string {
	return s.SampleLibraryShrink
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) SetApplicationConfigShrink(v string) *UpdateAndPublishAgentSelectiveShrinkRequest {
	s.ApplicationConfigShrink = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) SetInstructions(v string) *UpdateAndPublishAgentSelectiveShrinkRequest {
	s.Instructions = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) SetModelId(v string) *UpdateAndPublishAgentSelectiveShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) SetName(v string) *UpdateAndPublishAgentSelectiveShrinkRequest {
	s.Name = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) SetSampleLibraryShrink(v string) *UpdateAndPublishAgentSelectiveShrinkRequest {
	s.SampleLibraryShrink = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveShrinkRequest) Validate() error {
	return dara.Validate(s)
}
