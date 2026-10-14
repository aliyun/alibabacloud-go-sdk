// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelC2paBucketSignTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *CancelC2paBucketSignTaskRequest
	GetTaskId() *string
}

type CancelC2paBucketSignTaskRequest struct {
	// The parent task ID of the bucket returned by CreateC2paBucketSignTask.
	//
	// This parameter is required.
	//
	// example:
	//
	// 047b8305-9572-4eed-b7fb-8588b989b6e5
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelC2paBucketSignTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelC2paBucketSignTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelC2paBucketSignTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CancelC2paBucketSignTaskRequest) SetTaskId(v string) *CancelC2paBucketSignTaskRequest {
	s.TaskId = &v
	return s
}

func (s *CancelC2paBucketSignTaskRequest) Validate() error {
	return dara.Validate(s)
}
