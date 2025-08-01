// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGovernanceKubernetesClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGovernanceKubernetesClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGovernanceKubernetesClusterResponse
	GetStatusCode() *int32
	SetBody(v *QueryGovernanceKubernetesClusterResponseBody) *QueryGovernanceKubernetesClusterResponse
	GetBody() *QueryGovernanceKubernetesClusterResponseBody
}

type QueryGovernanceKubernetesClusterResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGovernanceKubernetesClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGovernanceKubernetesClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGovernanceKubernetesClusterResponse) GoString() string {
	return s.String()
}

func (s *QueryGovernanceKubernetesClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGovernanceKubernetesClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGovernanceKubernetesClusterResponse) GetBody() *QueryGovernanceKubernetesClusterResponseBody {
	return s.Body
}

func (s *QueryGovernanceKubernetesClusterResponse) SetHeaders(v map[string]*string) *QueryGovernanceKubernetesClusterResponse {
	s.Headers = v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponse) SetStatusCode(v int32) *QueryGovernanceKubernetesClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponse) SetBody(v *QueryGovernanceKubernetesClusterResponseBody) *QueryGovernanceKubernetesClusterResponse {
	s.Body = v
	return s
}

func (s *QueryGovernanceKubernetesClusterResponse) Validate() error {
	return dara.Validate(s)
}
