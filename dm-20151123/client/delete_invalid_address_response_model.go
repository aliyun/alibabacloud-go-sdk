// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInvalidAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInvalidAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInvalidAddressResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInvalidAddressResponseBody) *DeleteInvalidAddressResponse
	GetBody() *DeleteInvalidAddressResponseBody
}

type DeleteInvalidAddressResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInvalidAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInvalidAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInvalidAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteInvalidAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInvalidAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInvalidAddressResponse) GetBody() *DeleteInvalidAddressResponseBody {
	return s.Body
}

func (s *DeleteInvalidAddressResponse) SetHeaders(v map[string]*string) *DeleteInvalidAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteInvalidAddressResponse) SetStatusCode(v int32) *DeleteInvalidAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInvalidAddressResponse) SetBody(v *DeleteInvalidAddressResponseBody) *DeleteInvalidAddressResponse {
	s.Body = v
	return s
}

func (s *DeleteInvalidAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
