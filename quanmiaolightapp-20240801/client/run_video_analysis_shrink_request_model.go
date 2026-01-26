// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoAnalysisShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddDocumentParamShrink(v string) *RunVideoAnalysisShrinkRequest
	GetAddDocumentParamShrink() *string
	SetAutoRoleRecognitionVideoUrl(v string) *RunVideoAnalysisShrinkRequest
	GetAutoRoleRecognitionVideoUrl() *string
	SetExcludeGenerateOptionsShrink(v string) *RunVideoAnalysisShrinkRequest
	GetExcludeGenerateOptionsShrink() *string
	SetFaceIdentitySimilarityMinScore(v float32) *RunVideoAnalysisShrinkRequest
	GetFaceIdentitySimilarityMinScore() *float32
	SetFrameSampleMethodShrink(v string) *RunVideoAnalysisShrinkRequest
	GetFrameSampleMethodShrink() *string
	SetGenerateOptionsShrink(v string) *RunVideoAnalysisShrinkRequest
	GetGenerateOptionsShrink() *string
	SetLanguage(v string) *RunVideoAnalysisShrinkRequest
	GetLanguage() *string
	SetModelCustomPromptTemplate(v string) *RunVideoAnalysisShrinkRequest
	GetModelCustomPromptTemplate() *string
	SetModelCustomPromptTemplateId(v string) *RunVideoAnalysisShrinkRequest
	GetModelCustomPromptTemplateId() *string
	SetModelId(v string) *RunVideoAnalysisShrinkRequest
	GetModelId() *string
	SetOriginalSessionId(v string) *RunVideoAnalysisShrinkRequest
	GetOriginalSessionId() *string
	SetSnapshotInterval(v float64) *RunVideoAnalysisShrinkRequest
	GetSnapshotInterval() *float64
	SetSplitInterval(v int32) *RunVideoAnalysisShrinkRequest
	GetSplitInterval() *int32
	SetSplitType(v string) *RunVideoAnalysisShrinkRequest
	GetSplitType() *string
	SetTaskId(v string) *RunVideoAnalysisShrinkRequest
	GetTaskId() *string
	SetTextProcessTasksShrink(v string) *RunVideoAnalysisShrinkRequest
	GetTextProcessTasksShrink() *string
	SetVideoCaptionInfoShrink(v string) *RunVideoAnalysisShrinkRequest
	GetVideoCaptionInfoShrink() *string
	SetVideoExtraInfo(v string) *RunVideoAnalysisShrinkRequest
	GetVideoExtraInfo() *string
	SetVideoModelCustomPromptTemplate(v string) *RunVideoAnalysisShrinkRequest
	GetVideoModelCustomPromptTemplate() *string
	SetVideoModelId(v string) *RunVideoAnalysisShrinkRequest
	GetVideoModelId() *string
	SetVideoRolesShrink(v string) *RunVideoAnalysisShrinkRequest
	GetVideoRolesShrink() *string
	SetVideoShotFaceIdentityCount(v int32) *RunVideoAnalysisShrinkRequest
	GetVideoShotFaceIdentityCount() *int32
	SetVideoUrl(v string) *RunVideoAnalysisShrinkRequest
	GetVideoUrl() *string
	SetVideoUrlsShrink(v string) *RunVideoAnalysisShrinkRequest
	GetVideoUrlsShrink() *string
}

type RunVideoAnalysisShrinkRequest struct {
	AddDocumentParamShrink         *string  `json:"addDocumentParam,omitempty" xml:"addDocumentParam,omitempty"`
	AutoRoleRecognitionVideoUrl    *string  `json:"autoRoleRecognitionVideoUrl,omitempty" xml:"autoRoleRecognitionVideoUrl,omitempty"`
	ExcludeGenerateOptionsShrink   *string  `json:"excludeGenerateOptions,omitempty" xml:"excludeGenerateOptions,omitempty"`
	FaceIdentitySimilarityMinScore *float32 `json:"faceIdentitySimilarityMinScore,omitempty" xml:"faceIdentitySimilarityMinScore,omitempty"`
	FrameSampleMethodShrink        *string  `json:"frameSampleMethod,omitempty" xml:"frameSampleMethod,omitempty"`
	GenerateOptionsShrink          *string  `json:"generateOptions,omitempty" xml:"generateOptions,omitempty"`
	// example:
	//
	// english
	Language                  *string `json:"language,omitempty" xml:"language,omitempty"`
	ModelCustomPromptTemplate *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	// example:
	//
	// PlotDetail
	ModelCustomPromptTemplateId *string `json:"modelCustomPromptTemplateId,omitempty" xml:"modelCustomPromptTemplateId,omitempty"`
	// example:
	//
	// qwen-max
	ModelId *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5ax
	OriginalSessionId *string  `json:"originalSessionId,omitempty" xml:"originalSessionId,omitempty"`
	SnapshotInterval  *float64 `json:"snapshotInterval,omitempty" xml:"snapshotInterval,omitempty"`
	// example:
	//
	// 10
	SplitInterval *int32 `json:"splitInterval,omitempty" xml:"splitInterval,omitempty"`
	// example:
	//
	// fixDuration
	SplitType *string `json:"splitType,omitempty" xml:"splitType,omitempty"`
	// example:
	//
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId                         *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TextProcessTasksShrink         *string `json:"textProcessTasks,omitempty" xml:"textProcessTasks,omitempty"`
	VideoCaptionInfoShrink         *string `json:"videoCaptionInfo,omitempty" xml:"videoCaptionInfo,omitempty"`
	VideoExtraInfo                 *string `json:"videoExtraInfo,omitempty" xml:"videoExtraInfo,omitempty"`
	VideoModelCustomPromptTemplate *string `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max
	VideoModelId               *string `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	VideoRolesShrink           *string `json:"videoRoles,omitempty" xml:"videoRoles,omitempty"`
	VideoShotFaceIdentityCount *int32  `json:"videoShotFaceIdentityCount,omitempty" xml:"videoShotFaceIdentityCount,omitempty"`
	// example:
	//
	// http://xxxx.mp4
	VideoUrl        *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
	VideoUrlsShrink *string `json:"videoUrls,omitempty" xml:"videoUrls,omitempty"`
}

