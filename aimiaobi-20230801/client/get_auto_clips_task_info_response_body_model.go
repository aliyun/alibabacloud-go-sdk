// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoClipsTaskInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAutoClipsTaskInfoResponseBody
	GetCode() *string
	SetData(v *GetAutoClipsTaskInfoResponseBodyData) *GetAutoClipsTaskInfoResponseBody
	GetData() *GetAutoClipsTaskInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetAutoClipsTaskInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAutoClipsTaskInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAutoClipsTaskInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAutoClipsTaskInfoResponseBody
	GetSuccess() *bool
}

type GetAutoClipsTaskInfoResponseBody struct {
	// example:
	//
	// successful
	Code           *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetAutoClipsTaskInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1813ceee-7fe5-41b4-87e5-982a4d18cca5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAutoClipsTaskInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoClipsTaskInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoClipsTaskInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAutoClipsTaskInfoResponseBody) GetData() *GetAutoClipsTaskInfoResponseBodyData {
	return s.Data
}

func (s *GetAutoClipsTaskInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAutoClipsTaskInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAutoClipsTaskInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoClipsTaskInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAutoClipsTaskInfoResponseBody) SetCode(v string) *GetAutoClipsTaskInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBody) SetData(v *GetAutoClipsTaskInfoResponseBodyData) *GetAutoClipsTaskInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBody) SetHttpStatusCode(v int32) *GetAutoClipsTaskInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBody) SetMessage(v string) *GetAutoClipsTaskInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBody) SetRequestId(v string) *GetAutoClipsTaskInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBody) SetSuccess(v bool) *GetAutoClipsTaskInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAutoClipsTaskInfoResponseBodyData struct {
	ColorWords         []*GetAutoClipsTaskInfoResponseBodyDataColorWords `json:"ColorWords,omitempty" xml:"ColorWords,omitempty" type:"Repeated"`
	Content            *string                                           `json:"Content,omitempty" xml:"Content,omitempty"`
	ErrorMessage       *string                                           `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MediaCloudTimeline *string                                           `json:"MediaCloudTimeline,omitempty" xml:"MediaCloudTimeline,omitempty"`
	MusicStyle         *string                                           `json:"MusicStyle,omitempty" xml:"MusicStyle,omitempty"`
	// example:
	//
	// http://music.mp4
	MusicUrl *string `json:"MusicUrl,omitempty" xml:"MusicUrl,omitempty"`
	// example:
	//
	// 5
	MusicVolume *int32 `json:"MusicVolume,omitempty" xml:"MusicVolume,omitempty"`
	// example:
	//
	// http://output.mp4
	OutputVideoUrl *string `json:"OutputVideoUrl,omitempty" xml:"OutputVideoUrl,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// clips
	Step *string `json:"Step,omitempty" xml:"Step,omitempty"`
	// example:
	//
	// 5
	SubtitleFontSize *int32 `json:"SubtitleFontSize,omitempty" xml:"SubtitleFontSize,omitempty"`
	// example:
	//
	// e5a1a59c82d0454fad6454e8a04d0093
	TaskId     *string                                          `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Timelines  []*GetAutoClipsTaskInfoResponseBodyDataTimelines `json:"Timelines,omitempty" xml:"Timelines,omitempty" type:"Repeated"`
	VoiceStyle *string                                          `json:"VoiceStyle,omitempty" xml:"VoiceStyle,omitempty"`
	// example:
	//
	// 5
	VoiceVolume *int32 `json:"VoiceVolume,omitempty" xml:"VoiceVolume,omitempty"`
}

func (s GetAutoClipsTaskInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAutoClipsTaskInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetColorWords() []*GetAutoClipsTaskInfoResponseBodyDataColorWords {
	return s.ColorWords
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetMediaCloudTimeline() *string {
	return s.MediaCloudTimeline
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetMusicStyle() *string {
	return s.MusicStyle
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetMusicUrl() *string {
	return s.MusicUrl
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetMusicVolume() *int32 {
	return s.MusicVolume
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetOutputVideoUrl() *string {
	return s.OutputVideoUrl
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetStep() *string {
	return s.Step
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetSubtitleFontSize() *int32 {
	return s.SubtitleFontSize
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetTimelines() []*GetAutoClipsTaskInfoResponseBodyDataTimelines {
	return s.Timelines
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetVoiceStyle() *string {
	return s.VoiceStyle
}

func (s *GetAutoClipsTaskInfoResponseBodyData) GetVoiceVolume() *int32 {
	return s.VoiceVolume
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetColorWords(v []*GetAutoClipsTaskInfoResponseBodyDataColorWords) *GetAutoClipsTaskInfoResponseBodyData {
	s.ColorWords = v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetContent(v string) *GetAutoClipsTaskInfoResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetErrorMessage(v string) *GetAutoClipsTaskInfoResponseBodyData {
	s.ErrorMessage = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetMediaCloudTimeline(v string) *GetAutoClipsTaskInfoResponseBodyData {
	s.MediaCloudTimeline = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetMusicStyle(v string) *GetAutoClipsTaskInfoResponseBodyData {
	s.MusicStyle = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetMusicUrl(v string) *GetAutoClipsTaskInfoResponseBodyData {
	s.MusicUrl = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetMusicVolume(v int32) *GetAutoClipsTaskInfoResponseBodyData {
	s.MusicVolume = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetOutputVideoUrl(v string) *GetAutoClipsTaskInfoResponseBodyData {
	s.OutputVideoUrl = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetStatus(v int32) *GetAutoClipsTaskInfoResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetStep(v string) *GetAutoClipsTaskInfoResponseBodyData {
	s.Step = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetSubtitleFontSize(v int32) *GetAutoClipsTaskInfoResponseBodyData {
	s.SubtitleFontSize = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetTaskId(v string) *GetAutoClipsTaskInfoResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetTimelines(v []*GetAutoClipsTaskInfoResponseBodyDataTimelines) *GetAutoClipsTaskInfoResponseBodyData {
	s.Timelines = v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetVoiceStyle(v string) *GetAutoClipsTaskInfoResponseBodyData {
	s.VoiceStyle = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) SetVoiceVolume(v int32) *GetAutoClipsTaskInfoResponseBodyData {
	s.VoiceVolume = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyData) Validate() error {
	if s.ColorWords != nil {
		for _, item := range s.ColorWords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Timelines != nil {
		for _, item := range s.Timelines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAutoClipsTaskInfoResponseBodyDataColorWords struct {
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// CS0002-000008
	EffectColorStyle *string `json:"EffectColorStyle,omitempty" xml:"EffectColorStyle,omitempty"`
	// example:
	//
	// 5
	FontSize *int32 `json:"FontSize,omitempty" xml:"FontSize,omitempty"`
	// example:
	//
	// 0
	TimelineIn *int32 `json:"TimelineIn,omitempty" xml:"TimelineIn,omitempty"`
	// example:
	//
	// 5
	TimelineOut *int32 `json:"TimelineOut,omitempty" xml:"TimelineOut,omitempty"`
	// example:
	//
	// 0.2
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.5
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s GetAutoClipsTaskInfoResponseBodyDataColorWords) String() string {
	return dara.Prettify(s)
}

func (s GetAutoClipsTaskInfoResponseBodyDataColorWords) GoString() string {
	return s.String()
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) GetContent() *string {
	return s.Content
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) GetEffectColorStyle() *string {
	return s.EffectColorStyle
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) GetFontSize() *int32 {
	return s.FontSize
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) GetTimelineIn() *int32 {
	return s.TimelineIn
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) GetTimelineOut() *int32 {
	return s.TimelineOut
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) GetX() *float32 {
	return s.X
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) GetY() *float32 {
	return s.Y
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) SetContent(v string) *GetAutoClipsTaskInfoResponseBodyDataColorWords {
	s.Content = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) SetEffectColorStyle(v string) *GetAutoClipsTaskInfoResponseBodyDataColorWords {
	s.EffectColorStyle = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) SetFontSize(v int32) *GetAutoClipsTaskInfoResponseBodyDataColorWords {
	s.FontSize = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) SetTimelineIn(v int32) *GetAutoClipsTaskInfoResponseBodyDataColorWords {
	s.TimelineIn = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) SetTimelineOut(v int32) *GetAutoClipsTaskInfoResponseBodyDataColorWords {
	s.TimelineOut = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) SetX(v float32) *GetAutoClipsTaskInfoResponseBodyDataColorWords {
	s.X = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) SetY(v float32) *GetAutoClipsTaskInfoResponseBodyDataColorWords {
	s.Y = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataColorWords) Validate() error {
	return dara.Validate(s)
}

type GetAutoClipsTaskInfoResponseBodyDataTimelines struct {
	Clips   []*GetAutoClipsTaskInfoResponseBodyDataTimelinesClips `json:"Clips,omitempty" xml:"Clips,omitempty" type:"Repeated"`
	Content *string                                               `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 20774ebd9abc71ef80486632b68f0102
	TimelineId *string `json:"TimelineId,omitempty" xml:"TimelineId,omitempty"`
}

