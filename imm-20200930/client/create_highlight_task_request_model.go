// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHighlightTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfig(v *CredentialConfig) *CreateHighlightTaskRequest
	GetCredentialConfig() *CredentialConfig
	SetEdit(v *CreateHighlightTaskRequestEdit) *CreateHighlightTaskRequest
	GetEdit() *CreateHighlightTaskRequestEdit
	SetHighlight(v *CreateHighlightTaskRequestHighlight) *CreateHighlightTaskRequest
	GetHighlight() *CreateHighlightTaskRequestHighlight
	SetMode(v string) *CreateHighlightTaskRequest
	GetMode() *string
	SetNotification(v *Notification) *CreateHighlightTaskRequest
	GetNotification() *Notification
	SetOutput(v *CreateHighlightTaskRequestOutput) *CreateHighlightTaskRequest
	GetOutput() *CreateHighlightTaskRequestOutput
	SetProjectName(v string) *CreateHighlightTaskRequest
	GetProjectName() *string
	SetSources(v []*CreateHighlightTaskRequestSources) *CreateHighlightTaskRequest
	GetSources() []*CreateHighlightTaskRequestSources
	SetTags(v map[string]interface{}) *CreateHighlightTaskRequest
	GetTags() map[string]interface{}
	SetType(v string) *CreateHighlightTaskRequest
	GetType() *string
	SetUserData(v string) *CreateHighlightTaskRequest
	GetUserData() *string
}

type CreateHighlightTaskRequest struct {
	// The chained authorization configuration. **Leave this parameter empty unless otherwise required.**
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The editing configuration.
	Edit *CreateHighlightTaskRequestEdit `json:"Edit,omitempty" xml:"Edit,omitempty" type:"Struct"`
	// The highlight configuration.
	Highlight *CreateHighlightTaskRequestHighlight `json:"Highlight,omitempty" xml:"Highlight,omitempty" type:"Struct"`
	// The highlight recognition mode. Valid values:
	//
	// - Scene: scene and frame recognition
	//
	// - Average: average clip recognition
	//
	// Default value: Average.
	//
	// example:
	//
	// Average
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The message notification configuration. For more information, see Notification. For the format of asynchronous notification messages, see [Asynchronous notification message format](https://www.alibabacloud.com/help/en/imm/developer-reference/asynchronous-notification-message-examples).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The output configuration.
	//
	// This parameter is required.
	Output *CreateHighlightTaskRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The list of media resources to be processed. You can specify up to 10 videos.
	//
	// This parameter is required.
	Sources []*CreateHighlightTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The custom tags used to search and filter asynchronous tasks.
	//
	// example:
	//
	// {"test":"val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The processing type. Valid values:
	//
	// - Retrieval: highlight extraction
	//
	// - Concat: video composition
	//
	// - Compose: one-click video creation
	//
	// This parameter is required.
	//
	// example:
	//
	// Retrieval
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The custom user data, which is returned in asynchronous message notifications.
	//
	// example:
	//
	// {"ID": "testuid","Name": "test-user","Avatar": "http://test.com/testuid"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateHighlightTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *CreateHighlightTaskRequest) GetEdit() *CreateHighlightTaskRequestEdit {
	return s.Edit
}

func (s *CreateHighlightTaskRequest) GetHighlight() *CreateHighlightTaskRequestHighlight {
	return s.Highlight
}

func (s *CreateHighlightTaskRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateHighlightTaskRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *CreateHighlightTaskRequest) GetOutput() *CreateHighlightTaskRequestOutput {
	return s.Output
}

func (s *CreateHighlightTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateHighlightTaskRequest) GetSources() []*CreateHighlightTaskRequestSources {
	return s.Sources
}

func (s *CreateHighlightTaskRequest) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *CreateHighlightTaskRequest) GetType() *string {
	return s.Type
}

