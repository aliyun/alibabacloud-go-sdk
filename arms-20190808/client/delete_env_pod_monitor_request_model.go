// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvPodMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *DeleteEnvPodMonitorRequest
	GetEnvironmentId() *string
	SetNamespace(v string) *DeleteEnvPodMonitorRequest
	GetNamespace() *string
	SetPodMonitorName(v string) *DeleteEnvPodMonitorRequest
	GetPodMonitorName() *string
	SetRegionId(v string) *DeleteEnvPodMonitorRequest
	GetRegionId() *string
}

type DeleteEnvPodMonitorRequest struct {
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The namespace where the PodMonitor is located.
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

func (s DeleteEnvPodMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvPodMonitorRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnvPodMonitorRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DeleteEnvPodMonitorRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteEnvPodMonitorRequest) GetPodMonitorName() *string {
	return s.PodMonitorName
}

func (s *DeleteEnvPodMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEnvPodMonitorRequest) SetEnvironmentId(v string) *DeleteEnvPodMonitorRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DeleteEnvPodMonitorRequest) SetNamespace(v string) *DeleteEnvPodMonitorRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteEnvPodMonitorRequest) SetPodMonitorName(v string) *DeleteEnvPodMonitorRequest {
	s.PodMonitorName = &v
	return s
}

func (s *DeleteEnvPodMonitorRequest) SetRegionId(v string) *DeleteEnvPodMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEnvPodMonitorRequest) Validate() error {
	return dara.Validate(s)
}
