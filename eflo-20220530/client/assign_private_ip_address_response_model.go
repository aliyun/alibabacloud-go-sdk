// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignPrivateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignPrivateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignPrivateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *AssignPrivateIpAddressResponseBody) *AssignPrivateIpAddressResponse
	GetBody() *AssignPrivateIpAddressResponseBody
}

type AssignPrivateIpAddressResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignPrivateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *AssignPrivateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignPrivateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignPrivateIpAddressResponse) GetBody() *AssignPrivateIpAddressResponseBody {
	return s.Body
}

func (s *AssignPrivateIpAddressResponse) SetHeaders(v map[string]*string) *AssignPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *AssignPrivateIpAddressResponse) SetStatusCode(v int32) *AssignPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignPrivateIpAddressResponse) SetBody(v *AssignPrivateIpAddressResponseBody) *AssignPrivateIpAddressResponse {
	s.Body = v
	return s
}

func (s *AssignPrivateIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
