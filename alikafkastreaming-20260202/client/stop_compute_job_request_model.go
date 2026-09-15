// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StopComputeJobRequest
	GetInstanceId() *string
	SetJobName(v string) *StopComputeJobRequest
	GetJobName() *string
	SetRegionId(v string) *StopComputeJobRequest
	GetRegionId() *string
}

type StopComputeJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// alikafka_streaming-cn-pe333xxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// order_enrichment
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s StopComputeJobRequest) GoString() string {
	return s.String()
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
