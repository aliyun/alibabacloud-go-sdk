// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCACertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCACertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCACertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCACertificateResponseBody) *DeleteCACertificateResponse
	GetBody() *DeleteCACertificateResponseBody
}

type DeleteCACertificateResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCACertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCACertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCACertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteCACertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCACertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCACertificateResponse) GetBody() *DeleteCACertificateResponseBody {
	return s.Body
}

func (s *DeleteCACertificateResponse) SetHeaders(v map[string]*string) *DeleteCACertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteCACertificateResponse) SetStatusCode(v int32) *DeleteCACertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCACertificateResponse) SetBody(v *DeleteCACertificateResponseBody) *DeleteCACertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteCACertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
