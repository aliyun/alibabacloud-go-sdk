// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvServiceMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *DeleteEnvServiceMonitorRequest
	GetEnvironmentId() *string
	SetNamespace(v string) *DeleteEnvServiceMonitorRequest
	GetNamespace() *string
	SetRegionId(v string) *DeleteEnvServiceMonitorRequest
	GetRegionId() *string
	SetServiceMonitorName(v string) *DeleteEnvServiceMonitorRequest
	GetServiceMonitorName() *string
}

type DeleteEnvServiceMonitorRequest struct {
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The namespace where the ServiceMonitor is located.
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

func (s DeleteEnvServiceMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvServiceMonitorRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnvServiceMonitorRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DeleteEnvServiceMonitorRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteEnvServiceMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEnvServiceMonitorRequest) GetServiceMonitorName() *string {
	return s.ServiceMonitorName
}

func (s *DeleteEnvServiceMonitorRequest) SetEnvironmentId(v string) *DeleteEnvServiceMonitorRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DeleteEnvServiceMonitorRequest) SetNamespace(v string) *DeleteEnvServiceMonitorRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteEnvServiceMonitorRequest) SetRegionId(v string) *DeleteEnvServiceMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEnvServiceMonitorRequest) SetServiceMonitorName(v string) *DeleteEnvServiceMonitorRequest {
	s.ServiceMonitorName = &v
	return s
}

func (s *DeleteEnvServiceMonitorRequest) Validate() error {
	return dara.Validate(s)
}
