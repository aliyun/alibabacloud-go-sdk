// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateK8sSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateK8sSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateK8sSecretResponse
	GetStatusCode() *int32
	SetBody(v *UpdateK8sSecretResponseBody) *UpdateK8sSecretResponse
	GetBody() *UpdateK8sSecretResponseBody
}

type UpdateK8sSecretResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateK8sSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateK8sSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateK8sSecretResponse) GoString() string {
	return s.String()
}

func (s *UpdateK8sSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateK8sSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateK8sSecretResponse) GetBody() *UpdateK8sSecretResponseBody {
	return s.Body
}

func (s *UpdateK8sSecretResponse) SetHeaders(v map[string]*string) *UpdateK8sSecretResponse {
	s.Headers = v
	return s
}

func (s *UpdateK8sSecretResponse) SetStatusCode(v int32) *UpdateK8sSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateK8sSecretResponse) SetBody(v *UpdateK8sSecretResponseBody) *UpdateK8sSecretResponse {
	s.Body = v
	return s
}

func (s *UpdateK8sSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
