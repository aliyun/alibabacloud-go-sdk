// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRollbackTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *CreateRollbackTaskRequest
	GetJobId() *int64
	SetWorkerId(v int64) *CreateRollbackTaskRequest
	GetWorkerId() *int64
}

type CreateRollbackTaskRequest struct {
	// The ID of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 436493
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the deployment worker.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4197913
	WorkerId *int64 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s CreateRollbackTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRollbackTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateRollbackTaskRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *CreateRollbackTaskRequest) GetWorkerId() *int64 {
	return s.WorkerId
}

func (s *CreateRollbackTaskRequest) SetJobId(v int64) *CreateRollbackTaskRequest {
	s.JobId = &v
	return s
}

func (s *CreateRollbackTaskRequest) SetWorkerId(v int64) *CreateRollbackTaskRequest {
	s.WorkerId = &v
	return s
}

func (s *CreateRollbackTaskRequest) Validate() error {
	return dara.Validate(s)
}
