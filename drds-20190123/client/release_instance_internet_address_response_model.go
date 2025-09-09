// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceInternetAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseInstanceInternetAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseInstanceInternetAddressResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseInstanceInternetAddressResponseBody) *ReleaseInstanceInternetAddressResponse
	GetBody() *ReleaseInstanceInternetAddressResponseBody
}

type ReleaseInstanceInternetAddressResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseInstanceInternetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseInstanceInternetAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceInternetAddressResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceInternetAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseInstanceInternetAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseInstanceInternetAddressResponse) GetBody() *ReleaseInstanceInternetAddressResponseBody {
	return s.Body
}

func (s *ReleaseInstanceInternetAddressResponse) SetHeaders(v map[string]*string) *ReleaseInstanceInternetAddressResponse {
	s.Headers = v
	return s
}

func (s *ReleaseInstanceInternetAddressResponse) SetStatusCode(v int32) *ReleaseInstanceInternetAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseInstanceInternetAddressResponse) SetBody(v *ReleaseInstanceInternetAddressResponseBody) *ReleaseInstanceInternetAddressResponse {
	s.Body = v
	return s
}

func (s *ReleaseInstanceInternetAddressResponse) Validate() error {
	return dara.Validate(s)
}
