// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLniPrivateIpAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLniPrivateIpAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLniPrivateIpAddressResponse
	GetStatusCode() *int32
	SetBody(v *ListLniPrivateIpAddressResponseBody) *ListLniPrivateIpAddressResponse
	GetBody() *ListLniPrivateIpAddressResponseBody
}

type ListLniPrivateIpAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLniPrivateIpAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLniPrivateIpAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLniPrivateIpAddressResponse) GoString() string {
	return s.String()
}

func (s *ListLniPrivateIpAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLniPrivateIpAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLniPrivateIpAddressResponse) GetBody() *ListLniPrivateIpAddressResponseBody {
	return s.Body
}

func (s *ListLniPrivateIpAddressResponse) SetHeaders(v map[string]*string) *ListLniPrivateIpAddressResponse {
	s.Headers = v
	return s
}

func (s *ListLniPrivateIpAddressResponse) SetStatusCode(v int32) *ListLniPrivateIpAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLniPrivateIpAddressResponse) SetBody(v *ListLniPrivateIpAddressResponseBody) *ListLniPrivateIpAddressResponse {
	s.Body = v
	return s
}

func (s *ListLniPrivateIpAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
