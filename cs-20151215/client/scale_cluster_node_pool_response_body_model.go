// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleClusterNodePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *ScaleClusterNodePoolResponseBody
	GetTaskId() *string
}

type ScaleClusterNodePoolResponseBody struct {
	// The task ID.
	//
	// example:
	//
	// T-5faa48fb31b6b8078d00****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ScaleClusterNodePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleClusterNodePoolResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleClusterNodePoolResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ScaleClusterNodePoolResponseBody) SetTaskId(v string) *ScaleClusterNodePoolResponseBody {
	s.TaskId = &v
	return s
}

func (s *ScaleClusterNodePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
