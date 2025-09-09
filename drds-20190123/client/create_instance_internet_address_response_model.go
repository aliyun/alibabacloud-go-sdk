// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceInternetAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceInternetAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceInternetAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceInternetAddressResponseBody) *CreateInstanceInternetAddressResponse
	GetBody() *CreateInstanceInternetAddressResponseBody
}

type CreateInstanceInternetAddressResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceInternetAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceInternetAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceInternetAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceInternetAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceInternetAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceInternetAddressResponse) GetBody() *CreateInstanceInternetAddressResponseBody {
	return s.Body
}

func (s *CreateInstanceInternetAddressResponse) SetHeaders(v map[string]*string) *CreateInstanceInternetAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceInternetAddressResponse) SetStatusCode(v int32) *CreateInstanceInternetAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceInternetAddressResponse) SetBody(v *CreateInstanceInternetAddressResponseBody) *CreateInstanceInternetAddressResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceInternetAddressResponse) Validate() error {
	return dara.Validate(s)
}
