// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateProjectRequest
	GetDescription() *string
	SetEngines(v *CreateProjectRequestEngines) *CreateProjectRequest
	GetEngines() *CreateProjectRequestEngines
	SetInstructionPrompt(v string) *CreateProjectRequest
	GetInstructionPrompt() *string
	SetName(v string) *CreateProjectRequest
	GetName() *string
	SetSource(v *CreateProjectRequestSource) *CreateProjectRequest
	GetSource() *CreateProjectRequestSource
}

type CreateProjectRequest struct {
	// The description.
	//
	// example:
	//
	// This is default function description by fc-deploy component
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The engine switches for the project or scan snapshot. Only SAST and SCA are supported.
	Engines *CreateProjectRequestEngines `json:"engines,omitempty" xml:"engines,omitempty" type:"Struct"`
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
	Source *CreateProjectRequestSource `json:"source,omitempty" xml:"source,omitempty" type:"Struct"`
}

func (s CreateProjectRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectRequest) GetEngines() *CreateProjectRequestEngines {
	return s.Engines
}

func (s *CreateProjectRequest) GetInstructionPrompt() *string {
	return s.InstructionPrompt
}

func (s *CreateProjectRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectRequest) GetSource() *CreateProjectRequestSource {
	return s.Source
}

func (s *CreateProjectRequest) SetDescription(v string) *CreateProjectRequest {
	s.Description = &v
	return s
}

func (s *CreateProjectRequest) SetEngines(v *CreateProjectRequestEngines) *CreateProjectRequest {
	s.Engines = v
	return s
}

func (s *CreateProjectRequest) SetInstructionPrompt(v string) *CreateProjectRequest {
	s.InstructionPrompt = &v
	return s
}

func (s *CreateProjectRequest) SetName(v string) *CreateProjectRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectRequest) SetSource(v *CreateProjectRequestSource) *CreateProjectRequest {
	s.Source = v
	return s
}

func (s *CreateProjectRequest) Validate() error {
	if s.Engines != nil {
		if err := s.Engines.Validate(); err != nil {
			return err
		}
	}
	if s.Source != nil {
		if err := s.Source.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectRequestEngines struct {
	// Specifies whether SAST is supported.
	//
	// example:
	//
	// true
	Sast *bool `json:"sast,omitempty" xml:"sast,omitempty"`
	// The engine-level configuration.
	SastConfig *CreateProjectRequestEnginesSastConfig `json:"sastConfig,omitempty" xml:"sastConfig,omitempty" type:"Struct"`
	// Specifies whether SCA is supported.
	//
	// example:
	//
	// false
	Sca *bool `json:"sca,omitempty" xml:"sca,omitempty"`
	// The engine-level configuration.
	ScaConfig *CreateProjectRequestEnginesScaConfig `json:"scaConfig,omitempty" xml:"scaConfig,omitempty" type:"Struct"`
}

func (s CreateProjectRequestEngines) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequestEngines) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestEngines) GetSast() *bool {
	return s.Sast
}

func (s *CreateProjectRequestEngines) GetSastConfig() *CreateProjectRequestEnginesSastConfig {
	return s.SastConfig
}

func (s *CreateProjectRequestEngines) GetSca() *bool {
	return s.Sca
}

func (s *CreateProjectRequestEngines) GetScaConfig() *CreateProjectRequestEnginesScaConfig {
	return s.ScaConfig
}

func (s *CreateProjectRequestEngines) SetSast(v bool) *CreateProjectRequestEngines {
	s.Sast = &v
	return s
}

func (s *CreateProjectRequestEngines) SetSastConfig(v *CreateProjectRequestEnginesSastConfig) *CreateProjectRequestEngines {
	s.SastConfig = v
	return s
}

func (s *CreateProjectRequestEngines) SetSca(v bool) *CreateProjectRequestEngines {
	s.Sca = &v
	return s
}

func (s *CreateProjectRequestEngines) SetScaConfig(v *CreateProjectRequestEnginesScaConfig) *CreateProjectRequestEngines {
	s.ScaConfig = v
	return s
}

func (s *CreateProjectRequestEngines) Validate() error {
	if s.SastConfig != nil {
		if err := s.SastConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ScaConfig != nil {
		if err := s.ScaConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectRequestEnginesSastConfig struct {
	// Specifies whether to generate remediation suggestions.
	Remediation *bool `json:"remediation,omitempty" xml:"remediation,omitempty"`
}

func (s CreateProjectRequestEnginesSastConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequestEnginesSastConfig) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestEnginesSastConfig) GetRemediation() *bool {
	return s.Remediation
}

func (s *CreateProjectRequestEnginesSastConfig) SetRemediation(v bool) *CreateProjectRequestEnginesSastConfig {
	s.Remediation = &v
	return s
}

func (s *CreateProjectRequestEnginesSastConfig) Validate() error {
	return dara.Validate(s)
}

type CreateProjectRequestEnginesScaConfig struct {
	// Specifies whether to generate remediation suggestions.
	Remediation *bool `json:"remediation,omitempty" xml:"remediation,omitempty"`
}

func (s CreateProjectRequestEnginesScaConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequestEnginesScaConfig) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestEnginesScaConfig) GetRemediation() *bool {
	return s.Remediation
}

func (s *CreateProjectRequestEnginesScaConfig) SetRemediation(v bool) *CreateProjectRequestEnginesScaConfig {
	s.Remediation = &v
	return s
}

func (s *CreateProjectRequestEnginesScaConfig) Validate() error {
	return dara.Validate(s)
}

type CreateProjectRequestSource struct {
	// The project type.
	//
	// example:
	//
	// api
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateProjectRequestSource) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectRequestSource) GoString() string {
	return s.String()
}

func (s *CreateProjectRequestSource) GetType() *string {
	return s.Type
}

func (s *CreateProjectRequestSource) SetType(v string) *CreateProjectRequestSource {
	s.Type = &v
	return s
}

func (s *CreateProjectRequestSource) Validate() error {
	return dara.Validate(s)
}
