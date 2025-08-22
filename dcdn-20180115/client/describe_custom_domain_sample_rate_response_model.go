// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomDomainSampleRateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomDomainSampleRateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomDomainSampleRateResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomDomainSampleRateResponseBody) *DescribeCustomDomainSampleRateResponse
	GetBody() *DescribeCustomDomainSampleRateResponseBody
}

type DescribeCustomDomainSampleRateResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomDomainSampleRateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomDomainSampleRateResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomDomainSampleRateResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomDomainSampleRateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomDomainSampleRateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomDomainSampleRateResponse) GetBody() *DescribeCustomDomainSampleRateResponseBody {
	return s.Body
}

func (s *DescribeCustomDomainSampleRateResponse) SetHeaders(v map[string]*string) *DescribeCustomDomainSampleRateResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomDomainSampleRateResponse) SetStatusCode(v int32) *DescribeCustomDomainSampleRateResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomDomainSampleRateResponse) SetBody(v *DescribeCustomDomainSampleRateResponseBody) *DescribeCustomDomainSampleRateResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomDomainSampleRateResponse) Validate() error {
	return dara.Validate(s)
}
