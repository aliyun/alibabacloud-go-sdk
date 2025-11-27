// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainPvUvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainPvUvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainPvUvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainPvUvDataResponseBody) *DescribeVsDomainPvUvDataResponse
	GetBody() *DescribeVsDomainPvUvDataResponseBody
}

type DescribeVsDomainPvUvDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainPvUvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainPvUvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainPvUvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainPvUvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainPvUvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainPvUvDataResponse) GetBody() *DescribeVsDomainPvUvDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainPvUvDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainPvUvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainPvUvDataResponse) SetStatusCode(v int32) *DescribeVsDomainPvUvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainPvUvDataResponse) SetBody(v *DescribeVsDomainPvUvDataResponseBody) *DescribeVsDomainPvUvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainPvUvDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
