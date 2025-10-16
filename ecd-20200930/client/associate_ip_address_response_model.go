// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssociateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssociateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *AssociateIpAddressResponseBody) *AssociateIpAddressResponse
	GetBody() *AssociateIpAddressResponseBody
}

type AssociateIpAddressResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssociateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssociateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AssociateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *AssociateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssociateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssociateIpAddressResponse) GetBody() *AssociateIpAddressResponseBody {
	return s.Body
}

func (s *AssociateIpAddressResponse) SetHeaders(v map[string]*string) *AssociateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *AssociateIpAddressResponse) SetStatusCode(v int32) *AssociateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AssociateIpAddressResponse) SetBody(v *AssociateIpAddressResponseBody) *AssociateIpAddressResponse {
	s.Body = v
	return s
}

func (s *AssociateIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
