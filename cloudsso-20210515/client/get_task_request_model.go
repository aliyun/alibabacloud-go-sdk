// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetTaskRequest
	GetDirectoryId() *string
	SetTaskId(v string) *GetTaskRequest
	GetTaskId() *string
}

type GetTaskRequest struct {
	// The directory ID.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-shfqw1u1edszvxw5****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskRequest) SetDirectoryId(v string) *GetTaskRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetTaskRequest) SetTaskId(v string) *GetTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskRequest) Validate() error {
	return dara.Validate(s)
}
