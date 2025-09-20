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
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The OSS path of the master playlist.
	//
	// The OSS path must be in the oss://${Bucket}/${Object} format. ${Bucket} specifies the name of the OSS bucket that is in the same region as the current project. ${Object} specifies the full path of the file that is suffixed with .m3u8.
	//
	// >  If a playlist contains subtitles or multiple outputs, the MasterURI parameter is required and the URI of subtitle files or outputs must be in the directory specified by the MasterURI parameter or its subdirectory.
	//
	// example:
	//
	// oss://bucket/object/master.m3u8
	MasterURI *string `json:"MasterURI,omitempty" xml:"MasterURI,omitempty"`
	// The notification settings. To view details, click Notification. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The overwrite policy when the media playlist exists. Valid values:
	//
	// 	- overwrite (default): overwrites an existing media playlist.
	//
	// 	- skip-existing: skips generation and retains the existing media playlist.
	//
	// example:
	//
	// overwrite
	OverwritePolicy *string `json:"OverwritePolicy,omitempty" xml:"OverwritePolicy,omitempty"`
	// The project name.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The period of time during which the playlist is generated. Unit: seconds.
	//
	// 	- If you set this parameter to 0 (default) or leave this parameter empty, a playlist is generated until the end time of the source video.
	//
	// 	- If you set this parameter to a value greater than 0, a playlist is generated for the specified period of time from the start time that you specify.
	//
	// >  If you set this parameter to a value that exceeds the end time of a source video, use the default value.
	//
	// example:
	//
	// 0
	SourceDuration *float32 `json:"SourceDuration,omitempty" xml:"SourceDuration,omitempty"`
	// The time when the playlist starts to generate. Unit: seconds.
	//
	// 	- If you set this parameter to 0 (default) or leave this parameter empty, the start time of the source video is used as the time when a playlist starts to generate.
	//
	// 	- If you set this parameter to a value greater than 0, the time when a playlist starts to generate is the specified point in time.
	//
	// >  If you use this parameter together with the **SourceDuration*	- parameter, a playlist can be generated based on the partial content of a source video.
	//
	// example:
	//
	// 0
	SourceStartTime *float32 `json:"SourceStartTime,omitempty" xml:"SourceStartTime,omitempty"`
	// The subtitle files. By default, this parameter is left empty. Up to two subtitle files are supported.
	SourceSubtitles []*GenerateVideoPlaylistRequestSourceSubtitles `json:"SourceSubtitles,omitempty" xml:"SourceSubtitles,omitempty" type:"Repeated"`
	// The OSS path of the video file.
	//
	// The OSS path must be in the oss://${Bucket}/${Object} format. ${Bucket} specifies the name of the OSS bucket that is in the same region as the current project. ${Object} specifies the full path of the file that contains the file name extension.
	//
	// >  Only OSS buckets of the Standard storage class are supported. OSS buckets for which hotlink protection whitelists are configured are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://imm-test/testcases/video.mp4
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The [tags](https://help.aliyun.com/document_detail/106678.html) that you want to add to a TS file in OSS. You can use tags to manage the lifecycles of TS files in OSS.
	//
	// example:
	//
	// {"key1": "value1", "key2": "value2"}
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The live transcoding playlists. Up to 6 playlists are supported. Each output corresponds to at most one video media playlist and one or more subtitle media playlists.
	//
	// >  If more than one output is configured, the **MasterURI*	- parameter is required.
	//
	// This parameter is required.
	Targets []*GenerateVideoPlaylistRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The custom user information, which is returned in asynchronous notifications to help you handle the notifications in the system. The maximum length of a notification is 2048 bytes.
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
	return dara.Validate(s)
}