func (s *CreateHighlightTaskRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateHighlightTaskRequest) SetCredentialConfig(v *CredentialConfig) *CreateHighlightTaskRequest {
	s.CredentialConfig = v
	return s
}

func (s *CreateHighlightTaskRequest) SetEdit(v *CreateHighlightTaskRequestEdit) *CreateHighlightTaskRequest {
	s.Edit = v
	return s
}

func (s *CreateHighlightTaskRequest) SetHighlight(v *CreateHighlightTaskRequestHighlight) *CreateHighlightTaskRequest {
	s.Highlight = v
	return s
}

func (s *CreateHighlightTaskRequest) SetMode(v string) *CreateHighlightTaskRequest {
	s.Mode = &v
	return s
}

func (s *CreateHighlightTaskRequest) SetNotification(v *Notification) *CreateHighlightTaskRequest {
	s.Notification = v
	return s
}

func (s *CreateHighlightTaskRequest) SetOutput(v *CreateHighlightTaskRequestOutput) *CreateHighlightTaskRequest {
	s.Output = v
	return s
}

func (s *CreateHighlightTaskRequest) SetProjectName(v string) *CreateHighlightTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateHighlightTaskRequest) SetSources(v []*CreateHighlightTaskRequestSources) *CreateHighlightTaskRequest {
	s.Sources = v
	return s
}

func (s *CreateHighlightTaskRequest) SetTags(v map[string]interface{}) *CreateHighlightTaskRequest {
	s.Tags = v
	return s
}

func (s *CreateHighlightTaskRequest) SetType(v string) *CreateHighlightTaskRequest {
	s.Type = &v
	return s
}

func (s *CreateHighlightTaskRequest) SetUserData(v string) *CreateHighlightTaskRequest {
	s.UserData = &v
	return s
}

