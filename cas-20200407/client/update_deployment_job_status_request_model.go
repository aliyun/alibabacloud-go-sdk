// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentJobStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *UpdateDeploymentJobStatusRequest
	GetJobId() *int64
	SetStatus(v string) *UpdateDeploymentJobStatusRequest
	GetStatus() *string
}

type UpdateDeploymentJobStatusRequest struct {
	// The ID of the deployment task.
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
	// 	- pending
	//
	// 	- scheduling
	//
	// 	- editing
	//
	// This parameter is required.
	//
	// example:
	//
	// editing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateDeploymentJobStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentJobStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentJobStatusRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *UpdateDeploymentJobStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateDeploymentJobStatusRequest) SetJobId(v int64) *UpdateDeploymentJobStatusRequest {
	s.JobId = &v
	return s
}

func (s *UpdateDeploymentJobStatusRequest) SetStatus(v string) *UpdateDeploymentJobStatusRequest {
	s.Status = &v
	return s
}

func (s *UpdateDeploymentJobStatusRequest) Validate() error {
	return dara.Validate(s)
}
