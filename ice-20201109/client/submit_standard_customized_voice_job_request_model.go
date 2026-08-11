// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitStandardCustomizedVoiceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudios(v string) *SubmitStandardCustomizedVoiceJobRequest
	GetAudios() *string
	SetAuthentication(v string) *SubmitStandardCustomizedVoiceJobRequest
	GetAuthentication() *string
	SetDemoAudioMediaURL(v string) *SubmitStandardCustomizedVoiceJobRequest
	GetDemoAudioMediaURL() *string
	SetGender(v string) *SubmitStandardCustomizedVoiceJobRequest
	GetGender() *string
	SetVoiceName(v string) *SubmitStandardCustomizedVoiceJobRequest
	GetVoiceName() *string
}

type SubmitStandardCustomizedVoiceJobRequest struct {
	// The list of media asset IDs for training audio materials. Separate multiple media asset IDs with commas (,).
	//
	// 	Notice: The total duration of all materials must be between 15 and 30 minutes, and the duration of each individual material must be greater than 1 minute.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****,****571c704445f9a0ee011406c2****,****571c704445f9a0ee011406c2****
	Audios *string `json:"Audios,omitempty" xml:"Audios,omitempty"`
	// The media asset ID of the authentication audio. Upload an audio clip to verify your identity. The task fails if the voiceprint does not match the training audio.
	//
	// 	Notice: Read and record the following statement clearly: I confirm that I am initiating voice cloning customization. The training audio is provided by me. I commit to being responsible for the customized content and guarantee that no illegal or non-compliant content will be created.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****
	Authentication *string `json:"Authentication,omitempty" xml:"Authentication,omitempty"`
	// The audio output address for the sample.
	//
	// - If you specify this parameter, a sample audio file is generated at the specified OSS address after training succeeds.
	//
	// - If you do not specify this parameter, no sample audio is generated.
	//
	// 	Notice: The address must be a valid public OSS URL under your account.
	//
	// example:
	//
	// https://your-bucket.oss-cn-shanghai.aliyuncs.com/demo.mp3
	DemoAudioMediaURL *string `json:"DemoAudioMediaURL,omitempty" xml:"DemoAudioMediaURL,omitempty"`
	// The gender. Valid values:
	//
	// - female
	//
	// - male
	//
	// example:
	//
	// female
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The voice name. The name can be up to 32 characters in length.
	//
	// example:
	//
	// Basic
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s SubmitStandardCustomizedVoiceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitStandardCustomizedVoiceJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitStandardCustomizedVoiceJobRequest) GetAudios() *string {
	return s.Audios
}

func (s *SubmitStandardCustomizedVoiceJobRequest) GetAuthentication() *string {
	return s.Authentication
}

func (s *SubmitStandardCustomizedVoiceJobRequest) GetDemoAudioMediaURL() *string {
	return s.DemoAudioMediaURL
}

func (s *SubmitStandardCustomizedVoiceJobRequest) GetGender() *string {
	return s.Gender
}

func (s *SubmitStandardCustomizedVoiceJobRequest) GetVoiceName() *string {
	return s.VoiceName
}

func (s *SubmitStandardCustomizedVoiceJobRequest) SetAudios(v string) *SubmitStandardCustomizedVoiceJobRequest {
	s.Audios = &v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobRequest) SetAuthentication(v string) *SubmitStandardCustomizedVoiceJobRequest {
	s.Authentication = &v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobRequest) SetDemoAudioMediaURL(v string) *SubmitStandardCustomizedVoiceJobRequest {
	s.DemoAudioMediaURL = &v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobRequest) SetGender(v string) *SubmitStandardCustomizedVoiceJobRequest {
	s.Gender = &v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobRequest) SetVoiceName(v string) *SubmitStandardCustomizedVoiceJobRequest {
	s.VoiceName = &v
	return s
}

func (s *SubmitStandardCustomizedVoiceJobRequest) Validate() error {
	return dara.Validate(s)
}
