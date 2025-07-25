// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDnsGtmAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDnsGtmAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDnsGtmAddressPoolResponseBody) *UpdateDnsGtmAddressPoolResponse
	GetBody() *UpdateDnsGtmAddressPoolResponseBody
}

type UpdateDnsGtmAddressPoolResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDnsGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDnsGtmAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDnsGtmAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDnsGtmAddressPoolResponse) GetBody() *UpdateDnsGtmAddressPoolResponseBody {
	return s.Body
}

func (s *UpdateDnsGtmAddressPoolResponse) SetHeaders(v map[string]*string) *UpdateDnsGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *UpdateDnsGtmAddressPoolResponse) SetStatusCode(v int32) *UpdateDnsGtmAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDnsGtmAddressPoolResponse) SetBody(v *UpdateDnsGtmAddressPoolResponseBody) *UpdateDnsGtmAddressPoolResponse {
	s.Body = v
	return s
}

func (s *UpdateDnsGtmAddressPoolResponse) Validate() error {
	return dara.Validate(s)
}
