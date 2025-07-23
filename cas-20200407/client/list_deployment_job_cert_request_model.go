// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentJobCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *ListDeploymentJobCertRequest
	GetJobId() *int64
}

type ListDeploymentJobCertRequest struct {
	// The ID of the deployment task. You can call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) operation to obtain the ID of a deployment task from the **JobId*	- parameter. You can also call the [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation to obtain the ID of a deployment task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ListDeploymentJobCertRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentJobCertRequest) GoString() string {
	return s.String()
}

func (s *ListDeploymentJobCertRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListDeploymentJobCertRequest) SetJobId(v int64) *ListDeploymentJobCertRequest {
	s.JobId = &v
	return s
}

func (s *ListDeploymentJobCertRequest) Validate() error {
	return dara.Validate(s)
}
