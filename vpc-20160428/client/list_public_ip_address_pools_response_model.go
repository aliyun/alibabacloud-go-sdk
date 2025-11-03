// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicIpAddressPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublicIpAddressPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublicIpAddressPoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListPublicIpAddressPoolsResponseBody) *ListPublicIpAddressPoolsResponse
	GetBody() *ListPublicIpAddressPoolsResponseBody
}

type ListPublicIpAddressPoolsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicIpAddressPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicIpAddressPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublicIpAddressPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListPublicIpAddressPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublicIpAddressPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublicIpAddressPoolsResponse) GetBody() *ListPublicIpAddressPoolsResponseBody {
	return s.Body
}

func (s *ListPublicIpAddressPoolsResponse) SetHeaders(v map[string]*string) *ListPublicIpAddressPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListPublicIpAddressPoolsResponse) SetStatusCode(v int32) *ListPublicIpAddressPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicIpAddressPoolsResponse) SetBody(v *ListPublicIpAddressPoolsResponseBody) *ListPublicIpAddressPoolsResponse {
	s.Body = v
	return s
}

func (s *ListPublicIpAddressPoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
