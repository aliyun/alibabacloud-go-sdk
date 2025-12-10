// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAddressResponse
	GetStatusCode() *int32
}

type CreateAddressResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAddressResponse) SetHeaders(v map[string]*string) *CreateAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateAddressResponse) SetStatusCode(v int32) *CreateAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAddressResponse) Validate() error {
	return dara.Validate(s)
}
