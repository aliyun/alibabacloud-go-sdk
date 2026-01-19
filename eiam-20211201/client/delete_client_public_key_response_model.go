// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClientPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClientPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClientPublicKeyResponseBody) *DeleteClientPublicKeyResponse
	GetBody() *DeleteClientPublicKeyResponseBody
}

type DeleteClientPublicKeyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClientPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClientPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteClientPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClientPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClientPublicKeyResponse) GetBody() *DeleteClientPublicKeyResponseBody {
	return s.Body
}

func (s *DeleteClientPublicKeyResponse) SetHeaders(v map[string]*string) *DeleteClientPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteClientPublicKeyResponse) SetStatusCode(v int32) *DeleteClientPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClientPublicKeyResponse) SetBody(v *DeleteClientPublicKeyResponseBody) *DeleteClientPublicKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteClientPublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
