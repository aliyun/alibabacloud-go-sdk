// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkAvailableZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsPoc(v bool) *GetNetworkAvailableZonesRequest
	GetIsPoc() *bool
	SetNetworkRegionId(v string) *GetNetworkAvailableZonesRequest
	GetNetworkRegionId() *string
	SetPrivateVpcConnectionMode(v string) *GetNetworkAvailableZonesRequest
	GetPrivateVpcConnectionMode() *string
	SetServiceId(v string) *GetNetworkAvailableZonesRequest
	GetServiceId() *string
	SetServiceInstanceEndpointServiceType(v string) *GetNetworkAvailableZonesRequest
	GetServiceInstanceEndpointServiceType() *string
	SetServiceRegionId(v string) *GetNetworkAvailableZonesRequest
	GetServiceRegionId() *string
	SetServiceVersion(v string) *GetNetworkAvailableZonesRequest
	GetServiceVersion() *string
	SetZoneId(v string) *GetNetworkAvailableZonesRequest
	GetZoneId() *string
}

type GetNetworkAvailableZonesRequest struct {
	// example:
	//
	// true
	IsPoc *bool `json:"IsPoc,omitempty" xml:"IsPoc,omitempty"`
	// example:
	//
	// cn-hangzhou
	NetworkRegionId *string `json:"NetworkRegionId,omitempty" xml:"NetworkRegionId,omitempty"`
	// example:
	//
	// PrivateLink
	PrivateVpcConnectionMode *string `json:"PrivateVpcConnectionMode,omitempty" xml:"PrivateVpcConnectionMode,omitempty"`
	// example:
	//
	// service-55fc2eabbce647fa976b
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// Forward
	ServiceInstanceEndpointServiceType *string `json:"ServiceInstanceEndpointServiceType,omitempty" xml:"ServiceInstanceEndpointServiceType,omitempty"`
	// example:
	//
	// cn-hangzhou
	ServiceRegionId *string `json:"ServiceRegionId,omitempty" xml:"ServiceRegionId,omitempty"`
	// example:
	//
	// 1
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetNetworkAvailableZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkAvailableZonesRequest) GoString() string {
	return s.String()
}

func (s *GetNetworkAvailableZonesRequest) GetIsPoc() *bool {
	return s.IsPoc
}

func (s *GetNetworkAvailableZonesRequest) GetNetworkRegionId() *string {
	return s.NetworkRegionId
}

func (s *GetNetworkAvailableZonesRequest) GetPrivateVpcConnectionMode() *string {
	return s.PrivateVpcConnectionMode
}

func (s *GetNetworkAvailableZonesRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetNetworkAvailableZonesRequest) GetServiceInstanceEndpointServiceType() *string {
	return s.ServiceInstanceEndpointServiceType
}

func (s *GetNetworkAvailableZonesRequest) GetServiceRegionId() *string {
	return s.ServiceRegionId
}

func (s *GetNetworkAvailableZonesRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *GetNetworkAvailableZonesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetNetworkAvailableZonesRequest) SetIsPoc(v bool) *GetNetworkAvailableZonesRequest {
	s.IsPoc = &v
	return s
}

func (s *GetNetworkAvailableZonesRequest) SetNetworkRegionId(v string) *GetNetworkAvailableZonesRequest {
	s.NetworkRegionId = &v
	return s
}

func (s *GetNetworkAvailableZonesRequest) SetPrivateVpcConnectionMode(v string) *GetNetworkAvailableZonesRequest {
	s.PrivateVpcConnectionMode = &v
	return s
}

func (s *GetNetworkAvailableZonesRequest) SetServiceId(v string) *GetNetworkAvailableZonesRequest {
	s.ServiceId = &v
	return s
}

func (s *GetNetworkAvailableZonesRequest) SetServiceInstanceEndpointServiceType(v string) *GetNetworkAvailableZonesRequest {
	s.ServiceInstanceEndpointServiceType = &v
	return s
}

func (s *GetNetworkAvailableZonesRequest) SetServiceRegionId(v string) *GetNetworkAvailableZonesRequest {
	s.ServiceRegionId = &v
	return s
}

func (s *GetNetworkAvailableZonesRequest) SetServiceVersion(v string) *GetNetworkAvailableZonesRequest {
	s.ServiceVersion = &v
	return s
}

func (s *GetNetworkAvailableZonesRequest) SetZoneId(v string) *GetNetworkAvailableZonesRequest {
	s.ZoneId = &v
	return s
}

func (s *GetNetworkAvailableZonesRequest) Validate() error {
	return dara.Validate(s)
}
