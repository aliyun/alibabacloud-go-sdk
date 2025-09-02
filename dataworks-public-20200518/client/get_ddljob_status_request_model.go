// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDDLJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetDDLJobStatusRequest
	GetTaskId() *string
}

type GetDDLJobStatusRequest struct {
	// The ID of the DDL task.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetDDLJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDDLJobStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDDLJobStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetDDLJobStatusRequest) SetTaskId(v string) *GetDDLJobStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetDDLJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
