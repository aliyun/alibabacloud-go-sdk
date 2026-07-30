// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoRenderJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScript(v string) *SubmitVideoRenderJobRequest
	GetScript() *string
	SetSettings(v string) *SubmitVideoRenderJobRequest
	GetSettings() *string
	SetUserData(v string) *SubmitVideoRenderJobRequest
	GetUserData() *string
}

type SubmitVideoRenderJobRequest struct {
	// example:
	//
	// {
	//
	// "schemaVersion":"creative/v1",
	//
	// "algoResult":{...},
	//
	// "extraInfo":{...}
	//
	// }
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
	// example:
	//
	// {
	//
	//   "VoiceoverLanguage": "zh",
	//
	//   "Resolution": "1080P",
	//
	//   "AspectRatio": "9:16",
	//
	//   "TTS": {
	//
	//     "VoiceUrl": "http://xxx.mp3"
	//
	//   },
	//
	//   "WithSubtitles": true,
	//
	//   "Bgm": "http://xxx.mp3"
	//
	// }
	Settings *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// example:
	//
	// {"NotifyAddress": "http://xxx.callback.url"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitVideoRenderJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoRenderJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoRenderJobRequest) GetScript() *string {
	return s.Script
}

func (s *SubmitVideoRenderJobRequest) GetSettings() *string {
	return s.Settings
}

func (s *SubmitVideoRenderJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitVideoRenderJobRequest) SetScript(v string) *SubmitVideoRenderJobRequest {
	s.Script = &v
	return s
}

func (s *SubmitVideoRenderJobRequest) SetSettings(v string) *SubmitVideoRenderJobRequest {
	s.Settings = &v
	return s
}

func (s *SubmitVideoRenderJobRequest) SetUserData(v string) *SubmitVideoRenderJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitVideoRenderJobRequest) Validate() error {
	return dara.Validate(s)
}
