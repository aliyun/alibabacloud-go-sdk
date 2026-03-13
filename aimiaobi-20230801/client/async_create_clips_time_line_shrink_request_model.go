// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTimeLineShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalContent(v string) *AsyncCreateClipsTimeLineShrinkRequest
	GetAdditionalContent() *string
	SetCustomContent(v string) *AsyncCreateClipsTimeLineShrinkRequest
	GetCustomContent() *string
	SetHighLightConfigShrink(v string) *AsyncCreateClipsTimeLineShrinkRequest
	GetHighLightConfigShrink() *string
	SetNoRefVideo(v bool) *AsyncCreateClipsTimeLineShrinkRequest
	GetNoRefVideo() *bool
	SetProcessPrompt(v string) *AsyncCreateClipsTimeLineShrinkRequest
	GetProcessPrompt() *string
	SetRecommendAudio(v bool) *AsyncCreateClipsTimeLineShrinkRequest
	GetRecommendAudio() *bool
	SetTaskId(v string) *AsyncCreateClipsTimeLineShrinkRequest
	GetTaskId() *string
	SetTimelineScene(v int32) *AsyncCreateClipsTimeLineShrinkRequest
	GetTimelineScene() *int32
	SetWorkspaceId(v string) *AsyncCreateClipsTimeLineShrinkRequest
	GetWorkspaceId() *string
}

type AsyncCreateClipsTimeLineShrinkRequest struct {
	// example:
	//
	// 素材附加信息
	AdditionalContent *string `json:"AdditionalContent,omitempty" xml:"AdditionalContent,omitempty"`
	// example:
	//
	// 自定义口播内容
	CustomContent         *string `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	HighLightConfigShrink *string `json:"HighLightConfig,omitempty" xml:"HighLightConfig,omitempty"`
	// example:
	//
	// 默认开启
	NoRefVideo *bool `json:"NoRefVideo,omitempty" xml:"NoRefVideo,omitempty"`
	// example:
	//
	// 口播内容是乌镇旅游宣传广告，口播内容时长约为1分钟，开头要描述乌镇是千年文化传承的江南水乡，之后要体现乌镇的传统手工艺、美食和美景，最后要号召大家来乌镇旅游
	ProcessPrompt *string `json:"ProcessPrompt,omitempty" xml:"ProcessPrompt,omitempty"`
	// example:
	//
	// false
	RecommendAudio *bool `json:"RecommendAudio,omitempty" xml:"RecommendAudio,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// null - 通用口播
	//
	//  0-通用口播
	//
	//  1- 高光
	TimelineScene *int32 `json:"TimelineScene,omitempty" xml:"TimelineScene,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncCreateClipsTimeLineShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTimeLineShrinkRequest) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) GetAdditionalContent() *string {
	return s.AdditionalContent
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) GetCustomContent() *string {
	return s.CustomContent
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) GetHighLightConfigShrink() *string {
	return s.HighLightConfigShrink
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) GetNoRefVideo() *bool {
	return s.NoRefVideo
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) GetProcessPrompt() *string {
	return s.ProcessPrompt
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) GetRecommendAudio() *bool {
	return s.RecommendAudio
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) GetTimelineScene() *int32 {
	return s.TimelineScene
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) SetAdditionalContent(v string) *AsyncCreateClipsTimeLineShrinkRequest {
	s.AdditionalContent = &v
	return s
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) SetCustomContent(v string) *AsyncCreateClipsTimeLineShrinkRequest {
	s.CustomContent = &v
	return s
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) SetHighLightConfigShrink(v string) *AsyncCreateClipsTimeLineShrinkRequest {
	s.HighLightConfigShrink = &v
	return s
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) SetNoRefVideo(v bool) *AsyncCreateClipsTimeLineShrinkRequest {
	s.NoRefVideo = &v
	return s
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) SetProcessPrompt(v string) *AsyncCreateClipsTimeLineShrinkRequest {
	s.ProcessPrompt = &v
	return s
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) SetRecommendAudio(v bool) *AsyncCreateClipsTimeLineShrinkRequest {
	s.RecommendAudio = &v
	return s
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) SetTaskId(v string) *AsyncCreateClipsTimeLineShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) SetTimelineScene(v int32) *AsyncCreateClipsTimeLineShrinkRequest {
	s.TimelineScene = &v
	return s
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) SetWorkspaceId(v string) *AsyncCreateClipsTimeLineShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncCreateClipsTimeLineShrinkRequest) Validate() error {
	return dara.Validate(s)
}
