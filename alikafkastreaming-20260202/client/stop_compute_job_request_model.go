// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StopComputeJobRequest
	GetClientToken() *string
	SetInstanceId(v string) *StopComputeJobRequest
	GetInstanceId() *string
	SetJobName(v string) *StopComputeJobRequest
	GetJobName() *string
	SetRegionId(v string) *StopComputeJobRequest
	GetRegionId() *string
}

type StopComputeJobRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopComputeJobRequest) GoString() string {
	return s.String()
}

func (s *StopComputeJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopComputeJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopComputeJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *StopComputeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopComputeJobRequest) SetClientToken(v string) *StopComputeJobRequest {
	s.ClientToken = &v
	return s
}

func (s *StopComputeJobRequest) SetInstanceId(v string) *StopComputeJobRequest {
	s.InstanceId = &v
	return s
}

func (s *StopComputeJobRequest) SetJobName(v string) *StopComputeJobRequest {
	s.JobName = &v
	return s
}

func (s *StopComputeJobRequest) SetRegionId(v string) *StopComputeJobRequest {
	s.RegionId = &v
	return s
}

func (s *StopComputeJobRequest) Validate() error {
	return dara.Validate(s)
}
