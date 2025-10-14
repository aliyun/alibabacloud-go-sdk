// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsGtmAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDnsGtmAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDnsGtmAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *AddDnsGtmAddressPoolResponseBody) *AddDnsGtmAddressPoolResponse
	GetBody() *AddDnsGtmAddressPoolResponseBody
}

type AddDnsGtmAddressPoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDnsGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDnsGtmAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDnsGtmAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDnsGtmAddressPoolResponse) GetBody() *AddDnsGtmAddressPoolResponseBody {
	return s.Body
}

func (s *AddDnsGtmAddressPoolResponse) SetHeaders(v map[string]*string) *AddDnsGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *AddDnsGtmAddressPoolResponse) SetStatusCode(v int32) *AddDnsGtmAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDnsGtmAddressPoolResponse) SetBody(v *AddDnsGtmAddressPoolResponseBody) *AddDnsGtmAddressPoolResponse {
	s.Body = v
	return s
}

func (s *AddDnsGtmAddressPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
