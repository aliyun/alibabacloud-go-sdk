// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLeniPrivateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLeniPrivateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLeniPrivateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLeniPrivateIpAddressResponseBody) *UpdateLeniPrivateIpAddressResponse
	GetBody() *UpdateLeniPrivateIpAddressResponseBody
}

type UpdateLeniPrivateIpAddressResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLeniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLeniPrivateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLeniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *UpdateLeniPrivateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLeniPrivateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLeniPrivateIpAddressResponse) GetBody() *UpdateLeniPrivateIpAddressResponseBody {
	return s.Body
}

func (s *UpdateLeniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *UpdateLeniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponse) SetStatusCode(v int32) *UpdateLeniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponse) SetBody(v *UpdateLeniPrivateIpAddressResponseBody) *UpdateLeniPrivateIpAddressResponse {
	s.Body = v
	return s
}

func (s *UpdateLeniPrivateIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
