// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAppKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAppKeyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAppKeyResponseBody) *DeleteAppKeyResponse
	GetBody() *DeleteAppKeyResponseBody
}

type DeleteAppKeyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAppKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAppKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppKeyResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAppKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAppKeyResponse) GetBody() *DeleteAppKeyResponseBody {
	return s.Body
}

func (s *DeleteAppKeyResponse) SetHeaders(v map[string]*string) *DeleteAppKeyResponse {
	s.Headers = v
	return s
}

func (s *DeleteAppKeyResponse) SetStatusCode(v int32) *DeleteAppKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAppKeyResponse) SetBody(v *DeleteAppKeyResponseBody) *DeleteAppKeyResponse {
	s.Body = v
	return s
}

func (s *DeleteAppKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
