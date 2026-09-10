// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*int64) *DeleteDataCheckTaskRequest
	GetTaskIds() []*int64
}

type DeleteDataCheckTaskRequest struct {
	// The list of task IDs. Batch deletion is supported.
	//
	// This parameter is required.
	TaskIds []*int64 `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s DeleteDataCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataCheckTaskRequest) GetTaskIds() []*int64 {
	return s.TaskIds
}

func (s *DeleteDataCheckTaskRequest) SetTaskIds(v []*int64) *DeleteDataCheckTaskRequest {
	s.TaskIds = v
	return s
}

func (s *DeleteDataCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
