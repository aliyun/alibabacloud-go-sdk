// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageOsSdkVersionDistributionStatDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUsageOsSdkVersionDistributionStatDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUsageOsSdkVersionDistributionStatDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) *DescribeUsageOsSdkVersionDistributionStatDataResponse
	GetBody() *DescribeUsageOsSdkVersionDistributionStatDataResponseBody
}

type DescribeUsageOsSdkVersionDistributionStatDataResponse struct {
	Headers    map[string]*string                                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageOsSdkVersionDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageOsSdkVersionDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) GetBody() *DescribeUsageOsSdkVersionDistributionStatDataResponseBody {
	return s.Body
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeUsageOsSdkVersionDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) SetStatusCode(v int32) *DescribeUsageOsSdkVersionDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) SetBody(v *DescribeUsageOsSdkVersionDistributionStatDataResponseBody) *DescribeUsageOsSdkVersionDistributionStatDataResponse {
	s.Body = v
	return s
}

func (s *DescribeUsageOsSdkVersionDistributionStatDataResponse) Validate() error {
	return dara.Validate(s)
}
