// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvPodMonitorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *ListEnvPodMonitorsRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *ListEnvPodMonitorsRequest
	GetRegionId() *string
}

type ListEnvPodMonitorsRequest struct {
	// The environment ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEnvPodMonitorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvPodMonitorsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvPodMonitorsRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvPodMonitorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvPodMonitorsRequest) SetEnvironmentId(v string) *ListEnvPodMonitorsRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvPodMonitorsRequest) SetRegionId(v string) *ListEnvPodMonitorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvPodMonitorsRequest) Validate() error {
	return dara.Validate(s)
}