func (s *CreateHighlightTaskRequest) Validate() error {
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Edit != nil {
		if err := s.Edit.Validate(); err != nil {
			return err
		}
	}
	if s.Highlight != nil {
		if err := s.Highlight.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.Output != nil {
		if err := s.Output.Validate(); err != nil {
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
	return nil
}

type CreateHighlightTaskRequestEdit struct {
	// The background music mode. Valid values:
	//
	// - Random: custom background music, randomly selected based on weights
	//
	// - Sequential: custom background music, applied in sequence
	//
	// - Closed: no background music
	//
	// Default value: Closed.
	//
	// example:
	//
	// Closed
	BackgroundMusicMode *string `json:"BackgroundMusicMode,omitempty" xml:"BackgroundMusicMode,omitempty"`
	// The background music. This parameter is valid only when BackgroundMusicMode is set to Random or Sequential. **The current maximum number of background music tracks is 1.**
	BackgroundMusics []*CreateHighlightTaskRequestEditBackgroundMusics `json:"BackgroundMusics,omitempty" xml:"BackgroundMusics,omitempty" type:"Repeated"`
	// The editing mode. Valid values:
	//
	// - Sequential: sequential mode
	//
	// This parameter is required.
	//
	// example:
	//
	// Sequential
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The transition mode. Valid values:
	//
	// - Auto: automatic transition
	//
	// - Random: custom transition, randomly selected based on weights
	//
	// - Sequential: custom transition, applied in sequence
	//
	// - Closed: no transition
	//
	// Default value: Closed.
	//
	// example:
	//
	// Closed
	TransitionMode *string `json:"TransitionMode,omitempty" xml:"TransitionMode,omitempty"`
	// The transition effects. This parameter is valid only when TransitionMode is set to Random or Sequential. You can specify up to 10 transition effects.
	Transitions []*CreateHighlightTaskRequestEditTransitions `json:"Transitions,omitempty" xml:"Transitions,omitempty" type:"Repeated"`
	// The visual effect mode. Valid values:
	//
	// - Auto: automatic visual effect
	//
	// - Random: custom visual effect, randomly selected based on weights
	//
	// - Sequential: custom visual effect, applied in sequence
	//
	// - Closed: no visual effect
	//
	// Default value: Closed.
	//
	// example:
	//
	// Closed
	VfxEffectMode *string `json:"VfxEffectMode,omitempty" xml:"VfxEffectMode,omitempty"`
	// The visual effects. This parameter is valid only when VfxEffectMode is set to Random or Sequential. You can specify up to 10 visual effects.
	VfxEffects []*CreateHighlightTaskRequestEditVfxEffects `json:"VfxEffects,omitempty" xml:"VfxEffects,omitempty" type:"Repeated"`
}

func (s CreateHighlightTaskRequestEdit) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskRequestEdit) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskRequestEdit) GetBackgroundMusicMode() *string {
	return s.BackgroundMusicMode
}

func (s *CreateHighlightTaskRequestEdit) GetBackgroundMusics() []*CreateHighlightTaskRequestEditBackgroundMusics {
	return s.BackgroundMusics
}

func (s *CreateHighlightTaskRequestEdit) GetMode() *string {
	return s.Mode
}

func (s *CreateHighlightTaskRequestEdit) GetTransitionMode() *string {
	return s.TransitionMode
}

func (s *CreateHighlightTaskRequestEdit) GetTransitions() []*CreateHighlightTaskRequestEditTransitions {
	return s.Transitions
}

func (s *CreateHighlightTaskRequestEdit) GetVfxEffectMode() *string {
	return s.VfxEffectMode
}

func (s *CreateHighlightTaskRequestEdit) GetVfxEffects() []*CreateHighlightTaskRequestEditVfxEffects {
	return s.VfxEffects
}

func (s *CreateHighlightTaskRequestEdit) SetBackgroundMusicMode(v string) *CreateHighlightTaskRequestEdit {
	s.BackgroundMusicMode = &v
	return s
}

func (s *CreateHighlightTaskRequestEdit) SetBackgroundMusics(v []*CreateHighlightTaskRequestEditBackgroundMusics) *CreateHighlightTaskRequestEdit {
	s.BackgroundMusics = v
	return s
}

func (s *CreateHighlightTaskRequestEdit) SetMode(v string) *CreateHighlightTaskRequestEdit {
	s.Mode = &v
	return s
}

func (s *CreateHighlightTaskRequestEdit) SetTransitionMode(v string) *CreateHighlightTaskRequestEdit {
	s.TransitionMode = &v
	return s
}

func (s *CreateHighlightTaskRequestEdit) SetTransitions(v []*CreateHighlightTaskRequestEditTransitions) *CreateHighlightTaskRequestEdit {
	s.Transitions = v
	return s
}

func (s *CreateHighlightTaskRequestEdit) SetVfxEffectMode(v string) *CreateHighlightTaskRequestEdit {
	s.VfxEffectMode = &v
	return s
}

func (s *CreateHighlightTaskRequestEdit) SetVfxEffects(v []*CreateHighlightTaskRequestEditVfxEffects) *CreateHighlightTaskRequestEdit {
	s.VfxEffects = v
	return s
}

func (s *CreateHighlightTaskRequestEdit) Validate() error {
	if s.BackgroundMusics != nil {
		for _, item := range s.BackgroundMusics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Transitions != nil {
		for _, item := range s.Transitions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VfxEffects != nil {
		for _, item := range s.VfxEffects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateHighlightTaskRequestEditBackgroundMusics struct {
	// The URI of the background music, which is an OSS URI. Only audio files are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object/test.mp3
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The volume of the background music. Valid values: [0, 10]. Default value: 0.2. A value of 1 indicates the original volume.
	//
	// example:
	//
	// 0.2
	Volume *float64 `json:"Volume,omitempty" xml:"Volume,omitempty"`
}

func (s CreateHighlightTaskRequestEditBackgroundMusics) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskRequestEditBackgroundMusics) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskRequestEditBackgroundMusics) GetURI() *string {
	return s.URI
}

func (s *CreateHighlightTaskRequestEditBackgroundMusics) GetVolume() *float64 {
	return s.Volume
}

func (s *CreateHighlightTaskRequestEditBackgroundMusics) SetURI(v string) *CreateHighlightTaskRequestEditBackgroundMusics {
	s.URI = &v
	return s
}

func (s *CreateHighlightTaskRequestEditBackgroundMusics) SetVolume(v float64) *CreateHighlightTaskRequestEditBackgroundMusics {
	s.Volume = &v
	return s
}

func (s *CreateHighlightTaskRequestEditBackgroundMusics) Validate() error {
	return dara.Validate(s)
}

type CreateHighlightTaskRequestEditTransitions struct {
	// The duration of the transition. Unit: seconds. If the transition duration is greater than the clip duration minus 1, the transition effect on the clip does not take effect. Valid values: [0, 5].
	//
	// example:
	//
	// 0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The transition effect. For more information, see [Transition effects](https://www.alibabacloud.com/help/en/imm/developer-reference/transition-effect).
	//
	// This parameter is required.
	//
	// example:
	//
	// fade
	Transition *string `json:"Transition,omitempty" xml:"Transition,omitempty"`
	// The weight of the transition. Valid values: [1, 100]. Default value: 50. This parameter is valid only when TransitionMode is set to Random.
	//
	// example:
	//
	// 50
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateHighlightTaskRequestEditTransitions) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskRequestEditTransitions) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskRequestEditTransitions) GetDuration() *float64 {
	return s.Duration
}

func (s *CreateHighlightTaskRequestEditTransitions) GetTransition() *string {
	return s.Transition
}

func (s *CreateHighlightTaskRequestEditTransitions) GetWeight() *int64 {
	return s.Weight
}

func (s *CreateHighlightTaskRequestEditTransitions) SetDuration(v float64) *CreateHighlightTaskRequestEditTransitions {
	s.Duration = &v
	return s
}

func (s *CreateHighlightTaskRequestEditTransitions) SetTransition(v string) *CreateHighlightTaskRequestEditTransitions {
	s.Transition = &v
	return s
}

func (s *CreateHighlightTaskRequestEditTransitions) SetWeight(v int64) *CreateHighlightTaskRequestEditTransitions {
	s.Weight = &v
	return s
}

func (s *CreateHighlightTaskRequestEditTransitions) Validate() error {
	return dara.Validate(s)
}

type CreateHighlightTaskRequestEditVfxEffects struct {
	// The visual effect. For more information, see [Visual effects](https://www.alibabacloud.com/help/en/imm/developer-reference/effects).
	//
	// This parameter is required.
	//
	// example:
	//
	// letterboxed
	VfxEffect *string `json:"VfxEffect,omitempty" xml:"VfxEffect,omitempty"`
	// The weight of the visual effect. Valid values: [1, 100]. Default value: 50. This parameter is valid only when VfxEffectMode is set to Random.
	//
	// example:
	//
	// 50
	Weight *int64 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s CreateHighlightTaskRequestEditVfxEffects) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskRequestEditVfxEffects) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskRequestEditVfxEffects) GetVfxEffect() *string {
	return s.VfxEffect
}

