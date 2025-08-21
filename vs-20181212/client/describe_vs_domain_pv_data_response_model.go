// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainPvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainPvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainPvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainPvDataResponseBody) *DescribeVsDomainPvDataResponse
	GetBody() *DescribeVsDomainPvDataResponseBody
}

type DescribeVsDomainPvDataResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainPvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainPvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainPvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainPvDataResponse) GetBody() *DescribeVsDomainPvDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainPvDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainPvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainPvDataResponse) SetStatusCode(v int32) *DescribeVsDomainPvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainPvDataResponse) SetBody(v *DescribeVsDomainPvDataResponseBody) *DescribeVsDomainPvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainPvDataResponse) Validate() error {
	return dara.Validate(s)
}
