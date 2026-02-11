// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociatePocTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProjectId(v int64) *AssociatePocTaskRequest
	GetProjectId() *int64
	SetTaskId(v int64) *AssociatePocTaskRequest
	GetTaskId() *int64
}

type AssociatePocTaskRequest struct {
	// Project ID
	//
	// example:
	//
	// 01
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// Retrospective task ID.
	//
	// example:
	//
	// 01
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s AssociatePocTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociatePocTaskRequest) GoString() string {
	return s.String()
}

func (s *AssociatePocTaskRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *AssociatePocTaskRequest) GetTaskId() *int64 {
	return s.TaskId
}

func (s *AssociatePocTaskRequest) SetProjectId(v int64) *AssociatePocTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *AssociatePocTaskRequest) SetTaskId(v int64) *AssociatePocTaskRequest {
	s.TaskId = &v
	return s
}

func (s *AssociatePocTaskRequest) Validate() error {
	return dara.Validate(s)
}
