// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopySceneTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetCopySceneTaskStatusRequest
	GetTaskId() *string
}

type GetCopySceneTaskStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yuywey****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetCopySceneTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCopySceneTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetCopySceneTaskStatusRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetCopySceneTaskStatusRequest) SetTaskId(v string) *GetCopySceneTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetCopySceneTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
