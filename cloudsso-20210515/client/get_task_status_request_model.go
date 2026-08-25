// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *GetTaskStatusRequest
	GetDirectoryId() *string
	SetTaskId(v string) *GetTaskStatusRequest
	GetTaskId() *string
}

type GetTaskStatusRequest struct {
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

func (s GetTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskStatusRequest) SetDirectoryId(v string) *GetTaskStatusRequest {
	s.DirectoryId = &v
	return s
}

func (s *GetTaskStatusRequest) SetTaskId(v string) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
