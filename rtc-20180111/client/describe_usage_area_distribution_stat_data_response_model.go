// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsageAreaDistributionStatDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUsageAreaDistributionStatDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUsageAreaDistributionStatDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUsageAreaDistributionStatDataResponseBody) *DescribeUsageAreaDistributionStatDataResponse
	GetBody() *DescribeUsageAreaDistributionStatDataResponseBody
}

type DescribeUsageAreaDistributionStatDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUsageAreaDistributionStatDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUsageAreaDistributionStatDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsageAreaDistributionStatDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUsageAreaDistributionStatDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUsageAreaDistributionStatDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUsageAreaDistributionStatDataResponse) GetBody() *DescribeUsageAreaDistributionStatDataResponseBody {
	return s.Body
}

func (s *DescribeUsageAreaDistributionStatDataResponse) SetHeaders(v map[string]*string) *DescribeUsageAreaDistributionStatDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponse) SetStatusCode(v int32) *DescribeUsageAreaDistributionStatDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponse) SetBody(v *DescribeUsageAreaDistributionStatDataResponseBody) *DescribeUsageAreaDistributionStatDataResponse {
	s.Body = v
	return s
}

func (s *DescribeUsageAreaDistributionStatDataResponse) Validate() error {
	return dara.Validate(s)
}
