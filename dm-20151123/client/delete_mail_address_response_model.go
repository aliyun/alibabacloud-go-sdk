// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMailAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMailAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMailAddressResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMailAddressResponseBody) *DeleteMailAddressResponse
	GetBody() *DeleteMailAddressResponseBody
}

type DeleteMailAddressResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMailAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMailAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMailAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteMailAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMailAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMailAddressResponse) GetBody() *DeleteMailAddressResponseBody {
	return s.Body
}

func (s *DeleteMailAddressResponse) SetHeaders(v map[string]*string) *DeleteMailAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteMailAddressResponse) SetStatusCode(v int32) *DeleteMailAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMailAddressResponse) SetBody(v *DeleteMailAddressResponseBody) *DeleteMailAddressResponse {
	s.Body = v
	return s
}

func (s *DeleteMailAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
