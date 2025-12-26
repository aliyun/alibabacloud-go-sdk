// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAddressResponse
	GetStatusCode() *int32
	SetBody(v *AddAddressResponseBody) *AddAddressResponse
	GetBody() *AddAddressResponseBody
}

type AddAddressResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAddressResponse) GoString() string {
	return s.String()
}

func (s *AddAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAddressResponse) GetBody() *AddAddressResponseBody {
	return s.Body
}

func (s *AddAddressResponse) SetHeaders(v map[string]*string) *AddAddressResponse {
	s.Headers = v
	return s
}

func (s *AddAddressResponse) SetStatusCode(v int32) *AddAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAddressResponse) SetBody(v *AddAddressResponseBody) *AddAddressResponse {
	s.Body = v
	return s
}

func (s *AddAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
