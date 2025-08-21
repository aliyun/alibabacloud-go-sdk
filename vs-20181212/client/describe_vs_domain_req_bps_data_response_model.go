// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsDomainReqBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsDomainReqBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsDomainReqBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsDomainReqBpsDataResponseBody) *DescribeVsDomainReqBpsDataResponse
	GetBody() *DescribeVsDomainReqBpsDataResponseBody
}

type DescribeVsDomainReqBpsDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsDomainReqBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsDomainReqBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsDomainReqBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsDomainReqBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsDomainReqBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsDomainReqBpsDataResponse) GetBody() *DescribeVsDomainReqBpsDataResponseBody {
	return s.Body
}

func (s *DescribeVsDomainReqBpsDataResponse) SetHeaders(v map[string]*string) *DescribeVsDomainReqBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponse) SetStatusCode(v int32) *DescribeVsDomainReqBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponse) SetBody(v *DescribeVsDomainReqBpsDataResponseBody) *DescribeVsDomainReqBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVsDomainReqBpsDataResponse) Validate() error {
	return dara.Validate(s)
}
