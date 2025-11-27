// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeletePublicKeyResponseBody) *DeletePublicKeyResponse
	GetBody() *DeletePublicKeyResponseBody
}

type DeletePublicKeyResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePublicKeyResponse) GoString() string {
	return s.String()
}

func (s *DeletePublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePublicKeyResponse) GetBody() *DeletePublicKeyResponseBody {
	return s.Body
}

func (s *DeletePublicKeyResponse) SetHeaders(v map[string]*string) *DeletePublicKeyResponse {
	s.Headers = v
	return s
}

func (s *DeletePublicKeyResponse) SetStatusCode(v int32) *DeletePublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePublicKeyResponse) SetBody(v *DeletePublicKeyResponseBody) *DeletePublicKeyResponse {
	s.Body = v
	return s
}

func (s *DeletePublicKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
