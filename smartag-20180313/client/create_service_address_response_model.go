// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateServiceAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateServiceAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateServiceAddressResponseBody) *CreateServiceAddressResponse
	GetBody() *CreateServiceAddressResponseBody
}

type CreateServiceAddressResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateServiceAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateServiceAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateServiceAddressResponse) GetBody() *CreateServiceAddressResponseBody {
	return s.Body
}

func (s *CreateServiceAddressResponse) SetHeaders(v map[string]*string) *CreateServiceAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceAddressResponse) SetStatusCode(v int32) *CreateServiceAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceAddressResponse) SetBody(v *CreateServiceAddressResponseBody) *CreateServiceAddressResponse {
	s.Body = v
	return s
}

func (s *CreateServiceAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