type GenerateVideoPlaylistRequestSourceSubtitles struct {
	// The subtitle language. If you configure this parameter, the value must comply with the ISO 639-2 standard. By default, this parameter is left empty.
	//
	// example:
	//
	// eng
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The OSS path of the subtitle file.
	//
	// The OSS path must be in the oss://${Bucket}/${Object} format. ${Bucket} specifies the name of the OSS bucket that is in the same region as the current project. ${Object} specifies the full path of the file.
	//
	// >  The **MasterURI*	- parameter cannot be left empty, and the OSS path `oss://${Bucket}/${Object}` of a subtitle file must be in the directory specified by the **MasterURI*	- parameter or its subdirectory.
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
	// The audio processing configuration. If you set this parameter to null (default), audio processing is disabled. The generated TS files do not contain audio streams.
	//
	// >  The Audio and Subtitle parameters in the same output are mutually exclusive. If the Audio parameter is configured, the Subtitle parameter is ignored. The Audio and Video parameters can be configured at the same time. You can also configure only the Audio parameter to generate only audio information.
	Audio *TargetAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The playback duration of a single TS file. Unit: seconds. Default value: 10. Valid values: 5 to 15.
	//
	// example:
	//
	// 5
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The array of the durations of the pre-transcoded TS files. The array can contain the durations of up to six pre-transcoded TS files. By default, this parameter is left empty. This parameter is independent of the **Duration*	- parameter.
	InitialSegments []*float32 `json:"InitialSegments,omitempty" xml:"InitialSegments,omitempty" type:"Repeated"`
	// The pre-transcoding duration. Unit: seconds. Default value: 30.
	//
	// 	- If you set this parameter to 0, pre-transcoding is disabled.
	//
	// 	- If you set this parameter to a value that is less than 0 or greater than the duration of a source video, the entire video is pre-transcoded.
	//
	// 	- If you set this parameter to a value that is within the middle of the playback duration of a TS file, the transcoding continues until the end of the playback duration.
	//
	// >  This parameter is used to reduce the time spent in waiting for the initial playback of a video and improve the playback experience. If you want to replace the traditional video on demand (VOD) business scenario, you can try to pre-transcode the entire video.
	//
	// example:
	//
	// 30.0
	InitialTranscode *float32 `json:"InitialTranscode,omitempty" xml:"InitialTranscode,omitempty"`
	// The subtitle processing configuration.
	//
	// >  The Subtitle and Video or Audio parameters in the same output are mutually exclusive. You must configure the Subtitle parameter independently to generate subtitles.
	Subtitle *TargetSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// The [tags](https://help.aliyun.com/document_detail/106678.html) that you want to add to a TS file in OSS. You can use tags to manage the lifecycles of TS files in OSS.
	//
	// >  The combination of the value of the Tags parameter and the value of the Tags parameter in the upper level is used as the tag value of the current output. If the value of the Tags parameter in the current level is the same as the value of the Tags parameter in the upper level, use the value of the Tags parameter in the current level.
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The number of TS files that are pre-transcoded when the live transcoding is triggered. By default, a 2-minute video is pre-transcoded.
	//
	// 	- Example: If you set the **Duration*	- parameter to 10, the value of the **TranscodeAhead*	- parameter is 12 by default. You can configure this parameter to manage the number of pre-transcoded files in an asynchronous manner. Valid values: 10 to 30.
	//
	// example:
	//
	// 3
	TranscodeAhead *int32 `json:"TranscodeAhead,omitempty" xml:"TranscodeAhead,omitempty"`
	// The prefix of the OSS path that is used to store the live transcoding files. The live transcoding files include a M3U8 file and multiple TS files.
	//
	// The OSS path must be in the oss://${Bucket}/${Object} format. ${Bucket} specifies the name of the OSS bucket that is in the same region as the current project. ${Object} specifies the prefix of the full path of the file that does not contain the file name extension.
	//
	// 	- Example: If the URI is oss://test-bucket/test-object/output-video, the output-video.m3u8 file and multiple output-video-${token}-${index}.ts files are generated in the oss://test-bucket/test-object/ directory. ${token} is a unique string generated based on the transcoding parameters. The ${token} parameter is included in the response of the operation. ${index} is the serial number of the generated TS files that are numbered starting from 0.
	//
	// >  If the **MasterURI*	- parameter is not left empty, the URI specified by this parameter must be in the directory specified by the **MasterURI*	- parameter or its subdirectory.
	//
	// example:
	//
	// oss://imm-test/testcases/video
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The video processing configuration. If you set this parameter to null (default), video processing is disabled. The generated TS files do not contain video streams.
	//
	// >  The Video and Subtitle parameters in the same output are mutually exclusive. If the Video parameter is configured, the Subtitle parameter is ignored.
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
	return dara.Validate(s)
}
