// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateApplicationTrafficConfigInput interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalVersionWeight(v map[string]*float32) *UpdateApplicationTrafficConfigInput
	GetAdditionalVersionWeight() map[string]*float32
	SetResolvePolicy(v string) *UpdateApplicationTrafficConfigInput
	GetResolvePolicy() *string
	SetRoutePolicy(v *RoutePolicy) *UpdateApplicationTrafficConfigInput
	GetRoutePolicy() *RoutePolicy
	SetVersionId(v string) *UpdateApplicationTrafficConfigInput
	GetVersionId() *string
}

type UpdateApplicationTrafficConfigInput struct {
	AdditionalVersionWeight map[string]*float32 `json:"additionalVersionWeight,omitempty" xml:"additionalVersionWeight,omitempty"`
	ResolvePolicy           *string             `json:"resolvePolicy,omitempty" xml:"resolvePolicy,omitempty"`
	RoutePolicy             *RoutePolicy        `json:"routePolicy,omitempty" xml:"routePolicy,omitempty"`
	VersionId               *string             `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s UpdateApplicationTrafficConfigInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateApplicationTrafficConfigInput) GoString() string {
	return s.String()
}

func (s *UpdateApplicationTrafficConfigInput) GetAdditionalVersionWeight() map[string]*float32 {
	return s.AdditionalVersionWeight
}

func (s *UpdateApplicationTrafficConfigInput) GetResolvePolicy() *string {
	return s.ResolvePolicy
}

func (s *UpdateApplicationTrafficConfigInput) GetRoutePolicy() *RoutePolicy {
	return s.RoutePolicy
}

func (s *UpdateApplicationTrafficConfigInput) GetVersionId() *string {
	return s.VersionId
}

func (s *UpdateApplicationTrafficConfigInput) SetAdditionalVersionWeight(v map[string]*float32) *UpdateApplicationTrafficConfigInput {
	s.AdditionalVersionWeight = v
	return s
}

func (s *UpdateApplicationTrafficConfigInput) SetResolvePolicy(v string) *UpdateApplicationTrafficConfigInput {
	s.ResolvePolicy = &v
	return s
}

func (s *UpdateApplicationTrafficConfigInput) SetRoutePolicy(v *RoutePolicy) *UpdateApplicationTrafficConfigInput {
	s.RoutePolicy = v
	return s
}

func (s *UpdateApplicationTrafficConfigInput) SetVersionId(v string) *UpdateApplicationTrafficConfigInput {
	s.VersionId = &v
	return s
}

func (s *UpdateApplicationTrafficConfigInput) Validate() error {
	if s.RoutePolicy != nil {
		if err := s.RoutePolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}
