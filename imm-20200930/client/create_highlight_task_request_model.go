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
	CredentialConfig *CredentialConfig                    `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	Edit             *CreateHighlightTaskRequestEdit      `json:"Edit,omitempty" xml:"Edit,omitempty" type:"Struct"`
	Highlight        *CreateHighlightTaskRequestHighlight `json:"Highlight,omitempty" xml:"Highlight,omitempty" type:"Struct"`
	// example:
	//
	// Average
	Mode         *string       `json:"Mode,omitempty" xml:"Mode,omitempty"`
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// This parameter is required.
	Output *CreateHighlightTaskRequestOutput `json:"Output,omitempty" xml:"Output,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	Sources []*CreateHighlightTaskRequestSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// example:
	//
	// {"test":"val1"}
	Tags map[string]interface{} `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Retrieval
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// example:
	//
	// Closed
	BackgroundMusicMode *string                                           `json:"BackgroundMusicMode,omitempty" xml:"BackgroundMusicMode,omitempty"`
	BackgroundMusics    []*CreateHighlightTaskRequestEditBackgroundMusics `json:"BackgroundMusics,omitempty" xml:"BackgroundMusics,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// Sequential
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// Closed
	TransitionMode *string                                      `json:"TransitionMode,omitempty" xml:"TransitionMode,omitempty"`
	Transitions    []*CreateHighlightTaskRequestEditTransitions `json:"Transitions,omitempty" xml:"Transitions,omitempty" type:"Repeated"`
	// example:
	//
	// Closed
	VfxEffectMode *string                                     `json:"VfxEffectMode,omitempty" xml:"VfxEffectMode,omitempty"`
	VfxEffects    []*CreateHighlightTaskRequestEditVfxEffects `json:"VfxEffects,omitempty" xml:"VfxEffects,omitempty" type:"Repeated"`
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
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
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
	// example:
	//
	// 0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// directional
	Transition *string `json:"Transition,omitempty" xml:"Transition,omitempty"`
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
	// This parameter is required.
	//
	// example:
	//
	// open
	VfxEffect *string `json:"VfxEffect,omitempty" xml:"VfxEffect,omitempty"`
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
	Audio *TargetAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	// example:
	//
	// mp4
	Container *string `json:"Container,omitempty" xml:"Container,omitempty"`
	// example:
	//
	// 10.0
	MaxDuration *float64                                 `json:"MaxDuration,omitempty" xml:"MaxDuration,omitempty"`
	Segment     *CreateHighlightTaskRequestOutputSegment `json:"Segment,omitempty" xml:"Segment,omitempty" type:"Struct"`
	// example:
	//
	// 1.0
	Speed *float64 `json:"Speed,omitempty" xml:"Speed,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-target-object.mp4
	URI   *string      `json:"URI,omitempty" xml:"URI,omitempty"`
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
	// example:
	//
	// 1
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// hls
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
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
	// example:
	//
	// 0
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 0
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
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