func (s GetAutoClipsTaskInfoResponseBodyDataTimelines) String() string {
	return dara.Prettify(s)
}

func (s GetAutoClipsTaskInfoResponseBodyDataTimelines) GoString() string {
	return s.String()
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelines) GetClips() []*GetAutoClipsTaskInfoResponseBodyDataTimelinesClips {
	return s.Clips
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelines) GetContent() *string {
	return s.Content
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelines) GetTimelineId() *string {
	return s.TimelineId
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelines) SetClips(v []*GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) *GetAutoClipsTaskInfoResponseBodyDataTimelines {
	s.Clips = v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelines) SetContent(v string) *GetAutoClipsTaskInfoResponseBodyDataTimelines {
	s.Content = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelines) SetTimelineId(v string) *GetAutoClipsTaskInfoResponseBodyDataTimelines {
	s.TimelineId = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelines) Validate() error {
	if s.Clips != nil {
		for _, item := range s.Clips {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAutoClipsTaskInfoResponseBodyDataTimelinesClips struct {
	// example:
	//
	// 20774ebd9abc71ef80486632b68f0102
	ClipId       *string `json:"ClipId,omitempty" xml:"ClipId,omitempty"`
	ContentInner *string `json:"ContentInner,omitempty" xml:"ContentInner,omitempty"`
	// example:
	//
	// 0
	In   *int32   `json:"In,omitempty" xml:"In,omitempty"`
	InEx *float32 `json:"InEx,omitempty" xml:"InEx,omitempty"`
	// example:
	//
	// 5
	Out   *int32   `json:"Out,omitempty" xml:"Out,omitempty"`
	OutEx *float32 `json:"OutEx,omitempty" xml:"OutEx,omitempty"`
	// example:
	//
	// 20774ebd9abc71ef80486632b68f0102
	VideoId *string `json:"VideoId,omitempty" xml:"VideoId,omitempty"`
	// example:
	//
	// 123.mp4
	VideoName *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
}

func (s GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) String() string {
	return dara.Prettify(s)
}

func (s GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) GoString() string {
	return s.String()
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) GetClipId() *string {
	return s.ClipId
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) GetContentInner() *string {
	return s.ContentInner
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) GetIn() *int32 {
	return s.In
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) GetInEx() *float32 {
	return s.InEx
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) GetOut() *int32 {
	return s.Out
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) GetOutEx() *float32 {
	return s.OutEx
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) GetVideoId() *string {
	return s.VideoId
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) GetVideoName() *string {
	return s.VideoName
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) SetClipId(v string) *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips {
	s.ClipId = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) SetContentInner(v string) *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips {
	s.ContentInner = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) SetIn(v int32) *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips {
	s.In = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) SetInEx(v float32) *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips {
	s.InEx = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) SetOut(v int32) *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips {
	s.Out = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) SetOutEx(v float32) *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips {
	s.OutEx = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) SetVideoId(v string) *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips {
	s.VideoId = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) SetVideoName(v string) *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips {
	s.VideoName = &v
	return s
}

func (s *GetAutoClipsTaskInfoResponseBodyDataTimelinesClips) Validate() error {
	return dara.Validate(s)
}
