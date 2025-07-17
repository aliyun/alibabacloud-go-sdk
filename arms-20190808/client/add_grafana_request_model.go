// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGrafanaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AddGrafanaRequest
	GetClusterId() *string
	SetIntegration(v string) *AddGrafanaRequest
	GetIntegration() *string
	SetRegionId(v string) *AddGrafanaRequest
	GetRegionId() *string
}

type AddGrafanaRequest struct {
	// The ID of the Container Service for Kubernetes (ACK) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The abbreviation of the software that is supported by Application Real-Time Monitoring Service (ARMS). Valid values (case-insensitive): `ASM`, `IoT`, and `Flink`.
	//
	// This parameter is required.
	//
	// example:
	//
	// asm
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddGrafanaRequest) String() string {
	return dara.Prettify(s)
}

func (s AddGrafanaRequest) GoString() string {
	return s.String()
}

func (s *AddGrafanaRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddGrafanaRequest) GetIntegration() *string {
	return s.Integration
}

func (s *AddGrafanaRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddGrafanaRequest) SetClusterId(v string) *AddGrafanaRequest {
	s.ClusterId = &v
	return s
}

func (s *AddGrafanaRequest) SetIntegration(v string) *AddGrafanaRequest {
	s.Integration = &v
	return s
}

func (s *AddGrafanaRequest) SetRegionId(v string) *AddGrafanaRequest {
	s.RegionId = &v
	return s
}

func (s *AddGrafanaRequest) Validate() error {
	return dara.Validate(s)
}
