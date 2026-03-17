// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableServiceAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableServiceAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableServiceAddressResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableServiceAddressResponseBody) *ListAvailableServiceAddressResponse
	GetBody() *ListAvailableServiceAddressResponseBody
}

type ListAvailableServiceAddressResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableServiceAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableServiceAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableServiceAddressResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableServiceAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableServiceAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableServiceAddressResponse) GetBody() *ListAvailableServiceAddressResponseBody {
	return s.Body
}

func (s *ListAvailableServiceAddressResponse) SetHeaders(v map[string]*string) *ListAvailableServiceAddressResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableServiceAddressResponse) SetStatusCode(v int32) *ListAvailableServiceAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableServiceAddressResponse) SetBody(v *ListAvailableServiceAddressResponseBody) *ListAvailableServiceAddressResponse {
	s.Body = v
	return s
}

func (s *ListAvailableServiceAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
