// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGovernanceKubernetesClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyGovernanceKubernetesClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyGovernanceKubernetesClusterResponse
	GetStatusCode() *int32
	SetBody(v *ModifyGovernanceKubernetesClusterResponseBody) *ModifyGovernanceKubernetesClusterResponse
	GetBody() *ModifyGovernanceKubernetesClusterResponseBody
}

type ModifyGovernanceKubernetesClusterResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyGovernanceKubernetesClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyGovernanceKubernetesClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyGovernanceKubernetesClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyGovernanceKubernetesClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyGovernanceKubernetesClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyGovernanceKubernetesClusterResponse) GetBody() *ModifyGovernanceKubernetesClusterResponseBody {
	return s.Body
}

func (s *ModifyGovernanceKubernetesClusterResponse) SetHeaders(v map[string]*string) *ModifyGovernanceKubernetesClusterResponse {
	s.Headers = v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponse) SetStatusCode(v int32) *ModifyGovernanceKubernetesClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponse) SetBody(v *ModifyGovernanceKubernetesClusterResponseBody) *ModifyGovernanceKubernetesClusterResponse {
	s.Body = v
	return s
}

func (s *ModifyGovernanceKubernetesClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
