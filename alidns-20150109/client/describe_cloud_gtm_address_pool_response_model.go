// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudGtmAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudGtmAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudGtmAddressPoolResponseBody) *DescribeCloudGtmAddressPoolResponse
	GetBody() *DescribeCloudGtmAddressPoolResponseBody
}

type DescribeCloudGtmAddressPoolResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudGtmAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudGtmAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudGtmAddressPoolResponse) GetBody() *DescribeCloudGtmAddressPoolResponseBody {
	return s.Body
}

func (s *DescribeCloudGtmAddressPoolResponse) SetHeaders(v map[string]*string) *DescribeCloudGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponse) SetStatusCode(v int32) *DescribeCloudGtmAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponse) SetBody(v *DescribeCloudGtmAddressPoolResponseBody) *DescribeCloudGtmAddressPoolResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
