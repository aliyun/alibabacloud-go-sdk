// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomizedVoiceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDemoAudioMediaURL(v string) *SubmitCustomizedVoiceJobRequest
	GetDemoAudioMediaURL() *string
	SetVoiceId(v string) *SubmitCustomizedVoiceJobRequest
	GetVoiceId() *string
}

type SubmitCustomizedVoiceJobRequest struct {
	// The OSS URL where the demo audio will be saved.
	//
	// - If specified, the service generates a demo audio file at the provided OSS URL after training completes.
	//
	// - 	Notice:
	//
	//   The URL must be a valid public address for an OSS object in your account.
	//
	// example:
	//
	// https://your-bucket.oss-cn-shanghai.aliyuncs.com/demo.MP3
	DemoAudioMediaURL *string `json:"DemoAudioMediaURL,omitempty" xml:"DemoAudioMediaURL,omitempty"`
	// The unique identifier for the voice.
	//
	// This parameter is required.
	//
	// example:
	//
	// xiaozhuan
	VoiceId *string `json:"VoiceId,omitempty" xml:"VoiceId,omitempty"`
}

func (s SubmitCustomizedVoiceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomizedVoiceJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitCustomizedVoiceJobRequest) GetDemoAudioMediaURL() *string {
	return s.DemoAudioMediaURL
}

func (s *SubmitCustomizedVoiceJobRequest) GetVoiceId() *string {
	return s.VoiceId
}

func (s *SubmitCustomizedVoiceJobRequest) SetDemoAudioMediaURL(v string) *SubmitCustomizedVoiceJobRequest {
	s.DemoAudioMediaURL = &v
	return s
}

func (s *SubmitCustomizedVoiceJobRequest) SetVoiceId(v string) *SubmitCustomizedVoiceJobRequest {
	s.VoiceId = &v
	return s
}

func (s *SubmitCustomizedVoiceJobRequest) Validate() error {
	return dara.Validate(s)
}