func (s *CreateHighlightTaskRequestEditVfxEffects) GetWeight() *int64 {
	return s.Weight
}

func (s *CreateHighlightTaskRequestEditVfxEffects) SetVfxEffect(v string) *CreateHighlightTaskRequestEditVfxEffects {
	s.VfxEffect = &v
	return s
}

func (s *CreateHighlightTaskRequestEditVfxEffects) SetWeight(v int64) *CreateHighlightTaskRequestEditVfxEffects {
	s.Weight = &v
	return s
}

func (s *CreateHighlightTaskRequestEditVfxEffects) Validate() error {
	return dara.Validate(s)
}

type CreateHighlightTaskRequestHighlight struct {
	// The highlight content. Valid values:
	//
	// - Pet
	//
	// - Person
	//
	// - Sports
	//
	// - Meeting
	//
	// The value cannot exceed 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// character
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
}

func (s CreateHighlightTaskRequestHighlight) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskRequestHighlight) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskRequestHighlight) GetContent() *string {
	return s.Content
}

func (s *CreateHighlightTaskRequestHighlight) SetContent(v string) *CreateHighlightTaskRequestHighlight {
	s.Content = &v
	return s
}

func (s *CreateHighlightTaskRequestHighlight) Validate() error {
	return dara.Validate(s)
}

