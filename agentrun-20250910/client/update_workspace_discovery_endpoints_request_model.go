// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceDiscoveryEndpointsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UpdateDiscoveryEndpointsInput) *UpdateWorkspaceDiscoveryEndpointsRequest
	GetBody() *UpdateDiscoveryEndpointsInput
}

type UpdateWorkspaceDiscoveryEndpointsRequest struct {
	Body *UpdateDiscoveryEndpointsInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceDiscoveryEndpointsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDiscoveryEndpointsRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDiscoveryEndpointsRequest) GetBody() *UpdateDiscoveryEndpointsInput {
	return s.Body
}

func (s *UpdateWorkspaceDiscoveryEndpointsRequest) SetBody(v *UpdateDiscoveryEndpointsInput) *UpdateWorkspaceDiscoveryEndpointsRequest {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceDiscoveryEndpointsRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
