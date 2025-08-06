// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMediaDistributionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMediaDistributionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMediaDistributionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMediaDistributionResponseBody) *DescribeMediaDistributionResponse
	GetBody() *DescribeMediaDistributionResponseBody
}

type DescribeMediaDistributionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMediaDistributionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMediaDistributionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMediaDistributionResponse) GoString() string {
	return s.String()
}

func (s *DescribeMediaDistributionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMediaDistributionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMediaDistributionResponse) GetBody() *DescribeMediaDistributionResponseBody {
	return s.Body
}

func (s *DescribeMediaDistributionResponse) SetHeaders(v map[string]*string) *DescribeMediaDistributionResponse {
	s.Headers = v
	return s
}

func (s *DescribeMediaDistributionResponse) SetStatusCode(v int32) *DescribeMediaDistributionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMediaDistributionResponse) SetBody(v *DescribeMediaDistributionResponseBody) *DescribeMediaDistributionResponse {
	s.Body = v
	return s
}

func (s *DescribeMediaDistributionResponse) Validate() error {
	return dara.Validate(s)
}
