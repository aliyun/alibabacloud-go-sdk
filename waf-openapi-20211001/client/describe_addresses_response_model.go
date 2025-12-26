// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAddressesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAddressesResponseBody) *DescribeAddressesResponse
	GetBody() *DescribeAddressesResponseBody
}

type DescribeAddressesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAddressesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAddressesResponse) GetBody() *DescribeAddressesResponseBody {
	return s.Body
}

func (s *DescribeAddressesResponse) SetHeaders(v map[string]*string) *DescribeAddressesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAddressesResponse) SetStatusCode(v int32) *DescribeAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAddressesResponse) SetBody(v *DescribeAddressesResponseBody) *DescribeAddressesResponse {
	s.Body = v
	return s
}

func (s *DescribeAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
