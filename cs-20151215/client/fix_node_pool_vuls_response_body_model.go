// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFixNodePoolVulsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *FixNodePoolVulsResponseBody
	GetTaskId() *string
}

type FixNodePoolVulsResponseBody struct {
	// The ID of the CVE patching task.
	//
	// example:
	//
	// T-60fea8ad2e277f087900****
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s FixNodePoolVulsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FixNodePoolVulsResponseBody) GoString() string {
	return s.String()
}

func (s *FixNodePoolVulsResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *FixNodePoolVulsResponseBody) SetTaskId(v string) *FixNodePoolVulsResponseBody {
	s.TaskId = &v
	return s
}

func (s *FixNodePoolVulsResponseBody) Validate() error {
	return dara.Validate(s)
}
