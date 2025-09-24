// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomerAddressListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCustomerAddressListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCustomerAddressListResponse
	GetStatusCode() *int32
	SetBody(v *QueryCustomerAddressListResponseBody) *QueryCustomerAddressListResponse
	GetBody() *QueryCustomerAddressListResponseBody
}

type QueryCustomerAddressListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCustomerAddressListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCustomerAddressListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomerAddressListResponse) GoString() string {
	return s.String()
}

func (s *QueryCustomerAddressListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCustomerAddressListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCustomerAddressListResponse) GetBody() *QueryCustomerAddressListResponseBody {
	return s.Body
}

func (s *QueryCustomerAddressListResponse) SetHeaders(v map[string]*string) *QueryCustomerAddressListResponse {
	s.Headers = v
	return s
}

func (s *QueryCustomerAddressListResponse) SetStatusCode(v int32) *QueryCustomerAddressListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCustomerAddressListResponse) SetBody(v *QueryCustomerAddressListResponseBody) *QueryCustomerAddressListResponse {
	s.Body = v
	return s
}

func (s *QueryCustomerAddressListResponse) Validate() error {
	return dara.Validate(s)
}
