// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteOriginCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteOriginCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteOriginCaCertificateResponseBody) *DeleteOriginCaCertificateResponse
	GetBody() *DeleteOriginCaCertificateResponseBody
}

type DeleteOriginCaCertificateResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteOriginCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteOriginCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteOriginCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteOriginCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteOriginCaCertificateResponse) GetBody() *DeleteOriginCaCertificateResponseBody {
	return s.Body
}

func (s *DeleteOriginCaCertificateResponse) SetHeaders(v map[string]*string) *DeleteOriginCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteOriginCaCertificateResponse) SetStatusCode(v int32) *DeleteOriginCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOriginCaCertificateResponse) SetBody(v *DeleteOriginCaCertificateResponseBody) *DeleteOriginCaCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteOriginCaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
