// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseC2paBucketSignTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *PauseC2paBucketSignTaskRequest
	GetTaskId() *string
}

type PauseC2paBucketSignTaskRequest struct {
	// The bucket parent task ID returned by CreateC2paBucketSignTask.
	//
	// This parameter is required.
	//
	// example:
	//
	// 047b8305-9572-4eed-b7fb-8588b989b6e5
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PauseC2paBucketSignTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseC2paBucketSignTaskRequest) GoString() string {
	return s.String()
}

func (s *PauseC2paBucketSignTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *PauseC2paBucketSignTaskRequest) SetTaskId(v string) *PauseC2paBucketSignTaskRequest {
	s.TaskId = &v
	return s
}

func (s *PauseC2paBucketSignTaskRequest) Validate() error {
	return dara.Validate(s)
}
