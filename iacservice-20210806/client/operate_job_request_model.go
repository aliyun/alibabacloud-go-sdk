// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *OperateJobRequest
	GetComment() *string
	SetTaskType(v string) *OperateJobRequest
	GetTaskType() *string
}

type OperateJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dasd
	Comment  *string `json:"comment,omitempty" xml:"comment,omitempty"`
	TaskType *string `json:"taskType,omitempty" xml:"taskType,omitempty"`
}

func (s OperateJobRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateJobRequest) GoString() string {
	return s.String()
}

func (s *OperateJobRequest) GetComment() *string {
	return s.Comment
}

func (s *OperateJobRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *OperateJobRequest) SetComment(v string) *OperateJobRequest {
	s.Comment = &v
	return s
}

func (s *OperateJobRequest) SetTaskType(v string) *OperateJobRequest {
	s.TaskType = &v
	return s
}

func (s *OperateJobRequest) Validate() error {
	return dara.Validate(s)
}
