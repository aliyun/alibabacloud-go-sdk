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
	// 	- The material assets IDs of the materials for training.
	//
	// 	- Separate multiple media IDs with commas (,).
	//
	// > : The total duration of all materials must be within 15 to 30 minutes. The duration of each material must be greater than 1 minute.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****,****571c704445f9a0ee011406c2****,****571c704445f9a0ee011406c2****
	Audios *string `json:"Audios,omitempty" xml:"Audios,omitempty"`
	// 	- The media asset ID of the authentication audio.
	//
	// 	- Upload an audio file for identity authentication. If the voiceprint extracted from the uploaded file differs from that of the training file, the job fails.
	//
	//     **
	//
	//     **Note**: Clearly read and record the following text: I confirm to customize human voice cloning and provide audio files that contain my voice for training. I promise that I am responsible for the customized content and that the content complies with laws and regulations.
	//
	// example:
	//
	// ****571c704445f9a0ee011406c2****
	Authentication *string `json:"Authentication,omitempty" xml:"Authentication,omitempty"`
	// The URL of the sample audio file.
	//
	// 	- If this parameter is specified, a sample audio file is generated at the specified Object Storage Service (OSS) URL after the training is complete.
	//
	// 	- If this parameter is not specified, no sample audio file is generated.
	//
	//     **
	//
	//     **Note**: The URL must be a valid public OSS URL within your Alibaba Cloud account.
	//
	// example:
	//
	// https://your-bucket.oss-cn-shanghai.aliyuncs.com/demo.mp3
	DemoAudioMediaURL *string `json:"DemoAudioMediaURL,omitempty" xml:"DemoAudioMediaURL,omitempty"`
	// The gender. Valid values:
	//
	// 	- female
	//
	// 	- male
	//
	// example:
	//
	// female
	Gender *string `json:"Gender,omitempty" xml:"Gender,omitempty"`
	// The voice name.
	//
	// 	- The name can be up to 32 characters in length.
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