type CreateHighlightTaskRequestOutput struct {
	// The audio processing parameter settings. 	Notice: If Audio is empty, the first audio stream (if any) is directly copied to the output file.</notice>
	Audio *TargetAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// The media container type. This parameter is required when Type is set to Concat or Compose. Valid values:
	//
	// - Audio and video containers: mp4, mkv, mov, asf, avi, mxf, ts, and flv
	//
	// 	Notice: You must specify both Container and URI.</notice>
	//
	// example:
	//
	// mp4
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// The maximum duration of the edited video. Unit: seconds.
	//
	// example:
	//
	// 10.0
	MaxDuration *float64 `json:"MaxDuration,omitempty" xml:"MaxDuration,omitempty"`
	// The media segmentation settings. By default, segmentation is not performed.
	Segment *CreateHighlightTaskRequestOutputSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// The playback speed multiplier for the media. Valid values: [0.5, 1.0]. Default value: 1.0.
	//
	// > The ratio of the default playback speed of the transcoded media file to that of the source media file. This is not speed-adjusted transcoding.
	//
	// example:
	//
	// 1.0
	Speed *float64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// The target duration of the video. Unit: seconds.
	//
	// example:
	//
	// 10.0
	TargetDuration *float64 `json:"TargetDuration,omitempty" xml:"TargetDuration,omitempty"`
	// The URI of the output file.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-target-object.mp4
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The video processing parameter settings. 	Notice: If Video is empty, the first video stream (if any) is directly copied to the output file.</notice>
	Video *TargetVideo `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s CreateHighlightTaskRequestOutput) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskRequestOutput) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskRequestOutput) GetAudio() *TargetAudio {
	return s.Audio
}

func (s *CreateHighlightTaskRequestOutput) GetContainer() *string {
	return s.Container
}

func (s *CreateHighlightTaskRequestOutput) GetMaxDuration() *float64 {
	return s.MaxDuration
}

func (s *CreateHighlightTaskRequestOutput) GetSegment() *CreateHighlightTaskRequestOutputSegment {
	return s.Segment
}

func (s *CreateHighlightTaskRequestOutput) GetSpeed() *float64 {
	return s.Speed
}

func (s *CreateHighlightTaskRequestOutput) GetTargetDuration() *float64 {
	return s.TargetDuration
}

func (s *CreateHighlightTaskRequestOutput) GetURI() *string {
	return s.URI
}

func (s *CreateHighlightTaskRequestOutput) GetVideo() *TargetVideo {
	return s.Video
}

func (s *CreateHighlightTaskRequestOutput) SetAudio(v *TargetAudio) *CreateHighlightTaskRequestOutput {
	s.Audio = v
	return s
}

func (s *CreateHighlightTaskRequestOutput) SetContainer(v string) *CreateHighlightTaskRequestOutput {
	s.Container = &v
	return s
}

func (s *CreateHighlightTaskRequestOutput) SetMaxDuration(v float64) *CreateHighlightTaskRequestOutput {
	s.MaxDuration = &v
	return s
}

func (s *CreateHighlightTaskRequestOutput) SetSegment(v *CreateHighlightTaskRequestOutputSegment) *CreateHighlightTaskRequestOutput {
	s.Segment = v
	return s
}

func (s *CreateHighlightTaskRequestOutput) SetSpeed(v float64) *CreateHighlightTaskRequestOutput {
	s.Speed = &v
	return s
}

func (s *CreateHighlightTaskRequestOutput) SetTargetDuration(v float64) *CreateHighlightTaskRequestOutput {
	s.TargetDuration = &v
	return s
}

func (s *CreateHighlightTaskRequestOutput) SetURI(v string) *CreateHighlightTaskRequestOutput {
	s.URI = &v
	return s
}

func (s *CreateHighlightTaskRequestOutput) SetVideo(v *TargetVideo) *CreateHighlightTaskRequestOutput {
	s.Video = v
	return s
}

func (s *CreateHighlightTaskRequestOutput) Validate() error {
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
	if s.Video != nil {
		if err := s.Video.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateHighlightTaskRequestOutputSegment struct {
	// The length of each segment. Unit: seconds.
	//
	// example:
	//
	// 1
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The media segmentation format. Valid values:
	//
	// - hls
	//
	// - dash
	//
	// example:
	//
	// hls
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The start number. This parameter is supported only for hls. Default value: 0.
	//
	// example:
	//
	// 0
	StartNumber *int64 `json:"StartNumber,omitempty" xml:"StartNumber,omitempty"`
}

func (s CreateHighlightTaskRequestOutputSegment) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskRequestOutputSegment) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskRequestOutputSegment) GetDuration() *float64 {
	return s.Duration
}

func (s *CreateHighlightTaskRequestOutputSegment) GetFormat() *string {
	return s.Format
}

func (s *CreateHighlightTaskRequestOutputSegment) GetStartNumber() *int64 {
	return s.StartNumber
}

func (s *CreateHighlightTaskRequestOutputSegment) SetDuration(v float64) *CreateHighlightTaskRequestOutputSegment {
	s.Duration = &v
	return s
}

func (s *CreateHighlightTaskRequestOutputSegment) SetFormat(v string) *CreateHighlightTaskRequestOutputSegment {
	s.Format = &v
	return s
}

func (s *CreateHighlightTaskRequestOutputSegment) SetStartNumber(v int64) *CreateHighlightTaskRequestOutputSegment {
	s.StartNumber = &v
	return s
}

func (s *CreateHighlightTaskRequestOutputSegment) Validate() error {
	return dara.Validate(s)
}

type CreateHighlightTaskRequestSources struct {
	// The duration of the media clip. Unit: seconds. Default value: 0, which indicates the end time of the video. This parameter is valid only when Type is set to Concat.
	//
	// example:
	//
	// 0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the media resource. Valid values: [0, video duration]. This parameter is valid only when Type is set to Concat. Unit: seconds.
	//
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The URI of the media resource, which is an OSS URI. Only videos are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s CreateHighlightTaskRequestSources) String() string {
	return dara.Prettify(s)
}

func (s CreateHighlightTaskRequestSources) GoString() string {
	return s.String()
}

func (s *CreateHighlightTaskRequestSources) GetDuration() *float64 {
	return s.Duration
}

func (s *CreateHighlightTaskRequestSources) GetStartTime() *float64 {
	return s.StartTime
}

func (s *CreateHighlightTaskRequestSources) GetURI() *string {
	return s.URI
}

func (s *CreateHighlightTaskRequestSources) SetDuration(v float64) *CreateHighlightTaskRequestSources {
	s.Duration = &v
	return s
}

func (s *CreateHighlightTaskRequestSources) SetStartTime(v float64) *CreateHighlightTaskRequestSources {
	s.StartTime = &v
	return s
}

func (s *CreateHighlightTaskRequestSources) SetURI(v string) *CreateHighlightTaskRequestSources {
	s.URI = &v
	return s
}

func (s *CreateHighlightTaskRequestSources) Validate() error {
	return dara.Validate(s)
}
