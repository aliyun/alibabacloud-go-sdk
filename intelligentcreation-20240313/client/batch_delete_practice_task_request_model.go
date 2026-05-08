// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeletePracticeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIdempotentId(v string) *BatchDeletePracticeTaskRequest
	GetIdempotentId() *string
	SetTaskIds(v []*string) *BatchDeletePracticeTaskRequest
	GetTaskIds() []*string
}

type BatchDeletePracticeTaskRequest struct {
	// example:
	//
	// 1234567890
	IdempotentId *string   `json:"idempotentId,omitempty" xml:"idempotentId,omitempty"`
	TaskIds      []*string `json:"taskIds,omitempty" xml:"taskIds,omitempty" type:"Repeated"`
}

func (s BatchDeletePracticeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeletePracticeTaskRequest) GoString() string {
	return s.String()
}

func (s *BatchDeletePracticeTaskRequest) GetIdempotentId() *string {
	return s.IdempotentId
}

func (s *BatchDeletePracticeTaskRequest) GetTaskIds() []*string {
	return s.TaskIds
}

func (s *BatchDeletePracticeTaskRequest) SetIdempotentId(v string) *BatchDeletePracticeTaskRequest {
	s.IdempotentId = &v
	return s
}

func (s *BatchDeletePracticeTaskRequest) SetTaskIds(v []*string) *BatchDeletePracticeTaskRequest {
	s.TaskIds = v
	return s
}

func (s *BatchDeletePracticeTaskRequest) Validate() error {
	return dara.Validate(s)
}
