// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSceneBatchEditingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutputConfig(v string) *SubmitSceneBatchEditingJobRequest
	GetOutputConfig() *string
	SetProjectIds(v string) *SubmitSceneBatchEditingJobRequest
	GetProjectIds() *string
	SetUserData(v string) *SubmitSceneBatchEditingJobRequest
	GetUserData() *string
}

type SubmitSceneBatchEditingJobRequest struct {
	// The output configuration. The structure is the same as the [OutputConfig](https://help.aliyun.com/zh/ims/use-cases/create-highlight-videos?spm=a2c4g.11186623.help-menu-193643.d_3_2_0_3.3af86997GreVu9\\&scm=20140722.H_2863940._.OR_help-T_cn~zh-V_1#4111a373d0xbz) for batch video generation, except that Count and GeneratePreviewOnly are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "MediaURL": "http://test-bucket.oss-cn-shanghai.aliyuncs.com/xxx_{index}.mp4",
	//
	//   "Width": 1080,
	//
	//   "Height": 1920
	//
	// }
	OutputConfig *string `json:"OutputConfig,omitempty" xml:"OutputConfig,omitempty"`
	// A comma-separated list of editing project IDs. The video is rendered based on the timeline from each project.
	//
	// This parameter is required.
	//
	// example:
	//
	// ****ae91539d46bb9000f74b40b80dd2,****ae91539000f74b40b80dd9d46bb
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	// Custom user data, including callback configurations. For more information, see [UserData](~~357745#section-urj-v3f-0s1~~).
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitSceneBatchEditingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSceneBatchEditingJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitSceneBatchEditingJobRequest) GetOutputConfig() *string {
	return s.OutputConfig
}

func (s *SubmitSceneBatchEditingJobRequest) GetProjectIds() *string {
	return s.ProjectIds
}

func (s *SubmitSceneBatchEditingJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitSceneBatchEditingJobRequest) SetOutputConfig(v string) *SubmitSceneBatchEditingJobRequest {
	s.OutputConfig = &v
	return s
}

func (s *SubmitSceneBatchEditingJobRequest) SetProjectIds(v string) *SubmitSceneBatchEditingJobRequest {
	s.ProjectIds = &v
	return s
}

func (s *SubmitSceneBatchEditingJobRequest) SetUserData(v string) *SubmitSceneBatchEditingJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitSceneBatchEditingJobRequest) Validate() error {
	return dara.Validate(s)
}
