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
	SetTargets(v []*CreateMediaConvertTaskRequestTargets) *CreateMediaConvertTaskRequest
	GetTargets() []*CreateMediaConvertTaskRequestTargets
	SetUserData(v string) *CreateMediaConvertTaskRequest
	GetUserData() *string
}

type CreateMediaConvertTaskRequest struct {
	// When performing media concatenation, the index of the primary media file (which provides the default transcoding parameters for `Video` and `Audio`, including resolution, frame rate, etc.) in the concatenation list. The default value is 0 (aligning with the first media file in the concatenation list).
	//
	// example:
	//
	// 0
	AlignmentIndex *int32 `json:"AlignmentIndex,omitempty" xml:"AlignmentIndex,omitempty"`
	// **If there are no special requirements, please leave this blank.**
	//
	// Chain authorization configuration. For more information, see [Using Chain Authorization to Access Other Entity Resources](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// Notification configuration. For details, click Notification. The format of asynchronous notification messages can be found in [Asynchronous Notification Message Format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. For how to obtain it, see [Creating a Project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// A list of media files. If the list contains more than one element, it indicates that the Concat (concatenation) function is enabled. The Concat order follows the sequence of the input video file URIs.
	//
	// This parameter is required.
	Sources []*CreateMediaConvertTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// Custom tags used for searching and filtering asynchronous tasks.
	//
	// example:
	//
	// {"test":"val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// List of media processing tasks, supporting multiple task configurations.
	//
	// This parameter is required.
	Targets []*CreateMediaConvertTaskRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// User-defined information that will be returned in asynchronous message notifications, used for convenient association and processing within your system. The maximum length is 2048 bytes.
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
	// The alignment strategy for adding audio and video streams, with the following value range:
	//
	// - false (default): No alignment.
	//
	// - loop: Loop the audio and video content to align.
	//
	// - pad: Align by padding silent frames and black video frames.
	//
	// > - Only valid when the Attached parameter is true.
	//
	// example:
	//
	// false
	AlignMode *string `json:"AlignMode,omitempty" xml:"AlignMode,omitempty"`
	// Add the current source media file as a synchronized audio or video stream to the output media file, with a default value of false.
	//
	// > - The AlignmentIndex parameter pointing to the Attached parameter of the Source cannot be true.
	//
	// example:
	//
	// false
	Attached *bool `json:"Attached,omitempty" xml:"Attached,omitempty"`
	// Whether to disable the audio in the source media file. The value range is as follows:
	//
	// - true: Disable.
	//
	// - false (default): Do not disable.
	//
	// example:
	//
	// false
	DisableAudio *bool `json:"DisableAudio,omitempty" xml:"DisableAudio,omitempty"`
	// Whether to disable the video in the source media file. The value range is as follows:
	//
	// - true: Disable.
	//
	// - false (default): Do not disable.
	//
	// example:
	//
	// false
	DisableVideo *bool `json:"DisableVideo,omitempty" xml:"DisableVideo,omitempty"`
	// The duration of media transcoding, in seconds. The default value is 0, indicating until the end of the video.
	//
	// example:
	//
	// 0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time for media transcoding, in seconds. The value range is as follows:
	//
	// - 0 (default): Start transcoding from the beginning of the media.
	//
	// - n (greater than 0): Start transcoding n seconds after the beginning of the media.
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// A list of subtitles to add, which is empty by default.
	Subtitles []*CreateMediaConvertTaskRequestSourcesSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The OSS address rule is `oss://${Bucket}/${Object}`, where `${Bucket}` is the name of the OSS Bucket in the same region (Region) as the current project, and `${Object}` is the complete path of the file including the file extension.
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
	// The language of the subtitle, referenced by ISO 639-2, with a default value of empty.
	//
	// example:
	//
	// eng
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The subtitle delay time, in seconds, with a default value of 0.
	//
	// example:
	//
	// 10.5
	TimeOffset *float64 `json:"TimeOffset,omitempty" xml:"TimeOffset,omitempty"`
	// The OSS address rule is `oss://${Bucket}/${Object}`, where `${Bucket}` is the name of the OSS Bucket in the same region (Region) as the current project, and `${Object}` is the complete path of the file including the file extension.
	//
	// Supported subtitle formats include: srt, vtt, mov_text, ass, dvd_sub, pgs.
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

type CreateMediaConvertTaskRequestTargets struct {
	// Audio processing parameter configuration.
	//
	// 	Notice: If Audio is null, the first audio stream (if present) will be directly copied to the output file.</notice>
	Audio *TargetAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// Media container type. Available container types are as follows:
	//
	// - Audio and video containers: mp4, mkv, mov, asf, avi, mxf, ts, flv
	//
	// - Audio containers: mp3, aac, flac, oga, ac3, opus
	//
	// 	Notice: Both Container and URI parameters need to be set. If only subtitle extraction, frame capture, sprite image capture, or media-to-gif conversion is performed, both Container and URI should be set to null, making the Segment, Video, Audio, and Speed parameters meaningless.</notice>
	//
	// example:
	//
	// mp4
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// Configuration for frame capture, sprite image capture, and media to animated image conversion.
	Image *TargetImage `json:"Image,omitempty" xml:"Image,omitempty"`
	// Media segment settings, no segmentation by default.
	Segment *CreateMediaConvertTaskRequestTargetsSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// Media playback speed setting, with a value range of [0.5,1.0], default is 1.0.
	//
	// > The ratio of the playback speed of the transcoded media file to the original media file, not a speed-up transcoding.
	//
	// example:
	//
	// 1.0
	Speed *float32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// Removes metadata from the media file, such as `title`, `album`, etc. The default value is false.
	StripMetadata *bool `json:"StripMetadata,omitempty" xml:"StripMetadata,omitempty"`
	// Subtitle processing parameter configuration.
	//
	// 	Notice: If Subtitle is null, the first subtitle stream (if present) will be directly copied to the output file.</notice>
	Subtitle *TargetSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// OSS address for the output file of media transcoding.
	//
	// The OSS address rule is `oss://${Bucket}/${Object}`, where `${Bucket}` is the name of the OSS Bucket in the same region (Region) as the current project, and `${Object}` is the complete path of the file including the file extension.
	//
	// - When **URI*	- has an extension, the OSS address for the transcoded media file will be **URI**. If there are multiple output files, they may overwrite each other.
	//
	// - When **URI*	- does not have an extension, the OSS address for the transcoded media file is determined by the **URI**, **Container**, and **Segment*	- parameters. For example, if **URI*	- is `oss://examplebucket/outputVideo`:
	//
	//    -  When **Container*	- is `mp4` and **Segment*	- is empty, the generated media file\\"s OSS address will be `oss://examplebucket/outputVideo.mp4`.
	//
	//    -  When **Container*	- is `ts` and **Segment**\\"s **Format*	- is `hls`, it will generate an m3u8 file with the OSS address `oss://examplebucket/outputVideo.m3u8` and multiple ts files with the prefix `oss://examplebucket/outputVideo`.
	//
	// example:
	//
	// oss://test-bucket/test-target-object.mp4
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// Video processing parameter configuration.
	//
	// 	Notice: If Video is null, the first video stream (if present) will be directly copied to the output file.</notice>
	Video *TargetVideo `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s CreateMediaConvertTaskRequestTargets) String() string {
	return dara.Prettify(s)
}

func (s CreateMediaConvertTaskRequestTargets) GoString() string {
	return s.String()
}

func (s *CreateMediaConvertTaskRequestTargets) GetAudio() *TargetAudio {
	return s.Audio
}

func (s *CreateMediaConvertTaskRequestTargets) GetContainer() *string {
	return s.Container
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

func (s *CreateMediaConvertTaskRequestTargets) SetAudio(v *TargetAudio) *CreateMediaConvertTaskRequestTargets {
	s.Audio = v
	return s
}

func (s *CreateMediaConvertTaskRequestTargets) SetContainer(v string) *CreateMediaConvertTaskRequestTargets {
	s.Container = &v
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
	if s.Audio != nil {
		if err := s.Audio.Validate(); err != nil {
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

type CreateMediaConvertTaskRequestTargetsSegment struct {
	// Segment length. Unit: seconds.
	//
	// example:
	//
	// 30
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Media slicing method. The value range is as follows:
	//
	// - hls
	//
	// - dash
	//
	// example:
	//
	// hls
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// Starting sequence number, supported only for hls, default is 0.
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
