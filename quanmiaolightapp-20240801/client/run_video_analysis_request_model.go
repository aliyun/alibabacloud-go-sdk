// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoAnalysisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExcludeGenerateOptions(v []*string) *RunVideoAnalysisRequest
	GetExcludeGenerateOptions() []*string
	SetFaceIdentitySimilarityMinScore(v float32) *RunVideoAnalysisRequest
	GetFaceIdentitySimilarityMinScore() *float32
	SetFrameSampleMethod(v *RunVideoAnalysisRequestFrameSampleMethod) *RunVideoAnalysisRequest
	GetFrameSampleMethod() *RunVideoAnalysisRequestFrameSampleMethod
	SetGenerateOptions(v []*string) *RunVideoAnalysisRequest
	GetGenerateOptions() []*string
	SetLanguage(v string) *RunVideoAnalysisRequest
	GetLanguage() *string
	SetModelCustomPromptTemplate(v string) *RunVideoAnalysisRequest
	GetModelCustomPromptTemplate() *string
	SetModelCustomPromptTemplateId(v string) *RunVideoAnalysisRequest
	GetModelCustomPromptTemplateId() *string
	SetModelId(v string) *RunVideoAnalysisRequest
	GetModelId() *string
	SetOriginalSessionId(v string) *RunVideoAnalysisRequest
	GetOriginalSessionId() *string
	SetSnapshotInterval(v float64) *RunVideoAnalysisRequest
	GetSnapshotInterval() *float64
	SetSplitInterval(v int32) *RunVideoAnalysisRequest
	GetSplitInterval() *int32
	SetTaskId(v string) *RunVideoAnalysisRequest
	GetTaskId() *string
	SetTextProcessTasks(v []*RunVideoAnalysisRequestTextProcessTasks) *RunVideoAnalysisRequest
	GetTextProcessTasks() []*RunVideoAnalysisRequestTextProcessTasks
	SetVideoCaptionInfo(v *RunVideoAnalysisRequestVideoCaptionInfo) *RunVideoAnalysisRequest
	GetVideoCaptionInfo() *RunVideoAnalysisRequestVideoCaptionInfo
	SetVideoExtraInfo(v string) *RunVideoAnalysisRequest
	GetVideoExtraInfo() *string
	SetVideoModelCustomPromptTemplate(v string) *RunVideoAnalysisRequest
	GetVideoModelCustomPromptTemplate() *string
	SetVideoModelId(v string) *RunVideoAnalysisRequest
	GetVideoModelId() *string
	SetVideoRoles(v []*RunVideoAnalysisRequestVideoRoles) *RunVideoAnalysisRequest
	GetVideoRoles() []*RunVideoAnalysisRequestVideoRoles
	SetVideoShotFaceIdentityCount(v int32) *RunVideoAnalysisRequest
	GetVideoShotFaceIdentityCount() *int32
	SetVideoUrl(v string) *RunVideoAnalysisRequest
	GetVideoUrl() *string
}

type RunVideoAnalysisRequest struct {
	ExcludeGenerateOptions         []*string                                 `json:"excludeGenerateOptions,omitempty" xml:"excludeGenerateOptions,omitempty" type:"Repeated"`
	FaceIdentitySimilarityMinScore *float32                                  `json:"faceIdentitySimilarityMinScore,omitempty" xml:"faceIdentitySimilarityMinScore,omitempty"`
	FrameSampleMethod              *RunVideoAnalysisRequestFrameSampleMethod `json:"frameSampleMethod,omitempty" xml:"frameSampleMethod,omitempty" type:"Struct"`
	GenerateOptions                []*string                                 `json:"generateOptions,omitempty" xml:"generateOptions,omitempty" type:"Repeated"`
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
	// a3d1c2ac-f086-4a21-9069-f5631542f5a2
	TaskId                         *string                                    `json:"taskId,omitempty" xml:"taskId,omitempty"`
	TextProcessTasks               []*RunVideoAnalysisRequestTextProcessTasks `json:"textProcessTasks,omitempty" xml:"textProcessTasks,omitempty" type:"Repeated"`
	VideoCaptionInfo               *RunVideoAnalysisRequestVideoCaptionInfo   `json:"videoCaptionInfo,omitempty" xml:"videoCaptionInfo,omitempty" type:"Struct"`
	VideoExtraInfo                 *string                                    `json:"videoExtraInfo,omitempty" xml:"videoExtraInfo,omitempty"`
	VideoModelCustomPromptTemplate *string                                    `json:"videoModelCustomPromptTemplate,omitempty" xml:"videoModelCustomPromptTemplate,omitempty"`
	// example:
	//
	// qwen-vl-max
	VideoModelId               *string                              `json:"videoModelId,omitempty" xml:"videoModelId,omitempty"`
	VideoRoles                 []*RunVideoAnalysisRequestVideoRoles `json:"videoRoles,omitempty" xml:"videoRoles,omitempty" type:"Repeated"`
	VideoShotFaceIdentityCount *int32                               `json:"videoShotFaceIdentityCount,omitempty" xml:"videoShotFaceIdentityCount,omitempty"`
	// example:
	//
	// http://xxxx.mp4
	VideoUrl *string `json:"videoUrl,omitempty" xml:"videoUrl,omitempty"`
}

