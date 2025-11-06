// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGovernanceKubernetesClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGovernanceKubernetesClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGovernanceKubernetesClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetGovernanceKubernetesClusterResponseBody) *GetGovernanceKubernetesClusterResponse
	GetBody() *GetGovernanceKubernetesClusterResponseBody
}

type GetGovernanceKubernetesClusterResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGovernanceKubernetesClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGovernanceKubernetesClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGovernanceKubernetesClusterResponse) GoString() string {
	return s.String()
}

func (s *GetGovernanceKubernetesClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGovernanceKubernetesClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGovernanceKubernetesClusterResponse) GetBody() *GetGovernanceKubernetesClusterResponseBody {
	return s.Body
}

func (s *GetGovernanceKubernetesClusterResponse) SetHeaders(v map[string]*string) *GetGovernanceKubernetesClusterResponse {
	s.Headers = v
	return s
}

func (s *GetGovernanceKubernetesClusterResponse) SetStatusCode(v int32) *GetGovernanceKubernetesClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGovernanceKubernetesClusterResponse) SetBody(v *GetGovernanceKubernetesClusterResponseBody) *GetGovernanceKubernetesClusterResponse {
	s.Body = v
	return s
}

func (s *GetGovernanceKubernetesClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
