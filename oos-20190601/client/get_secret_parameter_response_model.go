// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecretParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSecretParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSecretParameterResponse
	GetStatusCode() *int32
	SetBody(v *GetSecretParameterResponseBody) *GetSecretParameterResponse
	GetBody() *GetSecretParameterResponseBody
}

type GetSecretParameterResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSecretParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *GetSecretParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSecretParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSecretParameterResponse) GetBody() *GetSecretParameterResponseBody {
	return s.Body
}

func (s *GetSecretParameterResponse) SetHeaders(v map[string]*string) *GetSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *GetSecretParameterResponse) SetStatusCode(v int32) *GetSecretParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSecretParameterResponse) SetBody(v *GetSecretParameterResponseBody) *GetSecretParameterResponse {
	s.Body = v
	return s
}

func (s *GetSecretParameterResponse) Validate() error {
	return dara.Validate(s)
}
