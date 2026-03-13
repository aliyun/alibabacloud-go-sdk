// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateClipsTimeLineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalContent(v string) *AsyncCreateClipsTimeLineRequest
	GetAdditionalContent() *string
	SetCustomContent(v string) *AsyncCreateClipsTimeLineRequest
	GetCustomContent() *string
	SetHighLightConfig(v *AsyncCreateClipsTimeLineRequestHighLightConfig) *AsyncCreateClipsTimeLineRequest
	GetHighLightConfig() *AsyncCreateClipsTimeLineRequestHighLightConfig
	SetNoRefVideo(v bool) *AsyncCreateClipsTimeLineRequest
	GetNoRefVideo() *bool
	SetProcessPrompt(v string) *AsyncCreateClipsTimeLineRequest
	GetProcessPrompt() *string
	SetRecommendAudio(v bool) *AsyncCreateClipsTimeLineRequest
	GetRecommendAudio() *bool
	SetTaskId(v string) *AsyncCreateClipsTimeLineRequest
	GetTaskId() *string
	SetTimelineScene(v int32) *AsyncCreateClipsTimeLineRequest
	GetTimelineScene() *int32
	SetWorkspaceId(v string) *AsyncCreateClipsTimeLineRequest
	GetWorkspaceId() *string
}

type AsyncCreateClipsTimeLineRequest struct {
	// example:
	//
	// 素材附加信息
	AdditionalContent *string `json:"AdditionalContent,omitempty" xml:"AdditionalContent,omitempty"`
	// example:
	//
	// 自定义口播内容
	CustomContent   *string                                         `json:"CustomContent,omitempty" xml:"CustomContent,omitempty"`
	HighLightConfig *AsyncCreateClipsTimeLineRequestHighLightConfig `json:"HighLightConfig,omitempty" xml:"HighLightConfig,omitempty" type:"Struct"`
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

func (s AsyncCreateClipsTimeLineRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTimeLineRequest) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTimeLineRequest) GetAdditionalContent() *string {
	return s.AdditionalContent
}

func (s *AsyncCreateClipsTimeLineRequest) GetCustomContent() *string {
	return s.CustomContent
}

func (s *AsyncCreateClipsTimeLineRequest) GetHighLightConfig() *AsyncCreateClipsTimeLineRequestHighLightConfig {
	return s.HighLightConfig
}

func (s *AsyncCreateClipsTimeLineRequest) GetNoRefVideo() *bool {
	return s.NoRefVideo
}

func (s *AsyncCreateClipsTimeLineRequest) GetProcessPrompt() *string {
	return s.ProcessPrompt
}

func (s *AsyncCreateClipsTimeLineRequest) GetRecommendAudio() *bool {
	return s.RecommendAudio
}

func (s *AsyncCreateClipsTimeLineRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *AsyncCreateClipsTimeLineRequest) GetTimelineScene() *int32 {
	return s.TimelineScene
}

func (s *AsyncCreateClipsTimeLineRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncCreateClipsTimeLineRequest) SetAdditionalContent(v string) *AsyncCreateClipsTimeLineRequest {
	s.AdditionalContent = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetCustomContent(v string) *AsyncCreateClipsTimeLineRequest {
	s.CustomContent = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetHighLightConfig(v *AsyncCreateClipsTimeLineRequestHighLightConfig) *AsyncCreateClipsTimeLineRequest {
	s.HighLightConfig = v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetNoRefVideo(v bool) *AsyncCreateClipsTimeLineRequest {
	s.NoRefVideo = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetProcessPrompt(v string) *AsyncCreateClipsTimeLineRequest {
	s.ProcessPrompt = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetRecommendAudio(v bool) *AsyncCreateClipsTimeLineRequest {
	s.RecommendAudio = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetTaskId(v string) *AsyncCreateClipsTimeLineRequest {
	s.TaskId = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetTimelineScene(v int32) *AsyncCreateClipsTimeLineRequest {
	s.TimelineScene = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) SetWorkspaceId(v string) *AsyncCreateClipsTimeLineRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequest) Validate() error {
	if s.HighLightConfig != nil {
		if err := s.HighLightConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AsyncCreateClipsTimeLineRequestHighLightConfig struct {
	// example:
	//
	// false
	HtAnalyzeRhythm *bool `json:"HtAnalyzeRhythm,omitempty" xml:"HtAnalyzeRhythm,omitempty"`
	// example:
	//
	// [
	//
	//         "高清演员近景特写镜头【高优】",
	//
	//         "演出高潮-沉浸表演【高优】",
	//
	//         "演出高潮-近景表情【高优】"
	//
	// ]
	HtHighQualityLabel []*string `json:"HtHighQualityLabel,omitempty" xml:"HtHighQualityLabel,omitempty" type:"Repeated"`
	// example:
	//
	// [
	//
	//         "画面昏暗",
	//
	//         "采访画面",
	//
	//         "字幕画面"
	//
	// ]
	HtLowQualityLabel []*string `json:"HtLowQualityLabel,omitempty" xml:"HtLowQualityLabel,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	HtMaxTimeLength *int32 `json:"HtMaxTimeLength,omitempty" xml:"HtMaxTimeLength,omitempty"`
	// example:
	//
	// 10
	HtMinTimeLength *int32 `json:"HtMinTimeLength,omitempty" xml:"HtMinTimeLength,omitempty"`
	// example:
	//
	// 请从输入的音乐演出视频中，自动识别并提取出「最具视觉冲击力、情感爆发力或独特吸引力」...
	HtPrompt *string `json:"HtPrompt,omitempty" xml:"HtPrompt,omitempty"`
	// example:
	//
	// 1.5
	HtSingleShotLen *int32 `json:"HtSingleShotLen,omitempty" xml:"HtSingleShotLen,omitempty"`
}

func (s AsyncCreateClipsTimeLineRequestHighLightConfig) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateClipsTimeLineRequestHighLightConfig) GoString() string {
	return s.String()
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) GetHtAnalyzeRhythm() *bool {
	return s.HtAnalyzeRhythm
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) GetHtHighQualityLabel() []*string {
	return s.HtHighQualityLabel
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) GetHtLowQualityLabel() []*string {
	return s.HtLowQualityLabel
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) GetHtMaxTimeLength() *int32 {
	return s.HtMaxTimeLength
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) GetHtMinTimeLength() *int32 {
	return s.HtMinTimeLength
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) GetHtPrompt() *string {
	return s.HtPrompt
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) GetHtSingleShotLen() *int32 {
	return s.HtSingleShotLen
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) SetHtAnalyzeRhythm(v bool) *AsyncCreateClipsTimeLineRequestHighLightConfig {
	s.HtAnalyzeRhythm = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) SetHtHighQualityLabel(v []*string) *AsyncCreateClipsTimeLineRequestHighLightConfig {
	s.HtHighQualityLabel = v
	return s
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) SetHtLowQualityLabel(v []*string) *AsyncCreateClipsTimeLineRequestHighLightConfig {
	s.HtLowQualityLabel = v
	return s
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) SetHtMaxTimeLength(v int32) *AsyncCreateClipsTimeLineRequestHighLightConfig {
	s.HtMaxTimeLength = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) SetHtMinTimeLength(v int32) *AsyncCreateClipsTimeLineRequestHighLightConfig {
	s.HtMinTimeLength = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) SetHtPrompt(v string) *AsyncCreateClipsTimeLineRequestHighLightConfig {
	s.HtPrompt = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) SetHtSingleShotLen(v int32) *AsyncCreateClipsTimeLineRequestHighLightConfig {
	s.HtSingleShotLen = &v
	return s
}

func (s *AsyncCreateClipsTimeLineRequestHighLightConfig) Validate() error {
	return dara.Validate(s)
}
