// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *GetTaskAttributeRequest
	GetTaskId() *string
	SetTaskType(v string) *GetTaskAttributeRequest
	GetTaskType() *string
}

type GetTaskAttributeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// ApplyCertificate
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetTaskAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTaskAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetTaskAttributeRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetTaskAttributeRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetTaskAttributeRequest) SetTaskId(v string) *GetTaskAttributeRequest {
	s.TaskId = &v
	return s
}

func (s *GetTaskAttributeRequest) SetTaskType(v string) *GetTaskAttributeRequest {
	s.TaskType = &v
	return s
}

func (s *GetTaskAttributeRequest) Validate() error {
	return dara.Validate(s)
}