func (s RunVideoAnalysisRequest) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisRequest) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequest) GetExcludeGenerateOptions() []*string {
	return s.ExcludeGenerateOptions
}

func (s *RunVideoAnalysisRequest) GetFaceIdentitySimilarityMinScore() *float32 {
	return s.FaceIdentitySimilarityMinScore
}

func (s *RunVideoAnalysisRequest) GetFrameSampleMethod() *RunVideoAnalysisRequestFrameSampleMethod {
	return s.FrameSampleMethod
}

func (s *RunVideoAnalysisRequest) GetGenerateOptions() []*string {
	return s.GenerateOptions
}

func (s *RunVideoAnalysisRequest) GetLanguage() *string {
	return s.Language
}

func (s *RunVideoAnalysisRequest) GetModelCustomPromptTemplate() *string {
	return s.ModelCustomPromptTemplate
}

func (s *RunVideoAnalysisRequest) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *RunVideoAnalysisRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoAnalysisRequest) GetOriginalSessionId() *string {
	return s.OriginalSessionId
}

func (s *RunVideoAnalysisRequest) GetSnapshotInterval() *float64 {
	return s.SnapshotInterval
}

func (s *RunVideoAnalysisRequest) GetSplitInterval() *int32 {
	return s.SplitInterval
}

func (s *RunVideoAnalysisRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunVideoAnalysisRequest) GetTextProcessTasks() []*RunVideoAnalysisRequestTextProcessTasks {
	return s.TextProcessTasks
}

func (s *RunVideoAnalysisRequest) GetVideoCaptionInfo() *RunVideoAnalysisRequestVideoCaptionInfo {
	return s.VideoCaptionInfo
}

func (s *RunVideoAnalysisRequest) GetVideoExtraInfo() *string {
	return s.VideoExtraInfo
}

func (s *RunVideoAnalysisRequest) GetVideoModelCustomPromptTemplate() *string {
	return s.VideoModelCustomPromptTemplate
}

func (s *RunVideoAnalysisRequest) GetVideoModelId() *string {
	return s.VideoModelId
}

func (s *RunVideoAnalysisRequest) GetVideoRoles() []*RunVideoAnalysisRequestVideoRoles {
	return s.VideoRoles
}

func (s *RunVideoAnalysisRequest) GetVideoShotFaceIdentityCount() *int32 {
	return s.VideoShotFaceIdentityCount
}

func (s *RunVideoAnalysisRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *RunVideoAnalysisRequest) SetExcludeGenerateOptions(v []*string) *RunVideoAnalysisRequest {
	s.ExcludeGenerateOptions = v
	return s
}

func (s *RunVideoAnalysisRequest) SetFaceIdentitySimilarityMinScore(v float32) *RunVideoAnalysisRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetFrameSampleMethod(v *RunVideoAnalysisRequestFrameSampleMethod) *RunVideoAnalysisRequest {
	s.FrameSampleMethod = v
	return s
}

func (s *RunVideoAnalysisRequest) SetGenerateOptions(v []*string) *RunVideoAnalysisRequest {
	s.GenerateOptions = v
	return s
}

