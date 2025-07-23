// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *DeleteDeploymentJobRequest
	GetJobId() *int64
}

type DeleteDeploymentJobRequest struct {
	// The ID of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DeleteDeploymentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DeleteDeploymentJobRequest) SetJobId(v int64) *DeleteDeploymentJobRequest {
	s.JobId = &v
	return s
}

func (s *DeleteDeploymentJobRequest) Validate() error {
	return dara.Validate(s)
}
