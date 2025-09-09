// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKyuubiTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetKyuubiTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetKyuubiTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetKyuubiTokenResponseBody) *GetKyuubiTokenResponse
	GetBody() *GetKyuubiTokenResponseBody
}

type GetKyuubiTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetKyuubiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetKyuubiTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetKyuubiTokenResponse) GoString() string {
	return s.String()
}

func (s *GetKyuubiTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetKyuubiTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetKyuubiTokenResponse) GetBody() *GetKyuubiTokenResponseBody {
	return s.Body
}

func (s *GetKyuubiTokenResponse) SetHeaders(v map[string]*string) *GetKyuubiTokenResponse {
	s.Headers = v
	return s
}

func (s *GetKyuubiTokenResponse) SetStatusCode(v int32) *GetKyuubiTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetKyuubiTokenResponse) SetBody(v *GetKyuubiTokenResponseBody) *GetKyuubiTokenResponse {
	s.Body = v
	return s
}

func (s *GetKyuubiTokenResponse) Validate() error {
	return dara.Validate(s)
}