func (s *RunVideoAnalysisRequest) SetLanguage(v string) *RunVideoAnalysisRequest {
	s.Language = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetModelCustomPromptTemplate(v string) *RunVideoAnalysisRequest {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetModelCustomPromptTemplateId(v string) *RunVideoAnalysisRequest {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetModelId(v string) *RunVideoAnalysisRequest {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetOriginalSessionId(v string) *RunVideoAnalysisRequest {
	s.OriginalSessionId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetSnapshotInterval(v float64) *RunVideoAnalysisRequest {
	s.SnapshotInterval = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetSplitInterval(v int32) *RunVideoAnalysisRequest {
	s.SplitInterval = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetTaskId(v string) *RunVideoAnalysisRequest {
	s.TaskId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetTextProcessTasks(v []*RunVideoAnalysisRequestTextProcessTasks) *RunVideoAnalysisRequest {
	s.TextProcessTasks = v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoCaptionInfo(v *RunVideoAnalysisRequestVideoCaptionInfo) *RunVideoAnalysisRequest {
	s.VideoCaptionInfo = v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoExtraInfo(v string) *RunVideoAnalysisRequest {
	s.VideoExtraInfo = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoModelCustomPromptTemplate(v string) *RunVideoAnalysisRequest {
	s.VideoModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoModelId(v string) *RunVideoAnalysisRequest {
	s.VideoModelId = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoRoles(v []*RunVideoAnalysisRequestVideoRoles) *RunVideoAnalysisRequest {
	s.VideoRoles = v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoShotFaceIdentityCount(v int32) *RunVideoAnalysisRequest {
	s.VideoShotFaceIdentityCount = &v
	return s
}

func (s *RunVideoAnalysisRequest) SetVideoUrl(v string) *RunVideoAnalysisRequest {
	s.VideoUrl = &v
	return s
}

func (s *RunVideoAnalysisRequest) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisRequestFrameSampleMethod struct {
	Interval   *float64 `json:"interval,omitempty" xml:"interval,omitempty"`
	MethodName *string  `json:"methodName,omitempty" xml:"methodName,omitempty"`
	Pixel      *int32   `json:"pixel,omitempty" xml:"pixel,omitempty"`
}

func (s RunVideoAnalysisRequestFrameSampleMethod) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisRequestFrameSampleMethod) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) GetInterval() *float64 {
	return s.Interval
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) GetMethodName() *string {
	return s.MethodName
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) GetPixel() *int32 {
	return s.Pixel
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) SetInterval(v float64) *RunVideoAnalysisRequestFrameSampleMethod {
	s.Interval = &v
	return s
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) SetMethodName(v string) *RunVideoAnalysisRequestFrameSampleMethod {
	s.MethodName = &v
	return s
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) SetPixel(v int32) *RunVideoAnalysisRequestFrameSampleMethod {
	s.Pixel = &v
	return s
}

func (s *RunVideoAnalysisRequestFrameSampleMethod) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisRequestTextProcessTasks struct {
	ModelCustomPromptTemplate   *string `json:"modelCustomPromptTemplate,omitempty" xml:"modelCustomPromptTemplate,omitempty"`
	ModelCustomPromptTemplateId *string `json:"modelCustomPromptTemplateId,omitempty" xml:"modelCustomPromptTemplateId,omitempty"`
	ModelId                     *string `json:"modelId,omitempty" xml:"modelId,omitempty"`
}

func (s RunVideoAnalysisRequestTextProcessTasks) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisRequestTextProcessTasks) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequestTextProcessTasks) GetModelCustomPromptTemplate() *string {
	return s.ModelCustomPromptTemplate
}

func (s *RunVideoAnalysisRequestTextProcessTasks) GetModelCustomPromptTemplateId() *string {
	return s.ModelCustomPromptTemplateId
}

func (s *RunVideoAnalysisRequestTextProcessTasks) GetModelId() *string {
	return s.ModelId
}

func (s *RunVideoAnalysisRequestTextProcessTasks) SetModelCustomPromptTemplate(v string) *RunVideoAnalysisRequestTextProcessTasks {
	s.ModelCustomPromptTemplate = &v
	return s
}

func (s *RunVideoAnalysisRequestTextProcessTasks) SetModelCustomPromptTemplateId(v string) *RunVideoAnalysisRequestTextProcessTasks {
	s.ModelCustomPromptTemplateId = &v
	return s
}

func (s *RunVideoAnalysisRequestTextProcessTasks) SetModelId(v string) *RunVideoAnalysisRequestTextProcessTasks {
	s.ModelId = &v
	return s
}

func (s *RunVideoAnalysisRequestTextProcessTasks) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisRequestVideoCaptionInfo struct {
	// example:
	//
	// oss:// | http://
	VideoCaptionFileUrl *string                                                 `json:"videoCaptionFileUrl,omitempty" xml:"videoCaptionFileUrl,omitempty"`
	VideoCaptions       []*RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions `json:"videoCaptions,omitempty" xml:"videoCaptions,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisRequestVideoCaptionInfo) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisRequestVideoCaptionInfo) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequestVideoCaptionInfo) GetVideoCaptionFileUrl() *string {
	return s.VideoCaptionFileUrl
}

func (s *RunVideoAnalysisRequestVideoCaptionInfo) GetVideoCaptions() []*RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions {
	return s.VideoCaptions
}

func (s *RunVideoAnalysisRequestVideoCaptionInfo) SetVideoCaptionFileUrl(v string) *RunVideoAnalysisRequestVideoCaptionInfo {
	s.VideoCaptionFileUrl = &v
	return s
}

func (s *RunVideoAnalysisRequestVideoCaptionInfo) SetVideoCaptions(v []*RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) *RunVideoAnalysisRequestVideoCaptionInfo {
	s.VideoCaptions = v
	return s
}

func (s *RunVideoAnalysisRequestVideoCaptionInfo) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions struct {
	// example:
	//
	// 10000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 张三
	Speaker *string `json:"speaker,omitempty" xml:"speaker,omitempty"`
	// example:
	//
	// 1000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 你好
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) GetEndTime() *int64 {
	return s.EndTime
}

func (s *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) GetSpeaker() *string {
	return s.Speaker
}

func (s *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) GetStartTime() *int64 {
	return s.StartTime
}

func (s *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) GetText() *string {
	return s.Text
}

func (s *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) SetEndTime(v int64) *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions {
	s.EndTime = &v
	return s
}

func (s *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) SetSpeaker(v string) *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions {
	s.Speaker = &v
	return s
}

func (s *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) SetStartTime(v int64) *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions {
	s.StartTime = &v
	return s
}

func (s *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) SetText(v string) *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions {
	s.Text = &v
	return s
}

func (s *RunVideoAnalysisRequestVideoCaptionInfoVideoCaptions) Validate() error {
	return dara.Validate(s)
}

type RunVideoAnalysisRequestVideoRoles struct {
	RoleInfo *string   `json:"roleInfo,omitempty" xml:"roleInfo,omitempty"`
	RoleName *string   `json:"roleName,omitempty" xml:"roleName,omitempty"`
	Urls     []*string `json:"urls,omitempty" xml:"urls,omitempty" type:"Repeated"`
}

func (s RunVideoAnalysisRequestVideoRoles) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisRequestVideoRoles) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisRequestVideoRoles) GetRoleInfo() *string {
	return s.RoleInfo
}

func (s *RunVideoAnalysisRequestVideoRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *RunVideoAnalysisRequestVideoRoles) GetUrls() []*string {
	return s.Urls
}

func (s *RunVideoAnalysisRequestVideoRoles) SetRoleInfo(v string) *RunVideoAnalysisRequestVideoRoles {
	s.RoleInfo = &v
	return s
}

func (s *RunVideoAnalysisRequestVideoRoles) SetRoleName(v string) *RunVideoAnalysisRequestVideoRoles {
	s.RoleName = &v
	return s
}

func (s *RunVideoAnalysisRequestVideoRoles) SetUrls(v []*string) *RunVideoAnalysisRequestVideoRoles {
	s.Urls = v
	return s
}

func (s *RunVideoAnalysisRequestVideoRoles) Validate() error {
	return dara.Validate(s)
}
