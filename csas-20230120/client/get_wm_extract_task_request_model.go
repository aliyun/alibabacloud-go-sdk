// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWmExtractTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetWmExtractTaskRequest
	GetTaskId() *string
}

type GetWmExtractTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// wmt-9648c22d2eb2cb57bb855dcae7898464********
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetWmExtractTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWmExtractTaskRequest) GoString() string {
	return s.String()
}

func (s *GetWmExtractTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetWmExtractTaskRequest) SetTaskId(v string) *GetWmExtractTaskRequest {
	s.TaskId = &v
	return s
}

func (s *GetWmExtractTaskRequest) Validate() error {
	return dara.Validate(s)
}
