// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKyuubiTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateKyuubiTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateKyuubiTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateKyuubiTokenResponseBody) *CreateKyuubiTokenResponse
	GetBody() *CreateKyuubiTokenResponseBody
}

type CreateKyuubiTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateKyuubiTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateKyuubiTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateKyuubiTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateKyuubiTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateKyuubiTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateKyuubiTokenResponse) GetBody() *CreateKyuubiTokenResponseBody {
	return s.Body
}

func (s *CreateKyuubiTokenResponse) SetHeaders(v map[string]*string) *CreateKyuubiTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateKyuubiTokenResponse) SetStatusCode(v int32) *CreateKyuubiTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateKyuubiTokenResponse) SetBody(v *CreateKyuubiTokenResponseBody) *CreateKyuubiTokenResponse {
	s.Body = v
	return s
}

func (s *CreateKyuubiTokenResponse) Validate() error {
	return dara.Validate(s)
}
