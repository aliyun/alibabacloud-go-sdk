// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebuildTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskIds(v []*string) *RebuildTaskRequest
	GetTaskIds() []*string
}

type RebuildTaskRequest struct {
	// This parameter is required.
	TaskIds []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s RebuildTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RebuildTaskRequest) GoString() string {
	return s.String()
}

func (s *RebuildTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *RebuildTaskRequest) SetTaskIds(v []*string) *RebuildTaskRequest {
	s.TaskIds = v
	return s
}

func (s *RebuildTaskRequest) Validate() error {
	return dara.Validate(s)
}
