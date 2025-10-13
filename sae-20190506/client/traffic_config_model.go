// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrafficConfig interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalVersionWeight(v map[string]*float32) *TrafficConfig
	GetAdditionalVersionWeight() map[string]*float32
	SetCreatedTime(v string) *TrafficConfig
	GetCreatedTime() *string
	SetLastModifiedTime(v string) *TrafficConfig
	GetLastModifiedTime() *string
	SetRequestId(v string) *TrafficConfig
	GetRequestId() *string
	SetResolvePolicy(v string) *TrafficConfig
	GetResolvePolicy() *string
	SetRoutePolicy(v *RoutePolicy) *TrafficConfig
	GetRoutePolicy() *RoutePolicy
	SetVersionId(v string) *TrafficConfig
	GetVersionId() *string
}

type TrafficConfig struct {
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight,omitempty" xml:"additionalVersionWeight,omitempty"`
	CreatedTime             *string             `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	LastModifiedTime        *string             `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	RequestId               *string             `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResolvePolicy           *string             `json:"resolvePolicy,omitempty" xml:"resolvePolicy,omitempty"`
	RoutePolicy             *RoutePolicy        `json:"routePolicy,omitempty" xml:"routePolicy,omitempty"`
	VersionId               *string             `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s TrafficConfig) String() string {
	return dara.Prettify(s)
}

func (s TrafficConfig) GoString() string {
	return s.String()
}

func (s *TrafficConfig) GetAdditionalVersionWeight() map[string]*float32 {
	return s.AdditionalVersionWeight
}

func (s *TrafficConfig) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *TrafficConfig) GetLastModifiedTime() *string {
	return s.LastModifiedTime
}

func (s *TrafficConfig) GetRequestId() *string {
	return s.RequestId
}

func (s *TrafficConfig) GetResolvePolicy() *string {
	return s.ResolvePolicy
}

func (s *TrafficConfig) GetRoutePolicy() *RoutePolicy {
	return s.RoutePolicy
}

func (s *TrafficConfig) GetVersionId() *string {
	return s.VersionId
}

func (s *TrafficConfig) SetAdditionalVersionWeight(v map[string]*float32) *TrafficConfig {
	s.AdditionalVersionWeight = v
	return s
}

func (s *TrafficConfig) SetCreatedTime(v string) *TrafficConfig {
	s.CreatedTime = &v
	return s
}

func (s *TrafficConfig) SetLastModifiedTime(v string) *TrafficConfig {
	s.LastModifiedTime = &v
	return s
}

func (s *TrafficConfig) SetRequestId(v string) *TrafficConfig {
	s.RequestId = &v
	return s
}

func (s *TrafficConfig) SetResolvePolicy(v string) *TrafficConfig {
	s.ResolvePolicy = &v
	return s
}

func (s *TrafficConfig) SetRoutePolicy(v *RoutePolicy) *TrafficConfig {
	s.RoutePolicy = v
	return s
}

func (s *TrafficConfig) SetVersionId(v string) *TrafficConfig {
	s.VersionId = &v
	return s
}

func (s *TrafficConfig) Validate() error {
	if s.RoutePolicy != nil {
		if err := s.RoutePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
