// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainTrafficDataResponseBody) *DescribeVsDomainTrafficDataResponse
	GetBody() *DescribeVsDomainTrafficDataResponseBody
}

type DescribeVsDomainTrafficDataResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainTrafficDataResponse) GetBody() *DescribeVsDomainTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainTrafficDataResponse) SetStatusCode(v int32) *DescribeVsDomainTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainTrafficDataResponse) SetBody(v *DescribeVsDomainTrafficDataResponseBody) *DescribeVsDomainTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