func (s RunVideoAnalysisShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisShrinkRequest) GetAddDocumentParamShrink() *string {
	return s.AddDocumentParamShrink
}

func (s *RunVideoAnalysisShrinkRequest) GetAutoRoleRecognitionVideoUrl() *string {
	return s.AutoRoleRecognitionVideoUrl
}

func (s *RunVideoAnalysisShrinkRequest) GetExcludeGenerateOptionsShrink() *string {
	return s.ExcludeGenerateOptionsShrink
}

func (s *RunVideoAnalysisShrinkRequest) GetFaceIdentitySimilarityMinScore() *float32 {
	return s.FaceIdentitySimilarityMinScore
}

func (s *RunVideoAnalysisShrinkRequest) GetFrameSampleMethodShrink() *string {
	return s.FrameSampleMethodShrink
}

func (s *RunVideoAnalysisShrinkRequest) GetGenerateOptionsShrink() *string {
	return s.GenerateOptionsShrink
}

func (s *RunVideoAnalysisShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *RunVideoAnalysisShrinkRequest) GetModelCustomPromptTemplate() *string {
	return s.ModelCustomPromptTemplate
}

func (s *RunVideoAnalysisShrinkRequest) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *RunVideoAnalysisShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoAnalysisShrinkRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *RunVideoAnalysisShrinkRequest) GetSnapshotInterval() *float64 {
	return s.SnapshotInterval
}

func (s *RunVideoAnalysisShrinkRequest) GetSplitInterval() *int32 {
	return s.SplitInterval
}

func (s *RunVideoAnalysisShrinkRequest) GetSplitType() *string {
	return s.SplitType
}

func (s *RunVideoAnalysisShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunVideoAnalysisShrinkRequest) GetTextProcessTasksShrink() *string {
	return s.TextProcessTasksShrink
}

func (s *RunVideoAnalysisShrinkRequest) GetVideoCaptionInfoShrink() *string {
	return s.VideoCaptionInfoShrink
}

func (s *RunVideoAnalysisShrinkRequest) GetVideoExtraInfo() *string {
	return s.VideoExtraInfo
}

func (s *RunVideoAnalysisShrinkRequest) GetVideoModelCustomPromptTemplate() *string {
	return s.VideoModelCustomPromptTemplate
}

func (s *RunVideoAnalysisShrinkRequest) GetVideoModelId() *string {
	return s.VideoModelId
}

func (s *RunVideoAnalysisShrinkRequest) GetVideoRolesShrink() *string {
	return s.VideoRolesShrink
}

func (s *RunVideoAnalysisShrinkRequest) GetVideoShotFaceIdentityCount() *int32 {
	return s.VideoShotFaceIdentityCount
}

func (s *RunVideoAnalysisShrinkRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *RunVideoAnalysisShrinkRequest) GetVideoUrlsShrink() *string {
	return s.VideoUrlsShrink
}

func (s *RunVideoAnalysisShrinkRequest) SetAddDocumentParamShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.AddDocumentParamShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetAutoRoleRecognitionVideoUrl(v string) *RunVideoAnalysisShrinkRequest {
	s.AutoRoleRecognitionVideoUrl = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetExcludeGenerateOptionsShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.ExcludeGenerateOptionsShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetFaceIdentitySimilarityMinScore(v float32) *RunVideoAnalysisShrinkRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetFrameSampleMethodShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.FrameSampleMethodShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetGenerateOptionsShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.GenerateOptionsShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetLanguage(v string) *RunVideoAnalysisShrinkRequest {
	s.Language = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetModelCustomPromptTemplate(v string) *RunVideoAnalysisShrinkRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetModelCustomPromptTemplateId(v string) *RunVideoAnalysisShrinkRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetModelId(v string) *RunVideoAnalysisShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetOriginalSessionId(v string) *RunVideoAnalysisShrinkRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetSnapshotInterval(v float64) *RunVideoAnalysisShrinkRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetSplitInterval(v int32) *RunVideoAnalysisShrinkRequest {
	s.SplitInterval = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetSplitType(v string) *RunVideoAnalysisShrinkRequest {
	s.SplitType = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetTaskId(v string) *RunVideoAnalysisShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetTextProcessTasksShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.TextProcessTasksShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoCaptionInfoShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoCaptionInfoShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoExtraInfo(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoExtraInfo = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoModelCustomPromptTemplate(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoModelId(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoModelId = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoRolesShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoRolesShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoShotFaceIdentityCount(v int32) *RunVideoAnalysisShrinkRequest {
	s.VideoShotFaceIdentityCount = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoUrl(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoUrl = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) SetVideoUrlsShrink(v string) *RunVideoAnalysisShrinkRequest {
	s.VideoUrlsShrink = &v
	return s
}

func (s *RunVideoAnalysisShrinkRequest) Validate() error {
	return dara.Validate(s)
}
