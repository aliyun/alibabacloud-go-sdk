// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DissociateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DissociateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *DissociateIpAddressResponseBody) *DissociateIpAddressResponse
	GetBody() *DissociateIpAddressResponseBody
}

type DissociateIpAddressResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DissociateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DissociateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DissociateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *DissociateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DissociateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DissociateIpAddressResponse) GetBody() *DissociateIpAddressResponseBody {
	return s.Body
}

func (s *DissociateIpAddressResponse) SetHeaders(v map[string]*string) *DissociateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *DissociateIpAddressResponse) SetStatusCode(v int32) *DissociateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DissociateIpAddressResponse) SetBody(v *DissociateIpAddressResponseBody) *DissociateIpAddressResponse {
	s.Body = v
	return s
}

func (s *DissociateIpAddressResponse) Validate() error {
	return dara.Validate(s)
}
