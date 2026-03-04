// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceDiscoveryEndpointsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkspaceDiscoveryEndpointsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkspaceDiscoveryEndpointsResponse
	GetStatusCode() *int32
	SetBody(v *GetDiscoveryEndpointsResult) *GetWorkspaceDiscoveryEndpointsResponse
	GetBody() *GetDiscoveryEndpointsResult
}

type GetWorkspaceDiscoveryEndpointsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDiscoveryEndpointsResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkspaceDiscoveryEndpointsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceDiscoveryEndpointsResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceDiscoveryEndpointsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkspaceDiscoveryEndpointsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkspaceDiscoveryEndpointsResponse) GetBody() *GetDiscoveryEndpointsResult {
	return s.Body
}

func (s *GetWorkspaceDiscoveryEndpointsResponse) SetHeaders(v map[string]*string) *GetWorkspaceDiscoveryEndpointsResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceDiscoveryEndpointsResponse) SetStatusCode(v int32) *GetWorkspaceDiscoveryEndpointsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkspaceDiscoveryEndpointsResponse) SetBody(v *GetDiscoveryEndpointsResult) *GetWorkspaceDiscoveryEndpointsResponse {
	s.Body = v
	return s
}

func (s *GetWorkspaceDiscoveryEndpointsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
