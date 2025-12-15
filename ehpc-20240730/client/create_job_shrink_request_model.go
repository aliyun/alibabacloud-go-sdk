// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateJobShrinkRequest
	GetClusterId() *string
	SetJobName(v string) *CreateJobShrinkRequest
	GetJobName() *string
	SetJobSpecShrink(v string) *CreateJobShrinkRequest
	GetJobSpecShrink() *string
}

type CreateJobShrinkRequest struct {
	// The cluster ID.
	//
	// You can call the [ListClusters](https://help.aliyun.com/document_detail/87116.html) operation to query the cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ehpc-hz-FYUr32****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The job name.
	//
	// example:
	//
	// TestJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job configurations.
	JobSpecShrink *string `json:"JobSpec,omitempty" xml:"JobSpec,omitempty"`
}

func (s CreateJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateJobShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateJobShrinkRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateJobShrinkRequest) GetJobSpecShrink() *string {
	return s.JobSpecShrink
}

func (s *CreateJobShrinkRequest) SetClusterId(v string) *CreateJobShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobName(v string) *CreateJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateJobShrinkRequest) SetJobSpecShrink(v string) *CreateJobShrinkRequest {
	s.JobSpecShrink = &v
	return s
}

func (s *CreateJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
