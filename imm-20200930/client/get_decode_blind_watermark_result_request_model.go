// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDecodeBlindWatermarkResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectName(v string) *GetDecodeBlindWatermarkResultRequest
	GetProjectName() *string
	SetTaskId(v string) *GetDecodeBlindWatermarkResultRequest
	GetTaskId() *string
	SetTaskType(v string) *GetDecodeBlindWatermarkResultRequest
	GetTaskType() *string
}

type GetDecodeBlindWatermarkResultRequest struct {
	// The project name. For information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The task ID. You can obtain the task ID from the response parameters of the blind watermark extraction task creation operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// DecodeBlindWatermark-c09b0943-ed79-4983-8dbe-7a882574****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The task type.
	//
	// This parameter is required.
	//
	// example:
	//
	// DecodeBlindWatermark
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetDecodeBlindWatermarkResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDecodeBlindWatermarkResultRequest) GoString() string {
	return s.String()
}

func (s *GetDecodeBlindWatermarkResultRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetDecodeBlindWatermarkResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDecodeBlindWatermarkResultRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetDecodeBlindWatermarkResultRequest) SetProjectName(v string) *GetDecodeBlindWatermarkResultRequest {
	s.ProjectName = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultRequest) SetTaskId(v string) *GetDecodeBlindWatermarkResultRequest {
	s.TaskId = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultRequest) SetTaskType(v string) *GetDecodeBlindWatermarkResultRequest {
	s.TaskType = &v
	return s
}

func (s *GetDecodeBlindWatermarkResultRequest) Validate() error {
	return dara.Validate(s)
}
