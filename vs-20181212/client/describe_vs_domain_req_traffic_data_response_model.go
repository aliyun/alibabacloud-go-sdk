// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainReqTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainReqTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainReqTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainReqTrafficDataResponseBody) *DescribeVsDomainReqTrafficDataResponse
	GetBody() *DescribeVsDomainReqTrafficDataResponseBody
}

type DescribeVsDomainReqTrafficDataResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainReqTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainReqTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainReqTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainReqTrafficDataResponse) GetBody() *DescribeVsDomainReqTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainReqTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainReqTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponse) SetStatusCode(v int32) *DescribeVsDomainReqTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponse) SetBody(v *DescribeVsDomainReqTrafficDataResponseBody) *DescribeVsDomainReqTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainReqTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
