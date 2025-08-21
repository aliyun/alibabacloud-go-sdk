// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhiteIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateWhiteIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateWhiteIpsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateWhiteIpsResponseBody) *UpdateWhiteIpsResponse
	GetBody() *UpdateWhiteIpsResponseBody
}

type UpdateWhiteIpsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateWhiteIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdateWhiteIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateWhiteIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateWhiteIpsResponse) GetBody() *UpdateWhiteIpsResponseBody {
	return s.Body
}

func (s *UpdateWhiteIpsResponse) SetHeaders(v map[string]*string) *UpdateWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdateWhiteIpsResponse) SetStatusCode(v int32) *UpdateWhiteIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateWhiteIpsResponse) SetBody(v *UpdateWhiteIpsResponseBody) *UpdateWhiteIpsResponse {
	s.Body = v
	return s
}

func (s *UpdateWhiteIpsResponse) Validate() error {
	return dara.Validate(s)
}
