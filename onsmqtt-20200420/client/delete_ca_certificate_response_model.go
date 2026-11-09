// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCaCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCaCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCaCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCaCertificateResponseBody) *DeleteCaCertificateResponse
	GetBody() *DeleteCaCertificateResponseBody
}

type DeleteCaCertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCaCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCaCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCaCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCaCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCaCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCaCertificateResponse) GetBody() *DeleteCaCertificateResponseBody {
	return s.Body
}

func (s *DeleteCaCertificateResponse) SetHeaders(v map[string]*string) *DeleteCaCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCaCertificateResponse) SetStatusCode(v int32) *DeleteCaCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCaCertificateResponse) SetBody(v *DeleteCaCertificateResponseBody) *DeleteCaCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteCaCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
