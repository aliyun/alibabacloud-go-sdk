// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKyuubiTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteKyuubiTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteKyuubiTokenResponse
	GetStatusCode() *int32
	SetBody(v *DeleteKyuubiTokenResponseBody) *DeleteKyuubiTokenResponse
	GetBody() *DeleteKyuubiTokenResponseBody
}

type DeleteKyuubiTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteKyuubiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteKyuubiTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteKyuubiTokenResponse) GoString() string {
	return s.String()
}

func (s *DeleteKyuubiTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteKyuubiTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteKyuubiTokenResponse) GetBody() *DeleteKyuubiTokenResponseBody {
	return s.Body
}

func (s *DeleteKyuubiTokenResponse) SetHeaders(v map[string]*string) *DeleteKyuubiTokenResponse {
	s.Headers = v
	return s
}

func (s *DeleteKyuubiTokenResponse) SetStatusCode(v int32) *DeleteKyuubiTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKyuubiTokenResponse) SetBody(v *DeleteKyuubiTokenResponseBody) *DeleteKyuubiTokenResponse {
	s.Body = v
	return s
}

func (s *DeleteKyuubiTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
