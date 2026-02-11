// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDiscoveryEndpointsInput interface {
	dara.Model
	String() string
	GoString() string
	SetDiscoveryEndpoints(v []*DiscoveryEndpoint) *UpdateDiscoveryEndpointsInput
	GetDiscoveryEndpoints() []*DiscoveryEndpoint
}

type UpdateDiscoveryEndpointsInput struct {
	DiscoveryEndpoints []*DiscoveryEndpoint `json:"discoveryEndpoints,omitempty" xml:"discoveryEndpoints,omitempty" type:"Repeated"`
}

func (s UpdateDiscoveryEndpointsInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateDiscoveryEndpointsInput) GoString() string {
	return s.String()
}

func (s *UpdateDiscoveryEndpointsInput) GetDiscoveryEndpoints() []*DiscoveryEndpoint {
	return s.DiscoveryEndpoints
}

func (s *UpdateDiscoveryEndpointsInput) SetDiscoveryEndpoints(v []*DiscoveryEndpoint) *UpdateDiscoveryEndpointsInput {
	s.DiscoveryEndpoints = v
	return s
}

func (s *UpdateDiscoveryEndpointsInput) Validate() error {
	if s.DiscoveryEndpoints != nil {
		for _, item := range s.DiscoveryEndpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
