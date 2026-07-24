// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartComputeJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RestartComputeJobRequest
	GetClientToken() *string
	SetInstanceId(v string) *RestartComputeJobRequest
	GetInstanceId() *string
	SetJobName(v string) *RestartComputeJobRequest
	GetJobName() *string
	SetRegionId(v string) *RestartComputeJobRequest
	GetRegionId() *string
}

type RestartComputeJobRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RestartComputeJobRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartComputeJobRequest) GoString() string {
	return s.String()
}

func (s *RestartComputeJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RestartComputeJobRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartComputeJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *RestartComputeJobRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RestartComputeJobRequest) SetClientToken(v string) *RestartComputeJobRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartComputeJobRequest) SetInstanceId(v string) *RestartComputeJobRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartComputeJobRequest) SetJobName(v string) *RestartComputeJobRequest {
	s.JobName = &v
	return s
}

func (s *RestartComputeJobRequest) SetRegionId(v string) *RestartComputeJobRequest {
	s.RegionId = &v
	return s
}

func (s *RestartComputeJobRequest) Validate() error {
	return dara.Validate(s)
}
