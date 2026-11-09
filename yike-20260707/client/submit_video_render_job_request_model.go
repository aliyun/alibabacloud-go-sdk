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
	// The complete creative script confirmed or edited by the user, provided as a JSON string. The structure aligns with the JSON content in the Result file returned by the GetRemakeScriptJob API.
	//
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
	// The rendering settings, provided as a JSON string.
	//
	// - **Resolution*	- (String, required): The resolution. Valid values: `720P` and `1080P`.
	//
	//   - **AspectRatio*	- (String, optional): The video aspect ratio. Valid values: `9:16`, `16:9`, and `1:1`. Default value: `9:16`.
	//
	//   - **VoiceoverLanguage*	- (String, optional): The voiceover language. Valid values: `zh` (Chinese), `en` (English), `es` (Spanish), `pt` (Portuguese), `fr` (French), `de` (German), `ja` (Japanese), `ko` (Korean), and `ar` (Arabic). Default value: `zh`.
	//
	//   - **WithSubtitles*	- (Boolean, optional): Specifies whether to generate subtitles. Default value: `true`.
	//
	//   - **TTS*	- (Object, optional): The text-to-speech (TTS) configuration. If this parameter is not specified, the default voice is used. This parameter applies only to single-person scenarios and pure voiceovers.
	//
	//     - **VoiceUrl*	- (String, optional): The URL of the voice file, which must be an HTTP or HTTPS URL. If specified, this voice is used for the voiceover throughout the entire video.
	//
	//   - **Bgm*	- (String, optional): The URL or 32-character media asset ID of the background music.
	//
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
	// The user-defined parameters, provided as a JSON string.
	//
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
