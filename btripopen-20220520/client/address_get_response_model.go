// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddressGetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddressGetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddressGetResponse
	GetStatusCode() *int32
	SetBody(v *AddressGetResponseBody) *AddressGetResponse
	GetBody() *AddressGetResponseBody
}

type AddressGetResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddressGetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddressGetResponse) String() string {
	return dara.Prettify(s)
}

func (s AddressGetResponse) GoString() string {
	return s.String()
}

func (s *AddressGetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddressGetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddressGetResponse) GetBody() *AddressGetResponseBody {
	return s.Body
}

func (s *AddressGetResponse) SetHeaders(v map[string]*string) *AddressGetResponse {
	s.Headers = v
	return s
}

func (s *AddressGetResponse) SetStatusCode(v int32) *AddressGetResponse {
	s.StatusCode = &v
	return s
}

func (s *AddressGetResponse) SetBody(v *AddressGetResponseBody) *AddressGetResponse {
	s.Body = v
	return s
}

func (s *AddressGetResponse) Validate() error {
	return dara.Validate(s)
}
