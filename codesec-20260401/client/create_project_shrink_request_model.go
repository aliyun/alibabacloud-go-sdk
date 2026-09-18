// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProjectShrinkRequest
	GetDescription() *string
	SetEnginesShrink(v string) *CreateProjectShrinkRequest
	GetEnginesShrink() *string
	SetInstructionPrompt(v string) *CreateProjectShrinkRequest
	GetInstructionPrompt() *string
	SetName(v string) *CreateProjectShrinkRequest
	GetName() *string
	SetSourceShrink(v string) *CreateProjectShrinkRequest
	GetSourceShrink() *string
}

type CreateProjectShrinkRequest struct {
	// The description.
	//
	// example:
	//
	// This is default function description by fc-deploy component
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The engine switches for the project or scan snapshot. Only SAST and SCA are supported.
	EnginesShrink *string `json:"engines,omitempty" xml:"engines,omitempty"`
	// The natural language prompt that describes scanning or result processing preferences, such as ignoring low-risk vulnerabilities.
	//
	// example:
	//
	// such as ignoring low-severity vulnerabilities, etc.
	InstructionPrompt *string `json:"instructionPrompt,omitempty" xml:"instructionPrompt,omitempty"`
	// The project name.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_paswd_103
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The project source.
	SourceShrink *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s CreateProjectShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectShrinkRequest) GetEnginesShrink() *string {
	return s.EnginesShrink
}

func (s *CreateProjectShrinkRequest) GetInstructionPrompt() *string {
	return s.InstructionPrompt
}

func (s *CreateProjectShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectShrinkRequest) GetSourceShrink() *string {
	return s.SourceShrink
}

func (s *CreateProjectShrinkRequest) SetDescription(v string) *CreateProjectShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetEnginesShrink(v string) *CreateProjectShrinkRequest {
	s.EnginesShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetInstructionPrompt(v string) *CreateProjectShrinkRequest {
	s.InstructionPrompt = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetName(v string) *CreateProjectShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectShrinkRequest) SetSourceShrink(v string) *CreateProjectShrinkRequest {
	s.SourceShrink = &v
	return s
}

func (s *CreateProjectShrinkRequest) Validate() error {
	return dara.Validate(s)
}
