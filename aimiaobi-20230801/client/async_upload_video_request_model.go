// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncUploadVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdaptiveThreshold(v float32) *AsyncUploadVideoRequest
	GetAdaptiveThreshold() *float32
	SetAnlysisPrompt(v string) *AsyncUploadVideoRequest
	GetAnlysisPrompt() *string
	SetFaceIdentitySimilarityMinScore(v float64) *AsyncUploadVideoRequest
	GetFaceIdentitySimilarityMinScore() *float64
	SetReferenceVideo(v *AsyncUploadVideoRequestReferenceVideo) *AsyncUploadVideoRequest
	GetReferenceVideo() *AsyncUploadVideoRequestReferenceVideo
	SetRemoveSubtitle(v bool) *AsyncUploadVideoRequest
	GetRemoveSubtitle() *bool
	SetSourceVideos(v []*AsyncUploadVideoRequestSourceVideos) *AsyncUploadVideoRequest
	GetSourceVideos() []*AsyncUploadVideoRequestSourceVideos
	SetSplitInterval(v int32) *AsyncUploadVideoRequest
	GetSplitInterval() *int32
	SetTaskName(v string) *AsyncUploadVideoRequest
	GetTaskName() *string
	SetTaskType(v string) *AsyncUploadVideoRequest
	GetTaskType() *string
	SetVideoRoles(v []*AsyncUploadVideoRequestVideoRoles) *AsyncUploadVideoRequest
	GetVideoRoles() []*AsyncUploadVideoRequestVideoRoles
	SetVideoShotFaceIdentityCount(v int32) *AsyncUploadVideoRequest
	GetVideoShotFaceIdentityCount() *int32
	SetWorkspaceId(v string) *AsyncUploadVideoRequest
	GetWorkspaceId() *string
}

type AsyncUploadVideoRequest struct {
	// example:
	//
	// 3
	AdaptiveThreshold *float32 `json:"AdaptiveThreshold,omitempty" xml:"AdaptiveThreshold,omitempty"`
	// example:
	//
	// 重点理解视频中的风景信息
	AnlysisPrompt *string `json:"AnlysisPrompt,omitempty" xml:"AnlysisPrompt,omitempty"`
	// example:
	//
	// 0.7
	FaceIdentitySimilarityMinScore *float64                               `json:"FaceIdentitySimilarityMinScore,omitempty" xml:"FaceIdentitySimilarityMinScore,omitempty"`
	ReferenceVideo                 *AsyncUploadVideoRequestReferenceVideo `json:"ReferenceVideo,omitempty" xml:"ReferenceVideo,omitempty" type:"Struct"`
	RemoveSubtitle                 *bool                                  `json:"RemoveSubtitle,omitempty" xml:"RemoveSubtitle,omitempty"`
	// This parameter is required.
	SourceVideos []*AsyncUploadVideoRequestSourceVideos `json:"SourceVideos,omitempty" xml:"SourceVideos,omitempty" type:"Repeated"`
	// example:
	//
	// 默认1
	SplitInterval *int32 `json:"SplitInterval,omitempty" xml:"SplitInterval,omitempty"`
	// example:
	//
	// task001
	TaskName *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	// example:
	//
	// type001
	TaskType   *string                              `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	VideoRoles []*AsyncUploadVideoRequestVideoRoles `json:"VideoRoles,omitempty" xml:"VideoRoles,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	VideoShotFaceIdentityCount *int32 `json:"VideoShotFaceIdentityCount,omitempty" xml:"VideoShotFaceIdentityCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s AsyncUploadVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoRequest) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoRequest) GetAdaptiveThreshold() *float32 {
	return s.AdaptiveThreshold
}

func (s *AsyncUploadVideoRequest) GetAnlysisPrompt() *string {
	return s.AnlysisPrompt
}

func (s *AsyncUploadVideoRequest) GetFaceIdentitySimilarityMinScore() *float64 {
	return s.FaceIdentitySimilarityMinScore
}

func (s *AsyncUploadVideoRequest) GetReferenceVideo() *AsyncUploadVideoRequestReferenceVideo {
	return s.ReferenceVideo
}

func (s *AsyncUploadVideoRequest) GetRemoveSubtitle() *bool {
	return s.RemoveSubtitle
}

func (s *AsyncUploadVideoRequest) GetSourceVideos() []*AsyncUploadVideoRequestSourceVideos {
	return s.SourceVideos
}

func (s *AsyncUploadVideoRequest) GetSplitInterval() *int32 {
	return s.SplitInterval
}

