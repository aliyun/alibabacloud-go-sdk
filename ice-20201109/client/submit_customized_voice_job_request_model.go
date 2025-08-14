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
	// https://your-bucket.oss-cn-shanghai.aliyuncs.com/demo.MP3
	DemoAudioMediaURL *string `json:"DemoAudioMediaURL,omitempty" xml:"DemoAudioMediaURL,omitempty"`
	// The voice ID.
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
