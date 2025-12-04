// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignLeniPrivateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignLeniPrivateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignLeniPrivateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *AssignLeniPrivateIpAddressResponseBody) *AssignLeniPrivateIpAddressResponse
	GetBody() *AssignLeniPrivateIpAddressResponseBody
}

type AssignLeniPrivateIpAddressResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignLeniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignLeniPrivateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignLeniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *AssignLeniPrivateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignLeniPrivateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignLeniPrivateIpAddressResponse) GetBody() *AssignLeniPrivateIpAddressResponseBody {
	return s.Body
}

func (s *AssignLeniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *AssignLeniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *AssignLeniPrivateIpAddressResponse) SetStatusCode(v int32) *AssignLeniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignLeniPrivateIpAddressResponse) SetBody(v *AssignLeniPrivateIpAddressResponseBody) *AssignLeniPrivateIpAddressResponse {
	s.Body = v
	return s
}

func (s *AssignLeniPrivateIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
