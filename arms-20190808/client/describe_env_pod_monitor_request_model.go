// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvPodMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *DescribeEnvPodMonitorRequest
	GetEnvironmentId() *string
	SetNamespace(v string) *DescribeEnvPodMonitorRequest
	GetNamespace() *string
	SetPodMonitorName(v string) *DescribeEnvPodMonitorRequest
	GetPodMonitorName() *string
	SetRegionId(v string) *DescribeEnvPodMonitorRequest
	GetRegionId() *string
}

type DescribeEnvPodMonitorRequest struct {
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The namespace where the PodMonitor resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The name of the PodMonitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// arms-admin-pm1
	PodMonitorName *string `json:"PodMonitorName,omitempty" xml:"PodMonitorName,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeEnvPodMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvPodMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnvPodMonitorRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvPodMonitorRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeEnvPodMonitorRequest) GetPodMonitorName() *string {
	return s.PodMonitorName
}

func (s *DescribeEnvPodMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvPodMonitorRequest) SetEnvironmentId(v string) *DescribeEnvPodMonitorRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvPodMonitorRequest) SetNamespace(v string) *DescribeEnvPodMonitorRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeEnvPodMonitorRequest) SetPodMonitorName(v string) *DescribeEnvPodMonitorRequest {
	s.PodMonitorName = &v
	return s
}

func (s *DescribeEnvPodMonitorRequest) SetRegionId(v string) *DescribeEnvPodMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvPodMonitorRequest) Validate() error {
	return dara.Validate(s)
}
