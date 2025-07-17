// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnvServiceMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *DescribeEnvServiceMonitorRequest
	GetEnvironmentId() *string
	SetNamespace(v string) *DescribeEnvServiceMonitorRequest
	GetNamespace() *string
	SetRegionId(v string) *DescribeEnvServiceMonitorRequest
	GetRegionId() *string
	SetServiceMonitorName(v string) *DescribeEnvServiceMonitorRequest
	GetServiceMonitorName() *string
}

type DescribeEnvServiceMonitorRequest struct {
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The namespace where the ServiceMonitor resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// arms-prom
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the ServiceMonitor.
	//
	// This parameter is required.
	//
	// example:
	//
	// arms-admin1
	ServiceMonitorName *string `json:"ServiceMonitorName,omitempty" xml:"ServiceMonitorName,omitempty"`
}

func (s DescribeEnvServiceMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnvServiceMonitorRequest) GoString() string {
	return s.String()
}

func (s *DescribeEnvServiceMonitorRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DescribeEnvServiceMonitorRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeEnvServiceMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEnvServiceMonitorRequest) GetServiceMonitorName() *string {
	return s.ServiceMonitorName
}

func (s *DescribeEnvServiceMonitorRequest) SetEnvironmentId(v string) *DescribeEnvServiceMonitorRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DescribeEnvServiceMonitorRequest) SetNamespace(v string) *DescribeEnvServiceMonitorRequest {
	s.Namespace = &v
	return s
}

func (s *DescribeEnvServiceMonitorRequest) SetRegionId(v string) *DescribeEnvServiceMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEnvServiceMonitorRequest) SetServiceMonitorName(v string) *DescribeEnvServiceMonitorRequest {
	s.ServiceMonitorName = &v
	return s
}

func (s *DescribeEnvServiceMonitorRequest) Validate() error {
	return dara.Validate(s)
}
