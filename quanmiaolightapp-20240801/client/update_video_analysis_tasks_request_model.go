// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoAnalysisTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *UpdateVideoAnalysisTasksRequest
	GetTaskIds() []*string
	SetTaskStatus(v string) *UpdateVideoAnalysisTasksRequest
	GetTaskStatus() *string
}

type UpdateVideoAnalysisTasksRequest struct {
	// This parameter is required.
	TaskIds []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// CANCELED
	TaskStatus *string `json:"taskStatus,omitempty" xml:"taskStatus,omitempty"`
}

func (s UpdateVideoAnalysisTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoAnalysisTasksRequest) GoString() string {
	return s.String()
}

func (s *UpdateVideoAnalysisTasksRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *UpdateVideoAnalysisTasksRequest) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *UpdateVideoAnalysisTasksRequest) SetTaskIds(v []*string) *UpdateVideoAnalysisTasksRequest {
	s.TaskIds = v
	return s
}

func (s *UpdateVideoAnalysisTasksRequest) SetTaskStatus(v string) *UpdateVideoAnalysisTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *UpdateVideoAnalysisTasksRequest) Validate() error {
	return dara.Validate(s)
}
