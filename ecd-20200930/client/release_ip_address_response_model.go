// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseIpAddressResponseBody) *ReleaseIpAddressResponse
	GetBody() *ReleaseIpAddressResponseBody
}

type ReleaseIpAddressResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseIpAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseIpAddressResponse) GetBody() *ReleaseIpAddressResponseBody {
	return s.Body
}

func (s *ReleaseIpAddressResponse) SetHeaders(v map[string]*string) *ReleaseIpAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseIpAddressResponse) SetStatusCode(v int32) *ReleaseIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseIpAddressResponse) SetBody(v *ReleaseIpAddressResponseBody) *ReleaseIpAddressResponse {
	s.Body = v
	return s
}

func (s *ReleaseIpAddressResponse) Validate() error {
	return dara.Validate(s)
}
