// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAttestorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAttestorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAttestorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAttestorResponseBody) *DeleteAttestorResponse
	GetBody() *DeleteAttestorResponseBody
}

type DeleteAttestorResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAttestorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAttestorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAttestorResponse) GoString() string {
	return s.String()
}

func (s *DeleteAttestorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAttestorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAttestorResponse) GetBody() *DeleteAttestorResponseBody {
	return s.Body
}

func (s *DeleteAttestorResponse) SetHeaders(v map[string]*string) *DeleteAttestorResponse {
	s.Headers = v
	return s
}

func (s *DeleteAttestorResponse) SetStatusCode(v int32) *DeleteAttestorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAttestorResponse) SetBody(v *DeleteAttestorResponseBody) *DeleteAttestorResponse {
	s.Body = v
	return s
}

func (s *DeleteAttestorResponse) Validate() error {
	return dara.Validate(s)
}
