// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDnsGtmAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDnsGtmAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDnsGtmAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDnsGtmAddressPoolResponseBody) *DeleteDnsGtmAddressPoolResponse
	GetBody() *DeleteDnsGtmAddressPoolResponseBody
}

type DeleteDnsGtmAddressPoolResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDnsGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDnsGtmAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDnsGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteDnsGtmAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDnsGtmAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDnsGtmAddressPoolResponse) GetBody() *DeleteDnsGtmAddressPoolResponseBody {
	return s.Body
}

func (s *DeleteDnsGtmAddressPoolResponse) SetHeaders(v map[string]*string) *DeleteDnsGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteDnsGtmAddressPoolResponse) SetStatusCode(v int32) *DeleteDnsGtmAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDnsGtmAddressPoolResponse) SetBody(v *DeleteDnsGtmAddressPoolResponseBody) *DeleteDnsGtmAddressPoolResponse {
	s.Body = v
	return s
}

func (s *DeleteDnsGtmAddressPoolResponse) Validate() error {
	return dara.Validate(s)
}
