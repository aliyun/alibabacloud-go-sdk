// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPrimaryClientPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetPrimaryClientPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetPrimaryClientPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *SetPrimaryClientPublicKeyResponseBody) *SetPrimaryClientPublicKeyResponse
	GetBody() *SetPrimaryClientPublicKeyResponseBody
}

type SetPrimaryClientPublicKeyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPrimaryClientPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetPrimaryClientPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s SetPrimaryClientPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *SetPrimaryClientPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetPrimaryClientPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetPrimaryClientPublicKeyResponse) GetBody() *SetPrimaryClientPublicKeyResponseBody {
	return s.Body
}

func (s *SetPrimaryClientPublicKeyResponse) SetHeaders(v map[string]*string) *SetPrimaryClientPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *SetPrimaryClientPublicKeyResponse) SetStatusCode(v int32) *SetPrimaryClientPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPrimaryClientPublicKeyResponse) SetBody(v *SetPrimaryClientPublicKeyResponseBody) *SetPrimaryClientPublicKeyResponse {
	s.Body = v
	return s
}

func (s *SetPrimaryClientPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
