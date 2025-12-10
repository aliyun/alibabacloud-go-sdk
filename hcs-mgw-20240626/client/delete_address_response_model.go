// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAddressResponse
	GetStatusCode() *int32
}

type DeleteAddressResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAddressResponse) SetHeaders(v map[string]*string) *DeleteAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteAddressResponse) SetStatusCode(v int32) *DeleteAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAddressResponse) Validate() error {
	return dara.Validate(s)
}
