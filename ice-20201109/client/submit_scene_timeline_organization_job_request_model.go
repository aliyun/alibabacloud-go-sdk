// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneTimelineOrganizationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEditingConfig(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetEditingConfig() *string
	SetInputConfig(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetInputConfig() *string
	SetJobType(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetJobType() *string
	SetMediaSelectResult(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetMediaSelectResult() *string
	SetOutputConfig(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetOutputConfig() *string
	SetUserData(v string) *SubmitSceneTimelineOrganizationJobRequest
	GetUserData() *string
}

type SubmitSceneTimelineOrganizationJobRequest struct {
	// example:
	//
	// {
	//
	//   "MediaConfig": {
	//
	//       "Volume": 0
	//
	//   },
	//
	//   "SpeechConfig": {
	//
	//       "Volume": 1
	//
	//   },
	//
	//  "BackgroundMusicConfig": {
	//
	//       "Volume": 0.3
	//
	//   }
	//
	// }
	EditingConfig *string `json:"EditingConfig,omitempty" xml:"EditingConfig,omitempty"`
	// This parameter is required.
	InputConfig *string `json:"InputConfig,omitempty" xml:"InputConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Smart_Mix_Timeline_Organize
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// This parameter is required.
	MediaSelectResult *string `json:"MediaSelectResult,omitempty" xml:"MediaSelectResult,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "MediaURL": "http://test-bucket.oss-cn-shanghai.aliyuncs.com/xxx_{index}.mp4",
	//
	//   "Count": 1,
	//
	//   "Width": 1080,
	//
	//   "Height": 1920
	//
	// }
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	UserData     *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSceneTimelineOrganizationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneTimelineOrganizationJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetEditingConfig() *string {
	return s.EditingConfig
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetInputConfig() *string {
	return s.InputConfig
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetJobType() *string {
	return s.JobType
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetMediaSelectResult() *string {
	return s.MediaSelectResult
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitSceneTimelineOrganizationJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetEditingConfig(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.EditingConfig = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetInputConfig(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.InputConfig = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetJobType(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.JobType = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetMediaSelectResult(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.MediaSelectResult = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetOutputConfig(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) SetUserData(v string) *SubmitSceneTimelineOrganizationJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSceneTimelineOrganizationJobRequest) Validate() error {
	return dara.Validate(s)
}
