// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserPublicKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserPublicKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserPublicKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserPublicKeyResponseBody) *DeleteUserPublicKeyResponse
	GetBody() *DeleteUserPublicKeyResponseBody
}

type DeleteUserPublicKeyResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserPublicKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserPublicKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserPublicKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserPublicKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserPublicKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserPublicKeyResponse) GetBody() *DeleteUserPublicKeyResponseBody {
	return s.Body
}

func (s *DeleteUserPublicKeyResponse) SetHeaders(v map[string]*string) *DeleteUserPublicKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserPublicKeyResponse) SetStatusCode(v int32) *DeleteUserPublicKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserPublicKeyResponse) SetBody(v *DeleteUserPublicKeyResponseBody) *DeleteUserPublicKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteUserPublicKeyResponse) Validate() error {
	return dara.Validate(s)
}
