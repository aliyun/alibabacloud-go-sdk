// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *DescribeDeploymentJobRequest
	GetJobId() *int64
}

type DescribeDeploymentJobRequest struct {
	// The ID of the deployment job. To get the job ID, call the [CreateDeploymentJob](https://help.aliyun.com/document_detail/2712234.html) or [ListDeploymentJob](https://help.aliyun.com/document_detail/2712223.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 8888
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s DescribeDeploymentJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentJobRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *DescribeDeploymentJobRequest) SetJobId(v int64) *DescribeDeploymentJobRequest {
	s.JobId = &v
	return s
}

func (s *DescribeDeploymentJobRequest) Validate() error {
	return dara.Validate(s)
}
