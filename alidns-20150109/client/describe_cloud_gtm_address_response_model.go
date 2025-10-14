// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmAddressResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmAddressResponseBody) *DescribeCloudGtmAddressResponse
	GetBody() *DescribeCloudGtmAddressResponseBody
}

type DescribeCloudGtmAddressResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmAddressResponse) GetBody() *DescribeCloudGtmAddressResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmAddressResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmAddressResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmAddressResponse) SetStatusCode(v int32) *DescribeCloudGtmAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmAddressResponse) SetBody(v *DescribeCloudGtmAddressResponseBody) *DescribeCloudGtmAddressResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
