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
	// This parameter is required.
	//
	// example:
	//
	// ****ae91539d46bb9000f74b40b80dd2,****ae91539000f74b40b80dd9d46bb
	ProjectIds *string `json:"ProjectIds,omitempty" xml:"ProjectIds,omitempty"`
	UserData   *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
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
