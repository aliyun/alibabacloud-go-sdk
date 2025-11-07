// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecretParameterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSecretParameterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSecretParameterResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSecretParameterResponseBody) *UpdateSecretParameterResponse
	GetBody() *UpdateSecretParameterResponseBody
}

type UpdateSecretParameterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSecretParameterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSecretParameterResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecretParameterResponse) GoString() string {
	return s.String()
}

func (s *UpdateSecretParameterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSecretParameterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSecretParameterResponse) GetBody() *UpdateSecretParameterResponseBody {
	return s.Body
}

func (s *UpdateSecretParameterResponse) SetHeaders(v map[string]*string) *UpdateSecretParameterResponse {
	s.Headers = v
	return s
}

func (s *UpdateSecretParameterResponse) SetStatusCode(v int32) *UpdateSecretParameterResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSecretParameterResponse) SetBody(v *UpdateSecretParameterResponseBody) *UpdateSecretParameterResponse {
	s.Body = v
	return s
}

func (s *UpdateSecretParameterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
