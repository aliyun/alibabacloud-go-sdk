// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMediaConvertTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlignmentIndex(v int32) *CreateMediaConvertTaskRequest
	GetAlignmentIndex() *int32
	SetCredentialConfig(v *CredentialConfig) *CreateMediaConvertTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetNotification(v *Notification) *CreateMediaConvertTaskRequest
	GetNotification() *Notification
	SetProjectName(v string) *CreateMediaConvertTaskRequest
	GetProjectName() *string
	SetSources(v []*CreateMediaConvertTaskRequestSources) *CreateMediaConvertTaskRequest
	GetSources() []*CreateMediaConvertTaskRequestSources
	SetTags(v map[string]interface{}) *CreateMediaConvertTaskRequest
	GetTags() map[string]interface{}
	SetTargetGroups(v []*CreateMediaConvertTaskRequestTargetGroups) *CreateMediaConvertTaskRequest
	GetTargetGroups() []*CreateMediaConvertTaskRequestTargetGroups
	SetTargets(v []*CreateMediaConvertTaskRequestTargets) *CreateMediaConvertTaskRequest
	GetTargets() []*CreateMediaConvertTaskRequestTargets
	SetUserData(v string) *CreateMediaConvertTaskRequest
	GetUserData() *string
}

type CreateMediaConvertTaskRequest struct {
	// When concatenating media files, this specifies the index of the primary file in the Sources list. The default transcoding parameters (such as resolution and frame rate from the `Video` and `Audio` objects) are taken from this primary file. The default index is 0.
	//
	// example:
	//
	// 0
	AlignmentIndex *int32 `json:"AlignmentIndex,omitempty" xml:"AlignmentIndex,omitempty"`
	// **You can leave this parameter empty if you do not have special requirements.**
	//
	// The chained authorization configuration. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The message notification settings. For more information, click Notification. For information about the format of asynchronous notifications, see [Asynchronous notification format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// A list of media files. If you provide more than one file, they are concatenated in the order of their URIs.
	//
	// This parameter is required.
	Sources []*CreateMediaConvertTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// Custom tags for searching and filtering asynchronous tasks.
	//
	// example:
	//
	// {"test":"val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// A list of media packaging tasks to convert and package the input media into HLS outputs. Each TargetGroup corresponds to one HLS master playlist.
	TargetGroups []*CreateMediaConvertTaskRequestTargetGroups `json:"TargetGroups,omitempty" xml:"TargetGroups,omitempty" type:"Repeated"`
	// A list of media processing tasks.
	Targets []*CreateMediaConvertTaskRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The custom user data. This data is returned in the asynchronous notification, allowing you to associate the notification with your internal system. The maximum length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "testuid","Name": "test-user","Avatar": "http://test.com/testuid"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateMediaConvertTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequest) GetAlignmentIndex() *int32 {
	return s.AlignmentIndex
}

func (s *CreateMediaConvertTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateMediaConvertTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateMediaConvertTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateMediaConvertTaskRequest) GetSources() []*CreateMediaConvertTaskRequestSources {
	return s.Sources
}

func (s *CreateMediaConvertTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateMediaConvertTaskRequest) GetTargetGroups() []*CreateMediaConvertTaskRequestTargetGroups {
	return s.TargetGroups
}

func (s *CreateMediaConvertTaskRequest) GetTargets() []*CreateMediaConvertTaskRequestTargets {
	return s.Targets
}

