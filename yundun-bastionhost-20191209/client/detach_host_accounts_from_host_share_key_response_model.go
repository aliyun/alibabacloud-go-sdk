// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachHostAccountsFromHostShareKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetachHostAccountsFromHostShareKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetachHostAccountsFromHostShareKeyResponse
	GetStatusCode() *int32
	SetBody(v *DetachHostAccountsFromHostShareKeyResponseBody) *DetachHostAccountsFromHostShareKeyResponse
	GetBody() *DetachHostAccountsFromHostShareKeyResponseBody
}

type DetachHostAccountsFromHostShareKeyResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetachHostAccountsFromHostShareKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetachHostAccountsFromHostShareKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s DetachHostAccountsFromHostShareKeyResponse) GoString() string {
	return s.String()
}

func (s *DetachHostAccountsFromHostShareKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetachHostAccountsFromHostShareKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetachHostAccountsFromHostShareKeyResponse) GetBody() *DetachHostAccountsFromHostShareKeyResponseBody {
	return s.Body
}

func (s *DetachHostAccountsFromHostShareKeyResponse) SetHeaders(v map[string]*string) *DetachHostAccountsFromHostShareKeyResponse {
	s.Headers = v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponse) SetStatusCode(v int32) *DetachHostAccountsFromHostShareKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponse) SetBody(v *DetachHostAccountsFromHostShareKeyResponseBody) *DetachHostAccountsFromHostShareKeyResponse {
	s.Body = v
	return s
}

func (s *DetachHostAccountsFromHostShareKeyResponse) Validate() error {
	return dara.Validate(s)
}
