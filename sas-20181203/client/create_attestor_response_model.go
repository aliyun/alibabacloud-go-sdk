// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttestorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAttestorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAttestorResponse
	GetStatusCode() *int32
	SetBody(v *CreateAttestorResponseBody) *CreateAttestorResponse
	GetBody() *CreateAttestorResponseBody
}

type CreateAttestorResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAttestorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAttestorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAttestorResponse) GoString() string {
	return s.String()
}

func (s *CreateAttestorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAttestorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAttestorResponse) GetBody() *CreateAttestorResponseBody {
	return s.Body
}

func (s *CreateAttestorResponse) SetHeaders(v map[string]*string) *CreateAttestorResponse {
	s.Headers = v
	return s
}

func (s *CreateAttestorResponse) SetStatusCode(v int32) *CreateAttestorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAttestorResponse) SetBody(v *CreateAttestorResponseBody) *CreateAttestorResponse {
	s.Body = v
	return s
}

func (s *CreateAttestorResponse) Validate() error {
	return dara.Validate(s)
}
