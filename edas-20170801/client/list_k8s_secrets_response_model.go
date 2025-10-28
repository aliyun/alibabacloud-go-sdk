// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sSecretsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListK8sSecretsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListK8sSecretsResponse
	GetStatusCode() *int32
	SetBody(v *ListK8sSecretsResponseBody) *ListK8sSecretsResponse
	GetBody() *ListK8sSecretsResponseBody
}

type ListK8sSecretsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListK8sSecretsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListK8sSecretsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListK8sSecretsResponse) GoString() string {
	return s.String()
}

func (s *ListK8sSecretsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListK8sSecretsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListK8sSecretsResponse) GetBody() *ListK8sSecretsResponseBody {
	return s.Body
}

func (s *ListK8sSecretsResponse) SetHeaders(v map[string]*string) *ListK8sSecretsResponse {
	s.Headers = v
	return s
}

func (s *ListK8sSecretsResponse) SetStatusCode(v int32) *ListK8sSecretsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListK8sSecretsResponse) SetBody(v *ListK8sSecretsResponseBody) *ListK8sSecretsResponse {
	s.Body = v
	return s
}

func (s *ListK8sSecretsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
