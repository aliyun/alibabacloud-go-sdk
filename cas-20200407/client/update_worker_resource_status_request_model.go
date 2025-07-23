// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkerResourceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *UpdateWorkerResourceStatusRequest
	GetJobId() *int64
	SetStatus(v string) *UpdateWorkerResourceStatusRequest
	GetStatus() *string
	SetWorkerId(v int64) *UpdateWorkerResourceStatusRequest
	GetWorkerId() *int64
}

type UpdateWorkerResourceStatusRequest struct {
	// The ID of the deployment task. You can call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation to obtain the ID of a deployment task from the **JobId*	- parameter. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID of a deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The desired status.
	//
	// Valid values:
	//
	// 	- rollback
	//
	// This parameter is required.
	//
	// example:
	//
	// rollback
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the worker task. You can call the [ListWorkerResource](https://help.aliyun.com/document_detail/2712224.html) operation to obtain the ID of a worker task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	WorkerId *int64 `json:"WorkerId,omitempty" xml:"WorkerId,omitempty"`
}

func (s UpdateWorkerResourceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkerResourceStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkerResourceStatusRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *UpdateWorkerResourceStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateWorkerResourceStatusRequest) GetWorkerId() *int64 {
	return s.WorkerId
}

func (s *UpdateWorkerResourceStatusRequest) SetJobId(v int64) *UpdateWorkerResourceStatusRequest {
	s.JobId = &v
	return s
}

func (s *UpdateWorkerResourceStatusRequest) SetStatus(v string) *UpdateWorkerResourceStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateWorkerResourceStatusRequest) SetWorkerId(v int64) *UpdateWorkerResourceStatusRequest {
	s.WorkerId = &v
	return s
}

func (s *UpdateWorkerResourceStatusRequest) Validate() error {
	return dara.Validate(s)
}