func (s *AsyncUploadVideoRequest) GetTaskName() *string {
	return s.TaskName
}

func (s *AsyncUploadVideoRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *AsyncUploadVideoRequest) GetVideoRoles() []*AsyncUploadVideoRequestVideoRoles {
	return s.VideoRoles
}

func (s *AsyncUploadVideoRequest) GetVideoShotFaceIdentityCount() *int32 {
	return s.VideoShotFaceIdentityCount
}

func (s *AsyncUploadVideoRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *AsyncUploadVideoRequest) SetAdaptiveThreshold(v float32) *AsyncUploadVideoRequest {
	s.AdaptiveThreshold = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetAnlysisPrompt(v string) *AsyncUploadVideoRequest {
	s.AnlysisPrompt = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetFaceIdentitySimilarityMinScore(v float64) *AsyncUploadVideoRequest {
	s.FaceIdentitySimilarityMinScore = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetReferenceVideo(v *AsyncUploadVideoRequestReferenceVideo) *AsyncUploadVideoRequest {
	s.ReferenceVideo = v
	return s
}

func (s *AsyncUploadVideoRequest) SetRemoveSubtitle(v bool) *AsyncUploadVideoRequest {
	s.RemoveSubtitle = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetSourceVideos(v []*AsyncUploadVideoRequestSourceVideos) *AsyncUploadVideoRequest {
	s.SourceVideos = v
	return s
}

func (s *AsyncUploadVideoRequest) SetSplitInterval(v int32) *AsyncUploadVideoRequest {
	s.SplitInterval = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetTaskName(v string) *AsyncUploadVideoRequest {
	s.TaskName = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetTaskType(v string) *AsyncUploadVideoRequest {
	s.TaskType = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetVideoRoles(v []*AsyncUploadVideoRequestVideoRoles) *AsyncUploadVideoRequest {
	s.VideoRoles = v
	return s
}

func (s *AsyncUploadVideoRequest) SetVideoShotFaceIdentityCount(v int32) *AsyncUploadVideoRequest {
	s.VideoShotFaceIdentityCount = &v
	return s
}

func (s *AsyncUploadVideoRequest) SetWorkspaceId(v string) *AsyncUploadVideoRequest {
	s.WorkspaceId = &v
	return s
}

func (s *AsyncUploadVideoRequest) Validate() error {
	if s.ReferenceVideo != nil {
		if err := s.ReferenceVideo.Validate(); err != nil {
			return err
		}
	}
	if s.SourceVideos != nil {
		for _, item := range s.SourceVideos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoRoles != nil {
		for _, item := range s.VideoRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AsyncUploadVideoRequestReferenceVideo struct {
	// example:
	//
	// 手机cpu采用3纳米技术
	VideoExtraInfo *string `json:"VideoExtraInfo,omitempty" xml:"VideoExtraInfo,omitempty"`
	// example:
	//
	// refvideo.mp4
	VideoName *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	// example:
	//
	// http://viapi-customer-pop.oss-cn-shanghai.aliyuncs.com/d71e_208334498220521996_51298e0a909d457693166eb883768c7a
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AsyncUploadVideoRequestReferenceVideo) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoRequestReferenceVideo) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoRequestReferenceVideo) GetVideoExtraInfo() *string {
	return s.VideoExtraInfo
}

func (s *AsyncUploadVideoRequestReferenceVideo) GetVideoName() *string {
	return s.VideoName
}

func (s *AsyncUploadVideoRequestReferenceVideo) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AsyncUploadVideoRequestReferenceVideo) SetVideoExtraInfo(v string) *AsyncUploadVideoRequestReferenceVideo {
	s.VideoExtraInfo = &v
	return s
}

func (s *AsyncUploadVideoRequestReferenceVideo) SetVideoName(v string) *AsyncUploadVideoRequestReferenceVideo {
	s.VideoName = &v
	return s
}

func (s *AsyncUploadVideoRequestReferenceVideo) SetVideoUrl(v string) *AsyncUploadVideoRequestReferenceVideo {
	s.VideoUrl = &v
	return s
}

func (s *AsyncUploadVideoRequestReferenceVideo) Validate() error {
	return dara.Validate(s)
}

type AsyncUploadVideoRequestSourceVideos struct {
	// example:
	//
	// 视频中有一个房子
	VideoExtraInfo *string `json:"VideoExtraInfo,omitempty" xml:"VideoExtraInfo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123.mp4
	VideoName *string `json:"VideoName,omitempty" xml:"VideoName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://123.mp4
	//
	// 目前该接口只进行视频理解额和分析，并没有对文件进行转存。请保证视频地址在任务执行期间有效。
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AsyncUploadVideoRequestSourceVideos) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoRequestSourceVideos) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoRequestSourceVideos) GetVideoExtraInfo() *string {
	return s.VideoExtraInfo
}

func (s *AsyncUploadVideoRequestSourceVideos) GetVideoName() *string {
	return s.VideoName
}

func (s *AsyncUploadVideoRequestSourceVideos) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AsyncUploadVideoRequestSourceVideos) SetVideoExtraInfo(v string) *AsyncUploadVideoRequestSourceVideos {
	s.VideoExtraInfo = &v
	return s
}

func (s *AsyncUploadVideoRequestSourceVideos) SetVideoName(v string) *AsyncUploadVideoRequestSourceVideos {
	s.VideoName = &v
	return s
}

func (s *AsyncUploadVideoRequestSourceVideos) SetVideoUrl(v string) *AsyncUploadVideoRequestSourceVideos {
	s.VideoUrl = &v
	return s
}

func (s *AsyncUploadVideoRequestSourceVideos) Validate() error {
	return dara.Validate(s)
}

type AsyncUploadVideoRequestVideoRoles struct {
	// example:
	//
	// 李晓明是一位警察
	RoleInfo *string `json:"RoleInfo,omitempty" xml:"RoleInfo,omitempty"`
	// example:
	//
	// 李晓明
	RoleName *string                                      `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	RoleUrls []*AsyncUploadVideoRequestVideoRolesRoleUrls `json:"RoleUrls,omitempty" xml:"RoleUrls,omitempty" type:"Repeated"`
}

func (s AsyncUploadVideoRequestVideoRoles) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoRequestVideoRoles) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoRequestVideoRoles) GetRoleInfo() *string {
	return s.RoleInfo
}

func (s *AsyncUploadVideoRequestVideoRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *AsyncUploadVideoRequestVideoRoles) GetRoleUrls() []*AsyncUploadVideoRequestVideoRolesRoleUrls {
	return s.RoleUrls
}

func (s *AsyncUploadVideoRequestVideoRoles) SetRoleInfo(v string) *AsyncUploadVideoRequestVideoRoles {
	s.RoleInfo = &v
	return s
}

func (s *AsyncUploadVideoRequestVideoRoles) SetRoleName(v string) *AsyncUploadVideoRequestVideoRoles {
	s.RoleName = &v
	return s
}

func (s *AsyncUploadVideoRequestVideoRoles) SetRoleUrls(v []*AsyncUploadVideoRequestVideoRolesRoleUrls) *AsyncUploadVideoRequestVideoRoles {
	s.RoleUrls = v
	return s
}

func (s *AsyncUploadVideoRequestVideoRoles) Validate() error {
	if s.RoleUrls != nil {
		for _, item := range s.RoleUrls {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AsyncUploadVideoRequestVideoRolesRoleUrls struct {
	// example:
	//
	// 王小明.jpeg
	RoleFileName *string `json:"RoleFileName,omitempty" xml:"RoleFileName,omitempty"`
	// example:
	//
	// http://xxx/xxx.jpeg
	RoleFileUrl *string `json:"RoleFileUrl,omitempty" xml:"RoleFileUrl,omitempty"`
}

func (s AsyncUploadVideoRequestVideoRolesRoleUrls) String() string {
	return dara.Prettify(s)
}

func (s AsyncUploadVideoRequestVideoRolesRoleUrls) GoString() string {
	return s.String()
}

func (s *AsyncUploadVideoRequestVideoRolesRoleUrls) GetRoleFileName() *string {
	return s.RoleFileName
}

func (s *AsyncUploadVideoRequestVideoRolesRoleUrls) GetRoleFileUrl() *string {
	return s.RoleFileUrl
}

func (s *AsyncUploadVideoRequestVideoRolesRoleUrls) SetRoleFileName(v string) *AsyncUploadVideoRequestVideoRolesRoleUrls {
	s.RoleFileName = &v
	return s
}

func (s *AsyncUploadVideoRequestVideoRolesRoleUrls) SetRoleFileUrl(v string) *AsyncUploadVideoRequestVideoRolesRoleUrls {
	s.RoleFileUrl = &v
	return s
}

func (s *AsyncUploadVideoRequestVideoRolesRoleUrls) Validate() error {
	return dara.Validate(s)
}
