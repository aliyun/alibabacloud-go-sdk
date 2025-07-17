// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntegrationStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetIntegrationStateRequest
	GetClusterId() *string
	SetIntegration(v string) *GetIntegrationStateRequest
	GetIntegration() *string
	SetRegionId(v string) *GetIntegrationStateRequest
	GetRegionId() *string
}

type GetIntegrationStateRequest struct {
	// The ID of the Container Service for Kubernetes (ACK) cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The abbreviation of the software that is supported by ARMS. Valid values (case-insensitive): `ASM`, `IoT`, and `Flink`.
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

func (s GetIntegrationStateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIntegrationStateRequest) GoString() string {
	return s.String()
}

func (s *GetIntegrationStateRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetIntegrationStateRequest) GetIntegration() *string {
	return s.Integration
}

func (s *GetIntegrationStateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetIntegrationStateRequest) SetClusterId(v string) *GetIntegrationStateRequest {
	s.ClusterId = &v
	return s
}

func (s *GetIntegrationStateRequest) SetIntegration(v string) *GetIntegrationStateRequest {
	s.Integration = &v
	return s
}

func (s *GetIntegrationStateRequest) SetRegionId(v string) *GetIntegrationStateRequest {
	s.RegionId = &v
	return s
}

func (s *GetIntegrationStateRequest) Validate() error {
	return dara.Validate(s)
}
