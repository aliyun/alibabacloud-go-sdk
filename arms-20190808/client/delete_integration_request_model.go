// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteIntegrationRequest
	GetClusterId() *string
	SetIntegration(v string) *DeleteIntegrationRequest
	GetIntegration() *string
	SetRegionId(v string) *DeleteIntegrationRequest
	GetRegionId() *string
}

type DeleteIntegrationRequest struct {
	// The ID of the ACK cluster.
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
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntegrationRequest) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteIntegrationRequest) GetIntegration() *string {
	return s.Integration
}

func (s *DeleteIntegrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteIntegrationRequest) SetClusterId(v string) *DeleteIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteIntegrationRequest) SetIntegration(v string) *DeleteIntegrationRequest {
	s.Integration = &v
	return s
}

func (s *DeleteIntegrationRequest) SetRegionId(v string) *DeleteIntegrationRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
