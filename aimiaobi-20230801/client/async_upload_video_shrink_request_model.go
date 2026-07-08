// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncUploadVideoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptiveThreshold(v float32) *AsyncUploadVideoShrinkRequest
	GetAdaptiveThreshold() *float32
	SetAnlysisPrompt(v string) *AsyncUploadVideoShrinkRequest
	GetAnlysisPrompt() *string
	SetFaceIdentitySimilarityMinScore(v float64) *AsyncUploadVideoShrinkRequest
	GetFaceIdentitySimilarityMinScore() *float64
	SetReferenceVideoShrink(v string) *AsyncUploadVideoShrinkRequest
	GetReferenceVideoShrink() *string
	SetRemoveSubtitle(v bool) *AsyncUploadVideoShrinkRequest
	GetRemoveSubtitle() *bool
	SetSourceVideosShrink(v string) *AsyncUploadVideoShrinkRequest
	GetSourceVideosShrink() *string
	SetSplitInterval(v int32) *AsyncUploadVideoShrinkRequest
	GetSplitInterval() *int32
	SetTaskName(v string) *AsyncUploadVideoShrinkRequest
	GetTaskName() *string
	SetTaskType(v string) *AsyncUploadVideoShrinkRequest
	GetTaskType() *string
	SetVideoRolesShrink(v string) *AsyncUploadVideoShrinkRequest
	GetVideoRolesShrink() *string
	SetVideoShotFaceIdentityCount(v int32) *AsyncUploadVideoShrinkRequest
	GetVideoShotFaceIdentityCount() *int32
	SetWorkspaceId(v string) *AsyncUploadVideoShrinkRequest
	GetWorkspaceId() *string
}

type AsyncUploadVideoShrinkRequest struct {
	// Shot segmentation threshold. A smaller value increases sensitivity. Valid range is 1 to 10. Default value is 3.
	//
	// example:
	//
	// 3.0
	AdaptiveThreshold *float32 `json:"AdaptiveThreshold,omitempty" xml:"AdaptiveThreshold,omitempty"`
	// The prompt for video understanding.
	//
	// example:
	//
	// 重点理解视频中的风景信息
	AnlysisPrompt *string `json:"AnlysisPrompt,omitempty" xml:"AnlysisPrompt,omitempty"`
	// The similarity threshold for character recognition.
	//
	// example:
	//
	// 0.7
	FaceIdentitySimilarityMinScore *float64 `json:"FaceIdentitySimilarityMinScore,omitempty" xml:"FaceIdentitySimilarityMinScore,omitempty"`
	// Information about the reference video.
	ReferenceVideoShrink *string `json:"ReferenceVideo,omitempty" xml:"ReferenceVideo,omitempty"`
	// Removes captions from the material.
	RemoveSubtitle *bool `json:"RemoveSubtitle,omitempty" xml:"RemoveSubtitle,omitempty"`
	// The structure of the video editing materials.
	//
	// This parameter is required.
	SourceVideosShrink *string `json:"SourceVideos,omitempty" xml:"SourceVideos,omitempty"`
	// The time interval for video understanding shots.
	//
	// example:
	//
	// 默认1
	SplitInterval *int32 `json:"SplitInterval,omitempty" xml:"SplitInterval,omitempty"`
	// Job name
	//
	// example:
	//
	// task001
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// Task Type
	//
	// example:
	//
	// type001
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Face information of the roles.
	VideoRolesShrink *string `json:"VideoRoles,omitempty" xml:"VideoRoles,omitempty"`
	// The number of frames sampled from a single shot for character matching.
	//
	// example:
	//
	// 2
	VideoShotFaceIdentityCount *int32 `json:"VideoShotFaceIdentityCount,omitempty" xml:"VideoShotFaceIdentityCount,omitempty"`
	// [The ID of the Alibaba Cloud Model Studio workspace.](https://help.aliyun.com/document_detail/2782167.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncUploadVideoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoShrinkRequest) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoShrinkRequest) GetAdaptiveThreshold() *float32 {
	return s.AdaptiveThreshold
}

func (s *AsyncUploadVideoShrinkRequest) GetAnlysisPrompt() *string {
	return s.AnlysisPrompt
}

func (s *AsyncUploadVideoShrinkRequest) GetFaceIdentitySimilarityMinScore() *float64 {
	return s.FaceIdentitySimilarityMinScore
}

func (s *AsyncUploadVideoShrinkRequest) GetReferenceVideoShrink() *string {
	return s.ReferenceVideoShrink
}

func (s *AsyncUploadVideoShrinkRequest) GetRemoveSubtitle() *bool {
	return s.RemoveSubtitle
}

func (s *AsyncUploadVideoShrinkRequest) GetSourceVideosShrink() *string {
	return s.SourceVideosShrink
}

func (s *AsyncUploadVideoShrinkRequest) GetSplitInterval() *int32 {
	return s.SplitInterval
}

func (s *AsyncUploadVideoShrinkRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *AsyncUploadVideoShrinkRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *AsyncUploadVideoShrinkRequest) GetVideoRolesShrink() *string {
	return s.VideoRolesShrink
}

func (s *AsyncUploadVideoShrinkRequest) GetVideoShotFaceIdentityCount() *int32 {
	return s.VideoShotFaceIdentityCount
}

func (s *AsyncUploadVideoShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncUploadVideoShrinkRequest) SetAdaptiveThreshold(v float32) *AsyncUploadVideoShrinkRequest {
	s.AdaptiveThreshold = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetAnlysisPrompt(v string) *AsyncUploadVideoShrinkRequest {
	s.AnlysisPrompt = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetFaceIdentitySimilarityMinScore(v float64) *AsyncUploadVideoShrinkRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetReferenceVideoShrink(v string) *AsyncUploadVideoShrinkRequest {
	s.ReferenceVideoShrink = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetRemoveSubtitle(v bool) *AsyncUploadVideoShrinkRequest {
	s.RemoveSubtitle = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetSourceVideosShrink(v string) *AsyncUploadVideoShrinkRequest {
	s.SourceVideosShrink = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetSplitInterval(v int32) *AsyncUploadVideoShrinkRequest {
	s.SplitInterval = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetTaskName(v string) *AsyncUploadVideoShrinkRequest {
	s.TaskName = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetTaskType(v string) *AsyncUploadVideoShrinkRequest {
	s.TaskType = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetVideoRolesShrink(v string) *AsyncUploadVideoShrinkRequest {
	s.VideoRolesShrink = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetVideoShotFaceIdentityCount(v int32) *AsyncUploadVideoShrinkRequest {
	s.VideoShotFaceIdentityCount = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) SetWorkspaceId(v string) *AsyncUploadVideoShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncUploadVideoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
