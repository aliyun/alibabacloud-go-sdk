// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetTaskRequest
	GetProjectName() *string
	SetRequestDefinition(v bool) *GetTaskRequest
	GetRequestDefinition() *bool
	SetTaskId(v string) *GetTaskRequest
	GetTaskId() *string
	SetTaskType(v string) *GetTaskRequest
	GetTaskType() *string
}

type GetTaskRequest struct {
	// The project name. For information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to return the original request parameters used to create the task. Valid values:
	//
	// - true
	//
	// - false (default)
	//
	// This parameter takes effect only for the following task types:
	//
	// - MediaConvert
	//
	// - VideoLabelClassification
	//
	// - FaceClustering
	//
	// - FileCompression
	//
	// - ArchiveFileInspection
	//
	// - FileUncompression
	//
	// - PointCloudCompress
	//
	// - ImageToPDF
	//
	// - StoryCreation
	//
	// - LocationDateClustering
	//
	// - ImageSplicing
	//
	// - FacesSearching.
	//
	// example:
	//
	// true
	RequestDefinition *bool `json:"RequestDefinition,omitempty" xml:"RequestDefinition,omitempty"`
	// The ID of the task that you want to query. This value is returned when you create the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// FileCompression-2f157087-91df-4fda-8c3e-232407ec*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task. For valid values, see [Task type list](https://help.aliyun.com/document_detail/2743993.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// FileCompression
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetTaskRequest) GetRequestDefinition() *bool {
	return s.RequestDefinition
}

func (s *GetTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTaskRequest) SetProjectName(v string) *GetTaskRequest {
	s.ProjectName = &v
	return s
}

func (s *GetTaskRequest) SetRequestDefinition(v bool) *GetTaskRequest {
	s.RequestDefinition = &v
	return s
}

func (s *GetTaskRequest) SetTaskId(v string) *GetTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskRequest) SetTaskType(v string) *GetTaskRequest {
	s.TaskType = &v
	return s
}

func (s *GetTaskRequest) Validate() error {
	return dara.Validate(s)
}
