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
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	Integration *string `json:"Integration,omitempty" xml:"Integration,omitempty"`
	// This parameter is required.
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
