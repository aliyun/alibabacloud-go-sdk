// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoAnalysisTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddDocumentParamShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetAddDocumentParamShrink() *string
	SetAutoRoleRecognitionVideoUrl(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetAutoRoleRecognitionVideoUrl() *string
	SetDeduplicationId(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetDeduplicationId() *string
	SetExcludeGenerateOptionsShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetExcludeGenerateOptionsShrink() *string
	SetFaceIdentitySimilarityMinScore(v float32) *SubmitVideoAnalysisTaskShrinkRequest
	GetFaceIdentitySimilarityMinScore() *float32
	SetFrameSampleMethodShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetFrameSampleMethodShrink() *string
	SetGenerateOptionsShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetGenerateOptionsShrink() *string
	SetLanguage(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetLanguage() *string
	SetModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetModelCustomPromptTemplate() *string
	SetModelCustomPromptTemplateId(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetModelCustomPromptTemplateId() *string
	SetModelId(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetModelId() *string
	SetSnapshotInterval(v float64) *SubmitVideoAnalysisTaskShrinkRequest
	GetSnapshotInterval() *float64
	SetSplitInterval(v int32) *SubmitVideoAnalysisTaskShrinkRequest
	GetSplitInterval() *int32
	SetSplitType(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetSplitType() *string
	SetTextProcessTasksShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetTextProcessTasksShrink() *string
	SetVideoCaptionInfoShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetVideoCaptionInfoShrink() *string
	SetVideoExtraInfo(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetVideoExtraInfo() *string
	SetVideoModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetVideoModelCustomPromptTemplate() *string
	SetVideoModelId(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetVideoModelId() *string
	SetVideoRolesShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetVideoRolesShrink() *string
	SetVideoShotFaceIdentityCount(v int32) *SubmitVideoAnalysisTaskShrinkRequest
	GetVideoShotFaceIdentityCount() *int32
	SetVideoUrl(v string) *SubmitVideoAnalysisTaskShrinkRequest
	GetVideoUrl() *string
}

type SubmitVideoAnalysisTaskShrinkRequest struct {
	AddDocumentParamShrink      *string `json:"addDocumentParam,omitempty" xml:"addDocumentParam,omitempty"`
	AutoRoleRecognitionVideoUrl *string `json:"autoRoleRecognitionVideoUrl,omitempty" xml:"autoRoleRecognitionVideoUrl,omitempty"`
	// example:
	//
	// 1
	DeduplicationId                *string  `json:"deduplicationId,omitempty" xml:"deduplicationId,omitempty"`
	ExcludeGenerateOptionsShrink   *string  `json:"excludeGenerateOptions,omitempty" xml:"excludeGenerateOptions,omitempty"`
	FaceIdentitySimilarityMinScore *float32 `json:"faceIdentitySimilarityMinScore,omitempty" xml:"faceIdentitySimilarityMinScore,omitempty"`
	FrameSampleMethodShrink        *string  `json:"frameSampleMethod,omitempty" xml:"frameSampleMethod,omitempty"`
	GenerateOptionsShrink          *string  `json:"generateOptions,omitempty" xml:"generateOptions,omitempty"`
	// example:
	//
	// chinese
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
	// 2
	SnapshotInterval *float64 `json:"snapshotInterval,omitempty" xml:"snapshotInterval,omitempty"`
	// example:
	//
	// 10
	SplitInterval *int32 `json:"splitInterval,omitempty" xml:"splitInterval,omitempty"`
	// example:
	//
	// fixDuration
	SplitType                      *string `json:"splitType,omitempty" xml:"splitType,omitempty"`
	TextProcessTasksShrink         *string `json:"textProcessTasks,omitempty" xml:"textProcessTasks,omitempty"`
	VideoCaptionInfoShrink         *string `json:"videoCaptionInfo,omitempty" xml:"videoCaptionInfo,omitempty"`
	VideoExtraInfo                 *string `json:"videoExtraInfo,omitempty" xml:"videoExtraInfo,omitempty"`
	VideoModelCustomPromptTemplate *string `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max-latest
	VideoModelId               *string `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	VideoRolesShrink           *string `json:"videoRoles,omitempty" xml:"videoRoles,omitempty"`
	VideoShotFaceIdentityCount *int32  `json:"videoShotFaceIdentityCount,omitempty" xml:"videoShotFaceIdentityCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://xxxx.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s SubmitVideoAnalysisTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoAnalysisTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetAddDocumentParamShrink() *string {
	return s.AddDocumentParamShrink
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetAutoRoleRecognitionVideoUrl() *string {
	return s.AutoRoleRecognitionVideoUrl
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetDeduplicationId() *string {
	return s.DeduplicationId
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetExcludeGenerateOptionsShrink() *string {
	return s.ExcludeGenerateOptionsShrink
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetFaceIdentitySimilarityMinScore() *float32 {
	return s.FaceIdentitySimilarityMinScore
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetFrameSampleMethodShrink() *string {
	return s.FrameSampleMethodShrink
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetGenerateOptionsShrink() *string {
	return s.GenerateOptionsShrink
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetLanguage() *string {
	return s.Language
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetModelCustomPromptTemplate() *string {
	return s.ModelCustomPromptTemplate
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetModelId() *string {
	return s.ModelId
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetSnapshotInterval() *float64 {
	return s.SnapshotInterval
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetSplitInterval() *int32 {
	return s.SplitInterval
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetSplitType() *string {
	return s.SplitType
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetTextProcessTasksShrink() *string {
	return s.TextProcessTasksShrink
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetVideoCaptionInfoShrink() *string {
	return s.VideoCaptionInfoShrink
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetVideoExtraInfo() *string {
	return s.VideoExtraInfo
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetVideoModelCustomPromptTemplate() *string {
	return s.VideoModelCustomPromptTemplate
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetVideoModelId() *string {
	return s.VideoModelId
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetVideoRolesShrink() *string {
	return s.VideoRolesShrink
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetVideoShotFaceIdentityCount() *int32 {
	return s.VideoShotFaceIdentityCount
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetAddDocumentParamShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.AddDocumentParamShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetAutoRoleRecognitionVideoUrl(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.AutoRoleRecognitionVideoUrl = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetDeduplicationId(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.DeduplicationId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetExcludeGenerateOptionsShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.ExcludeGenerateOptionsShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetFaceIdentitySimilarityMinScore(v float32) *SubmitVideoAnalysisTaskShrinkRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetFrameSampleMethodShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.FrameSampleMethodShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetGenerateOptionsShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.GenerateOptionsShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetLanguage(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.Language = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetModelCustomPromptTemplateId(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetModelId(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.ModelId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetSnapshotInterval(v float64) *SubmitVideoAnalysisTaskShrinkRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetSplitInterval(v int32) *SubmitVideoAnalysisTaskShrinkRequest {
	s.SplitInterval = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetSplitType(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.SplitType = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetTextProcessTasksShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.TextProcessTasksShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoCaptionInfoShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoCaptionInfoShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoExtraInfo(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoExtraInfo = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoModelCustomPromptTemplate(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoModelCustomPromptTemplate = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoModelId(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoModelId = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoRolesShrink(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoRolesShrink = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoShotFaceIdentityCount(v int32) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoShotFaceIdentityCount = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) SetVideoUrl(v string) *SubmitVideoAnalysisTaskShrinkRequest {
	s.VideoUrl = &v
	return s
}

func (s *SubmitVideoAnalysisTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
