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
	// The sequence number of the main media file in the concatenation list of media files. The main media file provides the default transcoding settings, such as the resolution and the frame rate, for videos and audios. Default value: `0`. A value of `0` specifies that the main media file is aligned with the first media file in the concatenation list.
	AlignmentIndex *int32 `json:"AlignmentIndex,omitempty" xml:"AlignmentIndex,omitempty"`
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The notification settings. For more information, see "Notification". For information about the asynchronous notification format, see [Asynchronous notification format](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The source media files. If multiple files exist at the same time, the Concat feature is enabled. The video files are concatenated in the order of their URI inputs.
	//
	// This parameter is required.
	Sources []*CreateMediaConvertTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The custom tags. You can search for or filter asynchronous tasks by custom tag.
	//
	// example:
	//
	// {"test":"val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The media processing tasks. You can specify multiple values for this parameter.
	//
	// This parameter is required.
	Targets []*CreateMediaConvertTaskRequestTargets `json:"Targets,omitempty" xml:"Targets,omitempty" type:"Repeated"`
	// The custom information, which is returned as asynchronous notifications to facilitate notification management in your system. The maximum information length is 2,048 bytes.
	//
	// example:
	//
	// {"ID": "user1","Name": "test-user1","Avatar": "http://example.com?id=user1"}
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
	return dara.Validate(s)
}

type CreateMediaConvertTaskRequestSources struct {
	AlignMode    *string `json:"AlignMode,omitempty" xml:"AlignMode,omitempty"`
	Attached     *bool   `json:"Attached,omitempty" xml:"Attached,omitempty"`
	DisableAudio *bool   `json:"DisableAudio,omitempty" xml:"DisableAudio,omitempty"`
	DisableVideo *bool   `json:"DisableVideo,omitempty" xml:"DisableVideo,omitempty"`
	// The transcoding duration of the media. Unit: seconds. Default value: 0. A value of 0 specifies that the transcoding duration lasts until the end of the video.
	//
	// example:
	//
	// 0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the media transcoding task. Unit: seconds. Valid values:
	//
	// 	- 0 (default): starts transcoding when the media starts playing.
	//
	// 	- n: starts transcoding n seconds after the media starts playing. n must be greater than 0.
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The subtitles. By default, this parameter is left empty.
	Subtitles []*CreateMediaConvertTaskRequestSourcesSubtitles `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The URI of the Object Storage Service (OSS) bucket. Specify the value in the `oss://${Bucket}/${Object}` format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region with the current project. `${Object}` specifies the complete path to the file whose name contains an extension.
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
	return dara.Validate(s)
}

type CreateMediaConvertTaskRequestSourcesSubtitles struct {
	// The subtitle language. If you specify this parameter, comply with the ISO 639-2 standard. This parameter is left empty by default.
	//
	// example:
	//
	// eng
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The time offset of the subtitle. Unit: seconds. Default value: 0.
	//
	// example:
	//
	// 10.5
	TimeOffset *float64 `json:"TimeOffset,omitempty" xml:"TimeOffset,omitempty"`
	// The URI of the Object Storage Service (OSS) bucket. Specify the value in the `oss://${Bucket}/${Object}` format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region with the current project. `${Object}` specifies the complete path to the file whose name contains an extension. The following subtitle formats are supported: srt, vtt, mov_text, ass, dvd_sub, and pgs.
	//
	// example:
	//
	// oss://test-bucket/subtitles
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
	// The audio processing settings.
	//
	// >  If you leave Audio empty and the first audio stream exists, the first audio stream is directly copied to the output file.
	Audio *TargetAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The type of the media container.
	//
	// 	- Valid values for audio and video containers: mp4, mkv, mov, asf, avi, mxf, ts, and flv.
	//
	// 	- Valid values only for audio containers: mp3, aac, flac, oga, ac3, and opus.
	//
	//     **
	//
	//     **Note*	- Specify Container and URI at the same time. If you want to extract subtitles, capture frames, capture image sprites, or rotate media images, set Container and URI to null. In this case, Segment, Video, Audio, and Speed do not take effect.
	//
	// example:
	//
	// mp4
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The frame capturing, sprite capturing, and media rotation settings.
	Image *TargetImage `json:"Image,omitempty" xml:"Image,omitempty"`
	// The media segmentation settings. By default, no segmentation is performed.
	Segment *CreateMediaConvertTaskRequestTargetsSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// The playback speed of the media. Valid values: 0.5 to 2. Default value: 1.0.
	//
	// >  This parameter specifies the ratio of the non-regular playback speed of the transcoded media file to the default playback speed of the source media file.
	//
	// example:
	//
	// 1.0
	Speed *float32 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// Specifies whether to remove the metadata, such as `title` and `album`, from the media file. Default value: false.
	StripMetadata *bool `json:"StripMetadata,omitempty" xml:"StripMetadata,omitempty"`
	// The subtitle processing settings.
	//
	// >  If you leave Subtitle empty and the first subtitle stream exists, the first subtitle stream is directly copied to the output file.
	Subtitle *TargetSubtitle `json:"Subtitle,omitempty" xml:"Subtitle,omitempty"`
	// The URI of the OSS bucket in which you want to store the media transcoding output file.
	//
	// Specify the value in the `oss://${Bucket}/${Object}` format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region with the current project. `${Object}` specifies the complete path to the file whose name contains an extension.
	//
	// 	- If the value of **URI*	- contains an extension, the endpoint of the OSS bucket matches the URI. If multiple media transcoding output files exist, the endpoints of the correspodning OSS buckets may be overwritten.****
	//
	// 	- If the value of **URI*	- does not contain an extension, the endpoint of the OSS bucket consists of the following parameters: **URI**, **Container**, and **Segment**. In the following examples, the value of **URI*	- is `oss://examplebucket/outputVideo`.
	//
	//     	- If the value of **Container*	- is `mp4` and the value of **Segment*	- is null, the endpoint of the OSS bucket is `oss://examplebucket/outputVideo.mp4`.
	//
	//     	- If the value of **Container*	- is `ts` and the value of **Format*	- in **Segment*	- is `hls`, the endpoint of the OSS bucket is `oss://examplebucket/outputVideo.m3u8`. In addition, multiple ts files prefixed with `oss://examplebucket/outputVideo` are generated.
	//
	// example:
	//
	// oss://test-bucket/targets
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The video processing settings.
	//
	// >  If you leave Video empty and the first video stream exists, the first video stream is directly copied to the output file.
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
	return dara.Validate(s)
}

type CreateMediaConvertTaskRequestTargetsSegment struct {
	// The duration of the segment. Unit: seconds.
	//
	// example:
	//
	// 30
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The media segmentation mode. Valid values:
	//
	// 	- hls
	//
	// 	- dash
	//
	// example:
	//
	// hls
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The start sequence number. You can specify this parameter only if you set Format to hls. Default value: 0.
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
