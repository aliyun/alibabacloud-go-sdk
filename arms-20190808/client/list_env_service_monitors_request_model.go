// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvServiceMonitorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *ListEnvServiceMonitorsRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *ListEnvServiceMonitorsRequest
	GetRegionId() *string
}

type ListEnvServiceMonitorsRequest struct {
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

func (s ListEnvServiceMonitorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvServiceMonitorsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvServiceMonitorsRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvServiceMonitorsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvServiceMonitorsRequest) SetEnvironmentId(v string) *ListEnvServiceMonitorsRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvServiceMonitorsRequest) SetRegionId(v string) *ListEnvServiceMonitorsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvServiceMonitorsRequest) Validate() error {
	return dara.Validate(s)
}
