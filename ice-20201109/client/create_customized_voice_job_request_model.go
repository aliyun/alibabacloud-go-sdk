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
	// 	- female
	//
	// 	- male
	//
	// This parameter is required.
	//
	// example:
	//
	// female
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The scenario. Valid values:
	//
	// 	- story
	//
	// 	- interaction
	//
	// 	- navigation
	//
	// This parameter is required.
	//
	// example:
	//
	// story
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	// The voice description.
	//
	// 	- The description can be up to 256 characters in length.
	VoiceDesc *string `json:"VoiceDesc,omitempty" xml:"VoiceDesc,omitempty"`
	// The voice ID. It can be the English name or Chinese Pinyin of the voice.
	//
	// 	- The value must be a unique ID that is not used by other custom voices.
	//
	// 	- The ID can be up to 32 characters in length.
	//
	// 	- Only letters and digits are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
	// The voice name.
	//
	// 	- The name can be up to 32 characters in length.
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
