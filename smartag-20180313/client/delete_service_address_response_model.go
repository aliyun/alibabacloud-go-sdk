// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceAddressResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceAddressResponseBody) *DeleteServiceAddressResponse
	GetBody() *DeleteServiceAddressResponseBody
}

type DeleteServiceAddressResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceAddressResponse) GetBody() *DeleteServiceAddressResponseBody {
	return s.Body
}

func (s *DeleteServiceAddressResponse) SetHeaders(v map[string]*string) *DeleteServiceAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceAddressResponse) SetStatusCode(v int32) *DeleteServiceAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceAddressResponse) SetBody(v *DeleteServiceAddressResponseBody) *DeleteServiceAddressResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
