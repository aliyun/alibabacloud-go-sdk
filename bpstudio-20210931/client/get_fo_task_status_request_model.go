// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFoTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v int32) *GetFoTaskStatusRequest
	GetTaskId() *int32
}

type GetFoTaskStatusRequest struct {
	// The disaster recovery switchover task ID.
	//
	// example:
	//
	// 2615
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetFoTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFoTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetFoTaskStatusRequest) GetTaskId() *int32 {
	return s.TaskId
}

func (s *GetFoTaskStatusRequest) SetTaskId(v int32) *GetFoTaskStatusRequest {
	s.TaskId = &v
	return s
}

func (s *GetFoTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
