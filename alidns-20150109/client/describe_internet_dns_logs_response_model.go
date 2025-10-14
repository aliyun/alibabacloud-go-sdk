// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInternetDnsLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInternetDnsLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInternetDnsLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInternetDnsLogsResponseBody) *DescribeInternetDnsLogsResponse
	GetBody() *DescribeInternetDnsLogsResponseBody
}

type DescribeInternetDnsLogsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInternetDnsLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInternetDnsLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInternetDnsLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInternetDnsLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInternetDnsLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInternetDnsLogsResponse) GetBody() *DescribeInternetDnsLogsResponseBody {
	return s.Body
}

func (s *DescribeInternetDnsLogsResponse) SetHeaders(v map[string]*string) *DescribeInternetDnsLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInternetDnsLogsResponse) SetStatusCode(v int32) *DescribeInternetDnsLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInternetDnsLogsResponse) SetBody(v *DescribeInternetDnsLogsResponseBody) *DescribeInternetDnsLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeInternetDnsLogsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
