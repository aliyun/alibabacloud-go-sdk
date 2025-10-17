// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKyuubiTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKyuubiTokenResponse
	GetStatusCode() *int32
	SetBody(v *ListKyuubiTokenResponseBody) *ListKyuubiTokenResponse
	GetBody() *ListKyuubiTokenResponseBody
}

type ListKyuubiTokenResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKyuubiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKyuubiTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiTokenResponse) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKyuubiTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKyuubiTokenResponse) GetBody() *ListKyuubiTokenResponseBody {
	return s.Body
}

func (s *ListKyuubiTokenResponse) SetHeaders(v map[string]*string) *ListKyuubiTokenResponse {
	s.Headers = v
	return s
}

func (s *ListKyuubiTokenResponse) SetStatusCode(v int32) *ListKyuubiTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKyuubiTokenResponse) SetBody(v *ListKyuubiTokenResponseBody) *ListKyuubiTokenResponse {
	s.Body = v
	return s
}

func (s *ListKyuubiTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
