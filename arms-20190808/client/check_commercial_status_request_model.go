// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCommercialStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CheckCommercialStatusRequest
	GetRegionId() *string
	SetService(v string) *CheckCommercialStatusRequest
	GetService() *string
}

type CheckCommercialStatusRequest struct {
	// The region ID. Default value: ap-southeast-1.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The sub-product type. Valid values:
	//
	// - apm: Application Monitoring
	//
	// - rum: Real User Monitoring
	//
	// - prometheus: Managed Service for Prometheus
	//
	// - xtrace: Managed Service for OpenTelemetry
	//
	// This parameter is required.
	//
	// example:
	//
	// apm
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s CheckCommercialStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckCommercialStatusRequest) GoString() string {
	return s.String()
}

func (s *CheckCommercialStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckCommercialStatusRequest) GetService() *string {
	return s.Service
}

func (s *CheckCommercialStatusRequest) SetRegionId(v string) *CheckCommercialStatusRequest {
	s.RegionId = &v
	return s
}

func (s *CheckCommercialStatusRequest) SetService(v string) *CheckCommercialStatusRequest {
	s.Service = &v
	return s
}

func (s *CheckCommercialStatusRequest) Validate() error {
	return dara.Validate(s)
}
