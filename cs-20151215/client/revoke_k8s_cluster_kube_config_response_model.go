// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeK8sClusterKubeConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeK8sClusterKubeConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeK8sClusterKubeConfigResponse
	GetStatusCode() *int32
	SetBody(v *RevokeK8sClusterKubeConfigResponseBody) *RevokeK8sClusterKubeConfigResponse
	GetBody() *RevokeK8sClusterKubeConfigResponseBody
}

type RevokeK8sClusterKubeConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeK8sClusterKubeConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeK8sClusterKubeConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeK8sClusterKubeConfigResponse) GoString() string {
	return s.String()
}

func (s *RevokeK8sClusterKubeConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeK8sClusterKubeConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeK8sClusterKubeConfigResponse) GetBody() *RevokeK8sClusterKubeConfigResponseBody {
	return s.Body
}

func (s *RevokeK8sClusterKubeConfigResponse) SetHeaders(v map[string]*string) *RevokeK8sClusterKubeConfigResponse {
	s.Headers = v
	return s
}

func (s *RevokeK8sClusterKubeConfigResponse) SetStatusCode(v int32) *RevokeK8sClusterKubeConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeK8sClusterKubeConfigResponse) SetBody(v *RevokeK8sClusterKubeConfigResponseBody) *RevokeK8sClusterKubeConfigResponse {
	s.Body = v
	return s
}

func (s *RevokeK8sClusterKubeConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
