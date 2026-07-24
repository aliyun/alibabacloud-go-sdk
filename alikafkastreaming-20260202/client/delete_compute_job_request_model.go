// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteComputeJobRequest
	GetInstanceId() *string
	SetJobName(v string) *DeleteComputeJobRequest
	GetJobName() *string
	SetRegionId(v string) *DeleteComputeJobRequest
	GetRegionId() *string
}

type DeleteComputeJobRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteComputeJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteComputeJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteComputeJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *DeleteComputeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteComputeJobRequest) SetInstanceId(v string) *DeleteComputeJobRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteComputeJobRequest) SetJobName(v string) *DeleteComputeJobRequest {
	s.JobName = &v
	return s
}

func (s *DeleteComputeJobRequest) SetRegionId(v string) *DeleteComputeJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteComputeJobRequest) Validate() error {
	return dara.Validate(s)
}
