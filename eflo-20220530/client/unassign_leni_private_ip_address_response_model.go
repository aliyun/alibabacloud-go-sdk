// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassignLeniPrivateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnassignLeniPrivateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnassignLeniPrivateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *UnassignLeniPrivateIpAddressResponseBody) *UnassignLeniPrivateIpAddressResponse
	GetBody() *UnassignLeniPrivateIpAddressResponseBody
}

type UnassignLeniPrivateIpAddressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnassignLeniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnassignLeniPrivateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s UnassignLeniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *UnassignLeniPrivateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnassignLeniPrivateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnassignLeniPrivateIpAddressResponse) GetBody() *UnassignLeniPrivateIpAddressResponseBody {
	return s.Body
}

func (s *UnassignLeniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *UnassignLeniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponse) SetStatusCode(v int32) *UnassignLeniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponse) SetBody(v *UnassignLeniPrivateIpAddressResponseBody) *UnassignLeniPrivateIpAddressResponse {
	s.Body = v
	return s
}

func (s *UnassignLeniPrivateIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
