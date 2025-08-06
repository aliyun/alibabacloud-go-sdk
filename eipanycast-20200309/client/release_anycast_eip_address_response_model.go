// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseAnycastEipAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseAnycastEipAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseAnycastEipAddressResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseAnycastEipAddressResponseBody) *ReleaseAnycastEipAddressResponse
	GetBody() *ReleaseAnycastEipAddressResponseBody
}

type ReleaseAnycastEipAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseAnycastEipAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseAnycastEipAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseAnycastEipAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseAnycastEipAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseAnycastEipAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseAnycastEipAddressResponse) GetBody() *ReleaseAnycastEipAddressResponseBody {
	return s.Body
}

func (s *ReleaseAnycastEipAddressResponse) SetHeaders(v map[string]*string) *ReleaseAnycastEipAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseAnycastEipAddressResponse) SetStatusCode(v int32) *ReleaseAnycastEipAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseAnycastEipAddressResponse) SetBody(v *ReleaseAnycastEipAddressResponseBody) *ReleaseAnycastEipAddressResponse {
	s.Body = v
	return s
}

func (s *ReleaseAnycastEipAddressResponse) Validate() error {
	return dara.Validate(s)
}
