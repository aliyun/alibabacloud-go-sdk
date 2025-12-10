// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAddressResponse
	GetStatusCode() *int32
}

type UpdateAddressResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAddressResponse) GoString() string {
	return s.String()
}

func (s *UpdateAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAddressResponse) SetHeaders(v map[string]*string) *UpdateAddressResponse {
	s.Headers = v
	return s
}

func (s *UpdateAddressResponse) SetStatusCode(v int32) *UpdateAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAddressResponse) Validate() error {
	return dara.Validate(s)
}
