// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKyuubiTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKyuubiTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKyuubiTokenResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKyuubiTokenResponseBody) *UpdateKyuubiTokenResponse
	GetBody() *UpdateKyuubiTokenResponseBody
}

type UpdateKyuubiTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKyuubiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKyuubiTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKyuubiTokenResponse) GoString() string {
	return s.String()
}

func (s *UpdateKyuubiTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKyuubiTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKyuubiTokenResponse) GetBody() *UpdateKyuubiTokenResponseBody {
	return s.Body
}

func (s *UpdateKyuubiTokenResponse) SetHeaders(v map[string]*string) *UpdateKyuubiTokenResponse {
	s.Headers = v
	return s
}

func (s *UpdateKyuubiTokenResponse) SetStatusCode(v int32) *UpdateKyuubiTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKyuubiTokenResponse) SetBody(v *UpdateKyuubiTokenResponseBody) *UpdateKyuubiTokenResponse {
	s.Body = v
	return s
}

func (s *UpdateKyuubiTokenResponse) Validate() error {
	return dara.Validate(s)
}
