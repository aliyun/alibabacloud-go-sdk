// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAppSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAppSecretResponse
	GetStatusCode() *int32
	SetBody(v *CreateAppSecretResponseBody) *CreateAppSecretResponse
	GetBody() *CreateAppSecretResponseBody
}

type CreateAppSecretResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAppSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateAppSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAppSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAppSecretResponse) GetBody() *CreateAppSecretResponseBody {
	return s.Body
}

func (s *CreateAppSecretResponse) SetHeaders(v map[string]*string) *CreateAppSecretResponse {
	s.Headers = v
	return s
}

func (s *CreateAppSecretResponse) SetStatusCode(v int32) *CreateAppSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppSecretResponse) SetBody(v *CreateAppSecretResponseBody) *CreateAppSecretResponse {
	s.Body = v
	return s
}

func (s *CreateAppSecretResponse) Validate() error {
	return dara.Validate(s)
}
