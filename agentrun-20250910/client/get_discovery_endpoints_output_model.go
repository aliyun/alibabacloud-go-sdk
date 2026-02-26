// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDiscoveryEndpointsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialName(v string) *GetDiscoveryEndpointsOutput
	GetCredentialName() *string
	SetDiscoveryEndpoints(v []*DiscoveryEndpoint) *GetDiscoveryEndpointsOutput
	GetDiscoveryEndpoints() []*DiscoveryEndpoint
}

type GetDiscoveryEndpointsOutput struct {
	CredentialName     *string              `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	DiscoveryEndpoints []*DiscoveryEndpoint `json:"discoveryEndpoints,omitempty" xml:"discoveryEndpoints,omitempty" type:"Repeated"`
}

func (s GetDiscoveryEndpointsOutput) String() string {
	return dara.Prettify(s)
}

func (s GetDiscoveryEndpointsOutput) GoString() string {
	return s.String()
}

func (s *GetDiscoveryEndpointsOutput) GetCredentialName() *string {
	return s.CredentialName
}

func (s *GetDiscoveryEndpointsOutput) GetDiscoveryEndpoints() []*DiscoveryEndpoint {
	return s.DiscoveryEndpoints
}

func (s *GetDiscoveryEndpointsOutput) SetCredentialName(v string) *GetDiscoveryEndpointsOutput {
	s.CredentialName = &v
	return s
}

func (s *GetDiscoveryEndpointsOutput) SetDiscoveryEndpoints(v []*DiscoveryEndpoint) *GetDiscoveryEndpointsOutput {
	s.DiscoveryEndpoints = v
	return s
}

func (s *GetDiscoveryEndpointsOutput) Validate() error {
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
