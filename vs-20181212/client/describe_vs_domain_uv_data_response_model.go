// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainUvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainUvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainUvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainUvDataResponseBody) *DescribeVsDomainUvDataResponse
	GetBody() *DescribeVsDomainUvDataResponseBody
}

type DescribeVsDomainUvDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainUvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainUvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainUvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainUvDataResponse) GetBody() *DescribeVsDomainUvDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainUvDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainUvDataResponse) SetStatusCode(v int32) *DescribeVsDomainUvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainUvDataResponse) SetBody(v *DescribeVsDomainUvDataResponseBody) *DescribeVsDomainUvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainUvDataResponse) Validate() error {
	return dara.Validate(s)
}
