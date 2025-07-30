// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationClientSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteApplicationClientSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteApplicationClientSecretResponse
	GetStatusCode() *int32
	SetBody(v *DeleteApplicationClientSecretResponseBody) *DeleteApplicationClientSecretResponse
	GetBody() *DeleteApplicationClientSecretResponseBody
}

type DeleteApplicationClientSecretResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteApplicationClientSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationClientSecretResponse) GoString() string {
	return s.String()
}

func (s *DeleteApplicationClientSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteApplicationClientSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteApplicationClientSecretResponse) GetBody() *DeleteApplicationClientSecretResponseBody {
	return s.Body
}

func (s *DeleteApplicationClientSecretResponse) SetHeaders(v map[string]*string) *DeleteApplicationClientSecretResponse {
	s.Headers = v
	return s
}

func (s *DeleteApplicationClientSecretResponse) SetStatusCode(v int32) *DeleteApplicationClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteApplicationClientSecretResponse) SetBody(v *DeleteApplicationClientSecretResponseBody) *DeleteApplicationClientSecretResponse {
	s.Body = v
	return s
}

func (s *DeleteApplicationClientSecretResponse) Validate() error {
	return dara.Validate(s)
}
