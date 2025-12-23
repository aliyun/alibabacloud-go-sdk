// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserCertificateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserCertificateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserCertificateResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserCertificateResponseBody) *DeleteUserCertificateResponse
	GetBody() *DeleteUserCertificateResponseBody
}

type DeleteUserCertificateResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserCertificateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserCertificateResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserCertificateResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserCertificateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserCertificateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserCertificateResponse) GetBody() *DeleteUserCertificateResponseBody {
	return s.Body
}

func (s *DeleteUserCertificateResponse) SetHeaders(v map[string]*string) *DeleteUserCertificateResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserCertificateResponse) SetStatusCode(v int32) *DeleteUserCertificateResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserCertificateResponse) SetBody(v *DeleteUserCertificateResponseBody) *DeleteUserCertificateResponse {
	s.Body = v
	return s
}

func (s *DeleteUserCertificateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
