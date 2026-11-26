// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectHttpApiConflictsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRestApiConfig(v *DetectHttpApiConflictsRequestRestApiConfig) *DetectHttpApiConflictsRequest
	GetRestApiConfig() *DetectHttpApiConflictsRequestRestApiConfig
	SetRouteId(v string) *DetectHttpApiConflictsRequest
	GetRouteId() *string
}

type DetectHttpApiConflictsRequest struct {
	RestApiConfig *DetectHttpApiConflictsRequestRestApiConfig `json:"restApiConfig,omitempty" xml:"restApiConfig,omitempty" type:"Struct"`
	RouteId       *string                                     `json:"routeId,omitempty" xml:"routeId,omitempty"`
}

func (s DetectHttpApiConflictsRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectHttpApiConflictsRequest) GoString() string {
	return s.String()
}

func (s *DetectHttpApiConflictsRequest) GetRestApiConfig() *DetectHttpApiConflictsRequestRestApiConfig {
	return s.RestApiConfig
}

func (s *DetectHttpApiConflictsRequest) GetRouteId() *string {
	return s.RouteId
}

func (s *DetectHttpApiConflictsRequest) SetRestApiConfig(v *DetectHttpApiConflictsRequestRestApiConfig) *DetectHttpApiConflictsRequest {
	s.RestApiConfig = v
	return s
}

func (s *DetectHttpApiConflictsRequest) SetRouteId(v string) *DetectHttpApiConflictsRequest {
	s.RouteId = &v
	return s
}

func (s *DetectHttpApiConflictsRequest) Validate() error {
	if s.RestApiConfig != nil {
		if err := s.RestApiConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectHttpApiConflictsRequestRestApiConfig struct {
	BackendScene    *string                                                     `json:"backendScene,omitempty" xml:"backendScene,omitempty"`
	CustomDomainIds []*string                                                   `json:"customDomainIds,omitempty" xml:"customDomainIds,omitempty" type:"Repeated"`
	EnvironmentId   *string                                                     `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	GatewayId       *string                                                     `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	ServiceConfigs  []*DetectHttpApiConflictsRequestRestApiConfigServiceConfigs `json:"serviceConfigs,omitempty" xml:"serviceConfigs,omitempty" type:"Repeated"`
}

func (s DetectHttpApiConflictsRequestRestApiConfig) String() string {
	return dara.Prettify(s)
}

func (s DetectHttpApiConflictsRequestRestApiConfig) GoString() string {
	return s.String()
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) GetBackendScene() *string {
	return s.BackendScene
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) GetCustomDomainIds() []*string {
	return s.CustomDomainIds
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) GetGatewayId() *string {
	return s.GatewayId
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) GetServiceConfigs() []*DetectHttpApiConflictsRequestRestApiConfigServiceConfigs {
	return s.ServiceConfigs
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) SetBackendScene(v string) *DetectHttpApiConflictsRequestRestApiConfig {
	s.BackendScene = &v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) SetCustomDomainIds(v []*string) *DetectHttpApiConflictsRequestRestApiConfig {
	s.CustomDomainIds = v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) SetEnvironmentId(v string) *DetectHttpApiConflictsRequestRestApiConfig {
	s.EnvironmentId = &v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) SetGatewayId(v string) *DetectHttpApiConflictsRequestRestApiConfig {
	s.GatewayId = &v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) SetServiceConfigs(v []*DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) *DetectHttpApiConflictsRequestRestApiConfig {
	s.ServiceConfigs = v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfig) Validate() error {
	if s.ServiceConfigs != nil {
		for _, item := range s.ServiceConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetectHttpApiConflictsRequestRestApiConfigServiceConfigs struct {
	Match     *HttpApiBackendMatchConditions `json:"match,omitempty" xml:"match,omitempty"`
	Port      *int32                         `json:"port,omitempty" xml:"port,omitempty"`
	Protocol  *string                        `json:"protocol,omitempty" xml:"protocol,omitempty"`
	ServiceId *string                        `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	Version   *string                        `json:"version,omitempty" xml:"version,omitempty"`
	Weight    *int32                         `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) String() string {
	return dara.Prettify(s)
}

func (s DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) GoString() string {
	return s.String()
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) GetMatch() *HttpApiBackendMatchConditions {
	return s.Match
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) GetPort() *int32 {
	return s.Port
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) GetProtocol() *string {
	return s.Protocol
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) GetServiceId() *string {
	return s.ServiceId
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) GetVersion() *string {
	return s.Version
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) GetWeight() *int32 {
	return s.Weight
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) SetMatch(v *HttpApiBackendMatchConditions) *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs {
	s.Match = v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) SetPort(v int32) *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs {
	s.Port = &v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) SetProtocol(v string) *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs {
	s.Protocol = &v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) SetServiceId(v string) *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs {
	s.ServiceId = &v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) SetVersion(v string) *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs {
	s.Version = &v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) SetWeight(v int32) *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs {
	s.Weight = &v
	return s
}

func (s *DetectHttpApiConflictsRequestRestApiConfigServiceConfigs) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	return nil
}
