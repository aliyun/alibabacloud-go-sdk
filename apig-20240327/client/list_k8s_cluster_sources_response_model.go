// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sClusterSourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListK8sClusterSourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListK8sClusterSourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListK8sClusterSourcesResponseBody) *ListK8sClusterSourcesResponse
	GetBody() *ListK8sClusterSourcesResponseBody
}

type ListK8sClusterSourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListK8sClusterSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListK8sClusterSourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListK8sClusterSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListK8sClusterSourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListK8sClusterSourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListK8sClusterSourcesResponse) GetBody() *ListK8sClusterSourcesResponseBody {
	return s.Body
}

func (s *ListK8sClusterSourcesResponse) SetHeaders(v map[string]*string) *ListK8sClusterSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListK8sClusterSourcesResponse) SetStatusCode(v int32) *ListK8sClusterSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListK8sClusterSourcesResponse) SetBody(v *ListK8sClusterSourcesResponseBody) *ListK8sClusterSourcesResponse {
	s.Body = v
	return s
}

func (s *ListK8sClusterSourcesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
