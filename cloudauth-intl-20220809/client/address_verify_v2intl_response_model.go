// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressVerifyV2IntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddressVerifyV2IntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddressVerifyV2IntlResponse
	GetStatusCode() *int32
	SetBody(v *AddressVerifyV2IntlResponseBody) *AddressVerifyV2IntlResponse
	GetBody() *AddressVerifyV2IntlResponseBody
}

type AddressVerifyV2IntlResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddressVerifyV2IntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddressVerifyV2IntlResponse) String() string {
	return dara.Prettify(s)
}

func (s AddressVerifyV2IntlResponse) GoString() string {
	return s.String()
}

func (s *AddressVerifyV2IntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddressVerifyV2IntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddressVerifyV2IntlResponse) GetBody() *AddressVerifyV2IntlResponseBody {
	return s.Body
}

func (s *AddressVerifyV2IntlResponse) SetHeaders(v map[string]*string) *AddressVerifyV2IntlResponse {
	s.Headers = v
	return s
}

func (s *AddressVerifyV2IntlResponse) SetStatusCode(v int32) *AddressVerifyV2IntlResponse {
	s.StatusCode = &v
	return s
}

func (s *AddressVerifyV2IntlResponse) SetBody(v *AddressVerifyV2IntlResponseBody) *AddressVerifyV2IntlResponse {
	s.Body = v
	return s
}

func (s *AddressVerifyV2IntlResponse) Validate() error {
	return dara.Validate(s)
}
