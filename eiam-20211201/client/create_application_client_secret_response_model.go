// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationClientSecretResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationClientSecretResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationClientSecretResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationClientSecretResponseBody) *CreateApplicationClientSecretResponse
	GetBody() *CreateApplicationClientSecretResponseBody
}

type CreateApplicationClientSecretResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationClientSecretResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationClientSecretResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationClientSecretResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationClientSecretResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationClientSecretResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationClientSecretResponse) GetBody() *CreateApplicationClientSecretResponseBody {
	return s.Body
}

func (s *CreateApplicationClientSecretResponse) SetHeaders(v map[string]*string) *CreateApplicationClientSecretResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationClientSecretResponse) SetStatusCode(v int32) *CreateApplicationClientSecretResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationClientSecretResponse) SetBody(v *CreateApplicationClientSecretResponseBody) *CreateApplicationClientSecretResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationClientSecretResponse) Validate() error {
	return dara.Validate(s)
}
