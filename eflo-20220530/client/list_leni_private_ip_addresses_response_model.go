// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLeniPrivateIpAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLeniPrivateIpAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLeniPrivateIpAddressesResponse
	GetStatusCode() *int32
	SetBody(v *ListLeniPrivateIpAddressesResponseBody) *ListLeniPrivateIpAddressesResponse
	GetBody() *ListLeniPrivateIpAddressesResponseBody
}

type ListLeniPrivateIpAddressesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLeniPrivateIpAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLeniPrivateIpAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLeniPrivateIpAddressesResponse) GoString() string {
	return s.String()
}

func (s *ListLeniPrivateIpAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLeniPrivateIpAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLeniPrivateIpAddressesResponse) GetBody() *ListLeniPrivateIpAddressesResponseBody {
	return s.Body
}

func (s *ListLeniPrivateIpAddressesResponse) SetHeaders(v map[string]*string) *ListLeniPrivateIpAddressesResponse {
	s.Headers = v
	return s
}

func (s *ListLeniPrivateIpAddressesResponse) SetStatusCode(v int32) *ListLeniPrivateIpAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLeniPrivateIpAddressesResponse) SetBody(v *ListLeniPrivateIpAddressesResponseBody) *ListLeniPrivateIpAddressesResponse {
	s.Body = v
	return s
}

func (s *ListLeniPrivateIpAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
