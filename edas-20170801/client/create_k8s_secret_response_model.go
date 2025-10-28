// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateK8sSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateK8sSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateK8sSecretResponse
	GetStatusCode() *int32
	SetBody(v *CreateK8sSecretResponseBody) *CreateK8sSecretResponse
	GetBody() *CreateK8sSecretResponseBody
}

type CreateK8sSecretResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateK8sSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateK8sSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateK8sSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateK8sSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateK8sSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateK8sSecretResponse) GetBody() *CreateK8sSecretResponseBody {
	return s.Body
}

func (s *CreateK8sSecretResponse) SetHeaders(v map[string]*string) *CreateK8sSecretResponse {
	s.Headers = v
	return s
}

func (s *CreateK8sSecretResponse) SetStatusCode(v int32) *CreateK8sSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateK8sSecretResponse) SetBody(v *CreateK8sSecretResponseBody) *CreateK8sSecretResponse {
	s.Body = v
	return s
}

func (s *CreateK8sSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
