// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomizedVoiceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGender(v string) *CreateCustomizedVoiceJobRequest
	GetGender() *string
	SetScenario(v string) *CreateCustomizedVoiceJobRequest
	GetScenario() *string
	SetVoiceDesc(v string) *CreateCustomizedVoiceJobRequest
	GetVoiceDesc() *string
	SetVoiceId(v string) *CreateCustomizedVoiceJobRequest
	GetVoiceId() *string
	SetVoiceName(v string) *CreateCustomizedVoiceJobRequest
	GetVoiceName() *string
}

type CreateCustomizedVoiceJobRequest struct {
	// The gender. Valid values:
	//
	// - female: female.
	//
	// - male: male.
	//
	// This parameter is required.
	//
	// example:
	//
	// female
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The scenario. Valid values:
	//
	// - story: story.
	//
	// - interaction: interaction.
	//
	// - navigation: navigation.
	//
	// This parameter is required.
	//
	// example:
	//
	// story
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The voice description.
	//
	// - The description cannot exceed 256 characters.
	//
	// example:
	//
	// This is a personalized voice
	VoiceDesc *string `json:"VoiceDesc,omitempty" xml:"VoiceDesc,omitempty"`
	// The custom voice ID (English name or pinyin of the voice).
	//
	// - The ID cannot be the same as any of your other custom voice IDs.
	//
	// - The ID cannot exceed 32 characters.
	//
	// - Only letters and numbers are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// The voice name (generally a Chinese name).
	//
	// - The name cannot exceed 32 characters.
	//
	// example:
	//
	// 小专
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s CreateCustomizedVoiceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomizedVoiceJobRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedVoiceJobRequest) GetGender() *string {
	return s.Gender
}

func (s *CreateCustomizedVoiceJobRequest) GetScenario() *string {
	return s.Scenario
}

func (s *CreateCustomizedVoiceJobRequest) GetVoiceDesc() *string {
	return s.VoiceDesc
}

func (s *CreateCustomizedVoiceJobRequest) GetVoiceId() *string {
	return s.VoiceId
}

func (s *CreateCustomizedVoiceJobRequest) GetVoiceName() *string {
	return s.VoiceName
}

func (s *CreateCustomizedVoiceJobRequest) SetGender(v string) *CreateCustomizedVoiceJobRequest {
	s.Gender = &v
	return s
}

func (s *CreateCustomizedVoiceJobRequest) SetScenario(v string) *CreateCustomizedVoiceJobRequest {
	s.Scenario = &v
	return s
}

func (s *CreateCustomizedVoiceJobRequest) SetVoiceDesc(v string) *CreateCustomizedVoiceJobRequest {
	s.VoiceDesc = &v
	return s
}

func (s *CreateCustomizedVoiceJobRequest) SetVoiceId(v string) *CreateCustomizedVoiceJobRequest {
	s.VoiceId = &v
	return s
}

func (s *CreateCustomizedVoiceJobRequest) SetVoiceName(v string) *CreateCustomizedVoiceJobRequest {
	s.VoiceName = &v
	return s
}

func (s *CreateCustomizedVoiceJobRequest) Validate() error {
	return dara.Validate(s)
}
