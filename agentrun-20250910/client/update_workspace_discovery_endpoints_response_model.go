// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceDiscoveryEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWorkspaceDiscoveryEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWorkspaceDiscoveryEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *GetDiscoveryEndpointsResult) *UpdateWorkspaceDiscoveryEndpointsResponse
	GetBody() *GetDiscoveryEndpointsResult
}

type UpdateWorkspaceDiscoveryEndpointsResponse struct {
	Headers    map[string]*string           `json:"headers" xml:"headers"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiscoveryEndpointsResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWorkspaceDiscoveryEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceDiscoveryEndpointsResponse) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceDiscoveryEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWorkspaceDiscoveryEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWorkspaceDiscoveryEndpointsResponse) GetBody() *GetDiscoveryEndpointsResult {
	return s.Body
}

func (s *UpdateWorkspaceDiscoveryEndpointsResponse) SetHeaders(v map[string]*string) *UpdateWorkspaceDiscoveryEndpointsResponse {
	s.Headers = v
	return s
}

func (s *UpdateWorkspaceDiscoveryEndpointsResponse) SetStatusCode(v int32) *UpdateWorkspaceDiscoveryEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWorkspaceDiscoveryEndpointsResponse) SetBody(v *GetDiscoveryEndpointsResult) *UpdateWorkspaceDiscoveryEndpointsResponse {
	s.Body = v
	return s
}

func (s *UpdateWorkspaceDiscoveryEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
