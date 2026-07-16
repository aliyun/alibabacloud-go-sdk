// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateVideoPlaylistRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *GenerateVideoPlaylistRequest
	GetCredentialConfig() *CredentialConfig
	SetMasterURI(v string) *GenerateVideoPlaylistRequest
	GetMasterURI() *string
	SetNotification(v *Notification) *GenerateVideoPlaylistRequest
	GetNotification() *Notification
	SetOverwritePolicy(v string) *GenerateVideoPlaylistRequest
	GetOverwritePolicy() *string
	SetProjectName(v string) *GenerateVideoPlaylistRequest
	GetProjectName() *string
	SetSourceDuration(v float32) *GenerateVideoPlaylistRequest
	GetSourceDuration() *float32
	SetSourceStartTime(v float32) *GenerateVideoPlaylistRequest
	GetSourceStartTime() *float32
	SetSourceSubtitles(v []*GenerateVideoPlaylistRequestSourceSubtitles) *GenerateVideoPlaylistRequest
	GetSourceSubtitles() []*GenerateVideoPlaylistRequestSourceSubtitles
	SetSourceURI(v string) *GenerateVideoPlaylistRequest
	GetSourceURI() *string
	SetTags(v map[string]*string) *GenerateVideoPlaylistRequest
	GetTags() map[string]*string
	SetTargets(v []*GenerateVideoPlaylistRequestTargets) *GenerateVideoPlaylistRequest
	GetTargets() []*GenerateVideoPlaylistRequestTargets
	SetUserData(v string) *GenerateVideoPlaylistRequest
	GetUserData() *string
}

type GenerateVideoPlaylistRequest struct {
	// **Leave this parameter empty unless you have specific requirements.**
	//
	// The China authorization configuration. This parameter is optional. For more information, see [Use Chinese authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The OSS URI of the Master Playlist.
	//
	// The OSS URI follows the format oss://${Bucket}/${Object}, where ${Bucket} is the name of the OSS bucket in the same region as the current project, and ${Object} is the full path of the file with the ".m3u8" extension.
	//
	// > If the playlist has subtitle input or multiple Target outputs, MasterURI is required. The subtitle URI or Target URI must be in the same directory as or a subdirectory of MasterURI.
	//
	// example:
	//
	// oss://test-bucket/test-object/master.m3u8
	MasterURI *string `json:"MasterURI,omitempty" xml:"MasterURI,omitempty"`
	// The message notification configuration. Click Notification for details. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The overwrite policy when a Media Playlist already exists. Valid values:
	//
	// - overwrite (default): overwrites the existing Media Playlist.
	//
	// - skip-existing: skips generation and retains the existing Media Playlist.
	//
	// example:
	//
	// overwrite
	OverwritePolicy *string `json:"OverwritePolicy,omitempty" xml:"OverwritePolicy,omitempty"`
	// The project name. For information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The duration for generating the playlist. Unit: seconds. Valid values:
	//
	// - 0 (default) or empty: continues until the end of the source video.
	//
	// - A value greater than 0: continues for the specified duration from the start time of the playlist.
	//
	// > If the time point corresponding to the specified parameter exceeds the end of the source video, the default value is used.
	//
	// example:
	//
	// 0
	SourceDuration *float32 `json:"SourceDuration,omitempty" xml:"SourceDuration,omitempty"`
	// The start time for generating the playlist. Unit: seconds. Valid values:
	//
	// - 0 (default) or empty: starts from the beginning of the source video.
	//
	// - A value greater than 0: starts from the specified time point in the source video.
	//
	// > You can set this parameter together with **SourceDuration*	- to generate a playlist for a specific portion of the source video.
	//
	// example:
	//
	// 0
	SourceStartTime *float32 `json:"SourceStartTime,omitempty" xml:"SourceStartTime,omitempty"`
	// The list of subtitles to add. This parameter is empty by default. A maximum of two subtitles are supported.
	SourceSubtitles []*GenerateVideoPlaylistRequestSourceSubtitles `json:"SourceSubtitles,omitempty" xml:"SourceSubtitles,omitempty" type:"Repeated"`
	// The OSS URI of the video.
	//
	// The OSS URI follows the format oss://${Bucket}/${Object}, where ${Bucket} is the name of the OSS bucket in the same region as the current project, and ${Object} is the full path of the file including the file name extension.
	//
	// > Only OSS buckets with Standard storage class are supported.
	//
	// > Buckets with hotlink protection whitelist configured are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-source-object/video.mp4
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The OSS object [tags](https://help.aliyun.com/document_detail/106678.html) to add to the generated TS files. You can use tags to control the lifecycle of OSS files.
	//
	// example:
	//
	// {"key1": "value1", "key2": "value2"}
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The array of just-in-time transcoding playlists. The maximum array length is 6. Each Target corresponds to at most one video Media Playlist and one or more subtitle Media Playlists.
	//
	// > If more than one Target is configured, the **MasterURI*	- parameter must not be empty.
	//
	// This parameter is required.
	Targets []*GenerateVideoPlaylistRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The custom information, which is returned in asynchronous message notifications. This allows you to associate message notifications with specific processes in your system. Maximum length: 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s GenerateVideoPlaylistRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoPlaylistRequest) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *GenerateVideoPlaylistRequest) GetMasterURI() *string {
	return s.MasterURI
}

func (s *GenerateVideoPlaylistRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *GenerateVideoPlaylistRequest) GetOverwritePolicy() *string {
	return s.OverwritePolicy
}

func (s *GenerateVideoPlaylistRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GenerateVideoPlaylistRequest) GetSourceDuration() *float32 {
	return s.SourceDuration
}

func (s *GenerateVideoPlaylistRequest) GetSourceStartTime() *float32 {
	return s.SourceStartTime
}

func (s *GenerateVideoPlaylistRequest) GetSourceSubtitles() []*GenerateVideoPlaylistRequestSourceSubtitles {
	return s.SourceSubtitles
}

func (s *GenerateVideoPlaylistRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *GenerateVideoPlaylistRequest) GetTags() map[string]*string {
	return s.Tags
}

func (s *GenerateVideoPlaylistRequest) GetTargets() []*GenerateVideoPlaylistRequestTargets {
	return s.Targets
}

func (s *GenerateVideoPlaylistRequest) GetUserData() *string {
	return s.UserData
}

func (s *GenerateVideoPlaylistRequest) SetCredentialConfig(v *CredentialConfig) *GenerateVideoPlaylistRequest {
	s.CredentialConfig = v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetMasterURI(v string) *GenerateVideoPlaylistRequest {
	s.MasterURI = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetNotification(v *Notification) *GenerateVideoPlaylistRequest {
	s.Notification = v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetOverwritePolicy(v string) *GenerateVideoPlaylistRequest {
	s.OverwritePolicy = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetProjectName(v string) *GenerateVideoPlaylistRequest {
	s.ProjectName = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetSourceDuration(v float32) *GenerateVideoPlaylistRequest {
	s.SourceDuration = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetSourceStartTime(v float32) *GenerateVideoPlaylistRequest {
	s.SourceStartTime = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetSourceSubtitles(v []*GenerateVideoPlaylistRequestSourceSubtitles) *GenerateVideoPlaylistRequest {
	s.SourceSubtitles = v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetSourceURI(v string) *GenerateVideoPlaylistRequest {
	s.SourceURI = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetTags(v map[string]*string) *GenerateVideoPlaylistRequest {
	s.Tags = v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetTargets(v []*GenerateVideoPlaylistRequestTargets) *GenerateVideoPlaylistRequest {
	s.Targets = v
	return s
}

func (s *GenerateVideoPlaylistRequest) SetUserData(v string) *GenerateVideoPlaylistRequest {
	s.UserData = &v
	return s
}

func (s *GenerateVideoPlaylistRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.SourceSubtitles != nil {
		for _, item := range s.SourceSubtitles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Targets != nil {
		for _, item := range s.Targets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GenerateVideoPlaylistRequestSourceSubtitles struct {
	// The subtitle language. The value follows the ISO 639-2 standard. This parameter is empty by default.
	//
	// example:
	//
	// eng
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The OSS URI of the subtitle to embed.
	//
	// The OSS URI follows the format oss://${Bucket}/${Object}, where ${Bucket} is the name of the OSS bucket in the same region as the current project, and ${Object} is the full path of the file.
	//
	// > The **MasterURI*	- parameter must not be empty, and the OSS URI `oss://${Bucket}/${Object}` of the subtitle must be in the same directory as or a subdirectory of the **MasterURI*	- parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object/subtitle/eng.vtt
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s GenerateVideoPlaylistRequestSourceSubtitles) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoPlaylistRequestSourceSubtitles) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistRequestSourceSubtitles) GetLanguage() *string {
	return s.Language
}

func (s *GenerateVideoPlaylistRequestSourceSubtitles) GetURI() *string {
	return s.URI
}

func (s *GenerateVideoPlaylistRequestSourceSubtitles) SetLanguage(v string) *GenerateVideoPlaylistRequestSourceSubtitles {
	s.Language = &v
	return s
}

func (s *GenerateVideoPlaylistRequestSourceSubtitles) SetURI(v string) *GenerateVideoPlaylistRequestSourceSubtitles {
	s.URI = &v
	return s
}

func (s *GenerateVideoPlaylistRequestSourceSubtitles) Validate() error {
	return dara.Validate(s)
}

type GenerateVideoPlaylistRequestTargets struct {
	// The audio processing parameter settings. An empty value (default) indicates that audio processing is disabled and the output TS files do not contain audio streams.
	//
	// > The Audio and Subtitle fields within the same Target are mutually exclusive. If the Audio field is set, the Subtitle field is ignored. Audio and Video can be set simultaneously. Audio specifies the audio information in the output video. You can also set only Audio to generate audio-only output.
	Audio     *TargetAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	Container *string      `json:"Container,omitempty" xml:"Container,omitempty"`
	// The playback duration of a single TS file. Unit: seconds. Default value: 10. Valid values: [5, 15].
	//
	// example:
	//
	// 10
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The array of initial transcoding TS file durations. The maximum array length is 6. This parameter is empty by default and is independent of the **Duration*	- parameter.
	InitialSegments []*float32 `json:"InitialSegments,omitempty" xml:"InitialSegments,omitempty" type:"Repeated"`
	// The initial transcoding duration. Unit: seconds. Default value: 30.
	//
	// - If the value is set to 0, no pre-transcoding is performed.
	//
	// - If the value is less than 0 or exceeds the source video length, the entire video is initially transcoded.
	//
	// - If the specified duration falls in the middle of a TS file, transcoding continues until the end of that TS file.
	//
	// > This parameter is primarily used to reduce the wait time for initial video playback and improve the playback experience. If you want to replace traditional VOD business scenarios, try initially transcoding the entire video.
	//
	// example:
	//
	// 30
	InitialTranscode *float32 `json:"InitialTranscode,omitempty" xml:"InitialTranscode,omitempty"`
	// The subtitle processing parameter settings.
	//
	// > The Subtitle field is mutually exclusive with the Video or Audio fields within the same Target. Subtitles are generated only when Subtitle is set independently.
	Subtitle *TargetSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// The OSS object [tags](https://help.aliyun.com/document_detail/106678.html) to add to the generated TS files. You can use OSS tags to control the lifecycle of OSS files.
	//
	// > The tag values at this level are merged with the Tags defined at the parent level to form the tag values for the current Target. If a tag with the same name exists, the value at this level takes precedence.
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The number of TS files to transcode ahead when just-in-time transcoding is triggered. By default, 2 minutes of video is transcoded ahead.
	//
	// - Example: If **Duration*	- is 10, the default value of **TranscodeAhead*	- is 12. You can specify this parameter to control the number of asynchronous ahead-of-time transcoding files. Valid values: [10, 30].
	//
	// example:
	//
	// 12
	TranscodeAhead *int32 `json:"TranscodeAhead,omitempty" xml:"TranscodeAhead,omitempty"`
	// The OSS URI prefix of the just-in-time transcoding output files, including M3U8 files and TS files.
	//
	// The OSS URI follows the format oss://${Bucket}/${Object}, where ${Bucket} is the name of the OSS bucket in the same region as the current project, and ${Object} is the full path prefix of the file without the file name extension.
	//
	// - Example: If URI is oss://test-bucket/test-object/output-video, an oss://test-bucket/test-object/output-video.m3u8 file and multiple oss://test-bucket/test-object/output-video-${token}-${index}.ts files are generated. ${token} is a unique string generated based on the transcoding parameters and is included in the API response. ${index} is the sequence number of the TS file starting from 0.
	//
	// > If the **MasterURI*	- parameter is not empty, the URI must be in the same directory as or a subdirectory of the **MasterURI*	- parameter.
	//
	// example:
	//
	// oss://test-bucket/test-object/output-video
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The video processing parameter settings. An empty value (default) indicates that video processing is disabled and the output TS files do not contain video streams.
	//
	// > The Video and Subtitle fields within the same Target are mutually exclusive. If the Video field is set, the Subtitle field is ignored.
	Video *TargetVideo `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s GenerateVideoPlaylistRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s GenerateVideoPlaylistRequestTargets) GoString() string {
	return s.String()
}

func (s *GenerateVideoPlaylistRequestTargets) GetAudio() *TargetAudio {
	return s.Audio
}

func (s *GenerateVideoPlaylistRequestTargets) GetContainer() *string {
	return s.Container
}

func (s *GenerateVideoPlaylistRequestTargets) GetDuration() *float32 {
	return s.Duration
}

func (s *GenerateVideoPlaylistRequestTargets) GetInitialSegments() []*float32 {
	return s.InitialSegments
}

func (s *GenerateVideoPlaylistRequestTargets) GetInitialTranscode() *float32 {
	return s.InitialTranscode
}

func (s *GenerateVideoPlaylistRequestTargets) GetSubtitle() *TargetSubtitle {
	return s.Subtitle
}

func (s *GenerateVideoPlaylistRequestTargets) GetTags() map[string]*string {
	return s.Tags
}

func (s *GenerateVideoPlaylistRequestTargets) GetTranscodeAhead() *int32 {
	return s.TranscodeAhead
}

func (s *GenerateVideoPlaylistRequestTargets) GetURI() *string {
	return s.URI
}

func (s *GenerateVideoPlaylistRequestTargets) GetVideo() *TargetVideo {
	return s.Video
}

func (s *GenerateVideoPlaylistRequestTargets) SetAudio(v *TargetAudio) *GenerateVideoPlaylistRequestTargets {
	s.Audio = v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetContainer(v string) *GenerateVideoPlaylistRequestTargets {
	s.Container = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetDuration(v float32) *GenerateVideoPlaylistRequestTargets {
	s.Duration = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetInitialSegments(v []*float32) *GenerateVideoPlaylistRequestTargets {
	s.InitialSegments = v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetInitialTranscode(v float32) *GenerateVideoPlaylistRequestTargets {
	s.InitialTranscode = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetSubtitle(v *TargetSubtitle) *GenerateVideoPlaylistRequestTargets {
	s.Subtitle = v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetTags(v map[string]*string) *GenerateVideoPlaylistRequestTargets {
	s.Tags = v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetTranscodeAhead(v int32) *GenerateVideoPlaylistRequestTargets {
	s.TranscodeAhead = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetURI(v string) *GenerateVideoPlaylistRequestTargets {
	s.URI = &v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) SetVideo(v *TargetVideo) *GenerateVideoPlaylistRequestTargets {
	s.Video = v
	return s
}

func (s *GenerateVideoPlaylistRequestTargets) Validate() error {
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Subtitle != nil {
		if err := s.Subtitle.Validate(); err != nil {
			return err
		}
	}
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}
