// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteK8sSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteK8sSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteK8sSecretResponse
	GetStatusCode() *int32
	SetBody(v *DeleteK8sSecretResponseBody) *DeleteK8sSecretResponse
	GetBody() *DeleteK8sSecretResponseBody
}

type DeleteK8sSecretResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteK8sSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteK8sSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteK8sSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteK8sSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteK8sSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteK8sSecretResponse) GetBody() *DeleteK8sSecretResponseBody {
	return s.Body
}

func (s *DeleteK8sSecretResponse) SetHeaders(v map[string]*string) *DeleteK8sSecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteK8sSecretResponse) SetStatusCode(v int32) *DeleteK8sSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteK8sSecretResponse) SetBody(v *DeleteK8sSecretResponseBody) *DeleteK8sSecretResponse {
	s.Body = v
	return s
}

func (s *DeleteK8sSecretResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
