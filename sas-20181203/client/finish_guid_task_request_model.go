// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishGuidTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskTypeName(v string) *FinishGuidTaskRequest
	GetTaskTypeName() *string
}

type FinishGuidTaskRequest struct {
	// The name of the task type.
	//
	// This parameter is required.
	//
	// example:
	//
	// guide_sub_task_config_add_collection
	TaskTypeName *string `json:"TaskTypeName,omitempty" xml:"TaskTypeName,omitempty"`
}

func (s FinishGuidTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s FinishGuidTaskRequest) GoString() string {
	return s.String()
}

func (s *FinishGuidTaskRequest) GetTaskTypeName() *string {
	return s.TaskTypeName
}

func (s *FinishGuidTaskRequest) SetTaskTypeName(v string) *FinishGuidTaskRequest {
	s.TaskTypeName = &v
	return s
}

func (s *FinishGuidTaskRequest) Validate() error {
	return dara.Validate(s)
}