func (s *CreateMediaConvertTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateMediaConvertTaskRequest) SetAlignmentIndex(v int32) *CreateMediaConvertTaskRequest {
	s.AlignmentIndex = &v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateMediaConvertTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetNotification(v *Notification) *CreateMediaConvertTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetProjectName(v string) *CreateMediaConvertTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetSources(v []*CreateMediaConvertTaskRequestSources) *CreateMediaConvertTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetTags(v map[string]interface{}) *CreateMediaConvertTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetTargetGroups(v []*CreateMediaConvertTaskRequestTargetGroups) *CreateMediaConvertTaskRequest {
	s.TargetGroups = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetTargets(v []*CreateMediaConvertTaskRequestTargets) *CreateMediaConvertTaskRequest {
	s.Targets = v
	return s
}

func (s *CreateMediaConvertTaskRequest) SetUserData(v string) *CreateMediaConvertTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateMediaConvertTaskRequest) Validate() error {
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
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TargetGroups != nil {
		for _, item := range s.TargetGroups {
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

type CreateMediaConvertTaskRequestSources struct {
	// The alignment mode for the added audio and video streams. Valid values include:
	//
	// - false (default): No alignment is performed.
	//
	// - loop: Aligns content by looping the audio or video.
	//
	// - pad: Aligns content by padding with silent frames or black frames.
	//
	// > 	- This parameter only takes effect if Attached is set to true.
	//
	// example:
	//
	// false
	AlignMode *string `json:"AlignMode,omitempty" xml:"AlignMode,omitempty"`
	// If true, adds the current source media file to the output as a synchronized audio stream or video stream. The default is false.
	//
	// > - You cannot set Attached to true for the source media file referenced by AlignmentIndex.
	//
	// example:
	//
	// false
	Attached *bool `json:"Attached,omitempty" xml:"Attached,omitempty"`
	// Specifies whether to disable the audio from the source media file. Valid values include:
	//
	// - true: Disables the audio.
	//
	// - false (default): Includes the audio.
	//
	// example:
	//
	// false
	DisableAudio *bool `json:"DisableAudio,omitempty" xml:"DisableAudio,omitempty"`
	// Specifies whether to disable the video from the source media file. Valid values include:
	//
	// - true: Disables the video.
	//
	// - false (default): Includes the video.
	//
	// example:
	//
	// false
	DisableVideo *bool `json:"DisableVideo,omitempty" xml:"DisableVideo,omitempty"`
	// The duration of media transcoding in seconds. The default value, 0, transcodes the media until its end.
	//
	// example:
	//
	// 0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of media transcoding, in seconds. Valid values include:
	//
	// - 0 (default): Transcoding starts from the beginning of the media file.
	//
	// - n (a value greater than 0): Transcoding starts n seconds into the media file.
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A list of subtitles to add.
	Subtitles []*CreateMediaConvertTaskRequestSourcesSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The OSS URI of the object. The URI must use the `oss://${Bucket}/${Object}` format, where `${Bucket}` is the name of an OSS bucket in the same region as the project, and `${Object}` is the full path to the object, including its file extension.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateMediaConvertTaskRequestSources) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestSources) GetAlignMode() *string {
	return s.AlignMode
}

func (s *CreateMediaConvertTaskRequestSources) GetAttached() *bool {
	return s.Attached
}

func (s *CreateMediaConvertTaskRequestSources) GetDisableAudio() *bool {
	return s.DisableAudio
}

func (s *CreateMediaConvertTaskRequestSources) GetDisableVideo() *bool {
	return s.DisableVideo
}

func (s *CreateMediaConvertTaskRequestSources) GetDuration() *float64 {
	return s.Duration
}

func (s *CreateMediaConvertTaskRequestSources) GetStartTime() *float64 {
	return s.StartTime
}

func (s *CreateMediaConvertTaskRequestSources) GetSubtitles() []*CreateMediaConvertTaskRequestSourcesSubtitles {
	return s.Subtitles
}

func (s *CreateMediaConvertTaskRequestSources) GetURI() *string {
	return s.URI
}

func (s *CreateMediaConvertTaskRequestSources) SetAlignMode(v string) *CreateMediaConvertTaskRequestSources {
	s.AlignMode = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetAttached(v bool) *CreateMediaConvertTaskRequestSources {
	s.Attached = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetDisableAudio(v bool) *CreateMediaConvertTaskRequestSources {
	s.DisableAudio = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetDisableVideo(v bool) *CreateMediaConvertTaskRequestSources {
	s.DisableVideo = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetDuration(v float64) *CreateMediaConvertTaskRequestSources {
	s.Duration = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetStartTime(v float64) *CreateMediaConvertTaskRequestSources {
	s.StartTime = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetSubtitles(v []*CreateMediaConvertTaskRequestSourcesSubtitles) *CreateMediaConvertTaskRequestSources {
	s.Subtitles = v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) SetURI(v string) *CreateMediaConvertTaskRequestSources {
	s.URI = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSources) Validate() error {
	if s.Subtitles != nil {
		for _, item := range s.Subtitles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateMediaConvertTaskRequestSourcesSubtitles struct {
	// The language of the subtitle. The value must comply with the ISO 639-2 standard.
	//
	// example:
	//
	// eng
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The subtitle delay, in seconds. The default value is 0.
	//
	// example:
	//
	// 10.5
	TimeOffset *float64 `json:"TimeOffset,omitempty" xml:"TimeOffset,omitempty"`
	// The OSS URI of the object. The URI must use the `oss://${Bucket}/${Object}` format, where `${Bucket}` is the name of an OSS bucket in the same region as the project, and `${Object}` is the full path to the object, including its file extension.
	//
	// Supported subtitle formats include: srt, vtt, mov_text, ass, dvd_sub, and pgs.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateMediaConvertTaskRequestSourcesSubtitles) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestSourcesSubtitles) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) GetLanguage() *string {
	return s.Language
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) GetTimeOffset() *float64 {
	return s.TimeOffset
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) GetURI() *string {
	return s.URI
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) SetLanguage(v string) *CreateMediaConvertTaskRequestSourcesSubtitles {
	s.Language = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) SetTimeOffset(v float64) *CreateMediaConvertTaskRequestSourcesSubtitles {
	s.TimeOffset = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) SetURI(v string) *CreateMediaConvertTaskRequestSourcesSubtitles {
	s.URI = &v
	return s
}

func (s *CreateMediaConvertTaskRequestSourcesSubtitles) Validate() error {
	return dara.Validate(s)
}

type CreateMediaConvertTaskRequestTargetGroups struct {
	// A list of media packaging subtasks. Each `Target` corresponds to a variant stream (`#EXT-X-STREAM-INF`) in the HLS master playlist and generates a corresponding HLS media playlist.
	Targets []*CreateMediaConvertTaskRequestTargetGroupsTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The OSS URI of the output HLS master playlist file for the packaging task.
	//
	// example:
	//
	// oss://test-bucket/test-object/master.m3u8
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetGroups) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetGroups) GetTargets() []*CreateMediaConvertTaskRequestTargetGroupsTargets {
	return s.Targets
}

func (s *CreateMediaConvertTaskRequestTargetGroups) GetURI() *string {
	return s.URI
}

func (s *CreateMediaConvertTaskRequestTargetGroups) SetTargets(v []*CreateMediaConvertTaskRequestTargetGroupsTargets) *CreateMediaConvertTaskRequestTargetGroups {
	s.Targets = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroups) SetURI(v string) *CreateMediaConvertTaskRequestTargetGroups {
	s.URI = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroups) Validate() error {
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

type CreateMediaConvertTaskRequestTargetGroupsTargets struct {
	// The audio processing parameters.
	//
	// 	Notice: If this parameter is left empty, the first audio stream, if it exists, is copied directly to the output file.
	Audio *TargetAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The packaging container type. Only `mp4` and `ts` are supported.
	//
	// example:
	//
	// mp4
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The media packaging settings.
	Segment *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// The playback speed of the output media. The value must be between 0.5 and 1.0, inclusive. The default value is 1.0.
	//
	// > This parameter specifies the default playback speed of the output file as a ratio of the source file\\"s speed. It does not perform speed-up transcoding.
	//
	// example:
	//
	// 1.0
	Speed *float32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// If true, removes metadata from the output file. The default is false.
	StripMetadata *bool `json:"StripMetadata,omitempty" xml:"StripMetadata,omitempty"`
	// The subtitle processing parameters.
	//
	// 	Notice: You must use the `Subtitle.ExtractSubtitle` parameter to package subtitle streams. The `URI` in `Subtitle.ExtractSubtitle` must be in the same directory as or a subdirectory of `TargetGroups.URI`. The `Format` in `Subtitle.ExtractSubtitle` must be `vtt`. You only need to configure this parameter in one `Target` to package all subtitle streams.
	Subtitle *TargetSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// The OSS URI of the output HLS media playlist file for the subtask.
	//
	// 	Notice: This URI must be in the same directory as or a subdirectory of `TargetGroups.URI`.
	//
	// example:
	//
	// oss://test-bucket/test-target-object.mp4
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The video processing parameters.
	//
	// 	Notice: If this parameter is left empty, the first video stream, if it exists, is copied directly to the output file.
	Video *TargetVideo `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetGroupsTargets) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetGroupsTargets) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) GetAudio() *TargetAudio {
	return s.Audio
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) GetContainer() *string {
	return s.Container
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) GetSegment() *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment {
	return s.Segment
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) GetSpeed() *float32 {
	return s.Speed
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) GetStripMetadata() *bool {
	return s.StripMetadata
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) GetSubtitle() *TargetSubtitle {
	return s.Subtitle
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) GetURI() *string {
	return s.URI
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) GetVideo() *TargetVideo {
	return s.Video
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) SetAudio(v *TargetAudio) *CreateMediaConvertTaskRequestTargetGroupsTargets {
	s.Audio = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) SetContainer(v string) *CreateMediaConvertTaskRequestTargetGroupsTargets {
	s.Container = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) SetSegment(v *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) *CreateMediaConvertTaskRequestTargetGroupsTargets {
	s.Segment = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) SetSpeed(v float32) *CreateMediaConvertTaskRequestTargetGroupsTargets {
	s.Speed = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) SetStripMetadata(v bool) *CreateMediaConvertTaskRequestTargetGroupsTargets {
	s.StripMetadata = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) SetSubtitle(v *TargetSubtitle) *CreateMediaConvertTaskRequestTargetGroupsTargets {
	s.Subtitle = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) SetURI(v string) *CreateMediaConvertTaskRequestTargetGroupsTargets {
	s.URI = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) SetVideo(v *TargetVideo) *CreateMediaConvertTaskRequestTargetGroupsTargets {
	s.Video = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargets) Validate() error {
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
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

type CreateMediaConvertTaskRequestTargetGroupsTargetsSegment struct {
	// The duration of each segment, in seconds.
	//
	// example:
	//
	// 30
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The media packaging format. Only `hls` is supported.
	//
	// example:
	//
	// hls
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The starting sequence number for segments. The default is 0.
	//
	// example:
	//
	// 5
	StartNumber *int32 `json:"StartNumber,omitempty" xml:"StartNumber,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) GetDuration() *float64 {
	return s.Duration
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) GetFormat() *string {
	return s.Format
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) GetStartNumber() *int32 {
	return s.StartNumber
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) SetDuration(v float64) *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment {
	s.Duration = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) SetFormat(v string) *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment {
	s.Format = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) SetStartNumber(v int32) *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment {
	s.StartNumber = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetGroupsTargetsSegment) Validate() error {
	return dara.Validate(s)
}

type CreateMediaConvertTaskRequestTargets struct {
	// Settings for retaining attached pictures.
	//
	// 	Notice: Retaining attached pictures is supported only when the `Container` parameter is set to `mp4` or `mkv`.
	AttachedPicture *CreateMediaConvertTaskRequestTargetsAttachedPicture `json:"AttachedPicture,omitempty" xml:"AttachedPicture,omitempty" type:"Struct"`
	// The audio processing parameters.
	//
	// 	Notice: If this parameter is left empty, the first audio stream, if it exists, is copied directly to the output file.
	Audio *TargetAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The media container type. Valid container types include:
	//
	// - Audio/video containers: mp4, mkv, mov, asf, avi, mxf, ts, flv
	//
	// - Audio-only containers: mp3, aac, flac, oga, ac3, opus
	//
	//
	//   	Notice:
	//
	//   The `Container` and `URI` parameters must be set together. To perform only subtitle extraction, frame capture, sprite generation, or animated image generation, leave both `Container` and `URI` empty. In this case, parameters such as `Segment`, `Video`, `Audio`, and `Speed` are ignored.
	//
	// example:
	//
	// mp4
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// Settings for retaining data streams.
	//
	// 	Notice: Retaining data streams is supported only when the `Container` parameter is set to `mp4`.
	Data *CreateMediaConvertTaskRequestTargetsData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The parameters for frame capture, sprite generation, and animated image generation.
	Image *TargetImage `json:"Image,omitempty" xml:"Image,omitempty"`
	// Settings for media segmentation.
	Segment *CreateMediaConvertTaskRequestTargetsSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// The playback speed of the output media. The value must be between 0.5 and 1.0, inclusive. The default value is 1.0.
	//
	// > This parameter specifies the default playback speed of the output file as a ratio of the source file\\"s speed. It does not perform speed-up transcoding.
	//
	// example:
	//
	// 1.0
	Speed *float32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// If true, removes metadata such as `title` and `album` from the media file. The default is false.
	StripMetadata *bool `json:"StripMetadata,omitempty" xml:"StripMetadata,omitempty"`
	// The subtitle processing parameters.
	//
	// 	Notice: If this parameter is left empty, the first subtitle stream, if it exists, is copied directly to the output file.
	Subtitle *TargetSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// The OSS URI of the output file for media transcoding.
	//
	// The URI must be in the `oss://${Bucket}/${Object}` format. In this format, `${Bucket}` is the name of the OSS bucket, which must be in the same region as the project, and `${Object}` is the full path to the object, including the file extension.
	//
	// - If the **URI*	- has a file extension, all output media files are saved to this **URI**. If multiple files are generated, they will overwrite each other.
	//
	// - If the **URI*	- does not have a file extension, the final output URI is generated based on the **URI**, **Container**, and **Segment*	- parameters. For example, if the **URI*	- is `oss://examplebucket/outputVideo`:
	//
	//   - If **Container*	- is `mp4` and **Segment*	- is empty, the OSS URI of the generated media file is `oss://examplebucket/outputVideo.mp4`.
	//
	//   - If **Container*	- is `ts` and **Format*	- in **Segment*	- is `hls`, the process generates an m3u8 file with the OSS URI `oss://examplebucket/outputVideo.m3u8` and multiple TS files prefixed with `oss://examplebucket/outputVideo`.
	//
	// example:
	//
	// oss://test-bucket/test-target-object.mp4
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The video processing parameters.
	//
	// 	Notice: If this parameter is left empty, the first video stream, if it exists, is copied directly to the output file.
	Video *TargetVideo `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargets) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargets) GetAttachedPicture() *CreateMediaConvertTaskRequestTargetsAttachedPicture {
	return s.AttachedPicture
}

func (s *CreateMediaConvertTaskRequestTargets) GetAudio() *TargetAudio {
	return s.Audio
}

func (s *CreateMediaConvertTaskRequestTargets) GetContainer() *string {
	return s.Container
}

func (s *CreateMediaConvertTaskRequestTargets) GetData() *CreateMediaConvertTaskRequestTargetsData {
	return s.Data
}

func (s *CreateMediaConvertTaskRequestTargets) GetImage() *TargetImage {
	return s.Image
}

func (s *CreateMediaConvertTaskRequestTargets) GetSegment() *CreateMediaConvertTaskRequestTargetsSegment {
	return s.Segment
}

func (s *CreateMediaConvertTaskRequestTargets) GetSpeed() *float32 {
	return s.Speed
}

func (s *CreateMediaConvertTaskRequestTargets) GetStripMetadata() *bool {
	return s.StripMetadata
}

func (s *CreateMediaConvertTaskRequestTargets) GetSubtitle() *TargetSubtitle {
	return s.Subtitle
}

func (s *CreateMediaConvertTaskRequestTargets) GetURI() *string {
	return s.URI
}

func (s *CreateMediaConvertTaskRequestTargets) GetVideo() *TargetVideo {
	return s.Video
}

func (s *CreateMediaConvertTaskRequestTargets) SetAttachedPicture(v *CreateMediaConvertTaskRequestTargetsAttachedPicture) *CreateMediaConvertTaskRequestTargets {
	s.AttachedPicture = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetAudio(v *TargetAudio) *CreateMediaConvertTaskRequestTargets {
	s.Audio = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetContainer(v string) *CreateMediaConvertTaskRequestTargets {
	s.Container = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetData(v *CreateMediaConvertTaskRequestTargetsData) *CreateMediaConvertTaskRequestTargets {
	s.Data = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetImage(v *TargetImage) *CreateMediaConvertTaskRequestTargets {
	s.Image = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetSegment(v *CreateMediaConvertTaskRequestTargetsSegment) *CreateMediaConvertTaskRequestTargets {
	s.Segment = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetSpeed(v float32) *CreateMediaConvertTaskRequestTargets {
	s.Speed = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetStripMetadata(v bool) *CreateMediaConvertTaskRequestTargets {
	s.StripMetadata = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetSubtitle(v *TargetSubtitle) *CreateMediaConvertTaskRequestTargets {
	s.Subtitle = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetURI(v string) *CreateMediaConvertTaskRequestTargets {
	s.URI = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetVideo(v *TargetVideo) *CreateMediaConvertTaskRequestTargets {
	s.Video = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) Validate() error {
	if s.AttachedPicture != nil {
		if err := s.AttachedPicture.Validate(); err != nil {
			return err
		}
	}
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	if s.Image != nil {
		if err := s.Image.Validate(); err != nil {
			return err
		}
	}
	if s.Segment != nil {
		if err := s.Segment.Validate(); err != nil {
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

type CreateMediaConvertTaskRequestTargetsAttachedPicture struct {
	// A list of indexes of the attached pictures in the source file to process. An empty list (default) indicates that no attached pictures are retained. An index of -1 indicates that all attached pictures are retained.
	//
	// - Example: `[0,1]` processes the attached pictures with index 0 and 1; `[1]` processes the attached picture with index 1; `[-1]` processes all attached pictures.
	//
	// > If a specified index does not correspond to an existing attached picture, it is ignored.
	Stream []*int32 `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
}

func (s CreateMediaConvertTaskRequestTargetsAttachedPicture) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsAttachedPicture) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsAttachedPicture) GetStream() []*int32 {
	return s.Stream
}

func (s *CreateMediaConvertTaskRequestTargetsAttachedPicture) SetStream(v []*int32) *CreateMediaConvertTaskRequestTargetsAttachedPicture {
	s.Stream = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsAttachedPicture) Validate() error {
	return dara.Validate(s)
}

type CreateMediaConvertTaskRequestTargetsData struct {
	// A list of indexes of the data streams in the source file to process. An empty list (default) indicates that no data streams are retained. An index of -1 indicates that all data streams are retained.
	//
	// - Example: `[0,1]` processes the data streams with index 0 and 1; `[1]` processes the data stream with index 1; `[-1]` processes all data streams.
	//
	// > If a specified index does not correspond to an existing data stream, it is ignored.
	Stream []*int32 `json:"Stream,omitempty" xml:"Stream,omitempty" type:"Repeated"`
}

func (s CreateMediaConvertTaskRequestTargetsData) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsData) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsData) GetStream() []*int32 {
	return s.Stream
}

func (s *CreateMediaConvertTaskRequestTargetsData) SetStream(v []*int32) *CreateMediaConvertTaskRequestTargetsData {
	s.Stream = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsData) Validate() error {
	return dara.Validate(s)
}

type CreateMediaConvertTaskRequestTargetsSegment struct {
	// The duration of each segment, in seconds.
	//
	// example:
	//
	// 30
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The segmentation method. Valid values include:
	//
	// - hls
	//
	// - dash
	//
	// example:
	//
	// hls
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The starting sequence number. This parameter is supported only for HLS. The default value is 0.
	//
	// example:
	//
	// 5
	StartNumber *int32 `json:"StartNumber,omitempty" xml:"StartNumber,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargetsSegment) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargetsSegment) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) GetDuration() *float64 {
	return s.Duration
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) GetFormat() *string {
	return s.Format
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) GetStartNumber() *int32 {
	return s.StartNumber
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) SetDuration(v float64) *CreateMediaConvertTaskRequestTargetsSegment {
	s.Duration = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) SetFormat(v string) *CreateMediaConvertTaskRequestTargetsSegment {
	s.Format = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) SetStartNumber(v int32) *CreateMediaConvertTaskRequestTargetsSegment {
	s.StartNumber = &v
	return s
}

func (s *CreateMediaConvertTaskRequestTargetsSegment) Validate() error {
	return dara.Validate(s)
}
