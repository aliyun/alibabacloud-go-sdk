// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *ListDeploymentJobResourceRequest
	GetJobId() *int64
}

type ListDeploymentJobResourceRequest struct {
	// The ID of the deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ListDeploymentJobResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobResourceRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobResourceRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListDeploymentJobResourceRequest) SetJobId(v int64) *ListDeploymentJobResourceRequest {
	s.JobId = &v
	return s
}

func (s *ListDeploymentJobResourceRequest) Validate() error {
	return dara.Validate(s)
}
