// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceDelete(v bool) *DeleteJobRequest
	GetForceDelete() *bool
}

type DeleteJobRequest struct {
	// Specifies whether to force delete the subtask. If the task has subtasks and you set this parameter to true, the task and its subtasks are forcibly deleted. If this parameter is set to false, the task and its subtasks fail to be deleted.
	//
	// example:
	//
	// true
	ForceDelete *bool `json:"forceDelete,omitempty" xml:"forceDelete,omitempty"`
}

func (s DeleteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRequest) GetForceDelete() *bool {
	return s.ForceDelete
}

func (s *DeleteJobRequest) SetForceDelete(v bool) *DeleteJobRequest {
	s.ForceDelete = &v
	return s
}

func (s *DeleteJobRequest) Validate() error {
	return dara.Validate(s)
}
