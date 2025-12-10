// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAddressResponse
	GetStatusCode() *int32
	SetBody(v *ListAddressResponseBody) *ListAddressResponse
	GetBody() *ListAddressResponseBody
}

type ListAddressResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAddressResponse) GoString() string {
	return s.String()
}

func (s *ListAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAddressResponse) GetBody() *ListAddressResponseBody {
	return s.Body
}

func (s *ListAddressResponse) SetHeaders(v map[string]*string) *ListAddressResponse {
	s.Headers = v
	return s
}

func (s *ListAddressResponse) SetStatusCode(v int32) *ListAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddressResponse) SetBody(v *ListAddressResponseBody) *ListAddressResponse {
	s.Body = v
	return s
}

func (s *ListAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
