// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserKeyResponseBody) *DeleteUserKeyResponse
	GetBody() *DeleteUserKeyResponseBody
}

type DeleteUserKeyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserKeyResponse) GetBody() *DeleteUserKeyResponseBody {
	return s.Body
}

func (s *DeleteUserKeyResponse) SetHeaders(v map[string]*string) *DeleteUserKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserKeyResponse) SetStatusCode(v int32) *DeleteUserKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserKeyResponse) SetBody(v *DeleteUserKeyResponseBody) *DeleteUserKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteUserKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
