// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *AddIntegrationRequest
	GetClusterId() *string
	SetIntegration(v string) *AddIntegrationRequest
	GetIntegration() *string
	SetRegionId(v string) *AddIntegrationRequest
	GetRegionId() *string
}

type AddIntegrationRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s AddIntegrationRequest) GoString() string {
	return s.String()
}

func (s *AddIntegrationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *AddIntegrationRequest) GetIntegration() *string {
	return s.Integration
}

func (s *AddIntegrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddIntegrationRequest) SetClusterId(v string) *AddIntegrationRequest {
	s.ClusterId = &v
	return s
}

func (s *AddIntegrationRequest) SetIntegration(v string) *AddIntegrationRequest {
	s.Integration = &v
	return s
}

func (s *AddIntegrationRequest) SetRegionId(v string) *AddIntegrationRequest {
	s.RegionId = &v
	return s
}

func (s *AddIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
