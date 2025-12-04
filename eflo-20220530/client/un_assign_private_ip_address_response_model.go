// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnAssignPrivateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnAssignPrivateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnAssignPrivateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *UnAssignPrivateIpAddressResponseBody) *UnAssignPrivateIpAddressResponse
	GetBody() *UnAssignPrivateIpAddressResponseBody
}

type UnAssignPrivateIpAddressResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnAssignPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnAssignPrivateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s UnAssignPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *UnAssignPrivateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnAssignPrivateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnAssignPrivateIpAddressResponse) GetBody() *UnAssignPrivateIpAddressResponseBody {
	return s.Body
}

func (s *UnAssignPrivateIpAddressResponse) SetHeaders(v map[string]*string) *UnAssignPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *UnAssignPrivateIpAddressResponse) SetStatusCode(v int32) *UnAssignPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UnAssignPrivateIpAddressResponse) SetBody(v *UnAssignPrivateIpAddressResponseBody) *UnAssignPrivateIpAddressResponse {
	s.Body = v
	return s
}

func (s *UnAssignPrivateIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
