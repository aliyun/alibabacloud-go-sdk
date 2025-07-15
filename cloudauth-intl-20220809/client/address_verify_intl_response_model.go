// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressVerifyIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddressVerifyIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddressVerifyIntlResponse
	GetStatusCode() *int32
	SetBody(v *AddressVerifyIntlResponseBody) *AddressVerifyIntlResponse
	GetBody() *AddressVerifyIntlResponseBody
}

type AddressVerifyIntlResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddressVerifyIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddressVerifyIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s AddressVerifyIntlResponse) GoString() string {
	return s.String()
}

func (s *AddressVerifyIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddressVerifyIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddressVerifyIntlResponse) GetBody() *AddressVerifyIntlResponseBody {
	return s.Body
}

func (s *AddressVerifyIntlResponse) SetHeaders(v map[string]*string) *AddressVerifyIntlResponse {
	s.Headers = v
	return s
}

func (s *AddressVerifyIntlResponse) SetStatusCode(v int32) *AddressVerifyIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *AddressVerifyIntlResponse) SetBody(v *AddressVerifyIntlResponseBody) *AddressVerifyIntlResponse {
	s.Body = v
	return s
}

func (s *AddressVerifyIntlResponse) Validate() error {
	return dara.Validate(s)
}
