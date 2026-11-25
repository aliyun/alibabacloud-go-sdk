// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitEditMediaManifestJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioTracks(v string) *SubmitEditMediaManifestJobRequest
	GetAudioTracks() *string
	SetAutoStreamJobId(v string) *SubmitEditMediaManifestJobRequest
	GetAutoStreamJobId() *string
	SetOutputConfig(v string) *SubmitEditMediaManifestJobRequest
	GetOutputConfig() *string
	SetRemoveTracks(v string) *SubmitEditMediaManifestJobRequest
	GetRemoveTracks() *string
	SetSubtitles(v string) *SubmitEditMediaManifestJobRequest
	GetSubtitles() *string
	SetTemplateId(v string) *SubmitEditMediaManifestJobRequest
	GetTemplateId() *string
	SetUserData(v string) *SubmitEditMediaManifestJobRequest
	GetUserData() *string
	SetVideoId(v string) *SubmitEditMediaManifestJobRequest
	GetVideoId() *string
	SetVideoTracks(v string) *SubmitEditMediaManifestJobRequest
	GetVideoTracks() *string
}

type SubmitEditMediaManifestJobRequest struct {
	// The list of audio tracks to be added. The value is a JSON string array. Each element contains the following fields:
	//
	// - **Language**: the language identifier of the audio track.
	//
	// - **Name**: the display name of the track. If this field is not specified, the default value is audio-{Language}.
	//
	// - **AudioUrl**: the URL of the audio file, which is an OSS object URL.
	//
	// example:
	//
	// [{"Language":"spa","Name":"audio-spa","AudioUrl":"https://your-bucket.oss-cn-shanghai.aliyuncs.com/audio/spa.aac"}]
	AudioTracks *string `json:"AudioTracks,omitempty" xml:"AudioTracks,omitempty"`
	// The job ID of the target AUTO stream.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2c6581d59240de5e258ffa4c71a3****
	AutoStreamJobId *string `json:"AutoStreamJobId,omitempty" xml:"AutoStreamJobId,omitempty"`
	// The output configuration. The value is a JSON string. It contains the following fields:
	//
	// - **OutputDir**: the output directory. By default, this field is set to the same directory as the source manifest.
	//
	// - **ManifestName**: the file name of the output master playlist. Default value: master. The .m3u8 extension is not allowed.
	//
	// example:
	//
	// {"OutputDir":"output/editmanifest/","ManifestName":"master"}
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// The configuration of tracks to be removed. The value is a JSON string. It contains the following fields:
	//
	// - **AudioLanguages**: the list of audio language identifiers for which the entire audio tracks are to be removed.
	//
	// - **SubtitleLanguages**: the list of subtitle language identifiers for which the entire subtitle tracks are to be removed.
	//
	// example:
	//
	// {"AudioLanguages":["fra"],"SubtitleLanguages":["ger"]}
	RemoveTracks *string `json:"RemoveTracks,omitempty" xml:"RemoveTracks,omitempty"`
	// The list of subtitle tracks to be added. The value is a JSON string array. Each element contains the following fields:
	//
	// - **Language**: the language identifier of the subtitle track. If this field is not specified, the subtitle track is ignored.
	//
	// - **SubtitleUrl**: the URL of the subtitle file in WebVTT format, which is an OSS object URL.
	//
	// example:
	//
	// [{"Language":"eng","SubtitleUrl":"https://your-bucket.oss-cn-shanghai.aliyuncs.com/sub/eng.vtt"}]
	Subtitles *string `json:"Subtitles,omitempty" xml:"Subtitles,omitempty"`
	// The manifest editing template ID, which is returned after the AddManifestEditTemplate operation is called. If this parameter is specified, the output tracks are orchestrated based on the fixed video, audio, and subtitle item configurations in the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9c6581d59240de5e258ffa4c71a3e01b
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The custom pass-through data.
	//
	// example:
	//
	// {"remark":"add-french-remove-english"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// **The target video ID**, used for authentication isolation and consistency verification.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2cd8a4f4a67cb7fb8e1236f847f4****
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	// The list of video tracks to be added. The value is a JSON string array. Each element contains the following fields:
	//
	// - **Name**: the display name of the track. If this field is not specified, the video track is ignored.
	//
	// - **VideoUrl**: the URL of the video stream manifest file, which is an OSS object URL.
	//
	// example:
	//
	// [{"Name":"uhd","VideoUrl":"https://your-bucket.oss-cn-shanghai.aliyuncs.com/video/uhd/uhd.m3u8"}]
	VideoTracks *string `json:"VideoTracks,omitempty" xml:"VideoTracks,omitempty"`
}

func (s SubmitEditMediaManifestJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitEditMediaManifestJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitEditMediaManifestJobRequest) GetAudioTracks() *string {
	return s.AudioTracks
}

func (s *SubmitEditMediaManifestJobRequest) GetAutoStreamJobId() *string {
	return s.AutoStreamJobId
}

func (s *SubmitEditMediaManifestJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitEditMediaManifestJobRequest) GetRemoveTracks() *string {
	return s.RemoveTracks
}

func (s *SubmitEditMediaManifestJobRequest) GetSubtitles() *string {
	return s.Subtitles
}

func (s *SubmitEditMediaManifestJobRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SubmitEditMediaManifestJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitEditMediaManifestJobRequest) GetVideoId() *string {
	return s.VideoId
}

func (s *SubmitEditMediaManifestJobRequest) GetVideoTracks() *string {
	return s.VideoTracks
}

func (s *SubmitEditMediaManifestJobRequest) SetAudioTracks(v string) *SubmitEditMediaManifestJobRequest {
	s.AudioTracks = &v
	return s
}

func (s *SubmitEditMediaManifestJobRequest) SetAutoStreamJobId(v string) *SubmitEditMediaManifestJobRequest {
	s.AutoStreamJobId = &v
	return s
}

func (s *SubmitEditMediaManifestJobRequest) SetOutputConfig(v string) *SubmitEditMediaManifestJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitEditMediaManifestJobRequest) SetRemoveTracks(v string) *SubmitEditMediaManifestJobRequest {
	s.RemoveTracks = &v
	return s
}

func (s *SubmitEditMediaManifestJobRequest) SetSubtitles(v string) *SubmitEditMediaManifestJobRequest {
	s.Subtitles = &v
	return s
}

func (s *SubmitEditMediaManifestJobRequest) SetTemplateId(v string) *SubmitEditMediaManifestJobRequest {
	s.TemplateId = &v
	return s
}

func (s *SubmitEditMediaManifestJobRequest) SetUserData(v string) *SubmitEditMediaManifestJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitEditMediaManifestJobRequest) SetVideoId(v string) *SubmitEditMediaManifestJobRequest {
	s.VideoId = &v
	return s
}

func (s *SubmitEditMediaManifestJobRequest) SetVideoTracks(v string) *SubmitEditMediaManifestJobRequest {
	s.VideoTracks = &v
	return s
}

func (s *SubmitEditMediaManifestJobRequest) Validate() error {
	return dara.Validate(s)
}
