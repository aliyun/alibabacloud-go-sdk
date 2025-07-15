// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecretParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSecretParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSecretParameterResponse
	GetStatusCode() *int32
	SetBody(v *CreateSecretParameterResponseBody) *CreateSecretParameterResponse
	GetBody() *CreateSecretParameterResponseBody
}

type CreateSecretParameterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSecretParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *CreateSecretParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSecretParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSecretParameterResponse) GetBody() *CreateSecretParameterResponseBody {
	return s.Body
}

func (s *CreateSecretParameterResponse) SetHeaders(v map[string]*string) *CreateSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *CreateSecretParameterResponse) SetStatusCode(v int32) *CreateSecretParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecretParameterResponse) SetBody(v *CreateSecretParameterResponseBody) *CreateSecretParameterResponse {
	s.Body = v
	return s
}

func (s *CreateSecretParameterResponse) Validate() error {
	return dara.Validate(s)
}
