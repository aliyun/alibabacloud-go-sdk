// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkerResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *DeleteWorkerResourceRequest
	GetJobId() *int64
	SetWorkerId(v int64) *DeleteWorkerResourceRequest
	GetWorkerId() *int64
}

type DeleteWorkerResourceRequest struct {
	// The ID of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the worker for the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13
	WorkerId *int64 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s DeleteWorkerResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkerResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkerResourceRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DeleteWorkerResourceRequest) GetWorkerId() *int64 {
	return s.WorkerId
}

func (s *DeleteWorkerResourceRequest) SetJobId(v int64) *DeleteWorkerResourceRequest {
	s.JobId = &v
	return s
}

func (s *DeleteWorkerResourceRequest) SetWorkerId(v int64) *DeleteWorkerResourceRequest {
	s.WorkerId = &v
	return s
}

func (s *DeleteWorkerResourceRequest) Validate() error {
	return dara.Validate(s)
}
