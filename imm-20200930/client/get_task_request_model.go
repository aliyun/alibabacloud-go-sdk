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
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Specifies whether to return original request parameters specified to create the task.
	//
	// 	- true
	//
	// 	- false (default)
	//
	// This parameter applies only to the following tasks:
	//
	// 	- MediaConvert
	//
	// 	- VideoLabelClassification
	//
	// 	- FaceClustering
	//
	// 	- FileCompression
	//
	// 	- ArchiveFileInspection
	//
	// 	- FileUncompression
	//
	// 	- PointCloudCompress
	//
	// 	- ImageToPDF
	//
	// 	- StoryCreation
	//
	// 	- LocationDateClustering
	//
	// 	- ImageSplicing
	//
	// 	- FacesSearching
	//
	// example:
	//
	// True
	RequestDefinition *bool `json:"RequestDefinition,omitempty" xml:"RequestDefinition,omitempty"`
	// The ID of the task. You can obtain the ID of a task after you create the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// c2b277b9-0d30-4882-ad6d-ad661382****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The type of the task. For information about valid values, see [Task types](https://help.aliyun.com/document_detail/2743993.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// VideoLabelClassification
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
