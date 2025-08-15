// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressCompareIntlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddressCompareIntlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddressCompareIntlResponse
	GetStatusCode() *int32
	SetBody(v *AddressCompareIntlResponseBody) *AddressCompareIntlResponse
	GetBody() *AddressCompareIntlResponseBody
}

type AddressCompareIntlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddressCompareIntlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddressCompareIntlResponse) String() string {
	return dara.Prettify(s)
}

func (s AddressCompareIntlResponse) GoString() string {
	return s.String()
}

func (s *AddressCompareIntlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddressCompareIntlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddressCompareIntlResponse) GetBody() *AddressCompareIntlResponseBody {
	return s.Body
}

func (s *AddressCompareIntlResponse) SetHeaders(v map[string]*string) *AddressCompareIntlResponse {
	s.Headers = v
	return s
}

func (s *AddressCompareIntlResponse) SetStatusCode(v int32) *AddressCompareIntlResponse {
	s.StatusCode = &v
	return s
}

func (s *AddressCompareIntlResponse) SetBody(v *AddressCompareIntlResponseBody) *AddressCompareIntlResponse {
	s.Body = v
	return s
}

func (s *AddressCompareIntlResponse) Validate() error {
	return dara.Validate(s)
}
