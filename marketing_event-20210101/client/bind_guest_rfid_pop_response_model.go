// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindGuestRfidPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindGuestRfidPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindGuestRfidPopResponse
	GetStatusCode() *int32
	SetBody(v *BindGuestRfidPopResponseBody) *BindGuestRfidPopResponse
	GetBody() *BindGuestRfidPopResponseBody
}

type BindGuestRfidPopResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindGuestRfidPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindGuestRfidPopResponse) String() string {
	return dara.Prettify(s)
}

func (s BindGuestRfidPopResponse) GoString() string {
	return s.String()
}

func (s *BindGuestRfidPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindGuestRfidPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindGuestRfidPopResponse) GetBody() *BindGuestRfidPopResponseBody {
	return s.Body
}

func (s *BindGuestRfidPopResponse) SetHeaders(v map[string]*string) *BindGuestRfidPopResponse {
	s.Headers = v
	return s
}

func (s *BindGuestRfidPopResponse) SetStatusCode(v int32) *BindGuestRfidPopResponse {
	s.StatusCode = &v
	return s
}

func (s *BindGuestRfidPopResponse) SetBody(v *BindGuestRfidPopResponseBody) *BindGuestRfidPopResponse {
	s.Body = v
	return s
}

func (s *BindGuestRfidPopResponse) Validate() error {
	return dara.Validate(s)
}
