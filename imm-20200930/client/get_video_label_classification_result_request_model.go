// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoLabelClassificationResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetVideoLabelClassificationResultRequest
	GetProjectName() *string
	SetTaskId(v string) *GetVideoLabelClassificationResultRequest
	GetTaskId() *string
	SetTaskType(v string) *GetVideoLabelClassificationResultRequest
	GetTaskType() *string
}

type GetVideoLabelClassificationResultRequest struct {
	// The name of the project. For more information, see [CreateProject](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The task ID, which is obtained from response parameters of [CreateVideoLabelClassificationTask](https://help.aliyun.com/document_detail/478223.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// VideoLabelClassification-2f157087-91df-4fda-8c3e-232407ec****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task. Valid values:
	//
	// This parameter is required.
	//
	// example:
	//
	// VideoLabelClassification
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetVideoLabelClassificationResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoLabelClassificationResultRequest) GoString() string {
	return s.String()
}

func (s *GetVideoLabelClassificationResultRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetVideoLabelClassificationResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetVideoLabelClassificationResultRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetVideoLabelClassificationResultRequest) SetProjectName(v string) *GetVideoLabelClassificationResultRequest {
	s.ProjectName = &v
	return s
}

func (s *GetVideoLabelClassificationResultRequest) SetTaskId(v string) *GetVideoLabelClassificationResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetVideoLabelClassificationResultRequest) SetTaskType(v string) *GetVideoLabelClassificationResultRequest {
	s.TaskType = &v
	return s
}

func (s *GetVideoLabelClassificationResultRequest) Validate() error {
	return dara.Validate(s)
}
