// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSecretParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSecretParameterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSecretParameterResponseBody) *DeleteSecretParameterResponse
	GetBody() *DeleteSecretParameterResponseBody
}

type DeleteSecretParameterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSecretParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecretParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSecretParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSecretParameterResponse) GetBody() *DeleteSecretParameterResponseBody {
	return s.Body
}

func (s *DeleteSecretParameterResponse) SetHeaders(v map[string]*string) *DeleteSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecretParameterResponse) SetStatusCode(v int32) *DeleteSecretParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecretParameterResponse) SetBody(v *DeleteSecretParameterResponseBody) *DeleteSecretParameterResponse {
	s.Body = v
	return s
}

func (s *DeleteSecretParameterResponse) Validate() error {
	return dara.Validate(s)
}
