// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCosyVoiceCustomizedVoiceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudios(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest
	GetAudios() *string
	SetDemoAudioMediaURL(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest
	GetDemoAudioMediaURL() *string
	SetGender(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest
	GetGender() *string
	SetModel(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest
	GetModel() *string
	SetVoiceName(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest
	GetVoiceName() *string
}

type SubmitCosyVoiceCustomizedVoiceJobRequest struct {
	// The media asset ID of the training audio material. Currently, only one audio material can be used for training.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****
	Audios *string `json:"Audios,omitempty" xml:"Audios,omitempty"`
	// The sample audio output address.
	//
	// - If you specify this parameter, a sample audio file is generated at the specified OSS address after training succeeds.
	//
	// 	Notice: The address must be a valid public OSS address under your account.
	//
	// example:
	//
	// https://your-bucket.oss-cn-shanghai.aliyuncs.com/demo.MP3
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
	// The voice cloning model. Valid values:
	//
	// - **cosyvoice-v3-plus**
	//
	// - **cosyvoice-v3-flash**
	//
	// - **cosyvoice-v3.5-plus**
	//
	// - **cosyvoice-v3.5-flash**
	//
	// example:
	//
	// cosyvoice-v3-plus
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// The voice name. The name can be up to 32 characters in length.
	//
	// example:
	//
	// 小专
	VoiceName *string `json:"VoiceName,omitempty" xml:"VoiceName,omitempty"`
}

func (s SubmitCosyVoiceCustomizedVoiceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCosyVoiceCustomizedVoiceJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) GetAudios() *string {
	return s.Audios
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) GetDemoAudioMediaURL() *string {
	return s.DemoAudioMediaURL
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) GetGender() *string {
	return s.Gender
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) GetModel() *string {
	return s.Model
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) GetVoiceName() *string {
	return s.VoiceName
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) SetAudios(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest {
	s.Audios = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) SetDemoAudioMediaURL(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest {
	s.DemoAudioMediaURL = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) SetGender(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest {
	s.Gender = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) SetModel(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest {
	s.Model = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) SetVoiceName(v string) *SubmitCosyVoiceCustomizedVoiceJobRequest {
	s.VoiceName = &v
	return s
}

func (s *SubmitCosyVoiceCustomizedVoiceJobRequest) Validate() error {
	return dara.Validate(s)
}
