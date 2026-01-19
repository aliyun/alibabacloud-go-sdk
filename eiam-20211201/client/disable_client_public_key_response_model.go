// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableClientPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableClientPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableClientPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *DisableClientPublicKeyResponseBody) *DisableClientPublicKeyResponse
	GetBody() *DisableClientPublicKeyResponseBody
}

type DisableClientPublicKeyResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableClientPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableClientPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableClientPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *DisableClientPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableClientPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableClientPublicKeyResponse) GetBody() *DisableClientPublicKeyResponseBody {
	return s.Body
}

func (s *DisableClientPublicKeyResponse) SetHeaders(v map[string]*string) *DisableClientPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *DisableClientPublicKeyResponse) SetStatusCode(v int32) *DisableClientPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableClientPublicKeyResponse) SetBody(v *DisableClientPublicKeyResponseBody) *DisableClientPublicKeyResponse {
	s.Body = v
	return s
}

func (s *DisableClientPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
