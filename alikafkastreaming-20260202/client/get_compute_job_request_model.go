// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetComputeJobRequest
	GetInstanceId() *string
	SetJobName(v string) *GetComputeJobRequest
	GetJobName() *string
	SetRegionId(v string) *GetComputeJobRequest
	GetRegionId() *string
}

type GetComputeJobRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetComputeJobRequest) GoString() string {
	return s.String()
}

func (s *GetComputeJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetComputeJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *GetComputeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetComputeJobRequest) SetInstanceId(v string) *GetComputeJobRequest {
	s.InstanceId = &v
	return s
}

func (s *GetComputeJobRequest) SetJobName(v string) *GetComputeJobRequest {
	s.JobName = &v
	return s
}

func (s *GetComputeJobRequest) SetRegionId(v string) *GetComputeJobRequest {
	s.RegionId = &v
	return s
}

func (s *GetComputeJobRequest) Validate() error {
	return dara.Validate(s)
}
